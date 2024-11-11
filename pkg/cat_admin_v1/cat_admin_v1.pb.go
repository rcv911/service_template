// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: cat_admin_v1/cat_admin_v1.proto

package cat_admin_v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Запрос на создание кота
type CreateCatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`   // Имя кота
	Age   int32  `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`    // Возраст кота
	Color string `protobuf:"bytes,3,opt,name=color,proto3" json:"color,omitempty"` // Цвет шерсти
}

func (x *CreateCatRequest) Reset() {
	*x = CreateCatRequest{}
	mi := &file_cat_admin_v1_cat_admin_v1_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateCatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCatRequest) ProtoMessage() {}

func (x *CreateCatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cat_admin_v1_cat_admin_v1_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCatRequest.ProtoReflect.Descriptor instead.
func (*CreateCatRequest) Descriptor() ([]byte, []int) {
	return file_cat_admin_v1_cat_admin_v1_proto_rawDescGZIP(), []int{0}
}

func (x *CreateCatRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateCatRequest) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *CreateCatRequest) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

// Запрос на получение кота по ID
type GetCatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"` // Идентификатор кота
}

func (x *GetCatRequest) Reset() {
	*x = GetCatRequest{}
	mi := &file_cat_admin_v1_cat_admin_v1_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCatRequest) ProtoMessage() {}

func (x *GetCatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cat_admin_v1_cat_admin_v1_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCatRequest.ProtoReflect.Descriptor instead.
func (*GetCatRequest) Descriptor() ([]byte, []int) {
	return file_cat_admin_v1_cat_admin_v1_proto_rawDescGZIP(), []int{1}
}

func (x *GetCatRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// Запрос на обновление информации о коте
type UpdateCatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`      // Идентификатор кота
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`   // Новое имя кота
	Age   int32  `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`    // Новый возраст кота
	Color string `protobuf:"bytes,4,opt,name=color,proto3" json:"color,omitempty"` // Новый цвет шерсти
}

func (x *UpdateCatRequest) Reset() {
	*x = UpdateCatRequest{}
	mi := &file_cat_admin_v1_cat_admin_v1_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateCatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCatRequest) ProtoMessage() {}

func (x *UpdateCatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cat_admin_v1_cat_admin_v1_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCatRequest.ProtoReflect.Descriptor instead.
func (*UpdateCatRequest) Descriptor() ([]byte, []int) {
	return file_cat_admin_v1_cat_admin_v1_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateCatRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateCatRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateCatRequest) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *UpdateCatRequest) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

// Запрос на удаление кота
type DeleteCatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"` // Идентификатор кота
}

func (x *DeleteCatRequest) Reset() {
	*x = DeleteCatRequest{}
	mi := &file_cat_admin_v1_cat_admin_v1_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteCatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCatRequest) ProtoMessage() {}

