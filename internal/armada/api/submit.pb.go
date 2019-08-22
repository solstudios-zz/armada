// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: internal/armada/api/submit.proto

package api

import (
	context "context"
	encoding_binary "encoding/binary"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	grpc "google.golang.org/grpc"
	io "io"
	v1 "k8s.io/api/core/v1"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type JobRequest struct {
	Queue    string      `protobuf:"bytes,1,opt,name=Queue,proto3" json:"Queue,omitempty"`
	JobSetId string      `protobuf:"bytes,2,opt,name=JobSetId,proto3" json:"JobSetId,omitempty"`
	Priority float64     `protobuf:"fixed64,3,opt,name=Priority,proto3" json:"Priority,omitempty"`
	PodSpec  *v1.PodSpec `protobuf:"bytes,4,opt,name=PodSpec,proto3" json:"PodSpec,omitempty"`
}

func (m *JobRequest) Reset()         { *m = JobRequest{} }
func (m *JobRequest) String() string { return proto.CompactTextString(m) }
func (*JobRequest) ProtoMessage()    {}
func (*JobRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_83bbfbf574fac779, []int{0}
}
func (m *JobRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *JobRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_JobRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *JobRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobRequest.Merge(m, src)
}
func (m *JobRequest) XXX_Size() int {
	return m.Size()
}
func (m *JobRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_JobRequest.DiscardUnknown(m)
}

var xxx_messageInfo_JobRequest proto.InternalMessageInfo

func (m *JobRequest) GetQueue() string {
	if m != nil {
		return m.Queue
	}
	return ""
}

func (m *JobRequest) GetJobSetId() string {
	if m != nil {
		return m.JobSetId
	}
	return ""
}

func (m *JobRequest) GetPriority() float64 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *JobRequest) GetPodSpec() *v1.PodSpec {
	if m != nil {
		return m.PodSpec
	}
	return nil
}

type JobSubmitResponse struct {
	JobId string `protobuf:"bytes,1,opt,name=JobId,proto3" json:"JobId,omitempty"`
}

func (m *JobSubmitResponse) Reset()         { *m = JobSubmitResponse{} }
func (m *JobSubmitResponse) String() string { return proto.CompactTextString(m) }
func (*JobSubmitResponse) ProtoMessage()    {}
func (*JobSubmitResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_83bbfbf574fac779, []int{1}
}
func (m *JobSubmitResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *JobSubmitResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_JobSubmitResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *JobSubmitResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobSubmitResponse.Merge(m, src)
}
func (m *JobSubmitResponse) XXX_Size() int {
	return m.Size()
}
func (m *JobSubmitResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_JobSubmitResponse.DiscardUnknown(m)
}

var xxx_messageInfo_JobSubmitResponse proto.InternalMessageInfo

func (m *JobSubmitResponse) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

