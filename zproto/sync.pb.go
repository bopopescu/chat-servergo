// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sync.proto

package zproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// /////////////////////////////////////////////////////////////////////
// ServerAuthReq ==> VoidRsp
// SERVER_AUTH_REQ
type ServerAuthReq struct {
	ServerId   int32  `protobuf:"varint,1,opt,name=server_id,json=serverId" json:"server_id,omitempty"`
	ServerName string `protobuf:"bytes,2,opt,name=server_name,json=serverName" json:"server_name,omitempty"`
}

func (m *ServerAuthReq) Reset()                    { *m = ServerAuthReq{} }
func (m *ServerAuthReq) String() string            { return proto.CompactTextString(m) }
func (*ServerAuthReq) ProtoMessage()               {}
func (*ServerAuthReq) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *ServerAuthReq) GetServerId() int32 {
	if m != nil {
		return m.ServerId
	}
	return 0
}

func (m *ServerAuthReq) GetServerName() string {
	if m != nil {
		return m.ServerName
	}
	return ""
}

type VoidRsp struct {
}

func (m *VoidRsp) Reset()                    { *m = VoidRsp{} }
func (m *VoidRsp) String() string            { return proto.CompactTextString(m) }
func (*VoidRsp) ProtoMessage()               {}
func (*VoidRsp) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

//
type DeliveryUpdatesToUsers struct {
	MyAuthKeyId       int64   `protobuf:"varint,1,opt,name=my_auth_key_id,json=myAuthKeyId" json:"my_auth_key_id,omitempty"`
	MySessionId       int64   `protobuf:"varint,2,opt,name=my_session_id,json=mySessionId" json:"my_session_id,omitempty"`
	MyNetlibSessionId int64   `protobuf:"varint,3,opt,name=my_netlib_session_id,json=myNetlibSessionId" json:"my_netlib_session_id,omitempty"`
	SendtoUserIdList  []int32 `protobuf:"varint,4,rep,packed,name=sendto_user_id_list,json=sendtoUserIdList" json:"sendto_user_id_list,omitempty"`
	// uint32 raw_data_header = 4;
	RawData []byte `protobuf:"bytes,5,opt,name=raw_data,json=rawData,proto3" json:"raw_data,omitempty"`
}

func (m *DeliveryUpdatesToUsers) Reset()                    { *m = DeliveryUpdatesToUsers{} }
func (m *DeliveryUpdatesToUsers) String() string            { return proto.CompactTextString(m) }
func (*DeliveryUpdatesToUsers) ProtoMessage()               {}
func (*DeliveryUpdatesToUsers) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

func (m *DeliveryUpdatesToUsers) GetMyAuthKeyId() int64 {
	if m != nil {
		return m.MyAuthKeyId
	}
	return 0
}

func (m *DeliveryUpdatesToUsers) GetMySessionId() int64 {
	if m != nil {
		return m.MySessionId
	}
	return 0
}

func (m *DeliveryUpdatesToUsers) GetMyNetlibSessionId() int64 {
	if m != nil {
		return m.MyNetlibSessionId
	}
	return 0
}

func (m *DeliveryUpdatesToUsers) GetSendtoUserIdList() []int32 {
	if m != nil {
		return m.SendtoUserIdList
	}
	return nil
}

func (m *DeliveryUpdatesToUsers) GetRawData() []byte {
	if m != nil {
		return m.RawData
	}
	return nil
}

type PushUpdatesData struct {
	AuthKeyId       int64  `protobuf:"varint,1,opt,name=auth_key_id,json=authKeyId" json:"auth_key_id,omitempty"`
	SessionId       int64  `protobuf:"varint,2,opt,name=session_id,json=sessionId" json:"session_id,omitempty"`
	NetlibSessionId int64  `protobuf:"varint,3,opt,name=netlib_session_id,json=netlibSessionId" json:"netlib_session_id,omitempty"`
	RawDataHeader   uint32 `protobuf:"varint,4,opt,name=raw_data_header,json=rawDataHeader" json:"raw_data_header,omitempty"`
	RawData         []byte `protobuf:"bytes,5,opt,name=raw_data,json=rawData,proto3" json:"raw_data,omitempty"`
}

