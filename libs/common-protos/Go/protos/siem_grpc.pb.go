// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.2
// source: siem.proto

package protos

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	LogService_SendLog_FullMethodName       = "/siem.LogService/SendLog"
	LogService_SendLogStream_FullMethodName = "/siem.LogService/SendLogStream"
)

// LogServiceClient is the client API for LogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// LogService handles log ingestion
type LogServiceClient interface {
	SendLog(ctx context.Context, in *LogEntry, opts ...grpc.CallOption) (*LogResponse, error)
	SendLogStream(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[LogEntry, LogResponse], error)
}

type logServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLogServiceClient(cc grpc.ClientConnInterface) LogServiceClient {
	return &logServiceClient{cc}
}

func (c *logServiceClient) SendLog(ctx context.Context, in *LogEntry, opts ...grpc.CallOption) (*LogResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LogResponse)
	err := c.cc.Invoke(ctx, LogService_SendLog_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logServiceClient) SendLogStream(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[LogEntry, LogResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &LogService_ServiceDesc.Streams[0], LogService_SendLogStream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[LogEntry, LogResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type LogService_SendLogStreamClient = grpc.ClientStreamingClient[LogEntry, LogResponse]

// LogServiceServer is the server API for LogService service.
// All implementations must embed UnimplementedLogServiceServer
// for forward compatibility.
//
// LogService handles log ingestion
type LogServiceServer interface {
	SendLog(context.Context, *LogEntry) (*LogResponse, error)
	SendLogStream(grpc.ClientStreamingServer[LogEntry, LogResponse]) error
	mustEmbedUnimplementedLogServiceServer()
}

// UnimplementedLogServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedLogServiceServer struct{}

func (UnimplementedLogServiceServer) SendLog(context.Context, *LogEntry) (*LogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendLog not implemented")
}
func (UnimplementedLogServiceServer) SendLogStream(grpc.ClientStreamingServer[LogEntry, LogResponse]) error {
	return status.Errorf(codes.Unimplemented, "method SendLogStream not implemented")
}
func (UnimplementedLogServiceServer) mustEmbedUnimplementedLogServiceServer() {}
func (UnimplementedLogServiceServer) testEmbeddedByValue()                    {}

// UnsafeLogServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LogServiceServer will
// result in compilation errors.
type UnsafeLogServiceServer interface {
	mustEmbedUnimplementedLogServiceServer()
}

func RegisterLogServiceServer(s grpc.ServiceRegistrar, srv LogServiceServer) {
	// If the following call pancis, it indicates UnimplementedLogServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&LogService_ServiceDesc, srv)
}

func _LogService_SendLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogEntry)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogServiceServer).SendLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LogService_SendLog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogServiceServer).SendLog(ctx, req.(*LogEntry))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogService_SendLogStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(LogServiceServer).SendLogStream(&grpc.GenericServerStream[LogEntry, LogResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type LogService_SendLogStreamServer = grpc.ClientStreamingServer[LogEntry, LogResponse]

// LogService_ServiceDesc is the grpc.ServiceDesc for LogService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LogService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "siem.LogService",
	HandlerType: (*LogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendLog",
			Handler:    _LogService_SendLog_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendLogStream",
			Handler:       _LogService_SendLogStream_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "siem.proto",
}

const (
	HealthService_ReportHealth_FullMethodName = "/siem.HealthService/ReportHealth"
)

// HealthServiceClient is the client API for HealthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// HealthService handles device health reporting
type HealthServiceClient interface {
	ReportHealth(ctx context.Context, in *DeviceHealth, opts ...grpc.CallOption) (*HealthResponse, error)
}

type healthServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHealthServiceClient(cc grpc.ClientConnInterface) HealthServiceClient {
	return &healthServiceClient{cc}
}

func (c *healthServiceClient) ReportHealth(ctx context.Context, in *DeviceHealth, opts ...grpc.CallOption) (*HealthResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(HealthResponse)
	err := c.cc.Invoke(ctx, HealthService_ReportHealth_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HealthServiceServer is the server API for HealthService service.
// All implementations must embed UnimplementedHealthServiceServer
// for forward compatibility.
//
// HealthService handles device health reporting
type HealthServiceServer interface {
	ReportHealth(context.Context, *DeviceHealth) (*HealthResponse, error)
	mustEmbedUnimplementedHealthServiceServer()
}

// UnimplementedHealthServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedHealthServiceServer struct{}

func (UnimplementedHealthServiceServer) ReportHealth(context.Context, *DeviceHealth) (*HealthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportHealth not implemented")
}
func (UnimplementedHealthServiceServer) mustEmbedUnimplementedHealthServiceServer() {}
func (UnimplementedHealthServiceServer) testEmbeddedByValue()                       {}

// UnsafeHealthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HealthServiceServer will
// result in compilation errors.
type UnsafeHealthServiceServer interface {
	mustEmbedUnimplementedHealthServiceServer()
}

func RegisterHealthServiceServer(s grpc.ServiceRegistrar, srv HealthServiceServer) {
	// If the following call pancis, it indicates UnimplementedHealthServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&HealthService_ServiceDesc, srv)
}

func _HealthService_ReportHealth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeviceHealth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthServiceServer).ReportHealth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HealthService_ReportHealth_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthServiceServer).ReportHealth(ctx, req.(*DeviceHealth))
	}
	return interceptor(ctx, in, info, handler)
}

