// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: vcs_connection_internal.proto

package pb

import (
	types "alfred/protos/types"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type VCSConnection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//id generated by uuid
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// label of vcs connection
	Label string `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
	// Git Providers
	Provider types.VCSProviders `protobuf:"varint,3,opt,name=provider,proto3,enum=alfred.types.VCSProviders" json:"provider,omitempty"`
	// unique connection_id
	ConnectionId string `protobuf:"bytes,4,opt,name=connection_id,json=connectionId,proto3" json:"connection_id,omitempty"`
	// access token
	AccessToken string `protobuf:"bytes,5,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	// refresh token
	RefreshToken       string               `protobuf:"bytes,6,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	AccessTokenExpiry  *timestamp.Timestamp `protobuf:"bytes,7,opt,name=access_token_expiry,json=accessTokenExpiry,proto3" json:"access_token_expiry,omitempty"`
	RefreshTokenExpiry *timestamp.Timestamp `protobuf:"bytes,8,opt,name=refresh_token_expiry,json=refreshTokenExpiry,proto3" json:"refresh_token_expiry,omitempty"`
	// user can revoke the vcs access
	Revoked bool `protobuf:"varint,9,opt,name=revoked,proto3" json:"revoked,omitempty"`
	// vcs connection must belong to a account_id
	AccountId string `protobuf:"bytes,10,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	//specific user-name required while fetching the repo
	UserName string `protobuf:"bytes,11,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
}

func (x *VCSConnection) Reset() {
	*x = VCSConnection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vcs_connection_internal_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VCSConnection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VCSConnection) ProtoMessage() {}

func (x *VCSConnection) ProtoReflect() protoreflect.Message {
	mi := &file_vcs_connection_internal_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VCSConnection.ProtoReflect.Descriptor instead.
func (*VCSConnection) Descriptor() ([]byte, []int) {
	return file_vcs_connection_internal_proto_rawDescGZIP(), []int{0}
}

func (x *VCSConnection) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *VCSConnection) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *VCSConnection) GetProvider() types.VCSProviders {
	if x != nil {
		return x.Provider
	}
	return types.VCSProviders_UNKNOWN
}

func (x *VCSConnection) GetConnectionId() string {
	if x != nil {
		return x.ConnectionId
	}
	return ""
}

func (x *VCSConnection) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *VCSConnection) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

func (x *VCSConnection) GetAccessTokenExpiry() *timestamp.Timestamp {
	if x != nil {
		return x.AccessTokenExpiry
	}
	return nil
}

func (x *VCSConnection) GetRefreshTokenExpiry() *timestamp.Timestamp {
	if x != nil {
		return x.RefreshTokenExpiry
	}
	return nil
}

func (x *VCSConnection) GetRevoked() bool {
	if x != nil {
		return x.Revoked
	}
	return false
}

func (x *VCSConnection) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *VCSConnection) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

type AccountVCSConnection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Provider  types.VCSProviders `protobuf:"varint,1,opt,name=provider,proto3,enum=alfred.types.VCSProviders" json:"provider,omitempty"`
	AccountId string             `protobuf:"bytes,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	Label     string             `protobuf:"bytes,3,opt,name=label,proto3" json:"label,omitempty"`
}

func (x *AccountVCSConnection) Reset() {
	*x = AccountVCSConnection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vcs_connection_internal_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountVCSConnection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountVCSConnection) ProtoMessage() {}

func (x *AccountVCSConnection) ProtoReflect() protoreflect.Message {
	mi := &file_vcs_connection_internal_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountVCSConnection.ProtoReflect.Descriptor instead.
func (*AccountVCSConnection) Descriptor() ([]byte, []int) {
	return file_vcs_connection_internal_proto_rawDescGZIP(), []int{1}
}

func (x *AccountVCSConnection) GetProvider() types.VCSProviders {
	if x != nil {
		return x.Provider
	}
	return types.VCSProviders_UNKNOWN
}

func (x *AccountVCSConnection) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *AccountVCSConnection) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

type RevokeVCSTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Provider types.VCSProviders `protobuf:"varint,2,opt,name=provider,proto3,enum=alfred.types.VCSProviders" json:"provider,omitempty"`
	VcsId    string             `protobuf:"bytes,1,opt,name=vcs_id,json=vcsId,proto3" json:"vcs_id,omitempty"`
}

