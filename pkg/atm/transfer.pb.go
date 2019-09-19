// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/atm/transfer.proto

package atm

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Transfer struct {
	Sender               string   `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Recipient            string   `protobuf:"bytes,2,opt,name=recipient,proto3" json:"recipient,omitempty"`
	Amount               float32  `protobuf:"fixed32,3,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Transfer) Reset()         { *m = Transfer{} }
func (m *Transfer) String() string { return proto.CompactTextString(m) }
func (*Transfer) ProtoMessage()    {}
func (*Transfer) Descriptor() ([]byte, []int) {
	return fileDescriptor_80db3ba6719778af, []int{0}
}

func (m *Transfer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Transfer.Unmarshal(m, b)
}
func (m *Transfer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Transfer.Marshal(b, m, deterministic)
}
func (m *Transfer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transfer.Merge(m, src)
}
func (m *Transfer) XXX_Size() int {
	return xxx_messageInfo_Transfer.Size(m)
}
func (m *Transfer) XXX_DiscardUnknown() {
	xxx_messageInfo_Transfer.DiscardUnknown(m)
}

var xxx_messageInfo_Transfer proto.InternalMessageInfo

func (m *Transfer) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *Transfer) GetRecipient() string {
	if m != nil {
		return m.Recipient
	}
	return ""
}