func (x *DeleteCatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cat_admin_v1_cat_admin_v1_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCatRequest.ProtoReflect.Descriptor instead.
func (*DeleteCatRequest) Descriptor() ([]byte, []int) {
	return file_cat_admin_v1_cat_admin_v1_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteCatRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// Ответ на успешное удаление кота
type DeleteCatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"` // Удален ли кот успешно
}

func (x *DeleteCatResponse) Reset() {
	*x = DeleteCatResponse{}
	mi := &file_cat_admin_v1_cat_admin_v1_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteCatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCatResponse) ProtoMessage() {}

func (x *DeleteCatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cat_admin_v1_cat_admin_v1_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCatResponse.ProtoReflect.Descriptor instead.
func (*DeleteCatResponse) Descriptor() ([]byte, []int) {
	return file_cat_admin_v1_cat_admin_v1_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteCatResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

// Запрос на получение списка котов
type ListCatsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page     int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`                         // Номер страницы (для пагинации)
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"` // Количество котов на странице
}

func (x *ListCatsRequest) Reset() {
	*x = ListCatsRequest{}
	mi := &file_cat_admin_v1_cat_admin_v1_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListCatsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCatsRequest) ProtoMessage() {}

func (x *ListCatsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cat_admin_v1_cat_admin_v1_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCatsRequest.ProtoReflect.Descriptor instead.
func (*ListCatsRequest) Descriptor() ([]byte, []int) {
	return file_cat_admin_v1_cat_admin_v1_proto_rawDescGZIP(), []int{5}
}

func (x *ListCatsRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListCatsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

// Ответ с данными о коте
type CatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`      // Идентификатор кота
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`   // Имя кота
	Age   int32  `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`    // Возраст кота
	Color string `protobuf:"bytes,4,opt,name=color,proto3" json:"color,omitempty"` // Цвет шерсти
}

func (x *CatResponse) Reset() {
	*x = CatResponse{}
	mi := &file_cat_admin_v1_cat_admin_v1_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CatResponse) ProtoMessage() {}

func (x *CatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cat_admin_v1_cat_admin_v1_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CatResponse.ProtoReflect.Descriptor instead.
func (*CatResponse) Descriptor() ([]byte, []int) {
	return file_cat_admin_v1_cat_admin_v1_proto_rawDescGZIP(), []int{6}
}

func (x *CatResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CatResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CatResponse) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *CatResponse) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

// Ответ со списком котов
type ListCatsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cats       []*CatResponse `protobuf:"bytes,1,rep,name=cats,proto3" json:"cats,omitempty"`                                // Список котов
	TotalCount int32          `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"` // Общее количество котов
}

func (x *ListCatsResponse) Reset() {
	*x = ListCatsResponse{}
	mi := &file_cat_admin_v1_cat_admin_v1_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListCatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCatsResponse) ProtoMessage() {}

func (x *ListCatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cat_admin_v1_cat_admin_v1_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCatsResponse.ProtoReflect.Descriptor instead.
func (*ListCatsResponse) Descriptor() ([]byte, []int) {
	return file_cat_admin_v1_cat_admin_v1_proto_rawDescGZIP(), []int{7}
}

func (x *ListCatsResponse) GetCats() []*CatResponse {
	if x != nil {
		return x.Cats
	}
	return nil
}

func (x *ListCatsResponse) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

var File_cat_admin_v1_cat_admin_v1_proto protoreflect.FileDescriptor

var file_cat_admin_v1_cat_admin_v1_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x63, 0x61, 0x74, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x76, 0x31, 0x2f, 0x63,
	0x61, 0x74, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0c, 0x63, 0x61, 0x74, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x76, 0x31, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4e, 0x0a,
	0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x22, 0x1f, 0x0a,
	0x0d, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x5e,
	0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x22, 0x22,
	0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x2d, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x22, 0x42, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65,
	0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x59, 0x0a, 0x0b, 0x43, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f,
	0x6c, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72,
	0x22, 0x62, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x04, 0x63, 0x61, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x61, 0x74, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x76,
	0x31, 0x2e, 0x43, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04, 0x63,
	0x61, 0x74, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x32, 0xeb, 0x03, 0x0a, 0x0f, 0x43, 0x61, 0x74, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5b, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x43, 0x61, 0x74, 0x12, 0x1e, 0x2e, 0x63, 0x61, 0x74, 0x5f, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x5f, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x63, 0x61, 0x74, 0x5f, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x5f, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x3a, 0x01, 0x2a, 0x22, 0x08, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x61, 0x74, 0x73, 0x12, 0x57, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x12,
	0x1b, 0x2e, 0x63, 0x61, 0x74, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x43, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x63,
	0x61, 0x74, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12,
	0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x74, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x60,
	0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x12, 0x1e, 0x2e, 0x63, 0x61,
	0x74, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x43, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x63, 0x61,
	0x74, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x3a, 0x01,
	0x2a, 0x1a, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x74, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d,
	0x12, 0x63, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x74, 0x12, 0x1e, 0x2e,
	0x63, 0x61, 0x74, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x43, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e,
	0x63, 0x61, 0x74, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x43, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x2a, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x74, 0x73,
	0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x5b, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x74,
	0x73, 0x12, 0x1d, 0x2e, 0x63, 0x61, 0x74, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1e, 0x2e, 0x63, 0x61, 0x74, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x10, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x12, 0x08, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61,
	0x74, 0x73, 0x42, 0x42, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x72, 0x63, 0x76, 0x39, 0x31, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f,
	0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x63, 0x61, 0x74,
	0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x76, 0x31, 0x3b, 0x63, 0x61, 0x74, 0x5f, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x5f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cat_admin_v1_cat_admin_v1_proto_rawDescOnce sync.Once
	file_cat_admin_v1_cat_admin_v1_proto_rawDescData = file_cat_admin_v1_cat_admin_v1_proto_rawDesc
)

func file_cat_admin_v1_cat_admin_v1_proto_rawDescGZIP() []byte {
	file_cat_admin_v1_cat_admin_v1_proto_rawDescOnce.Do(func() {
		file_cat_admin_v1_cat_admin_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_cat_admin_v1_cat_admin_v1_proto_rawDescData)
	})
	return file_cat_admin_v1_cat_admin_v1_proto_rawDescData
}

var file_cat_admin_v1_cat_admin_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_cat_admin_v1_cat_admin_v1_proto_goTypes = []any{
	(*CreateCatRequest)(nil),  // 0: cat_admin_v1.CreateCatRequest
	(*GetCatRequest)(nil),     // 1: cat_admin_v1.GetCatRequest
	(*UpdateCatRequest)(nil),  // 2: cat_admin_v1.UpdateCatRequest
	(*DeleteCatRequest)(nil),  // 3: cat_admin_v1.DeleteCatRequest
	(*DeleteCatResponse)(nil), // 4: cat_admin_v1.DeleteCatResponse
	(*ListCatsRequest)(nil),   // 5: cat_admin_v1.ListCatsRequest
	(*CatResponse)(nil),       // 6: cat_admin_v1.CatResponse
	(*ListCatsResponse)(nil),  // 7: cat_admin_v1.ListCatsResponse
}
var file_cat_admin_v1_cat_admin_v1_proto_depIdxs = []int32{
	6, // 0: cat_admin_v1.ListCatsResponse.cats:type_name -> cat_admin_v1.CatResponse
	0, // 1: cat_admin_v1.CatAdminService.CreateCat:input_type -> cat_admin_v1.CreateCatRequest
	1, // 2: cat_admin_v1.CatAdminService.GetCat:input_type -> cat_admin_v1.GetCatRequest
	2, // 3: cat_admin_v1.CatAdminService.UpdateCat:input_type -> cat_admin_v1.UpdateCatRequest
	3, // 4: cat_admin_v1.CatAdminService.DeleteCat:input_type -> cat_admin_v1.DeleteCatRequest
	5, // 5: cat_admin_v1.CatAdminService.ListCats:input_type -> cat_admin_v1.ListCatsRequest
	6, // 6: cat_admin_v1.CatAdminService.CreateCat:output_type -> cat_admin_v1.CatResponse
	6, // 7: cat_admin_v1.CatAdminService.GetCat:output_type -> cat_admin_v1.CatResponse
	6, // 8: cat_admin_v1.CatAdminService.UpdateCat:output_type -> cat_admin_v1.CatResponse
	4, // 9: cat_admin_v1.CatAdminService.DeleteCat:output_type -> cat_admin_v1.DeleteCatResponse
	7, // 10: cat_admin_v1.CatAdminService.ListCats:output_type -> cat_admin_v1.ListCatsResponse
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_cat_admin_v1_cat_admin_v1_proto_init() }
func file_cat_admin_v1_cat_admin_v1_proto_init() {
	if File_cat_admin_v1_cat_admin_v1_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cat_admin_v1_cat_admin_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cat_admin_v1_cat_admin_v1_proto_goTypes,
		DependencyIndexes: file_cat_admin_v1_cat_admin_v1_proto_depIdxs,
		MessageInfos:      file_cat_admin_v1_cat_admin_v1_proto_msgTypes,
	}.Build()
	File_cat_admin_v1_cat_admin_v1_proto = out.File
	file_cat_admin_v1_cat_admin_v1_proto_rawDesc = nil
	file_cat_admin_v1_cat_admin_v1_proto_goTypes = nil
	file_cat_admin_v1_cat_admin_v1_proto_depIdxs = nil
}
