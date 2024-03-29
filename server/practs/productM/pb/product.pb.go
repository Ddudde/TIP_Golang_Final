// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: productM/pb/product.proto

package pb

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_productM_pb_product_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_productM_pb_product_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_productM_pb_product_proto_rawDescGZIP(), []int{0}
}

type GetProductResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status   int64     `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error    string    `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Products []*Prodct `protobuf:"bytes,3,rep,name=products,proto3" json:"products,omitempty"`
}

func (x *GetProductResponse) Reset() {
	*x = GetProductResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_productM_pb_product_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductResponse) ProtoMessage() {}

func (x *GetProductResponse) ProtoReflect() protoreflect.Message {
	mi := &file_productM_pb_product_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductResponse.ProtoReflect.Descriptor instead.
func (*GetProductResponse) Descriptor() ([]byte, []int) {
	return file_productM_pb_product_proto_rawDescGZIP(), []int{1}
}

func (x *GetProductResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *GetProductResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *GetProductResponse) GetProducts() []*Prodct {
	if x != nil {
		return x.Products
	}
	return nil
}

type Prodct struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Zag   string `protobuf:"bytes,1,opt,name=zag,proto3" json:"zag,omitempty"`
	Price int32  `protobuf:"varint,2,opt,name=price,proto3" json:"price,omitempty"`
	Gram  int32  `protobuf:"varint,3,opt,name=gram,proto3" json:"gram,omitempty"`
	Img   string `protobuf:"bytes,4,opt,name=img,proto3" json:"img,omitempty"`
}

func (x *Prodct) Reset() {
	*x = Prodct{}
	if protoimpl.UnsafeEnabled {
		mi := &file_productM_pb_product_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Prodct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Prodct) ProtoMessage() {}

func (x *Prodct) ProtoReflect() protoreflect.Message {
	mi := &file_productM_pb_product_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Prodct.ProtoReflect.Descriptor instead.
func (*Prodct) Descriptor() ([]byte, []int) {
	return file_productM_pb_product_proto_rawDescGZIP(), []int{2}
}

func (x *Prodct) GetZag() string {
	if x != nil {
		return x.Zag
	}
	return ""
}

func (x *Prodct) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Prodct) GetGram() int32 {
	if x != nil {
		return x.Gram
	}
	return 0
}

func (x *Prodct) GetImg() string {
	if x != nil {
		return x.Img
	}
	return ""
}

var File_productM_pb_product_proto protoreflect.FileDescriptor

var file_productM_pb_product_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4d, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x6f, 0x0a,
	0x12, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x12, 0x2b, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x50, 0x72,
	0x6f, 0x64, 0x63, 0x74, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x22, 0x56,
	0x0a, 0x06, 0x50, 0x72, 0x6f, 0x64, 0x63, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x7a, 0x61, 0x67, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x7a, 0x61, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x67, 0x72, 0x61, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x67, 0x72, 0x61, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x6d, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x69, 0x6d, 0x67, 0x32, 0x4d, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x4d, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_productM_pb_product_proto_rawDescOnce sync.Once
	file_productM_pb_product_proto_rawDescData = file_productM_pb_product_proto_rawDesc
)

func file_productM_pb_product_proto_rawDescGZIP() []byte {
	file_productM_pb_product_proto_rawDescOnce.Do(func() {
		file_productM_pb_product_proto_rawDescData = protoimpl.X.CompressGZIP(file_productM_pb_product_proto_rawDescData)
	})
	return file_productM_pb_product_proto_rawDescData
}

var file_productM_pb_product_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_productM_pb_product_proto_goTypes = []interface{}{
	(*Empty)(nil),              // 0: product.Empty
	(*GetProductResponse)(nil), // 1: product.GetProductResponse
	(*Prodct)(nil),             // 2: product.Prodct
}
var file_productM_pb_product_proto_depIdxs = []int32{
	2, // 0: product.GetProductResponse.products:type_name -> product.Prodct
	0, // 1: product.ProductService.GetProduct:input_type -> product.Empty
	1, // 2: product.ProductService.GetProduct:output_type -> product.GetProductResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_productM_pb_product_proto_init() }
func file_productM_pb_product_proto_init() {
	if File_productM_pb_product_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_productM_pb_product_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_productM_pb_product_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProductResponse); i {
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
		file_productM_pb_product_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Prodct); i {
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
			RawDescriptor: file_productM_pb_product_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_productM_pb_product_proto_goTypes,
		DependencyIndexes: file_productM_pb_product_proto_depIdxs,
		MessageInfos:      file_productM_pb_product_proto_msgTypes,
	}.Build()
	File_productM_pb_product_proto = out.File
	file_productM_pb_product_proto_rawDesc = nil
	file_productM_pb_product_proto_goTypes = nil
	file_productM_pb_product_proto_depIdxs = nil
}