func (m *Transfer) GetAmount() float32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type TransferHistory struct {
	List                 []*Transfer `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *TransferHistory) Reset()         { *m = TransferHistory{} }
func (m *TransferHistory) String() string { return proto.CompactTextString(m) }
func (*TransferHistory) ProtoMessage()    {}
func (*TransferHistory) Descriptor() ([]byte, []int) {
	return fileDescriptor_80db3ba6719778af, []int{1}
}

func (m *TransferHistory) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransferHistory.Unmarshal(m, b)
}
func (m *TransferHistory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransferHistory.Marshal(b, m, deterministic)
}
func (m *TransferHistory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransferHistory.Merge(m, src)
}
func (m *TransferHistory) XXX_Size() int {
	return xxx_messageInfo_TransferHistory.Size(m)
}
func (m *TransferHistory) XXX_DiscardUnknown() {
	xxx_messageInfo_TransferHistory.DiscardUnknown(m)
}

var xxx_messageInfo_TransferHistory proto.InternalMessageInfo

func (m *TransferHistory) GetList() []*Transfer {
	if m != nil {
		return m.List
	}
	return nil
}

type TransferResponse struct {
	Sender               string   `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Recipient            string   `protobuf:"bytes,2,opt,name=recipient,proto3" json:"recipient,omitempty"`
	Amount               float32  `protobuf:"fixed32,3,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransferResponse) Reset()         { *m = TransferResponse{} }
func (m *TransferResponse) String() string { return proto.CompactTextString(m) }
func (*TransferResponse) ProtoMessage()    {}
func (*TransferResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_80db3ba6719778af, []int{2}
}

func (m *TransferResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransferResponse.Unmarshal(m, b)
}
func (m *TransferResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransferResponse.Marshal(b, m, deterministic)
}
func (m *TransferResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransferResponse.Merge(m, src)
}
func (m *TransferResponse) XXX_Size() int {
	return xxx_messageInfo_TransferResponse.Size(m)
}
func (m *TransferResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TransferResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TransferResponse proto.InternalMessageInfo

func (m *TransferResponse) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *TransferResponse) GetRecipient() string {
	if m != nil {
		return m.Recipient
	}
	return ""
}

func (m *TransferResponse) GetAmount() float32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func init() {
	proto.RegisterType((*Transfer)(nil), "atm.Transfer")
	proto.RegisterType((*TransferHistory)(nil), "atm.TransferHistory")
	proto.RegisterType((*TransferResponse)(nil), "atm.TransferResponse")
}

func init() { proto.RegisterFile("pkg/atm/transfer.proto", fileDescriptor_80db3ba6719778af) }

var fileDescriptor_80db3ba6719778af = []byte{
	// 248 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x90, 0xcf, 0x4e, 0xf3, 0x30,
	0x10, 0xc4, 0x9b, 0xa6, 0xea, 0xf7, 0x65, 0x11, 0x02, 0x59, 0x10, 0x45, 0x81, 0x43, 0xf0, 0x29,
	0x27, 0x47, 0x2a, 0x5c, 0x78, 0x00, 0x24, 0xce, 0x81, 0x03, 0x47, 0xdc, 0x76, 0x1b, 0x59, 0xd4,
	0x7f, 0x64, 0x6f, 0x91, 0xfa, 0xf6, 0x08, 0xd7, 0x01, 0xe5, 0xce, 0x71, 0x67, 0x67, 0xec, 0xdf,
	0x0e, 0x94, 0xee, 0x63, 0xe8, 0x24, 0xe9, 0x8e, 0xbc, 0x34, 0x61, 0x87, 0x5e, 0x38, 0x6f, 0xc9,
	0xb2, 0x5c, 0x92, 0xae, 0x6f, 0x06, 0x6b, 0x87, 0x3d, 0x76, 0x51, 0x5a, 0x1f, 0x76, 0xdd, 0x93,
	0x76, 0x74, 0x3c, 0x39, 0xf8, 0x1b, 0xfc, 0x7f, 0x4d, 0x19, 0x56, 0xc2, 0x32, 0xa0, 0xd9, 0xa2,
	0xaf, 0xb2, 0x26, 0x6b, 0x8b, 0x3e, 0x4d, 0xec, 0x16, 0x0a, 0x8f, 0x1b, 0xe5, 0x14, 0x1a, 0xaa,
	0xe6, 0x71, 0xf5, 0x2b, 0x7c, 0xa7, 0xa4, 0xb6, 0x07, 0x43, 0x55, 0xde, 0x64, 0xed, 0xbc, 0x4f,
	0x13, 0x7f, 0x80, 0x8b, 0xf1, 0xe5, 0x67, 0x15, 0xc8, 0xfa, 0x23, 0xbb, 0x83, 0xc5, 0x5e, 0x05,
	0xaa, 0xb2, 0x26, 0x6f, 0xcf, 0x56, 0xe7, 0x42, 0x92, 0x16, 0xa3, 0xa7, 0x8f, 0x2b, 0xfe, 0x0e,
	0x97, 0x3f, 0x0a, 0x06, 0x67, 0x4d, 0xc0, 0xbf, 0xe5, 0x5a, 0x7d, 0x42, 0x31, 0xfe, 0x10, 0x98,
	0x80, 0xc5, 0x0b, 0x9a, 0x2d, 0x9b, 0xb2, 0xd4, 0xd7, 0x53, 0xb4, 0x04, 0xc2, 0x67, 0xec, 0x11,
	0xfe, 0xf5, 0xb8, 0x41, 0xe5, 0x88, 0x95, 0xe2, 0xd4, 0xab, 0x18, 0x7b, 0x15, 0xb1, 0xd7, 0xfa,
	0x6a, 0x92, 0x4d, 0xa7, 0xf3, 0xd9, 0x7a, 0x19, 0x7d, 0xf7, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x06, 0x1d, 0x3c, 0x2f, 0xac, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TransfersClient is the client API for Transfers service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TransfersClient interface {
	Send(ctx context.Context, in *Transfer, opts ...grpc.CallOption) (*TransferResponse, error)
	Receipt(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*TransferHistory, error)
}

type transfersClient struct {
	cc *grpc.ClientConn
}

func NewTransfersClient(cc *grpc.ClientConn) TransfersClient {
	return &transfersClient{cc}
}

func (c *transfersClient) Send(ctx context.Context, in *Transfer, opts ...grpc.CallOption) (*TransferResponse, error) {
	out := new(TransferResponse)
	err := c.cc.Invoke(ctx, "/atm.Transfers/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transfersClient) Receipt(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*TransferHistory, error) {
	out := new(TransferHistory)
	err := c.cc.Invoke(ctx, "/atm.Transfers/Receipt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransfersServer is the server API for Transfers service.
type TransfersServer interface {
	Send(context.Context, *Transfer) (*TransferResponse, error)
	Receipt(context.Context, *empty.Empty) (*TransferHistory, error)
}

// UnimplementedTransfersServer can be embedded to have forward compatible implementations.
type UnimplementedTransfersServer struct {
}

func (*UnimplementedTransfersServer) Send(ctx context.Context, req *Transfer) (*TransferResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (*UnimplementedTransfersServer) Receipt(ctx context.Context, req *empty.Empty) (*TransferHistory, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Receipt not implemented")
}

func RegisterTransfersServer(s *grpc.Server, srv TransfersServer) {
	s.RegisterService(&_Transfers_serviceDesc, srv)
}

func _Transfers_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transfer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransfersServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/atm.Transfers/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransfersServer).Send(ctx, req.(*Transfer))
	}
	return interceptor(ctx, in, info, handler)
}

func _Transfers_Receipt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransfersServer).Receipt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/atm.Transfers/Receipt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransfersServer).Receipt(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Transfers_serviceDesc = grpc.ServiceDesc{
	ServiceName: "atm.Transfers",
	HandlerType: (*TransfersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _Transfers_Send_Handler,
		},
		{
			MethodName: "Receipt",
			Handler:    _Transfers_Receipt_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/atm/transfer.proto",
}
