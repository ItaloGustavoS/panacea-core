// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: panacea/oracle/v2/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// MsgRegisterOracle defines the Msg/RegisterOracle request type.
type MsgRegisterOracle struct {
	OracleDetail *Oracle `protobuf:"bytes,1,opt,name=oracle_detail,json=oracleDetail,proto3" json:"oracle_detail,omitempty"`
}

func (m *MsgRegisterOracle) Reset()         { *m = MsgRegisterOracle{} }
func (m *MsgRegisterOracle) String() string { return proto.CompactTextString(m) }
func (*MsgRegisterOracle) ProtoMessage()    {}
func (*MsgRegisterOracle) Descriptor() ([]byte, []int) {
	return fileDescriptor_91b53ec53f5557f6, []int{0}
}
func (m *MsgRegisterOracle) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRegisterOracle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRegisterOracle.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRegisterOracle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRegisterOracle.Merge(m, src)
}
func (m *MsgRegisterOracle) XXX_Size() int {
	return m.Size()
}
func (m *MsgRegisterOracle) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRegisterOracle.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRegisterOracle proto.InternalMessageInfo

func (m *MsgRegisterOracle) GetOracleDetail() *Oracle {
	if m != nil {
		return m.OracleDetail
	}
	return nil
}

// MsgRegisterOracleResponse defines the Msg/RegisterOracle response type.
type MsgRegisterOracleResponse struct {
}

