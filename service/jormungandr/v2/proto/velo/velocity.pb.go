// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.3
// source: essence/velocity.proto

package velo

import (
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

type Property struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Delta *Delta  `protobuf:"bytes,1,opt,name=delta,proto3" json:"delta,omitempty"`
	X     float64 `protobuf:"fixed64,2,opt,name=x,proto3" json:"x,omitempty"`
	Y     float64 `protobuf:"fixed64,3,opt,name=y,proto3" json:"y,omitempty"`
	XV    float64 `protobuf:"fixed64,4,opt,name=x_v,json=xV,proto3" json:"x_v,omitempty"`
	YV    float64 `protobuf:"fixed64,5,opt,name=y_v,json=yV,proto3" json:"y_v,omitempty"`
}

func (x *Property) Reset() {
	*x = Property{}
	if protoimpl.UnsafeEnabled {
		mi := &file_essence_velocity_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Property) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Property) ProtoMessage() {}

func (x *Property) ProtoReflect() protoreflect.Message {
	mi := &file_essence_velocity_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Property.ProtoReflect.Descriptor instead.
func (*Property) Descriptor() ([]byte, []int) {
	return file_essence_velocity_proto_rawDescGZIP(), []int{0}
}

func (x *Property) GetDelta() *Delta {
	if x != nil {
		return x.Delta
	}
	return nil
}

func (x *Property) GetX() float64 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Property) GetY() float64 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *Property) GetXV() float64 {
	if x != nil {
		return x.XV
	}
	return 0
}

func (x *Property) GetYV() float64 {
	if x != nil {
		return x.YV
	}
	return 0
}

type Delta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	XA float64 `protobuf:"fixed64,1,opt,name=x_a,json=xA,proto3" json:"x_a,omitempty"`
	YA float64 `protobuf:"fixed64,2,opt,name=y_a,json=yA,proto3" json:"y_a,omitempty"`
}

func (x *Delta) Reset() {
	*x = Delta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_essence_velocity_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Delta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Delta) ProtoMessage() {}

func (x *Delta) ProtoReflect() protoreflect.Message {
	mi := &file_essence_velocity_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Delta.ProtoReflect.Descriptor instead.
func (*Delta) Descriptor() ([]byte, []int) {
	return file_essence_velocity_proto_rawDescGZIP(), []int{1}
}

func (x *Delta) GetXA() float64 {
	if x != nil {
		return x.XA
	}
	return 0
}

func (x *Delta) GetYA() float64 {
	if x != nil {
		return x.YA
	}
	return 0
}

var File_essence_velocity_proto protoreflect.FileDescriptor

var file_essence_velocity_proto_rawDesc = []byte{
	0x0a, 0x16, 0x65, 0x73, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x2f, 0x76, 0x65, 0x6c, 0x6f, 0x63, 0x69,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x76, 0x65, 0x6c, 0x6f, 0x22, 0x6b,
	0x0a, 0x08, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x12, 0x21, 0x0a, 0x05, 0x64, 0x65,
	0x6c, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x76, 0x65, 0x6c, 0x6f,
	0x2e, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x52, 0x05, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x12, 0x0c, 0x0a,
	0x01, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x79, 0x12, 0x0f, 0x0a, 0x03, 0x78, 0x5f, 0x76,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x02, 0x78, 0x56, 0x12, 0x0f, 0x0a, 0x03, 0x79, 0x5f,
	0x76, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x02, 0x79, 0x56, 0x22, 0x29, 0x0a, 0x05, 0x44,
	0x65, 0x6c, 0x74, 0x61, 0x12, 0x0f, 0x0a, 0x03, 0x78, 0x5f, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x02, 0x78, 0x41, 0x12, 0x0f, 0x0a, 0x03, 0x79, 0x5f, 0x61, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x02, 0x79, 0x41, 0x42, 0x1b, 0x5a, 0x19, 0x6a, 0x6f, 0x72, 0x6d, 0x75, 0x6e,
	0x67, 0x61, 0x6e, 0x64, 0x72, 0x2f, 0x76, 0x32, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76,
	0x65, 0x6c, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_essence_velocity_proto_rawDescOnce sync.Once
	file_essence_velocity_proto_rawDescData = file_essence_velocity_proto_rawDesc
)

func file_essence_velocity_proto_rawDescGZIP() []byte {
	file_essence_velocity_proto_rawDescOnce.Do(func() {
		file_essence_velocity_proto_rawDescData = protoimpl.X.CompressGZIP(file_essence_velocity_proto_rawDescData)
	})
	return file_essence_velocity_proto_rawDescData
}

var file_essence_velocity_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_essence_velocity_proto_goTypes = []interface{}{
	(*Property)(nil), // 0: velo.Property
	(*Delta)(nil),    // 1: velo.Delta
}
var file_essence_velocity_proto_depIdxs = []int32{
	1, // 0: velo.Property.delta:type_name -> velo.Delta
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_essence_velocity_proto_init() }
func file_essence_velocity_proto_init() {
	if File_essence_velocity_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_essence_velocity_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Property); i {
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
		file_essence_velocity_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Delta); i {
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
			RawDescriptor: file_essence_velocity_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_essence_velocity_proto_goTypes,
		DependencyIndexes: file_essence_velocity_proto_depIdxs,
		MessageInfos:      file_essence_velocity_proto_msgTypes,
	}.Build()
	File_essence_velocity_proto = out.File
	file_essence_velocity_proto_rawDesc = nil
	file_essence_velocity_proto_goTypes = nil
	file_essence_velocity_proto_depIdxs = nil
}