// HealthService_ServiceDesc is the grpc.ServiceDesc for HealthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HealthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "siem.HealthService",
	HandlerType: (*HealthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReportHealth",
			Handler:    _HealthService_ReportHealth_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "siem.proto",
}

const (
	AlertService_SendAlert_FullMethodName       = "/siem.AlertService/SendAlert"
	AlertService_SendAlertStream_FullMethodName = "/siem.AlertService/SendAlertStream"
)

// AlertServiceClient is the client API for AlertService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// AlertService handles alert notifications
type AlertServiceClient interface {
	SendAlert(ctx context.Context, in *Alert, opts ...grpc.CallOption) (*AlertResponse, error)
	SendAlertStream(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[Alert, AlertResponse], error)
}

type alertServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAlertServiceClient(cc grpc.ClientConnInterface) AlertServiceClient {
	return &alertServiceClient{cc}
}

func (c *alertServiceClient) SendAlert(ctx context.Context, in *Alert, opts ...grpc.CallOption) (*AlertResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AlertResponse)
	err := c.cc.Invoke(ctx, AlertService_SendAlert_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertServiceClient) SendAlertStream(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[Alert, AlertResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &AlertService_ServiceDesc.Streams[0], AlertService_SendAlertStream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[Alert, AlertResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type AlertService_SendAlertStreamClient = grpc.ClientStreamingClient[Alert, AlertResponse]

// AlertServiceServer is the server API for AlertService service.
// All implementations must embed UnimplementedAlertServiceServer
// for forward compatibility.
//
// AlertService handles alert notifications
type AlertServiceServer interface {
	SendAlert(context.Context, *Alert) (*AlertResponse, error)
	SendAlertStream(grpc.ClientStreamingServer[Alert, AlertResponse]) error
	mustEmbedUnimplementedAlertServiceServer()
}

// UnimplementedAlertServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAlertServiceServer struct{}

func (UnimplementedAlertServiceServer) SendAlert(context.Context, *Alert) (*AlertResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendAlert not implemented")
}
func (UnimplementedAlertServiceServer) SendAlertStream(grpc.ClientStreamingServer[Alert, AlertResponse]) error {
	return status.Errorf(codes.Unimplemented, "method SendAlertStream not implemented")
}
func (UnimplementedAlertServiceServer) mustEmbedUnimplementedAlertServiceServer() {}
func (UnimplementedAlertServiceServer) testEmbeddedByValue()                      {}

// UnsafeAlertServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AlertServiceServer will
// result in compilation errors.
type UnsafeAlertServiceServer interface {
	mustEmbedUnimplementedAlertServiceServer()
}

func RegisterAlertServiceServer(s grpc.ServiceRegistrar, srv AlertServiceServer) {
	// If the following call pancis, it indicates UnimplementedAlertServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AlertService_ServiceDesc, srv)
}

func _AlertService_SendAlert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Alert)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertServiceServer).SendAlert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AlertService_SendAlert_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertServiceServer).SendAlert(ctx, req.(*Alert))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlertService_SendAlertStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AlertServiceServer).SendAlertStream(&grpc.GenericServerStream[Alert, AlertResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type AlertService_SendAlertStreamServer = grpc.ClientStreamingServer[Alert, AlertResponse]

// AlertService_ServiceDesc is the grpc.ServiceDesc for AlertService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AlertService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "siem.AlertService",
	HandlerType: (*AlertServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendAlert",
			Handler:    _AlertService_SendAlert_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendAlertStream",
			Handler:       _AlertService_SendAlertStream_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "siem.proto",
}

const (
	AgentService_RegisterAgent_FullMethodName = "/siem.AgentService/RegisterAgent"
	AgentService_UpdateStatus_FullMethodName  = "/siem.AgentService/UpdateStatus"
)

// AgentServiceClient is the client API for AgentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// AgentService handles agent communication
type AgentServiceClient interface {
	RegisterAgent(ctx context.Context, in *AgentStatus, opts ...grpc.CallOption) (*AgentResponse, error)
	UpdateStatus(ctx context.Context, in *AgentStatus, opts ...grpc.CallOption) (*AgentResponse, error)
}

type agentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAgentServiceClient(cc grpc.ClientConnInterface) AgentServiceClient {
	return &agentServiceClient{cc}
}

func (c *agentServiceClient) RegisterAgent(ctx context.Context, in *AgentStatus, opts ...grpc.CallOption) (*AgentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AgentResponse)
	err := c.cc.Invoke(ctx, AgentService_RegisterAgent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentServiceClient) UpdateStatus(ctx context.Context, in *AgentStatus, opts ...grpc.CallOption) (*AgentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AgentResponse)
	err := c.cc.Invoke(ctx, AgentService_UpdateStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentServiceServer is the server API for AgentService service.
// All implementations must embed UnimplementedAgentServiceServer
// for forward compatibility.
//
// AgentService handles agent communication
type AgentServiceServer interface {
	RegisterAgent(context.Context, *AgentStatus) (*AgentResponse, error)
	UpdateStatus(context.Context, *AgentStatus) (*AgentResponse, error)
	mustEmbedUnimplementedAgentServiceServer()
}

// UnimplementedAgentServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAgentServiceServer struct{}

func (UnimplementedAgentServiceServer) RegisterAgent(context.Context, *AgentStatus) (*AgentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterAgent not implemented")
}
func (UnimplementedAgentServiceServer) UpdateStatus(context.Context, *AgentStatus) (*AgentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStatus not implemented")
}
func (UnimplementedAgentServiceServer) mustEmbedUnimplementedAgentServiceServer() {}
func (UnimplementedAgentServiceServer) testEmbeddedByValue()                      {}

// UnsafeAgentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AgentServiceServer will
// result in compilation errors.
type UnsafeAgentServiceServer interface {
	mustEmbedUnimplementedAgentServiceServer()
}

func RegisterAgentServiceServer(s grpc.ServiceRegistrar, srv AgentServiceServer) {
	// If the following call pancis, it indicates UnimplementedAgentServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AgentService_ServiceDesc, srv)
}

func _AgentService_RegisterAgent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgentStatus)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServiceServer).RegisterAgent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AgentService_RegisterAgent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServiceServer).RegisterAgent(ctx, req.(*AgentStatus))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentService_UpdateStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgentStatus)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServiceServer).UpdateStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AgentService_UpdateStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServiceServer).UpdateStatus(ctx, req.(*AgentStatus))
	}
	return interceptor(ctx, in, info, handler)
}

// AgentService_ServiceDesc is the grpc.ServiceDesc for AgentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AgentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "siem.AgentService",
	HandlerType: (*AgentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterAgent",
			Handler:    _AgentService_RegisterAgent_Handler,
		},
		{
			MethodName: "UpdateStatus",
			Handler:    _AgentService_UpdateStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "siem.proto",
}
