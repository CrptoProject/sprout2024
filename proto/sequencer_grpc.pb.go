// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.0
// source: sequencer.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Sequencer_Fetch_FullMethodName  = "/proto.Sequencer/Fetch"
	Sequencer_Report_FullMethodName = "/proto.Sequencer/Report"
)

// SequencerClient is the client API for Sequencer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SequencerClient interface {
	Fetch(ctx context.Context, in *FetchRequest, opts ...grpc.CallOption) (*FetchResponse, error)
	Report(ctx context.Context, in *ReportRequest, opts ...grpc.CallOption) (*ReportResponse, error)
}

type sequencerClient struct {
	cc grpc.ClientConnInterface
}

func NewSequencerClient(cc grpc.ClientConnInterface) SequencerClient {
	return &sequencerClient{cc}
}

func (c *sequencerClient) Fetch(ctx context.Context, in *FetchRequest, opts ...grpc.CallOption) (*FetchResponse, error) {
	out := new(FetchResponse)
	err := c.cc.Invoke(ctx, Sequencer_Fetch_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sequencerClient) Report(ctx context.Context, in *ReportRequest, opts ...grpc.CallOption) (*ReportResponse, error) {
	out := new(ReportResponse)
	err := c.cc.Invoke(ctx, Sequencer_Report_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SequencerServer is the server API for Sequencer service.
// All implementations must embed UnimplementedSequencerServer
// for forward compatibility
type SequencerServer interface {
	Fetch(context.Context, *FetchRequest) (*FetchResponse, error)
	Report(context.Context, *ReportRequest) (*ReportResponse, error)
	mustEmbedUnimplementedSequencerServer()
}

// UnimplementedSequencerServer must be embedded to have forward compatible implementations.
type UnimplementedSequencerServer struct {
}

func (UnimplementedSequencerServer) Fetch(context.Context, *FetchRequest) (*FetchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Fetch not implemented")
}
func (UnimplementedSequencerServer) Report(context.Context, *ReportRequest) (*ReportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Report not implemented")
}
func (UnimplementedSequencerServer) mustEmbedUnimplementedSequencerServer() {}

// UnsafeSequencerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SequencerServer will
// result in compilation errors.
type UnsafeSequencerServer interface {
	mustEmbedUnimplementedSequencerServer()
}

func RegisterSequencerServer(s grpc.ServiceRegistrar, srv SequencerServer) {
	s.RegisterService(&Sequencer_ServiceDesc, srv)
}

func _Sequencer_Fetch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SequencerServer).Fetch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sequencer_Fetch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SequencerServer).Fetch(ctx, req.(*FetchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sequencer_Report_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SequencerServer).Report(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sequencer_Report_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SequencerServer).Report(ctx, req.(*ReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Sequencer_ServiceDesc is the grpc.ServiceDesc for Sequencer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Sequencer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Sequencer",
	HandlerType: (*SequencerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Fetch",
			Handler:    _Sequencer_Fetch_Handler,
		},
		{
			MethodName: "Report",
			Handler:    _Sequencer_Report_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sequencer.proto",
}
