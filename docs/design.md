Disclaimer: This document is in progress, anything can change or can be out of date with current implementation

# High level design of Armada
Armada is job queueing system for multiple Kubernetes clusters.

## Goals for the project
- Easily handle large queues of jobs (million +)
- Enforce fair share over time
- Failure of few clusters should not cause an outage
- Reasonable latency between job submission and job start when there is resource available (in order of 10s)
- Maximize utilization of the cluster
- Smart queue instead of scheduler - implement only minimum logic needed on global level, let cluster scheduler do its own work.
- All components should be highly available

## Data model
### Job
Job is executable unit, currently contains Kubernetes Pod specification.

### Job Set
All jobs are grouped into Job Sets with user specified identifier. Job set represent project or other higher level unit of work. Users can observer events in Job Sets through api.

### Queue
All jobs needs to be placed into queues. Resources allocation is controlled using queues.

Queues has its own priority (lower number makes queue more important). Queue current priority is calculated from combination of resources used by jobs from the queue over time and queue priority. Current priority is used to decide which jobs to run first.

Usual setup maps users or teams one to one to queues to control resource usage.

## Proposed design
![Diagram](./batch-api.svg)

### Cluster Executor
Cluster executor is component running on each Kubernetes cluster. It keeps all pods and nodes information in memory and manages jobs within the cluster.
It proactively reports current state of the cluster and asks for jobs to run.
Executor can also refuse to execute assigned job and return it.

### Armada server
Armada server is central component which manages queues of jobs.
It stores all active jobs in Job database (current implementation use Redis).

#### Accounting
Executor periodically reports resource usage details to Armada server. 
Usage is recorded in database and used to update priorities of individual queues.

#### Job Leasing
Executor periodically ask server for jobs to run reporting available resources. Armada distributes these available resources among queues according to queue current priority. 
Jobs are taken from the top of each queue until the available resources is filled. These jobs are then returned to the executor to be executed on the cluster and marked as Leased with a timestamp to show when the lease began.

The executor must regularly renew the lease of all jobs it leases, otherwise leases expire and jobs will be considered failed and executed on different cluster.

#### Job Events
Jobs events are used to show when a Job reaches a new state, such as submitted, running, completed. They hold generic information about, such event created time along with state specific information, such as completed will hold an exit code.

Armada records events of all jobs against the job set the job belongs to. Events for a given job set are available through the api.

Armada records all necessary events to fully reconstruct state of the job at any time. This allows us to erase all job data from Jobs database after the job finishes and keep only the events.

Current implementation utilise Redis streams to store job events.