func (m *MsgRegisterOracleResponse) Reset()         { *m = MsgRegisterOracleResponse{} }
func (m *MsgRegisterOracleResponse) String() string { return proto.CompactTextString(m) }
func (*MsgRegisterOracleResponse) ProtoMessage()    {}
func (*MsgRegisterOracleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_91b53ec53f5557f6, []int{1}
}
func (m *MsgRegisterOracleResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRegisterOracleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRegisterOracleResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRegisterOracleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRegisterOracleResponse.Merge(m, src)
}
func (m *MsgRegisterOracleResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgRegisterOracleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRegisterOracleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRegisterOracleResponse proto.InternalMessageInfo

// MsgUpdateOracle defines the Msg/UpdateOracle request type.
type MsgUpdateOracle struct {
	Oracle   string `protobuf:"bytes,1,opt,name=oracle,proto3" json:"oracle,omitempty"`
	Endpoint string `protobuf:"bytes,2,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
}

func (m *MsgUpdateOracle) Reset()         { *m = MsgUpdateOracle{} }
func (m *MsgUpdateOracle) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateOracle) ProtoMessage()    {}
func (*MsgUpdateOracle) Descriptor() ([]byte, []int) {
	return fileDescriptor_91b53ec53f5557f6, []int{2}
}
func (m *MsgUpdateOracle) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateOracle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateOracle.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateOracle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateOracle.Merge(m, src)
}
func (m *MsgUpdateOracle) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateOracle) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateOracle.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateOracle proto.InternalMessageInfo

func (m *MsgUpdateOracle) GetOracle() string {
	if m != nil {
		return m.Oracle
	}
	return ""
}

func (m *MsgUpdateOracle) GetEndpoint() string {
	if m != nil {
		return m.Endpoint
	}
	return ""
}

// MsgUpdateOracleResponse defines the Msg/UpdateOracle response type.
type MsgUpdateOracleResponse struct {
}

func (m *MsgUpdateOracleResponse) Reset()         { *m = MsgUpdateOracleResponse{} }
func (m *MsgUpdateOracleResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateOracleResponse) ProtoMessage()    {}
func (*MsgUpdateOracleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_91b53ec53f5557f6, []int{3}
}
func (m *MsgUpdateOracleResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateOracleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateOracleResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateOracleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateOracleResponse.Merge(m, src)
}
func (m *MsgUpdateOracleResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateOracleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateOracleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateOracleResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgRegisterOracle)(nil), "panacea.oracle.v2.MsgRegisterOracle")
	proto.RegisterType((*MsgRegisterOracleResponse)(nil), "panacea.oracle.v2.MsgRegisterOracleResponse")
	proto.RegisterType((*MsgUpdateOracle)(nil), "panacea.oracle.v2.MsgUpdateOracle")
	proto.RegisterType((*MsgUpdateOracleResponse)(nil), "panacea.oracle.v2.MsgUpdateOracleResponse")
}

func init() { proto.RegisterFile("panacea/oracle/v2/tx.proto", fileDescriptor_91b53ec53f5557f6) }

var fileDescriptor_91b53ec53f5557f6 = []byte{
	// 304 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2a, 0x48, 0xcc, 0x4b,
	0x4c, 0x4e, 0x4d, 0xd4, 0xcf, 0x2f, 0x4a, 0x4c, 0xce, 0x49, 0xd5, 0x2f, 0x33, 0xd2, 0x2f, 0xa9,
	0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x84, 0xca, 0xe9, 0x41, 0xe4, 0xf4, 0xca, 0x8c,
	0xa4, 0xe4, 0x30, 0x95, 0x43, 0x25, 0xc1, 0x5a, 0x94, 0x82, 0xb9, 0x04, 0x7d, 0x8b, 0xd3, 0x83,
	0x52, 0xd3, 0x33, 0x8b, 0x4b, 0x52, 0x8b, 0xfc, 0xc1, 0x52, 0x42, 0x76, 0x5c, 0xbc, 0x10, 0x45,
	0xf1, 0x29, 0xa9, 0x25, 0x89, 0x99, 0x39, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0x92, 0x7a,
	0x18, 0xe6, 0xeb, 0x41, 0x74, 0x04, 0xf1, 0x40, 0x44, 0x5c, 0xc0, 0xca, 0x95, 0xa4, 0xb9, 0x24,
	0x31, 0x0c, 0x0d, 0x4a, 0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x55, 0x72, 0xe5, 0xe2, 0xf7, 0x2d,
	0x4e, 0x0f, 0x2d, 0x48, 0x49, 0x2c, 0x49, 0x85, 0xda, 0x27, 0xc6, 0xc5, 0x06, 0xd1, 0x0f, 0xb6,
	0x88, 0x33, 0x08, 0xca, 0x13, 0x92, 0xe2, 0xe2, 0x48, 0xcd, 0x4b, 0x29, 0xc8, 0xcf, 0xcc, 0x2b,
	0x91, 0x60, 0x02, 0xcb, 0xc0, 0xf9, 0x4a, 0x92, 0x5c, 0xe2, 0x68, 0xc6, 0xc0, 0x6c, 0x30, 0x3a,
	0xcd, 0xc8, 0xc5, 0xec, 0x5b, 0x9c, 0x2e, 0x94, 0xc2, 0xc5, 0x87, 0xe6, 0x31, 0x15, 0x2c, 0x3e,
	0xc0, 0x70, 0xa9, 0x94, 0x0e, 0x31, 0xaa, 0x60, 0xb6, 0x09, 0xc5, 0x71, 0xf1, 0xa0, 0x78, 0x46,
	0x09, 0xbb, 0x6e, 0x64, 0x35, 0x52, 0x5a, 0x84, 0xd5, 0xc0, 0xcc, 0x77, 0xf2, 0x3c, 0xf1, 0x48,
	0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0,
	0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28, 0xfd, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd,
	0xe4, 0xfc, 0x5c, 0xfd, 0xdc, 0xd4, 0x94, 0xcc, 0xa4, 0x9c, 0xfc, 0x64, 0x7d, 0xa8, 0xc1, 0xba,
	0xc9, 0xf9, 0x45, 0xa9, 0xfa, 0x15, 0xb0, 0x68, 0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0x4e, 0x62, 0x03,
	0xc7, 0xb9, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x00, 0x9b, 0xb9, 0x0c, 0x44, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// RegisterOracle defines a method for registration of oracle.
	RegisterOracle(ctx context.Context, in *MsgRegisterOracle, opts ...grpc.CallOption) (*MsgRegisterOracleResponse, error)
	// UpdateOracle defines a method for updating of oracle.
	UpdateOracle(ctx context.Context, in *MsgUpdateOracle, opts ...grpc.CallOption) (*MsgUpdateOracleResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) RegisterOracle(ctx context.Context, in *MsgRegisterOracle, opts ...grpc.CallOption) (*MsgRegisterOracleResponse, error) {
	out := new(MsgRegisterOracleResponse)
	err := c.cc.Invoke(ctx, "/panacea.oracle.v2.Msg/RegisterOracle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateOracle(ctx context.Context, in *MsgUpdateOracle, opts ...grpc.CallOption) (*MsgUpdateOracleResponse, error) {
	out := new(MsgUpdateOracleResponse)
	err := c.cc.Invoke(ctx, "/panacea.oracle.v2.Msg/UpdateOracle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// RegisterOracle defines a method for registration of oracle.
	RegisterOracle(context.Context, *MsgRegisterOracle) (*MsgRegisterOracleResponse, error)
	// UpdateOracle defines a method for updating of oracle.
	UpdateOracle(context.Context, *MsgUpdateOracle) (*MsgUpdateOracleResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) RegisterOracle(ctx context.Context, req *MsgRegisterOracle) (*MsgRegisterOracleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterOracle not implemented")
}
func (*UnimplementedMsgServer) UpdateOracle(ctx context.Context, req *MsgUpdateOracle) (*MsgUpdateOracleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOracle not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_RegisterOracle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRegisterOracle)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RegisterOracle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/panacea.oracle.v2.Msg/RegisterOracle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RegisterOracle(ctx, req.(*MsgRegisterOracle))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateOracle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateOracle)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateOracle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/panacea.oracle.v2.Msg/UpdateOracle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateOracle(ctx, req.(*MsgUpdateOracle))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "panacea.oracle.v2.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterOracle",
			Handler:    _Msg_RegisterOracle_Handler,
		},
		{
			MethodName: "UpdateOracle",
			Handler:    _Msg_UpdateOracle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "panacea/oracle/v2/tx.proto",
}

func (m *MsgRegisterOracle) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRegisterOracle) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRegisterOracle) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.OracleDetail != nil {
		{
			size, err := m.OracleDetail.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgRegisterOracleResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRegisterOracleResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRegisterOracleResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgUpdateOracle) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateOracle) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateOracle) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Endpoint) > 0 {
		i -= len(m.Endpoint)
		copy(dAtA[i:], m.Endpoint)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Endpoint)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Oracle) > 0 {
		i -= len(m.Oracle)
		copy(dAtA[i:], m.Oracle)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Oracle)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdateOracleResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateOracleResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateOracleResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgRegisterOracle) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.OracleDetail != nil {
		l = m.OracleDetail.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgRegisterOracleResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgUpdateOracle) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Oracle)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Endpoint)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgUpdateOracleResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgRegisterOracle) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgRegisterOracle: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRegisterOracle: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OracleDetail", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.OracleDetail == nil {
				m.OracleDetail = &Oracle{}
			}
			if err := m.OracleDetail.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgRegisterOracleResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgRegisterOracleResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRegisterOracleResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgUpdateOracle) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgUpdateOracle: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateOracle: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Oracle", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Oracle = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Endpoint", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Endpoint = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgUpdateOracleResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgUpdateOracleResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateOracleResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
