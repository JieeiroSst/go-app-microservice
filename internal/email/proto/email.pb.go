//protoc --go_out=. *.proto --go-grpc_out=. *.proto
//protoc --go_out=plugins=grpc:. *.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.5
// source: email.proto

package proto

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type Email struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EmailId     string                 `protobuf:"bytes,1,opt,name=email_id,json=emailId,proto3" json:"email_id,omitempty"`
	To          string                 `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	From        string                 `protobuf:"bytes,3,opt,name=from,proto3" json:"from,omitempty"`
	Body        string                 `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
	Subject     string                 `protobuf:"bytes,5,opt,name=subject,proto3" json:"subject,omitempty"`
	ContentType string                 `protobuf:"bytes,6,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Email) Reset() {
	*x = Email{}
	if protoimpl.UnsafeEnabled {
		mi := &file_email_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Email) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Email) ProtoMessage() {}

func (x *Email) ProtoReflect() protoreflect.Message {
	mi := &file_email_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Email.ProtoReflect.Descriptor instead.
func (*Email) Descriptor() ([]byte, []int) {
	return file_email_proto_rawDescGZIP(), []int{0}
}

func (x *Email) GetEmailId() string {
	if x != nil {
		return x.EmailId
	}
	return ""
}

func (x *Email) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *Email) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *Email) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *Email) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *Email) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

func (x *Email) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type SendEmailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	To      string `protobuf:"bytes,1,opt,name=to,proto3" json:"to,omitempty"`
	Subject string `protobuf:"bytes,2,opt,name=subject,proto3" json:"subject,omitempty"`
	Body    string `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *SendEmailRequest) Reset() {
	*x = SendEmailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_email_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendEmailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendEmailRequest) ProtoMessage() {}

func (x *SendEmailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_email_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendEmailRequest.ProtoReflect.Descriptor instead.
func (*SendEmailRequest) Descriptor() ([]byte, []int) {
	return file_email_proto_rawDescGZIP(), []int{1}
}

func (x *SendEmailRequest) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *SendEmailRequest) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *SendEmailRequest) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

type SendEmailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *SendEmailResponse) Reset() {
	*x = SendEmailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_email_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendEmailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendEmailResponse) ProtoMessage() {}