func (m *PushUpdatesData) Reset()                    { *m = PushUpdatesData{} }
func (m *PushUpdatesData) String() string            { return proto.CompactTextString(m) }
func (*PushUpdatesData) ProtoMessage()               {}
func (*PushUpdatesData) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

func (m *PushUpdatesData) GetAuthKeyId() int64 {
	if m != nil {
		return m.AuthKeyId
	}
	return 0
}

func (m *PushUpdatesData) GetSessionId() int64 {
	if m != nil {
		return m.SessionId
	}
	return 0
}

func (m *PushUpdatesData) GetNetlibSessionId() int64 {
	if m != nil {
		return m.NetlibSessionId
	}
	return 0
}

func (m *PushUpdatesData) GetRawDataHeader() uint32 {
	if m != nil {
		return m.RawDataHeader
	}
	return 0
}

func (m *PushUpdatesData) GetRawData() []byte {
	if m != nil {
		return m.RawData
	}
	return nil
}

func init() {
	proto.RegisterType((*ServerAuthReq)(nil), "zproto.ServerAuthReq")
	proto.RegisterType((*VoidRsp)(nil), "zproto.VoidRsp")
	proto.RegisterType((*DeliveryUpdatesToUsers)(nil), "zproto.DeliveryUpdatesToUsers")
	proto.RegisterType((*PushUpdatesData)(nil), "zproto.PushUpdatesData")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for RPCSync service

type RPCSyncClient interface {
	// frontend --> sync
	PushUpdatesStream(ctx context.Context, in *ServerAuthReq, opts ...grpc.CallOption) (RPCSync_PushUpdatesStreamClient, error)
	// rpc ServerAuth(ServerAuthReq) returns (VoidRsp);
	DeliveryUpdates(ctx context.Context, in *DeliveryUpdatesToUsers, opts ...grpc.CallOption) (*VoidRsp, error)
	DeliveryUpdatesNotMe(ctx context.Context, in *DeliveryUpdatesToUsers, opts ...grpc.CallOption) (*VoidRsp, error)
}

type rPCSyncClient struct {
	cc *grpc.ClientConn
}

func NewRPCSyncClient(cc *grpc.ClientConn) RPCSyncClient {
	return &rPCSyncClient{cc}
}

func (c *rPCSyncClient) PushUpdatesStream(ctx context.Context, in *ServerAuthReq, opts ...grpc.CallOption) (RPCSync_PushUpdatesStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_RPCSync_serviceDesc.Streams[0], c.cc, "/zproto.RPCSync/PushUpdatesStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &rPCSyncPushUpdatesStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCSync_PushUpdatesStreamClient interface {
	Recv() (*PushUpdatesData, error)
	grpc.ClientStream
}

type rPCSyncPushUpdatesStreamClient struct {
	grpc.ClientStream
}

func (x *rPCSyncPushUpdatesStreamClient) Recv() (*PushUpdatesData, error) {
	m := new(PushUpdatesData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *rPCSyncClient) DeliveryUpdates(ctx context.Context, in *DeliveryUpdatesToUsers, opts ...grpc.CallOption) (*VoidRsp, error) {
	out := new(VoidRsp)
	err := grpc.Invoke(ctx, "/zproto.RPCSync/DeliveryUpdates", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCSyncClient) DeliveryUpdatesNotMe(ctx context.Context, in *DeliveryUpdatesToUsers, opts ...grpc.CallOption) (*VoidRsp, error) {
	out := new(VoidRsp)
	err := grpc.Invoke(ctx, "/zproto.RPCSync/DeliveryUpdatesNotMe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RPCSync service

type RPCSyncServer interface {
	// frontend --> sync
	PushUpdatesStream(*ServerAuthReq, RPCSync_PushUpdatesStreamServer) error
	// rpc ServerAuth(ServerAuthReq) returns (VoidRsp);
	DeliveryUpdates(context.Context, *DeliveryUpdatesToUsers) (*VoidRsp, error)
	DeliveryUpdatesNotMe(context.Context, *DeliveryUpdatesToUsers) (*VoidRsp, error)
}

func RegisterRPCSyncServer(s *grpc.Server, srv RPCSyncServer) {
	s.RegisterService(&_RPCSync_serviceDesc, srv)
}

func _RPCSync_PushUpdatesStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ServerAuthReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RPCSyncServer).PushUpdatesStream(m, &rPCSyncPushUpdatesStreamServer{stream})
}

type RPCSync_PushUpdatesStreamServer interface {
	Send(*PushUpdatesData) error
	grpc.ServerStream
}

type rPCSyncPushUpdatesStreamServer struct {
	grpc.ServerStream
}

func (x *rPCSyncPushUpdatesStreamServer) Send(m *PushUpdatesData) error {
	return x.ServerStream.SendMsg(m)
}

func _RPCSync_DeliveryUpdates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeliveryUpdatesToUsers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCSyncServer).DeliveryUpdates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zproto.RPCSync/DeliveryUpdates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCSyncServer).DeliveryUpdates(ctx, req.(*DeliveryUpdatesToUsers))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCSync_DeliveryUpdatesNotMe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeliveryUpdatesToUsers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCSyncServer).DeliveryUpdatesNotMe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zproto.RPCSync/DeliveryUpdatesNotMe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCSyncServer).DeliveryUpdatesNotMe(ctx, req.(*DeliveryUpdatesToUsers))
	}
	return interceptor(ctx, in, info, handler)
}

