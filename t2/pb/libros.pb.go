// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: libros.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Escritura struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nombre string `protobuf:"bytes,1,opt,name=nombre,proto3" json:"nombre,omitempty"`
	Cant   int32  `protobuf:"varint,2,opt,name=cant,proto3" json:"cant,omitempty"`
	Orden  string `protobuf:"bytes,3,opt,name=orden,proto3" json:"orden,omitempty"`
}

func (x *Escritura) Reset() {
	*x = Escritura{}
	if protoimpl.UnsafeEnabled {
		mi := &file_libros_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Escritura) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Escritura) ProtoMessage() {}

func (x *Escritura) ProtoReflect() protoreflect.Message {
	mi := &file_libros_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Escritura.ProtoReflect.Descriptor instead.
func (*Escritura) Descriptor() ([]byte, []int) {
	return file_libros_proto_rawDescGZIP(), []int{0}
}

func (x *Escritura) GetNombre() string {
	if x != nil {
		return x.Nombre
	}
	return ""
}

func (x *Escritura) GetCant() int32 {
	if x != nil {
		return x.Cant
	}
	return 0
}

func (x *Escritura) GetOrden() string {
	if x != nil {
		return x.Orden
	}
	return ""
}

type ChunkLibro struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nombre string `protobuf:"bytes,1,opt,name=nombre,proto3" json:"nombre,omitempty"`
	Parte  int32  `protobuf:"varint,2,opt,name=parte,proto3" json:"parte,omitempty"`
	Chunk  []byte `protobuf:"bytes,3,opt,name=chunk,proto3" json:"chunk,omitempty"`
}

func (x *ChunkLibro) Reset() {
	*x = ChunkLibro{}
	if protoimpl.UnsafeEnabled {
		mi := &file_libros_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChunkLibro) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChunkLibro) ProtoMessage() {}

func (x *ChunkLibro) ProtoReflect() protoreflect.Message {
	mi := &file_libros_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChunkLibro.ProtoReflect.Descriptor instead.
func (*ChunkLibro) Descriptor() ([]byte, []int) {
	return file_libros_proto_rawDescGZIP(), []int{1}
}

func (x *ChunkLibro) GetNombre() string {
	if x != nil {
		return x.Nombre
	}
	return ""
}

func (x *ChunkLibro) GetParte() int32 {
	if x != nil {
		return x.Parte
	}
	return 0
}

func (x *ChunkLibro) GetChunk() []byte {
	if x != nil {
		return x.Chunk
	}
	return nil
}

type Mensaje struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *Mensaje) Reset() {
	*x = Mensaje{}
	if protoimpl.UnsafeEnabled {
		mi := &file_libros_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Mensaje) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Mensaje) ProtoMessage() {}

func (x *Mensaje) ProtoReflect() protoreflect.Message {
	mi := &file_libros_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Mensaje.ProtoReflect.Descriptor instead.
func (*Mensaje) Descriptor() ([]byte, []int) {
	return file_libros_proto_rawDescGZIP(), []int{2}
}

func (x *Mensaje) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type RespPropuesta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Respuesta string `protobuf:"bytes,1,opt,name=respuesta,proto3" json:"respuesta,omitempty"`
	Razon     string `protobuf:"bytes,2,opt,name=razon,proto3" json:"razon,omitempty"`
}

func (x *RespPropuesta) Reset() {
	*x = RespPropuesta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_libros_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RespPropuesta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RespPropuesta) ProtoMessage() {}