func (x *RevokeVCSTokenRequest) Reset() {
	*x = RevokeVCSTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vcs_connection_internal_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RevokeVCSTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RevokeVCSTokenRequest) ProtoMessage() {}

func (x *RevokeVCSTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vcs_connection_internal_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RevokeVCSTokenRequest.ProtoReflect.Descriptor instead.
func (*RevokeVCSTokenRequest) Descriptor() ([]byte, []int) {
	return file_vcs_connection_internal_proto_rawDescGZIP(), []int{2}
}

func (x *RevokeVCSTokenRequest) GetProvider() types.VCSProviders {
	if x != nil {
		return x.Provider
	}
	return types.VCSProviders_UNKNOWN
}

func (x *RevokeVCSTokenRequest) GetVcsId() string {
	if x != nil {
		return x.VcsId
	}
	return ""
}

type RenewVCSTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Provider types.VCSProviders `protobuf:"varint,1,opt,name=provider,proto3,enum=alfred.types.VCSProviders" json:"provider,omitempty"`
}

func (x *RenewVCSTokenRequest) Reset() {
	*x = RenewVCSTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vcs_connection_internal_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RenewVCSTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RenewVCSTokenRequest) ProtoMessage() {}

func (x *RenewVCSTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vcs_connection_internal_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RenewVCSTokenRequest.ProtoReflect.Descriptor instead.
func (*RenewVCSTokenRequest) Descriptor() ([]byte, []int) {
	return file_vcs_connection_internal_proto_rawDescGZIP(), []int{3}
}

func (x *RenewVCSTokenRequest) GetProvider() types.VCSProviders {
	if x != nil {
		return x.Provider
	}
	return types.VCSProviders_UNKNOWN
}

type CreateVCSConnectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VcsConnection *VCSConnection `protobuf:"bytes,1,opt,name=vcs_connection,json=vcsConnection,proto3" json:"vcs_connection,omitempty"`
}

func (x *CreateVCSConnectionRequest) Reset() {
	*x = CreateVCSConnectionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vcs_connection_internal_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateVCSConnectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateVCSConnectionRequest) ProtoMessage() {}

