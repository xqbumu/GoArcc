// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: pb/user_profile.proto

package pb

import (
	types "alfred/protos/types"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/ysugimoto/grpc-graphql-gateway/graphql"
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

type GetUserProfileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetUserProfileRequest) Reset() {
	*x = GetUserProfileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_user_profile_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserProfileRequest) ProtoMessage() {}

func (x *GetUserProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_user_profile_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserProfileRequest.ProtoReflect.Descriptor instead.
func (*GetUserProfileRequest) Descriptor() ([]byte, []int) {
	return file_pb_user_profile_proto_rawDescGZIP(), []int{0}
}

func (x *GetUserProfileRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetUserProfileBySubRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sub string `protobuf:"bytes,1,opt,name=sub,proto3" json:"sub,omitempty"`
}

func (x *GetUserProfileBySubRequest) Reset() {
	*x = GetUserProfileBySubRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_user_profile_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserProfileBySubRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserProfileBySubRequest) ProtoMessage() {}

func (x *GetUserProfileBySubRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_user_profile_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserProfileBySubRequest.ProtoReflect.Descriptor instead.
func (*GetUserProfileBySubRequest) Descriptor() ([]byte, []int) {
	return file_pb_user_profile_proto_rawDescGZIP(), []int{1}
}

func (x *GetUserProfileBySubRequest) GetSub() string {
	if x != nil {
		return x.Sub
	}
	return ""
}

type UserProfile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// external unique_id provided by vcs provider
	Sub string `protobuf:"bytes,2,opt,name=sub,proto3" json:"sub,omitempty"`
	// name of the user
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// a username provided by vcs provider
	UserName string `protobuf:"bytes,4,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	//email of user
	Email string `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	// phone of user
	PhoneNumber    string               `protobuf:"bytes,6,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	ExternalSource types.VCSProviders   `protobuf:"varint,7,opt,name=external_source,json=externalSource,proto3,enum=alfred.types.VCSProviders" json:"external_source,omitempty"`
	ProfilePicUrl  string               `protobuf:"bytes,8,opt,name=profile_pic_url,json=profilePicUrl,proto3" json:"profile_pic_url,omitempty"`
	TokenValidTill *timestamp.Timestamp `protobuf:"bytes,9,opt,name=token_valid_till,json=tokenValidTill,proto3" json:"token_valid_till,omitempty"`
}

func (x *UserProfile) Reset() {
	*x = UserProfile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_user_profile_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserProfile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserProfile) ProtoMessage() {}

func (x *UserProfile) ProtoReflect() protoreflect.Message {
	mi := &file_pb_user_profile_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserProfile.ProtoReflect.Descriptor instead.
func (*UserProfile) Descriptor() ([]byte, []int) {
	return file_pb_user_profile_proto_rawDescGZIP(), []int{2}
}

func (x *UserProfile) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserProfile) GetSub() string {
	if x != nil {
		return x.Sub
	}
	return ""
}

func (x *UserProfile) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserProfile) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *UserProfile) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserProfile) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *UserProfile) GetExternalSource() types.VCSProviders {
	if x != nil {
		return x.ExternalSource
	}
	return types.VCSProviders_UNKNOWN
}

func (x *UserProfile) GetProfilePicUrl() string {
	if x != nil {
		return x.ProfilePicUrl
	}
	return ""
}

func (x *UserProfile) GetTokenValidTill() *timestamp.Timestamp {
	if x != nil {
		return x.TokenValidTill
	}
	return nil
}

var File_pb_user_profile_proto protoreflect.FileDescriptor

var file_pb_user_profile_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x62, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x61, 0x6c, 0x66, 0x72, 0x65, 0x64, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x69, 0x6e,
	0x63, 0x6c, 0x75, 0x64, 0x65, 0x2f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x35, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x03, 0xba, 0x43,
	0x02, 0x08, 0x01, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3c, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x42, 0x79, 0x53, 0x75, 0x62, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x03, 0x73, 0x75, 0x62, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x0c, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x03, 0xba, 0x43, 0x02, 0x08, 0x01,
	0x52, 0x03, 0x73, 0x75, 0x62, 0x22, 0x86, 0x03, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x03, 0x73, 0x75, 0x62, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x0e, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x03, 0x18, 0x64, 0xba, 0x43, 0x02,
	0x08, 0x01, 0x52, 0x03, 0x73, 0x75, 0x62, 0x12, 0x22, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0e, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x01, 0x18, 0x64,
	0xba, 0x43, 0x02, 0x08, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0e, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x00,
	0xfa, 0x42, 0x04, 0x72, 0x02, 0x60, 0x01, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x21,
	0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x4d, 0x0a, 0x0f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x61, 0x6c, 0x66,
	0x72, 0x65, 0x64, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x56, 0x43, 0x53, 0x50, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x20, 0x00,
	0x52, 0x0e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x12, 0x26, 0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x69, 0x63, 0x5f,
	0x75, 0x72, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x50, 0x69, 0x63, 0x55, 0x72, 0x6c, 0x12, 0x44, 0x0a, 0x10, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x74, 0x69, 0x6c, 0x6c, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x54, 0x69, 0x6c, 0x6c, 0x32, 0xe5,
	0x02, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x12,
	0x9d, 0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x12, 0x2d, 0x2e, 0x61, 0x6c, 0x66, 0x72, 0x65, 0x64, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x23, 0x2e, 0x61, 0x6c, 0x66, 0x72, 0x65, 0x64, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x22, 0x37, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x28, 0x12, 0x26,
	0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x2f, 0x67, 0x65, 0x74, 0x2d, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0xba, 0x43, 0x06, 0x12, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12,
	0xb4, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x42, 0x79, 0x53, 0x75, 0x62, 0x12, 0x32, 0x2e, 0x61, 0x6c, 0x66, 0x72, 0x65, 0x64,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x42,
	0x79, 0x53, 0x75, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x61, 0x6c,
	0x66, 0x72, 0x65, 0x64, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x22, 0x44, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x30, 0x12, 0x2e, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x2d, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x67, 0x65, 0x74, 0x2d, 0x75,
	0x73, 0x65, 0x72, 0x2d, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2d, 0x62, 0x79, 0x2d, 0x73,
	0x75, 0x62, 0x2f, 0x7b, 0x73, 0x75, 0x62, 0x7d, 0xba, 0x43, 0x0b, 0x12, 0x09, 0x75, 0x73, 0x65,
	0x72, 0x42, 0x79, 0x53, 0x75, 0x62, 0x42, 0x05, 0x5a, 0x03, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_user_profile_proto_rawDescOnce sync.Once
	file_pb_user_profile_proto_rawDescData = file_pb_user_profile_proto_rawDesc
)

func file_pb_user_profile_proto_rawDescGZIP() []byte {
	file_pb_user_profile_proto_rawDescOnce.Do(func() {
		file_pb_user_profile_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_user_profile_proto_rawDescData)
	})
	return file_pb_user_profile_proto_rawDescData
}

var file_pb_user_profile_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pb_user_profile_proto_goTypes = []interface{}{
	(*GetUserProfileRequest)(nil),      // 0: alfred.user_profile.v1.GetUserProfileRequest
	(*GetUserProfileBySubRequest)(nil), // 1: alfred.user_profile.v1.GetUserProfileBySubRequest
	(*UserProfile)(nil),                // 2: alfred.user_profile.v1.UserProfile
	(types.VCSProviders)(0),            // 3: alfred.types.VCSProviders
	(*timestamp.Timestamp)(nil),        // 4: google.protobuf.Timestamp
}
var file_pb_user_profile_proto_depIdxs = []int32{
	3, // 0: alfred.user_profile.v1.UserProfile.external_source:type_name -> alfred.types.VCSProviders
	4, // 1: alfred.user_profile.v1.UserProfile.token_valid_till:type_name -> google.protobuf.Timestamp
	0, // 2: alfred.user_profile.v1.UserProfiles.GetUserProfile:input_type -> alfred.user_profile.v1.GetUserProfileRequest
	1, // 3: alfred.user_profile.v1.UserProfiles.GetUserProfileBySub:input_type -> alfred.user_profile.v1.GetUserProfileBySubRequest
	2, // 4: alfred.user_profile.v1.UserProfiles.GetUserProfile:output_type -> alfred.user_profile.v1.UserProfile
	2, // 5: alfred.user_profile.v1.UserProfiles.GetUserProfileBySub:output_type -> alfred.user_profile.v1.UserProfile
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pb_user_profile_proto_init() }
func file_pb_user_profile_proto_init() {
	if File_pb_user_profile_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_user_profile_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserProfileRequest); i {
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
		file_pb_user_profile_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserProfileBySubRequest); i {
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
		file_pb_user_profile_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserProfile); i {
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
			RawDescriptor: file_pb_user_profile_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_user_profile_proto_goTypes,
		DependencyIndexes: file_pb_user_profile_proto_depIdxs,
		MessageInfos:      file_pb_user_profile_proto_msgTypes,
	}.Build()
	File_pb_user_profile_proto = out.File
	file_pb_user_profile_proto_rawDesc = nil
	file_pb_user_profile_proto_goTypes = nil
	file_pb_user_profile_proto_depIdxs = nil
}