func (x *RespPropuesta) ProtoReflect() protoreflect.Message {
	mi := &file_libros_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RespPropuesta.ProtoReflect.Descriptor instead.
func (*RespPropuesta) Descriptor() ([]byte, []int) {
	return file_libros_proto_rawDescGZIP(), []int{3}
}

func (x *RespPropuesta) GetRespuesta() string {
	if x != nil {
		return x.Respuesta
	}
	return ""
}

func (x *RespPropuesta) GetRazon() string {
	if x != nil {
		return x.Razon
	}
	return ""
}

var File_libros_proto protoreflect.FileDescriptor

var file_libros_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6c, 0x69, 0x62, 0x72, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x70, 0x62, 0x22, 0x4d, 0x0a, 0x09, 0x45, 0x73, 0x63, 0x72, 0x69, 0x74, 0x75, 0x72, 0x61, 0x12,
	0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x61, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x61, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6f,
	0x72, 0x64, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65,
	0x6e, 0x22, 0x50, 0x0a, 0x0a, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x4c, 0x69, 0x62, 0x72, 0x6f, 0x12,
	0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x72, 0x74, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x61, 0x72, 0x74, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x63, 0x68,
	0x75, 0x6e, 0x6b, 0x22, 0x1b, 0x0a, 0x07, 0x4d, 0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67,
	0x22, 0x43, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x70, 0x50, 0x72, 0x6f, 0x70, 0x75, 0x65, 0x73, 0x74,
	0x61, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x73, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x12,
	0x14, 0x0a, 0x05, 0x72, 0x61, 0x7a, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x72, 0x61, 0x7a, 0x6f, 0x6e, 0x32, 0xa1, 0x03, 0x0a, 0x06, 0x4c, 0x69, 0x62, 0x72, 0x6f, 0x73,
	0x12, 0x2f, 0x0a, 0x0c, 0x47, 0x75, 0x61, 0x72, 0x64, 0x61, 0x72, 0x4c, 0x69, 0x62, 0x72, 0x6f,
	0x12, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x4c, 0x69, 0x62, 0x72, 0x6f,
	0x1a, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65, 0x22, 0x00, 0x28,
	0x01, 0x12, 0x2f, 0x0a, 0x0c, 0x52, 0x65, 0x63, 0x69, 0x62, 0x69, 0x72, 0x43, 0x68, 0x75, 0x6e,
	0x6b, 0x12, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x4c, 0x69, 0x62, 0x72,
	0x6f, 0x1a, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65, 0x22, 0x00,
	0x28, 0x01, 0x12, 0x2e, 0x0a, 0x0e, 0x45, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x6d, 0x65, 0x4c,
	0x69, 0x62, 0x72, 0x6f, 0x12, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x45, 0x73, 0x63, 0x72, 0x69, 0x74,
	0x75, 0x72, 0x61, 0x1a, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65,
	0x22, 0x00, 0x12, 0x2c, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x4c, 0x69, 0x62,
	0x72, 0x6f, 0x12, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65, 0x1a,
	0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65, 0x22, 0x00, 0x30, 0x01,
	0x12, 0x27, 0x0a, 0x09, 0x44, 0x61, 0x6d, 0x65, 0x4e, 0x6f, 0x64, 0x6f, 0x73, 0x12, 0x0b, 0x2e,
	0x70, 0x62, 0x2e, 0x4d, 0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65, 0x1a, 0x0b, 0x2e, 0x70, 0x62, 0x2e,
	0x4d, 0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x0e, 0x44, 0x65, 0x73,
	0x63, 0x61, 0x72, 0x67, 0x61, 0x72, 0x4c, 0x69, 0x62, 0x72, 0x6f, 0x12, 0x0b, 0x2e, 0x70, 0x62,
	0x2e, 0x4d, 0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x68,
	0x75, 0x6e, 0x6b, 0x4c, 0x69, 0x62, 0x72, 0x6f, 0x22, 0x00, 0x30, 0x01, 0x12, 0x2d, 0x0a, 0x09,
	0x50, 0x72, 0x6f, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x12, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x4d,
	0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65, 0x1a, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x50, 0x72, 0x6f, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x22, 0x00, 0x12, 0x28, 0x0a, 0x0a, 0x50,
	0x65, 0x64, 0x69, 0x72, 0x4c, 0x69, 0x62, 0x72, 0x6f, 0x12, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x4d,
	0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65, 0x1a, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x6e, 0x73,
	0x61, 0x6a, 0x65, 0x22, 0x00, 0x12, 0x22, 0x0a, 0x04, 0x56, 0x69, 0x76, 0x6f, 0x12, 0x0b, 0x2e,
	0x70, 0x62, 0x2e, 0x4d, 0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65, 0x1a, 0x0b, 0x2e, 0x70, 0x62, 0x2e,
	0x4d, 0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_libros_proto_rawDescOnce sync.Once
	file_libros_proto_rawDescData = file_libros_proto_rawDesc
)

func file_libros_proto_rawDescGZIP() []byte {
	file_libros_proto_rawDescOnce.Do(func() {
		file_libros_proto_rawDescData = protoimpl.X.CompressGZIP(file_libros_proto_rawDescData)
	})
	return file_libros_proto_rawDescData
}

var file_libros_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_libros_proto_goTypes = []interface{}{
	(*Escritura)(nil),     // 0: pb.Escritura
	(*ChunkLibro)(nil),    // 1: pb.ChunkLibro
	(*Mensaje)(nil),       // 2: pb.Mensaje
	(*RespPropuesta)(nil), // 3: pb.RespPropuesta
}
var file_libros_proto_depIdxs = []int32{
	1, // 0: pb.Libros.GuardarLibro:input_type -> pb.ChunkLibro
	1, // 1: pb.Libros.RecibirChunk:input_type -> pb.ChunkLibro
	0, // 2: pb.Libros.EscribemeLibro:input_type -> pb.Escritura
	2, // 3: pb.Libros.ListadoLibro:input_type -> pb.Mensaje
	2, // 4: pb.Libros.DameNodos:input_type -> pb.Mensaje
	2, // 5: pb.Libros.DescargarLibro:input_type -> pb.Mensaje
	2, // 6: pb.Libros.Propuesta:input_type -> pb.Mensaje
	2, // 7: pb.Libros.PedirLibro:input_type -> pb.Mensaje
	2, // 8: pb.Libros.Vivo:input_type -> pb.Mensaje
	2, // 9: pb.Libros.GuardarLibro:output_type -> pb.Mensaje
	2, // 10: pb.Libros.RecibirChunk:output_type -> pb.Mensaje
	2, // 11: pb.Libros.EscribemeLibro:output_type -> pb.Mensaje
	2, // 12: pb.Libros.ListadoLibro:output_type -> pb.Mensaje
	2, // 13: pb.Libros.DameNodos:output_type -> pb.Mensaje
	1, // 14: pb.Libros.DescargarLibro:output_type -> pb.ChunkLibro
	3, // 15: pb.Libros.Propuesta:output_type -> pb.RespPropuesta
	2, // 16: pb.Libros.PedirLibro:output_type -> pb.Mensaje
	2, // 17: pb.Libros.Vivo:output_type -> pb.Mensaje
	9, // [9:18] is the sub-list for method output_type
	0, // [0:9] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_libros_proto_init() }
func file_libros_proto_init() {
	if File_libros_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_libros_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Escritura); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_libros_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChunkLibro); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_libros_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Mensaje); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_libros_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RespPropuesta); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_libros_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_libros_proto_goTypes,
		DependencyIndexes: file_libros_proto_depIdxs,
		MessageInfos:      file_libros_proto_msgTypes,
	}.Build()
	File_libros_proto = out.File
	file_libros_proto_rawDesc = nil
	file_libros_proto_goTypes = nil
	file_libros_proto_depIdxs = nil
}