func (x *SendEmailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_email_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendEmailResponse.ProtoReflect.Descriptor instead.
func (*SendEmailResponse) Descriptor() ([]byte, []int) {
	return file_email_proto_rawDescGZIP(), []int{2}
}

func (x *SendEmailResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type FindEmailByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EmailUuid string `protobuf:"bytes,1,opt,name=email_uuid,json=emailUuid,proto3" json:"email_uuid,omitempty"`
}

func (x *FindEmailByIdRequest) Reset() {
	*x = FindEmailByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_email_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindEmailByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindEmailByIdRequest) ProtoMessage() {}

func (x *FindEmailByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_email_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindEmailByIdRequest.ProtoReflect.Descriptor instead.
func (*FindEmailByIdRequest) Descriptor() ([]byte, []int) {
	return file_email_proto_rawDescGZIP(), []int{3}
}

func (x *FindEmailByIdRequest) GetEmailUuid() string {
	if x != nil {
		return x.EmailUuid
	}
	return ""
}

type FindEmailByIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email *Email `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *FindEmailByIdResponse) Reset() {
	*x = FindEmailByIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_email_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindEmailByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindEmailByIdResponse) ProtoMessage() {}

func (x *FindEmailByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_email_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindEmailByIdResponse.ProtoReflect.Descriptor instead.
func (*FindEmailByIdResponse) Descriptor() ([]byte, []int) {
	return file_email_proto_rawDescGZIP(), []int{4}
}

func (x *FindEmailByIdResponse) GetEmail() *Email {
	if x != nil {
		return x.Email
	}
	return nil
}

type FindEmailsByReceiverRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReceiverEmail string `protobuf:"bytes,1,opt,name=receiver_email,json=receiverEmail,proto3" json:"receiver_email,omitempty"`
	Page          uint64 `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Size          uint64 `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *FindEmailsByReceiverRequest) Reset() {
	*x = FindEmailsByReceiverRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_email_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindEmailsByReceiverRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindEmailsByReceiverRequest) ProtoMessage() {}

func (x *FindEmailsByReceiverRequest) ProtoReflect() protoreflect.Message {
	mi := &file_email_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindEmailsByReceiverRequest.ProtoReflect.Descriptor instead.
func (*FindEmailsByReceiverRequest) Descriptor() ([]byte, []int) {
	return file_email_proto_rawDescGZIP(), []int{5}
}

func (x *FindEmailsByReceiverRequest) GetReceiverEmail() string {
	if x != nil {
		return x.ReceiverEmail
	}
	return ""
}

func (x *FindEmailsByReceiverRequest) GetPage() uint64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *FindEmailsByReceiverRequest) GetSize() uint64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type FindEmailsByReceiverResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Emails     []*Email `protobuf:"bytes,1,rep,name=emails,proto3" json:"emails,omitempty"`
	TotalPages uint64   `protobuf:"varint,2,opt,name=total_pages,json=totalPages,proto3" json:"total_pages,omitempty"`
	TotalCount uint64   `protobuf:"varint,3,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	HasMore    bool     `protobuf:"varint,4,opt,name=has_more,json=hasMore,proto3" json:"has_more,omitempty"`
	Page       uint64   `protobuf:"varint,5,opt,name=page,proto3" json:"page,omitempty"`
	Size       uint64   `protobuf:"varint,6,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *FindEmailsByReceiverResponse) Reset() {
	*x = FindEmailsByReceiverResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_email_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindEmailsByReceiverResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindEmailsByReceiverResponse) ProtoMessage() {}

func (x *FindEmailsByReceiverResponse) ProtoReflect() protoreflect.Message {
	mi := &file_email_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindEmailsByReceiverResponse.ProtoReflect.Descriptor instead.
func (*FindEmailsByReceiverResponse) Descriptor() ([]byte, []int) {
	return file_email_proto_rawDescGZIP(), []int{6}
}

func (x *FindEmailsByReceiverResponse) GetEmails() []*Email {
	if x != nil {
		return x.Emails
	}
	return nil
}

func (x *FindEmailsByReceiverResponse) GetTotalPages() uint64 {
	if x != nil {
		return x.TotalPages
	}
	return 0
}

func (x *FindEmailsByReceiverResponse) GetTotalCount() uint64 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

func (x *FindEmailsByReceiverResponse) GetHasMore() bool {
	if x != nil {
		return x.HasMore
	}
	return false
}

func (x *FindEmailsByReceiverResponse) GetPage() uint64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *FindEmailsByReceiverResponse) GetSize() uint64 {
	if x != nil {
		return x.Size
	}
	return 0
}

var File_email_proto protoreflect.FileDescriptor

var file_email_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd2, 0x01, 0x0a,
	0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x49,
	0x64, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74,
	0x6f, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x22, 0x50, 0x0a, 0x10, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62,
	0x6f, 0x64, 0x79, 0x22, 0x2b, 0x0a, 0x11, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0x35, 0x0a, 0x14, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x42, 0x79, 0x49,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x55, 0x75, 0x69, 0x64, 0x22, 0x42, 0x0a, 0x15, 0x46, 0x69, 0x6e, 0x64, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x29, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x6c, 0x0a, 0x1b, 0x46,
	0x69, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x42, 0x79, 0x52, 0x65, 0x63, 0x65, 0x69,
	0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65,
	0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0xd0, 0x01, 0x0a, 0x1c, 0x46, 0x69,
	0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x42, 0x79, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x06, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52,
	0x06, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x5f, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x50, 0x61, 0x67, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x68, 0x61, 0x73,
	0x5f, 0x6d, 0x6f, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x68, 0x61, 0x73,
	0x4d, 0x6f, 0x72, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x32, 0xa6, 0x02, 0x0a,
	0x0c, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4d, 0x0a,
	0x0a, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x1e, 0x2e, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x58, 0x0a, 0x0d,
	0x46, 0x69, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x42, 0x79, 0x49, 0x64, 0x12, 0x22, 0x2e,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x69, 0x6e,
	0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x23, 0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6d, 0x0a, 0x14, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x73, 0x42, 0x79, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x12, 0x29,
	0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x69,
	0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x42, 0x79, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x73, 0x42, 0x79, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_email_proto_rawDescOnce sync.Once
	file_email_proto_rawDescData = file_email_proto_rawDesc
)

func file_email_proto_rawDescGZIP() []byte {
	file_email_proto_rawDescOnce.Do(func() {
		file_email_proto_rawDescData = protoimpl.X.CompressGZIP(file_email_proto_rawDescData)
	})
	return file_email_proto_rawDescData
}

var file_email_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_email_proto_goTypes = []interface{}{
	(*Email)(nil),                        // 0: emailService.Email
	(*SendEmailRequest)(nil),             // 1: emailService.SendEmailRequest
	(*SendEmailResponse)(nil),            // 2: emailService.SendEmailResponse
	(*FindEmailByIdRequest)(nil),         // 3: emailService.FindEmailByIdRequest
	(*FindEmailByIdResponse)(nil),        // 4: emailService.FindEmailByIdResponse
	(*FindEmailsByReceiverRequest)(nil),  // 5: emailService.FindEmailsByReceiverRequest
	(*FindEmailsByReceiverResponse)(nil), // 6: emailService.FindEmailsByReceiverResponse
	(*timestamppb.Timestamp)(nil),        // 7: google.protobuf.Timestamp
}
var file_email_proto_depIdxs = []int32{
	7, // 0: emailService.Email.created_at:type_name -> google.protobuf.Timestamp
	0, // 1: emailService.FindEmailByIdResponse.email:type_name -> emailService.Email
	0, // 2: emailService.FindEmailsByReceiverResponse.emails:type_name -> emailService.Email
	1, // 3: emailService.EmailService.SendEmails:input_type -> emailService.SendEmailRequest
	3, // 4: emailService.EmailService.FindEmailById:input_type -> emailService.FindEmailByIdRequest
	5, // 5: emailService.EmailService.FindEmailsByReceiver:input_type -> emailService.FindEmailsByReceiverRequest
	2, // 6: emailService.EmailService.SendEmails:output_type -> emailService.SendEmailResponse
	4, // 7: emailService.EmailService.FindEmailById:output_type -> emailService.FindEmailByIdResponse
	6, // 8: emailService.EmailService.FindEmailsByReceiver:output_type -> emailService.FindEmailsByReceiverResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_email_proto_init() }
func file_email_proto_init() {
	if File_email_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_email_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Email); i {
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
		file_email_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendEmailRequest); i {
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
		file_email_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendEmailResponse); i {
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
		file_email_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindEmailByIdRequest); i {
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
		file_email_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindEmailByIdResponse); i {
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
		file_email_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindEmailsByReceiverRequest); i {
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
		file_email_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindEmailsByReceiverResponse); i {
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
			RawDescriptor: file_email_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_email_proto_goTypes,
		DependencyIndexes: file_email_proto_depIdxs,
		MessageInfos:      file_email_proto_msgTypes,
	}.Build()
	File_email_proto = out.File
	file_email_proto_rawDesc = nil
	file_email_proto_goTypes = nil
	file_email_proto_depIdxs = nil
}
