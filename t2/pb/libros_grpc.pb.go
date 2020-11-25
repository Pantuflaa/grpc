// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// LibrosClient is the client API for Libros service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LibrosClient interface {
	GuardarLibro(ctx context.Context, opts ...grpc.CallOption) (Libros_GuardarLibroClient, error)
	DescargarLibro(ctx context.Context, in *Mensaje, opts ...grpc.CallOption) (Libros_DescargarLibroClient, error)
	Propuesta(ctx context.Context, in *Mensaje, opts ...grpc.CallOption) (*RespPropuesta, error)
	PedirLibro(ctx context.Context, in *Mensaje, opts ...grpc.CallOption) (*Mensaje, error)
	Vivo(ctx context.Context, in *Mensaje, opts ...grpc.CallOption) (*Mensaje, error)
}

type librosClient struct {
	cc grpc.ClientConnInterface
}

func NewLibrosClient(cc grpc.ClientConnInterface) LibrosClient {
	return &librosClient{cc}
}

func (c *librosClient) GuardarLibro(ctx context.Context, opts ...grpc.CallOption) (Libros_GuardarLibroClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Libros_serviceDesc.Streams[0], "/pb.Libros/GuardarLibro", opts...)
	if err != nil {
		return nil, err
	}
	x := &librosGuardarLibroClient{stream}
	return x, nil
}

type Libros_GuardarLibroClient interface {
	Send(*ChunkLibro) error
	CloseAndRecv() (*Mensaje, error)
	grpc.ClientStream
}

type librosGuardarLibroClient struct {
	grpc.ClientStream
}

func (x *librosGuardarLibroClient) Send(m *ChunkLibro) error {
	return x.ClientStream.SendMsg(m)
}

func (x *librosGuardarLibroClient) CloseAndRecv() (*Mensaje, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Mensaje)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *librosClient) DescargarLibro(ctx context.Context, in *Mensaje, opts ...grpc.CallOption) (Libros_DescargarLibroClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Libros_serviceDesc.Streams[1], "/pb.Libros/DescargarLibro", opts...)
	if err != nil {
		return nil, err
	}
	x := &librosDescargarLibroClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Libros_DescargarLibroClient interface {
	Recv() (*ChunkLibro, error)
	grpc.ClientStream
}

type librosDescargarLibroClient struct {
	grpc.ClientStream
}

func (x *librosDescargarLibroClient) Recv() (*ChunkLibro, error) {
	m := new(ChunkLibro)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *librosClient) Propuesta(ctx context.Context, in *Mensaje, opts ...grpc.CallOption) (*RespPropuesta, error) {
	out := new(RespPropuesta)
	err := c.cc.Invoke(ctx, "/pb.Libros/Propuesta", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *librosClient) PedirLibro(ctx context.Context, in *Mensaje, opts ...grpc.CallOption) (*Mensaje, error) {
	out := new(Mensaje)
	err := c.cc.Invoke(ctx, "/pb.Libros/PedirLibro", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *librosClient) Vivo(ctx context.Context, in *Mensaje, opts ...grpc.CallOption) (*Mensaje, error) {
	out := new(Mensaje)
	err := c.cc.Invoke(ctx, "/pb.Libros/Vivo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LibrosServer is the server API for Libros service.
// All implementations must embed UnimplementedLibrosServer
// for forward compatibility
type LibrosServer interface {
	GuardarLibro(Libros_GuardarLibroServer) error
	DescargarLibro(*Mensaje, Libros_DescargarLibroServer) error
	Propuesta(context.Context, *Mensaje) (*RespPropuesta, error)
	PedirLibro(context.Context, *Mensaje) (*Mensaje, error)
	Vivo(context.Context, *Mensaje) (*Mensaje, error)
	mustEmbedUnimplementedLibrosServer()
}

// UnimplementedLibrosServer must be embedded to have forward compatible implementations.
type UnimplementedLibrosServer struct {
}

func (UnimplementedLibrosServer) GuardarLibro(Libros_GuardarLibroServer) error {
	return status.Errorf(codes.Unimplemented, "method GuardarLibro not implemented")
}
func (UnimplementedLibrosServer) DescargarLibro(*Mensaje, Libros_DescargarLibroServer) error {
	return status.Errorf(codes.Unimplemented, "method DescargarLibro not implemented")
}
func (UnimplementedLibrosServer) Propuesta(context.Context, *Mensaje) (*RespPropuesta, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Propuesta not implemented")
}
func (UnimplementedLibrosServer) PedirLibro(context.Context, *Mensaje) (*Mensaje, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PedirLibro not implemented")
}
func (UnimplementedLibrosServer) Vivo(context.Context, *Mensaje) (*Mensaje, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Vivo not implemented")
}
func (UnimplementedLibrosServer) mustEmbedUnimplementedLibrosServer() {}

// UnsafeLibrosServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LibrosServer will
// result in compilation errors.
type UnsafeLibrosServer interface {
	mustEmbedUnimplementedLibrosServer()
}

func RegisterLibrosServer(s *grpc.Server, srv LibrosServer) {
	s.RegisterService(&_Libros_serviceDesc, srv)
}

func _Libros_GuardarLibro_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(LibrosServer).GuardarLibro(&librosGuardarLibroServer{stream})
}

type Libros_GuardarLibroServer interface {
	SendAndClose(*Mensaje) error
	Recv() (*ChunkLibro, error)
	grpc.ServerStream
}

type librosGuardarLibroServer struct {
	grpc.ServerStream
}

func (x *librosGuardarLibroServer) SendAndClose(m *Mensaje) error {
	return x.ServerStream.SendMsg(m)
}

func (x *librosGuardarLibroServer) Recv() (*ChunkLibro, error) {
	m := new(ChunkLibro)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Libros_DescargarLibro_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Mensaje)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LibrosServer).DescargarLibro(m, &librosDescargarLibroServer{stream})
}

type Libros_DescargarLibroServer interface {
	Send(*ChunkLibro) error
	grpc.ServerStream
}

type librosDescargarLibroServer struct {
	grpc.ServerStream
}

func (x *librosDescargarLibroServer) Send(m *ChunkLibro) error {
	return x.ServerStream.SendMsg(m)
}

func _Libros_Propuesta_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Mensaje)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibrosServer).Propuesta(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Libros/Propuesta",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibrosServer).Propuesta(ctx, req.(*Mensaje))
	}
	return interceptor(ctx, in, info, handler)
}

func _Libros_PedirLibro_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Mensaje)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibrosServer).PedirLibro(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Libros/PedirLibro",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibrosServer).PedirLibro(ctx, req.(*Mensaje))
	}
	return interceptor(ctx, in, info, handler)
}

func _Libros_Vivo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Mensaje)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibrosServer).Vivo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Libros/Vivo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibrosServer).Vivo(ctx, req.(*Mensaje))
	}
	return interceptor(ctx, in, info, handler)
}

var _Libros_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Libros",
	HandlerType: (*LibrosServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Propuesta",
			Handler:    _Libros_Propuesta_Handler,
		},
		{
			MethodName: "PedirLibro",
			Handler:    _Libros_PedirLibro_Handler,
		},
		{
			MethodName: "Vivo",
			Handler:    _Libros_Vivo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GuardarLibro",
			Handler:       _Libros_GuardarLibro_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "DescargarLibro",
			Handler:       _Libros_DescargarLibro_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "libros.proto",
}
