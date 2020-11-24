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
		mi := &file_libros_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChunkLibro) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChunkLibro) ProtoMessage() {}

func (x *ChunkLibro) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ChunkLibro.ProtoReflect.Descriptor instead.
func (*ChunkLibro) Descriptor() ([]byte, []int) {
	return file_libros_proto_rawDescGZIP(), []int{0}
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
		mi := &file_libros_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Mensaje) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Mensaje) ProtoMessage() {}

func (x *Mensaje) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Mensaje.ProtoReflect.Descriptor instead.
func (*Mensaje) Descriptor() ([]byte, []int) {
	return file_libros_proto_rawDescGZIP(), []int{1}
}

func (x *Mensaje) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_libros_proto protoreflect.FileDescriptor

var file_libros_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6c, 0x69, 0x62, 0x72, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x70, 0x62, 0x22, 0x50, 0x0a, 0x0a, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x4c, 0x69, 0x62, 0x72, 0x6f,
	0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x72, 0x74,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x61, 0x72, 0x74, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x63,
	0x68, 0x75, 0x6e, 0x6b, 0x22, 0x1b, 0x0a, 0x07, 0x4d, 0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73,
	0x67, 0x32, 0x6c, 0x0a, 0x06, 0x4c, 0x69, 0x62, 0x72, 0x6f, 0x73, 0x12, 0x2f, 0x0a, 0x0c, 0x47,
	0x75, 0x61, 0x72, 0x64, 0x61, 0x72, 0x4c, 0x69, 0x62, 0x72, 0x6f, 0x12, 0x0e, 0x2e, 0x70, 0x62,
	0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x4c, 0x69, 0x62, 0x72, 0x6f, 0x1a, 0x0b, 0x2e, 0x70, 0x62,
	0x2e, 0x4d, 0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x31, 0x0a, 0x0e,
	0x44, 0x65, 0x73, 0x63, 0x61, 0x72, 0x67, 0x61, 0x72, 0x4c, 0x69, 0x62, 0x72, 0x6f, 0x12, 0x0b,
	0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65, 0x1a, 0x0e, 0x2e, 0x70, 0x62,
	0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x4c, 0x69, 0x62, 0x72, 0x6f, 0x22, 0x00, 0x30, 0x01, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_libros_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_libros_proto_goTypes = []interface{}{
	(*ChunkLibro)(nil), // 0: pb.ChunkLibro
	(*Mensaje)(nil),    // 1: pb.Mensaje
}
var file_libros_proto_depIdxs = []int32{
	0, // 0: pb.Libros.GuardarLibro:input_type -> pb.ChunkLibro
	1, // 1: pb.Libros.DescargarLibro:input_type -> pb.Mensaje
	1, // 2: pb.Libros.GuardarLibro:output_type -> pb.Mensaje
	0, // 3: pb.Libros.DescargarLibro:output_type -> pb.ChunkLibro
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
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
		file_libros_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_libros_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
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