var _RPCSync_serviceDesc = grpc.ServiceDesc{
	ServiceName: "zproto.RPCSync",
	HandlerType: (*RPCSyncServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeliveryUpdates",
			Handler:    _RPCSync_DeliveryUpdates_Handler,
		},
		{
			MethodName: "DeliveryUpdatesNotMe",
			Handler:    _RPCSync_DeliveryUpdatesNotMe_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PushUpdatesStream",
			Handler:       _RPCSync_PushUpdatesStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "sync.proto",
}

func init() { proto.RegisterFile("sync.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 434 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xdf, 0x6e, 0xd3, 0x30,
	0x18, 0xc5, 0xe5, 0x76, 0x5d, 0x97, 0xaf, 0x94, 0x50, 0x6f, 0x40, 0x18, 0x62, 0x44, 0x41, 0xa0,
	0x08, 0x89, 0x82, 0xe0, 0x09, 0x28, 0x43, 0xac, 0x82, 0x55, 0x95, 0xcb, 0xb8, 0xe0, 0xc6, 0x72,
	0xeb, 0x4f, 0x6a, 0x44, 0x9d, 0x14, 0xdb, 0xd9, 0x64, 0x5e, 0x8e, 0x17, 0x81, 0x77, 0x41, 0xf9,
	0x37, 0x46, 0x0b, 0x48, 0xbb, 0xfa, 0xa4, 0x73, 0x4e, 0x92, 0xf3, 0x53, 0x0e, 0x80, 0x71, 0xe9,
	0x62, 0xb8, 0xd6, 0x99, 0xcd, 0xe8, 0xee, 0xb7, 0xf2, 0x46, 0xa7, 0xd0, 0x9f, 0xa1, 0x3e, 0x47,
	0xfd, 0x3a, 0xb7, 0x4b, 0x86, 0x5f, 0xe9, 0x7d, 0xf0, 0x4c, 0x29, 0xf0, 0x44, 0x06, 0x24, 0x24,
	0x71, 0x87, 0xed, 0x55, 0xc2, 0x58, 0xd2, 0x87, 0xd0, 0xab, 0xcd, 0x54, 0x28, 0x0c, 0x5a, 0x21,
	0x89, 0x3d, 0x06, 0x95, 0x34, 0x11, 0x0a, 0x23, 0x0f, 0xba, 0x9f, 0xb2, 0x44, 0x32, 0xb3, 0x8e,
	0x7e, 0x10, 0xb8, 0x73, 0x8c, 0xab, 0xe4, 0x1c, 0xb5, 0x3b, 0x5b, 0x4b, 0x61, 0xd1, 0x7c, 0xcc,
	0xce, 0x0c, 0x6a, 0x43, 0x1f, 0xc1, 0x4d, 0xe5, 0xb8, 0xc8, 0xed, 0x92, 0x7f, 0x41, 0xd7, 0x7c,
	0xa8, 0xcd, 0x7a, 0xca, 0x15, 0x35, 0xde, 0xa3, 0x1b, 0x4b, 0x1a, 0x41, 0x5f, 0x39, 0x6e, 0xd0,
	0x98, 0x24, 0x4b, 0x8b, 0x4c, 0xab, 0xc9, 0xcc, 0x2a, 0x6d, 0x2c, 0xe9, 0x73, 0x38, 0x50, 0x8e,
	0xa7, 0x68, 0x57, 0xc9, 0xfc, 0x6a, 0xb4, 0x5d, 0x46, 0x07, 0xca, 0x4d, 0x4a, 0xeb, 0xf7, 0x03,
	0xcf, 0x60, 0xdf, 0x60, 0x2a, 0x6d, 0xc6, 0x73, 0x53, 0x22, 0xf2, 0x55, 0x62, 0x6c, 0xb0, 0x13,
	0xb6, 0xe3, 0x0e, 0xbb, 0x55, 0x59, 0x45, 0xc7, 0xb1, 0xfc, 0x90, 0x18, 0x4b, 0xef, 0xc1, 0x9e,
	0x16, 0x17, 0x5c, 0x0a, 0x2b, 0x82, 0x4e, 0x48, 0xe2, 0x1b, 0xac, 0xab, 0xc5, 0xc5, 0xb1, 0xb0,
	0x22, 0xfa, 0x4e, 0xc0, 0x9f, 0xe6, 0x66, 0x59, 0xa3, 0x15, 0x1a, 0x3d, 0x82, 0xde, 0x36, 0x94,
	0x27, 0x2e, 0x91, 0x1e, 0x00, 0x6c, 0xf1, 0x78, 0xe6, 0xb2, 0xdc, 0x53, 0x18, 0xfc, 0x0b, 0xc5,
	0x4f, 0x37, 0x40, 0x9e, 0x80, 0xdf, 0x34, 0xe3, 0x4b, 0x14, 0x12, 0x75, 0xb0, 0x13, 0x92, 0xb8,
	0xcf, 0xfa, 0x75, 0xc1, 0x93, 0x52, 0xfc, 0x0f, 0xc1, 0xcb, 0x9f, 0x04, 0xba, 0x6c, 0xfa, 0x66,
	0xe6, 0xd2, 0x05, 0x7d, 0x0b, 0x83, 0x2b, 0x30, 0x33, 0xab, 0x51, 0x28, 0x7a, 0x7b, 0x58, 0x8d,
	0x64, 0xf8, 0xc7, 0x42, 0x0e, 0xef, 0x36, 0xf2, 0x06, 0xfe, 0x0b, 0x42, 0x47, 0xe0, 0x6f, 0xfc,
	0x72, 0x7a, 0xd4, 0xa4, 0xff, 0xbe, 0x85, 0x43, 0xbf, 0xf1, 0xeb, 0xdd, 0xd0, 0x77, 0x70, 0xb0,
	0x11, 0x9d, 0x64, 0xf6, 0x14, 0xaf, 0xfd, 0xa2, 0xd1, 0x63, 0xd8, 0x5f, 0x64, 0x6a, 0x98, 0xe2,
	0x3c, 0x5f, 0x89, 0x44, 0xd5, 0xf6, 0x08, 0x3e, 0x4f, 0x8b, 0x5b, 0x60, 0x9f, 0xb4, 0xa6, 0x64,
	0xbe, 0x5b, 0xca, 0xaf, 0x7e, 0x05, 0x00, 0x00, 0xff, 0xff, 0x54, 0x03, 0x01, 0x04, 0x1e, 0x03,
	0x00, 0x00,
}