func (x *CreateVCSConnectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vcs_connection_internal_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateVCSConnectionRequest.ProtoReflect.Descriptor instead.
func (*CreateVCSConnectionRequest) Descriptor() ([]byte, []int) {
	return file_vcs_connection_internal_proto_rawDescGZIP(), []int{4}
}

func (x *CreateVCSConnectionRequest) GetVcsConnection() *VCSConnection {
	if x != nil {
		return x.VcsConnection
	}
	return nil
}

type GetVCSConnectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId string             `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	Provider  types.VCSProviders `protobuf:"varint,2,opt,name=provider,proto3,enum=alfred.types.VCSProviders" json:"provider,omitempty"`
}

func (x *GetVCSConnectionRequest) Reset() {
	*x = GetVCSConnectionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vcs_connection_internal_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVCSConnectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVCSConnectionRequest) ProtoMessage() {}

func (x *GetVCSConnectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vcs_connection_internal_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVCSConnectionRequest.ProtoReflect.Descriptor instead.
func (*GetVCSConnectionRequest) Descriptor() ([]byte, []int) {
	return file_vcs_connection_internal_proto_rawDescGZIP(), []int{5}
}

func (x *GetVCSConnectionRequest) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *GetVCSConnectionRequest) GetProvider() types.VCSProviders {
	if x != nil {
		return x.Provider
	}
	return types.VCSProviders_UNKNOWN
}

var File_vcs_connection_internal_proto protoreflect.FileDescriptor

var file_vcs_connection_internal_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x76, 0x63, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x21, 0x61, 0x6c, 0x66, 0x72, 0x65, 0x64, 0x2e, 0x76, 0x63, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x11, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xca, 0x03, 0x0a, 0x0d, 0x56, 0x43, 0x53, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x36, 0x0a, 0x08, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e,
	0x61, 0x6c, 0x66, 0x72, 0x65, 0x64, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x56, 0x43, 0x53,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x72,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x4a, 0x0a, 0x13, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x11, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x45, 0x78, 0x70, 0x69, 0x72, 0x79, 0x12, 0x4c, 0x0a, 0x14,
	0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x65, 0x78,
	0x70, 0x69, 0x72, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x12, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x45, 0x78, 0x70, 0x69, 0x72, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65,
	0x76, 0x6f, 0x6b, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x72, 0x65, 0x76,
	0x6f, 0x6b, 0x65, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x22, 0x83, 0x01, 0x0a, 0x14, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x56, 0x43, 0x53, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x36, 0x0a, 0x08, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x61, 0x6c,
	0x66, 0x72, 0x65, 0x64, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x56, 0x43, 0x53, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x22, 0x66, 0x0a, 0x15, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65,
	0x56, 0x43, 0x53, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x36, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1a, 0x2e, 0x61, 0x6c, 0x66, 0x72, 0x65, 0x64, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x56, 0x43, 0x53, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x52, 0x08, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x15, 0x0a, 0x06, 0x76, 0x63, 0x73, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x63, 0x73, 0x49, 0x64, 0x22, 0x4e,
	0x0a, 0x14, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x56, 0x43, 0x53, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x61, 0x6c, 0x66, 0x72, 0x65,
	0x64, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x56, 0x43, 0x53, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x73, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x22, 0x75,
	0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x43, 0x53, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x57, 0x0a, 0x0e,
	0x76, 0x63, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x61, 0x6c, 0x66, 0x72, 0x65, 0x64, 0x2e, 0x76, 0x63,
	0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x56, 0x43, 0x53, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x76, 0x63, 0x73, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x70, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x56, 0x43, 0x53, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x36, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1a, 0x2e, 0x61, 0x6c, 0x66, 0x72, 0x65, 0x64, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x56, 0x43, 0x53, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x52, 0x08, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x32, 0xe9, 0x03, 0x0a, 0x15, 0x56, 0x43, 0x53, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x12, 0x80, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x56, 0x43, 0x53, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3a, 0x2e, 0x61, 0x6c, 0x66, 0x72, 0x65, 0x64, 0x2e,
	0x76, 0x63, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x43,
	0x53, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x30, 0x2e, 0x61, 0x6c, 0x66, 0x72, 0x65, 0x64, 0x2e, 0x76, 0x63, 0x73, 0x5f,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x56, 0x43, 0x53, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x86, 0x01, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56,
	0x43, 0x53, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3d, 0x2e, 0x61,
	0x6c, 0x66, 0x72, 0x65, 0x64, 0x2e, 0x76, 0x63, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x43, 0x53, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x61, 0x6c,
	0x66, 0x72, 0x65, 0x64, 0x2e, 0x76, 0x63, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e,
	0x56, 0x43, 0x53, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x62, 0x0a,
	0x0e, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x56, 0x43, 0x53, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x38, 0x2e, 0x61, 0x6c, 0x66, 0x72, 0x65, 0x64, 0x2e, 0x76, 0x63, 0x73, 0x5f, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2e, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x56, 0x43, 0x53, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x12, 0x60, 0x0a, 0x0d, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x56, 0x43, 0x53, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x37, 0x2e, 0x61, 0x6c, 0x66, 0x72, 0x65, 0x64, 0x2e, 0x76, 0x63, 0x73, 0x5f,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x56, 0x43, 0x53, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x42, 0x05, 0x5a, 0x03, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_vcs_connection_internal_proto_rawDescOnce sync.Once
	file_vcs_connection_internal_proto_rawDescData = file_vcs_connection_internal_proto_rawDesc
)

func file_vcs_connection_internal_proto_rawDescGZIP() []byte {
	file_vcs_connection_internal_proto_rawDescOnce.Do(func() {
		file_vcs_connection_internal_proto_rawDescData = protoimpl.X.CompressGZIP(file_vcs_connection_internal_proto_rawDescData)
	})
	return file_vcs_connection_internal_proto_rawDescData
}

var file_vcs_connection_internal_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_vcs_connection_internal_proto_goTypes = []interface{}{
	(*VCSConnection)(nil),              // 0: alfred.vcs_connection.v1.internal.VCSConnection
	(*AccountVCSConnection)(nil),       // 1: alfred.vcs_connection.v1.internal.AccountVCSConnection
	(*RevokeVCSTokenRequest)(nil),      // 2: alfred.vcs_connection.v1.internal.RevokeVCSTokenRequest
	(*RenewVCSTokenRequest)(nil),       // 3: alfred.vcs_connection.v1.internal.RenewVCSTokenRequest
	(*CreateVCSConnectionRequest)(nil), // 4: alfred.vcs_connection.v1.internal.CreateVCSConnectionRequest
	(*GetVCSConnectionRequest)(nil),    // 5: alfred.vcs_connection.v1.internal.GetVCSConnectionRequest
	(types.VCSProviders)(0),            // 6: alfred.types.VCSProviders
	(*timestamp.Timestamp)(nil),        // 7: google.protobuf.Timestamp
	(*empty.Empty)(nil),                // 8: google.protobuf.Empty
}
var file_vcs_connection_internal_proto_depIdxs = []int32{
	6,  // 0: alfred.vcs_connection.v1.internal.VCSConnection.provider:type_name -> alfred.types.VCSProviders
	7,  // 1: alfred.vcs_connection.v1.internal.VCSConnection.access_token_expiry:type_name -> google.protobuf.Timestamp
	7,  // 2: alfred.vcs_connection.v1.internal.VCSConnection.refresh_token_expiry:type_name -> google.protobuf.Timestamp
	6,  // 3: alfred.vcs_connection.v1.internal.AccountVCSConnection.provider:type_name -> alfred.types.VCSProviders
	6,  // 4: alfred.vcs_connection.v1.internal.RevokeVCSTokenRequest.provider:type_name -> alfred.types.VCSProviders
	6,  // 5: alfred.vcs_connection.v1.internal.RenewVCSTokenRequest.provider:type_name -> alfred.types.VCSProviders
	0,  // 6: alfred.vcs_connection.v1.internal.CreateVCSConnectionRequest.vcs_connection:type_name -> alfred.vcs_connection.v1.internal.VCSConnection
	6,  // 7: alfred.vcs_connection.v1.internal.GetVCSConnectionRequest.provider:type_name -> alfred.types.VCSProviders
	5,  // 8: alfred.vcs_connection.v1.internal.VCSConnectionInternal.GetVCSConnection:input_type -> alfred.vcs_connection.v1.internal.GetVCSConnectionRequest
	4,  // 9: alfred.vcs_connection.v1.internal.VCSConnectionInternal.CreateVCSConnection:input_type -> alfred.vcs_connection.v1.internal.CreateVCSConnectionRequest
	2,  // 10: alfred.vcs_connection.v1.internal.VCSConnectionInternal.RevokeVCSToken:input_type -> alfred.vcs_connection.v1.internal.RevokeVCSTokenRequest
	3,  // 11: alfred.vcs_connection.v1.internal.VCSConnectionInternal.RenewVCSToken:input_type -> alfred.vcs_connection.v1.internal.RenewVCSTokenRequest
	0,  // 12: alfred.vcs_connection.v1.internal.VCSConnectionInternal.GetVCSConnection:output_type -> alfred.vcs_connection.v1.internal.VCSConnection
	0,  // 13: alfred.vcs_connection.v1.internal.VCSConnectionInternal.CreateVCSConnection:output_type -> alfred.vcs_connection.v1.internal.VCSConnection
	8,  // 14: alfred.vcs_connection.v1.internal.VCSConnectionInternal.RevokeVCSToken:output_type -> google.protobuf.Empty
	8,  // 15: alfred.vcs_connection.v1.internal.VCSConnectionInternal.RenewVCSToken:output_type -> google.protobuf.Empty
	12, // [12:16] is the sub-list for method output_type
	8,  // [8:12] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_vcs_connection_internal_proto_init() }
func file_vcs_connection_internal_proto_init() {
	if File_vcs_connection_internal_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_vcs_connection_internal_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VCSConnection); i {
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
		file_vcs_connection_internal_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountVCSConnection); i {
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
		file_vcs_connection_internal_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RevokeVCSTokenRequest); i {
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
		file_vcs_connection_internal_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RenewVCSTokenRequest); i {
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
		file_vcs_connection_internal_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateVCSConnectionRequest); i {
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
		file_vcs_connection_internal_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVCSConnectionRequest); i {
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
			RawDescriptor: file_vcs_connection_internal_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_vcs_connection_internal_proto_goTypes,
		DependencyIndexes: file_vcs_connection_internal_proto_depIdxs,
		MessageInfos:      file_vcs_connection_internal_proto_msgTypes,
	}.Build()
	File_vcs_connection_internal_proto = out.File
	file_vcs_connection_internal_proto_rawDesc = nil
	file_vcs_connection_internal_proto_goTypes = nil
	file_vcs_connection_internal_proto_depIdxs = nil
}