type Queue struct {
	Name     string  `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Priority float64 `protobuf:"fixed64,2,opt,name=Priority,proto3" json:"Priority,omitempty"`
}

func (m *Queue) Reset()         { *m = Queue{} }
func (m *Queue) String() string { return proto.CompactTextString(m) }
func (*Queue) ProtoMessage()    {}
func (*Queue) Descriptor() ([]byte, []int) {
	return fileDescriptor_83bbfbf574fac779, []int{2}
}
func (m *Queue) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Queue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Queue.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Queue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Queue.Merge(m, src)
}
func (m *Queue) XXX_Size() int {
	return m.Size()
}
func (m *Queue) XXX_DiscardUnknown() {
	xxx_messageInfo_Queue.DiscardUnknown(m)
}

var xxx_messageInfo_Queue proto.InternalMessageInfo

func (m *Queue) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Queue) GetPriority() float64 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func init() {
	proto.RegisterType((*JobRequest)(nil), "api.JobRequest")
	proto.RegisterType((*JobSubmitResponse)(nil), "api.JobSubmitResponse")
	proto.RegisterType((*Queue)(nil), "api.Queue")
}

func init() { proto.RegisterFile("internal/armada/api/submit.proto", fileDescriptor_83bbfbf574fac779) }

var fileDescriptor_83bbfbf574fac779 = []byte{
	// 341 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0x41, 0x4f, 0xfa, 0x30,
	0x18, 0xc6, 0x29, 0xf0, 0xe7, 0x2f, 0xe5, 0x60, 0x6c, 0x0c, 0x59, 0x46, 0xb2, 0x2c, 0x3b, 0xe1,
	0xa5, 0x0b, 0xa8, 0xd1, 0xb3, 0xc6, 0x03, 0x3b, 0x18, 0x1c, 0x9f, 0xa0, 0x63, 0xaf, 0xa4, 0x91,
	0xad, 0xa5, 0xeb, 0x48, 0xf8, 0x16, 0xfa, 0xad, 0x3c, 0x72, 0xf4, 0x68, 0xe0, 0x8b, 0x98, 0xb5,
	0x0c, 0x83, 0xb7, 0xf7, 0x79, 0xdf, 0xa7, 0x7d, 0x7f, 0xed, 0x83, 0x7d, 0x9e, 0x6b, 0x50, 0x39,
	0x5b, 0x86, 0x4c, 0x65, 0x2c, 0x65, 0x21, 0x93, 0x3c, 0x2c, 0xca, 0x24, 0xe3, 0x9a, 0x4a, 0x25,
	0xb4, 0x20, 0x2d, 0x26, 0xb9, 0x3b, 0x58, 0x08, 0xb1, 0x58, 0x42, 0x68, 0x5a, 0x49, 0xf9, 0x1a,
	0x42, 0x26, 0xf5, 0xc6, 0x3a, 0xdc, 0xe0, 0xed, 0xbe, 0xa0, 0x5c, 0x98, 0xa3, 0x73, 0xa1, 0x20,
	0x5c, 0x8f, 0xc2, 0x05, 0xe4, 0xa0, 0x98, 0x86, 0xd4, 0x7a, 0x82, 0x0f, 0x84, 0x71, 0x24, 0x92,
	0x18, 0x56, 0x25, 0x14, 0x9a, 0x5c, 0xe2, 0x7f, 0x2f, 0x25, 0x94, 0xe0, 0x20, 0x1f, 0x0d, 0xbb,
	0xb1, 0x15, 0xc4, 0xc5, 0x67, 0x91, 0x48, 0x66, 0xa0, 0x27, 0xa9, 0xd3, 0x34, 0x83, 0xa3, 0xae,
	0x66, 0x53, 0xc5, 0x85, 0xe2, 0x7a, 0xe3, 0xb4, 0x7c, 0x34, 0x44, 0xf1, 0x51, 0x93, 0x5b, 0xfc,
	0x7f, 0x2a, 0xd2, 0x99, 0x84, 0xb9, 0xd3, 0xf6, 0xd1, 0xb0, 0x37, 0x1e, 0x50, 0x8b, 0x44, 0x99,
	0xe4, 0xb4, 0x42, 0xa2, 0xeb, 0x11, 0x3d, 0x58, 0xe2, 0xda, 0x1b, 0x5c, 0xe1, 0x8b, 0xea, 0x7a,
	0xf3, 0xd8, 0x18, 0x0a, 0x29, 0xf2, 0x02, 0x2a, 0xb2, 0x48, 0x24, 0x93, 0xb4, 0x26, 0x33, 0x22,
	0xb8, 0x3b, 0xf0, 0x12, 0x82, 0xdb, 0xcf, 0x2c, 0xab, 0xb9, 0x4d, 0x7d, 0x82, 0xd6, 0x3c, 0x45,
	0x1b, 0xaf, 0x70, 0xc7, 0x2e, 0x20, 0x37, 0xb8, 0x6b, 0xab, 0x48, 0x24, 0xe4, 0xdc, 0x90, 0xfd,
	0x7e, 0x88, 0xdb, 0xaf, 0x1b, 0x7f, 0x70, 0x46, 0xb8, 0xf7, 0xa8, 0x80, 0x69, 0xb0, 0xeb, 0xb1,
	0xb1, 0x99, 0xda, 0xed, 0x53, 0x1b, 0x0a, 0xad, 0x43, 0xa1, 0x4f, 0x55, 0x28, 0x0f, 0xce, 0xe7,
	0xce, 0x43, 0xdb, 0x9d, 0x87, 0xbe, 0x77, 0x1e, 0x7a, 0xdf, 0x7b, 0x8d, 0xed, 0xde, 0x6b, 0x7c,
	0xed, 0xbd, 0x46, 0xd2, 0x31, 0xce, 0xeb, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb1, 0xcb, 0xc3,
	0xc8, 0xf5, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SubmitClient is the client API for Submit service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SubmitClient interface {
	SubmitJob(ctx context.Context, in *JobRequest, opts ...grpc.CallOption) (*JobSubmitResponse, error)
	CreateQueue(ctx context.Context, in *Queue, opts ...grpc.CallOption) (*types.Empty, error)
}

type submitClient struct {
	cc *grpc.ClientConn
}

func NewSubmitClient(cc *grpc.ClientConn) SubmitClient {
	return &submitClient{cc}
}

func (c *submitClient) SubmitJob(ctx context.Context, in *JobRequest, opts ...grpc.CallOption) (*JobSubmitResponse, error) {
	out := new(JobSubmitResponse)
	err := c.cc.Invoke(ctx, "/api.Submit/SubmitJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *submitClient) CreateQueue(ctx context.Context, in *Queue, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/api.Submit/CreateQueue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SubmitServer is the server API for Submit service.
type SubmitServer interface {
	SubmitJob(context.Context, *JobRequest) (*JobSubmitResponse, error)
	CreateQueue(context.Context, *Queue) (*types.Empty, error)
}

func RegisterSubmitServer(s *grpc.Server, srv SubmitServer) {
	s.RegisterService(&_Submit_serviceDesc, srv)
}

func _Submit_SubmitJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubmitServer).SubmitJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Submit/SubmitJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubmitServer).SubmitJob(ctx, req.(*JobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Submit_CreateQueue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Queue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubmitServer).CreateQueue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Submit/CreateQueue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubmitServer).CreateQueue(ctx, req.(*Queue))
	}
	return interceptor(ctx, in, info, handler)
}

var _Submit_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Submit",
	HandlerType: (*SubmitServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitJob",
			Handler:    _Submit_SubmitJob_Handler,
		},
		{
			MethodName: "CreateQueue",
			Handler:    _Submit_CreateQueue_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/armada/api/submit.proto",
}

func (m *JobRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *JobRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Queue) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSubmit(dAtA, i, uint64(len(m.Queue)))
		i += copy(dAtA[i:], m.Queue)
	}
	if len(m.JobSetId) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintSubmit(dAtA, i, uint64(len(m.JobSetId)))
		i += copy(dAtA[i:], m.JobSetId)
	}
	if m.Priority != 0 {
		dAtA[i] = 0x19
		i++
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Priority))))
		i += 8
	}
	if m.PodSpec != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintSubmit(dAtA, i, uint64(m.PodSpec.Size()))
		n1, err := m.PodSpec.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *JobSubmitResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *JobSubmitResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.JobId) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSubmit(dAtA, i, uint64(len(m.JobId)))
		i += copy(dAtA[i:], m.JobId)
	}
	return i, nil
}

func (m *Queue) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Queue) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSubmit(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if m.Priority != 0 {
		dAtA[i] = 0x11
		i++
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Priority))))
		i += 8
	}
	return i, nil
}

func encodeVarintSubmit(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *JobRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Queue)
	if l > 0 {
		n += 1 + l + sovSubmit(uint64(l))
	}
	l = len(m.JobSetId)
	if l > 0 {
		n += 1 + l + sovSubmit(uint64(l))
	}
	if m.Priority != 0 {
		n += 9
	}
	if m.PodSpec != nil {
		l = m.PodSpec.Size()
		n += 1 + l + sovSubmit(uint64(l))
	}
	return n
}

func (m *JobSubmitResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.JobId)
	if l > 0 {
		n += 1 + l + sovSubmit(uint64(l))
	}
	return n
}

func (m *Queue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovSubmit(uint64(l))
	}
	if m.Priority != 0 {
		n += 9
	}
	return n
}

func sovSubmit(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozSubmit(x uint64) (n int) {
	return sovSubmit(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *JobRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSubmit
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: JobRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: JobRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Queue", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubmit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSubmit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSubmit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Queue = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field JobSetId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubmit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSubmit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSubmit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.JobSetId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Priority", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Priority = float64(math.Float64frombits(v))
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PodSpec", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubmit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSubmit
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSubmit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PodSpec == nil {
				m.PodSpec = &v1.PodSpec{}
			}
			if err := m.PodSpec.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSubmit(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSubmit
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSubmit
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *JobSubmitResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSubmit
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: JobSubmitResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: JobSubmitResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field JobId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubmit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSubmit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSubmit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.JobId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSubmit(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSubmit
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSubmit
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Queue) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSubmit
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Queue: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Queue: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubmit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSubmit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSubmit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Priority", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Priority = float64(math.Float64frombits(v))
		default:
			iNdEx = preIndex
			skippy, err := skipSubmit(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSubmit
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSubmit
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipSubmit(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSubmit
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSubmit
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSubmit
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthSubmit
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthSubmit
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSubmit
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipSubmit(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthSubmit
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthSubmit = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSubmit   = fmt.Errorf("proto: integer overflow")
)