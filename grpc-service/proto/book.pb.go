// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.21.12
// source: proto/book.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Book struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title         string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Author        string                 `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
	Status        string                 `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Book) Reset() {
	*x = Book{}
	mi := &file_proto_book_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Book) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Book) ProtoMessage() {}

func (x *Book) ProtoReflect() protoreflect.Message {
	mi := &file_proto_book_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Book.ProtoReflect.Descriptor instead.
func (*Book) Descriptor() ([]byte, []int) {
	return file_proto_book_proto_rawDescGZIP(), []int{0}
}

func (x *Book) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Book) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Book) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *Book) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type BookID struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BookID) Reset() {
	*x = BookID{}
	mi := &file_proto_book_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BookID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookID) ProtoMessage() {}

func (x *BookID) ProtoReflect() protoreflect.Message {
	mi := &file_proto_book_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookID.ProtoReflect.Descriptor instead.
func (*BookID) Descriptor() ([]byte, []int) {
	return file_proto_book_proto_rawDescGZIP(), []int{1}
}

func (x *BookID) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type BookResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BookResponse) Reset() {
	*x = BookResponse{}
	mi := &file_proto_book_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BookResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookResponse) ProtoMessage() {}

func (x *BookResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_book_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookResponse.ProtoReflect.Descriptor instead.
func (*BookResponse) Descriptor() ([]byte, []int) {
	return file_proto_book_proto_rawDescGZIP(), []int{2}
}

func (x *BookResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteResponse) Reset() {
	*x = DeleteResponse{}
	mi := &file_proto_book_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResponse) ProtoMessage() {}

func (x *DeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_book_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResponse.ProtoReflect.Descriptor instead.
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return file_proto_book_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_proto_book_proto protoreflect.FileDescriptor

var file_proto_book_proto_rawDesc = string([]byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x22, 0x5c, 0x0a, 0x04, 0x42, 0x6f, 0x6f, 0x6b,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x18, 0x0a, 0x06, 0x42, 0x6f, 0x6f, 0x6b, 0x49, 0x44,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x1e, 0x0a, 0x0c, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x2a, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0xc0, 0x01, 0x0a,
	0x0b, 0x42, 0x6f, 0x6f, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2c, 0x0a, 0x0a,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x0a, 0x2e, 0x62, 0x6f, 0x6f,
	0x6b, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x1a, 0x12, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x42, 0x6f,
	0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x07, 0x47, 0x65,
	0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x0c, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x42, 0x6f, 0x6f,
	0x6b, 0x49, 0x44, 0x1a, 0x0a, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x12,
	0x2c, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x0a, 0x2e,
	0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x1a, 0x12, 0x2e, 0x62, 0x6f, 0x6f, 0x6b,
	0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a,
	0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x0c, 0x2e, 0x62, 0x6f,
	0x6f, 0x6b, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x49, 0x44, 0x1a, 0x14, 0x2e, 0x62, 0x6f, 0x6f, 0x6b,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x72,
	0x74, 0x72, 0x6f, 0x6d, 0x6f, 0x6e, 0x65, 0x2f, 0x6c, 0x63, 0x63, 0x73, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_proto_book_proto_rawDescOnce sync.Once
	file_proto_book_proto_rawDescData []byte
)

func file_proto_book_proto_rawDescGZIP() []byte {
	file_proto_book_proto_rawDescOnce.Do(func() {
		file_proto_book_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_book_proto_rawDesc), len(file_proto_book_proto_rawDesc)))
	})
	return file_proto_book_proto_rawDescData
}

var file_proto_book_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_book_proto_goTypes = []any{
	(*Book)(nil),           // 0: book.Book
	(*BookID)(nil),         // 1: book.BookID
	(*BookResponse)(nil),   // 2: book.BookResponse
	(*DeleteResponse)(nil), // 3: book.DeleteResponse
}
var file_proto_book_proto_depIdxs = []int32{
	0, // 0: book.BookService.CreateBook:input_type -> book.Book
	1, // 1: book.BookService.GetBook:input_type -> book.BookID
	0, // 2: book.BookService.UpdateBook:input_type -> book.Book
	1, // 3: book.BookService.DeleteBook:input_type -> book.BookID
	2, // 4: book.BookService.CreateBook:output_type -> book.BookResponse
	0, // 5: book.BookService.GetBook:output_type -> book.Book
	2, // 6: book.BookService.UpdateBook:output_type -> book.BookResponse
	3, // 7: book.BookService.DeleteBook:output_type -> book.DeleteResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_book_proto_init() }
func file_proto_book_proto_init() {
	if File_proto_book_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_book_proto_rawDesc), len(file_proto_book_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_book_proto_goTypes,
		DependencyIndexes: file_proto_book_proto_depIdxs,
		MessageInfos:      file_proto_book_proto_msgTypes,
	}.Build()
	File_proto_book_proto = out.File
	file_proto_book_proto_goTypes = nil
	file_proto_book_proto_depIdxs = nil
}
