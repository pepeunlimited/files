// Code generated by protoc-gen-go. DO NOT EDIT.
// source: spaces.proto

package rpc

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type CreateSpacesParams struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Endpoint             string   `protobuf:"bytes,2,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateSpacesParams) Reset()         { *m = CreateSpacesParams{} }
func (m *CreateSpacesParams) String() string { return proto.CompactTextString(m) }
func (*CreateSpacesParams) ProtoMessage()    {}
func (*CreateSpacesParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6da961f807523b7, []int{0}
}

func (m *CreateSpacesParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateSpacesParams.Unmarshal(m, b)
}
func (m *CreateSpacesParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateSpacesParams.Marshal(b, m, deterministic)
}
func (m *CreateSpacesParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateSpacesParams.Merge(m, src)
}
func (m *CreateSpacesParams) XXX_Size() int {
	return xxx_messageInfo_CreateSpacesParams.Size(m)
}
func (m *CreateSpacesParams) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateSpacesParams.DiscardUnknown(m)
}

var xxx_messageInfo_CreateSpacesParams proto.InternalMessageInfo

func (m *CreateSpacesParams) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateSpacesParams) GetEndpoint() string {
	if m != nil {
		return m.Endpoint
	}
	return ""
}

type CreateSpacesResponse struct {
	Endpoint             string   `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	CdnEndpoint          string   `protobuf:"bytes,2,opt,name=cdn_endpoint,json=cdnEndpoint,proto3" json:"cdn_endpoint,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	SpacesId             int64    `protobuf:"varint,4,opt,name=spaces_id,json=spacesId,proto3" json:"spaces_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateSpacesResponse) Reset()         { *m = CreateSpacesResponse{} }
func (m *CreateSpacesResponse) String() string { return proto.CompactTextString(m) }
func (*CreateSpacesResponse) ProtoMessage()    {}
func (*CreateSpacesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6da961f807523b7, []int{1}
}

func (m *CreateSpacesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateSpacesResponse.Unmarshal(m, b)
}
func (m *CreateSpacesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateSpacesResponse.Marshal(b, m, deterministic)
}
func (m *CreateSpacesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateSpacesResponse.Merge(m, src)
}
func (m *CreateSpacesResponse) XXX_Size() int {
	return xxx_messageInfo_CreateSpacesResponse.Size(m)
}
func (m *CreateSpacesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateSpacesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateSpacesResponse proto.InternalMessageInfo

func (m *CreateSpacesResponse) GetEndpoint() string {
	if m != nil {
		return m.Endpoint
	}
	return ""
}

func (m *CreateSpacesResponse) GetCdnEndpoint() string {
	if m != nil {
		return m.CdnEndpoint
	}
	return ""
}

func (m *CreateSpacesResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateSpacesResponse) GetSpacesId() int64 {
	if m != nil {
		return m.SpacesId
	}
	return 0
}

type GetFileParams struct {
	FileId               *wrappers.Int64Value `protobuf:"bytes,1,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
	Filename             *Filename            `protobuf:"bytes,2,opt,name=filename,proto3" json:"filename,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *GetFileParams) Reset()         { *m = GetFileParams{} }
func (m *GetFileParams) String() string { return proto.CompactTextString(m) }
func (*GetFileParams) ProtoMessage()    {}
func (*GetFileParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6da961f807523b7, []int{2}
}

func (m *GetFileParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFileParams.Unmarshal(m, b)
}
func (m *GetFileParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFileParams.Marshal(b, m, deterministic)
}
func (m *GetFileParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFileParams.Merge(m, src)
}
func (m *GetFileParams) XXX_Size() int {
	return xxx_messageInfo_GetFileParams.Size(m)
}
func (m *GetFileParams) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFileParams.DiscardUnknown(m)
}

var xxx_messageInfo_GetFileParams proto.InternalMessageInfo

func (m *GetFileParams) GetFileId() *wrappers.Int64Value {
	if m != nil {
		return m.FileId
	}
	return nil
}

func (m *GetFileParams) GetFilename() *Filename {
	if m != nil {
		return m.Filename
	}
	return nil
}

type GetFilesParams struct {
	PageSize             int32    `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageToken            int64    `protobuf:"varint,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFilesParams) Reset()         { *m = GetFilesParams{} }
func (m *GetFilesParams) String() string { return proto.CompactTextString(m) }
func (*GetFilesParams) ProtoMessage()    {}
func (*GetFilesParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6da961f807523b7, []int{3}
}

func (m *GetFilesParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFilesParams.Unmarshal(m, b)
}
func (m *GetFilesParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFilesParams.Marshal(b, m, deterministic)
}
func (m *GetFilesParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFilesParams.Merge(m, src)
}
func (m *GetFilesParams) XXX_Size() int {
	return xxx_messageInfo_GetFilesParams.Size(m)
}
func (m *GetFilesParams) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFilesParams.DiscardUnknown(m)
}

var xxx_messageInfo_GetFilesParams proto.InternalMessageInfo

func (m *GetFilesParams) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *GetFilesParams) GetPageToken() int64 {
	if m != nil {
		return m.PageToken
	}
	return 0
}

type GetFilesResponse struct {
	Files                []*File  `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFilesResponse) Reset()         { *m = GetFilesResponse{} }
func (m *GetFilesResponse) String() string { return proto.CompactTextString(m) }
func (*GetFilesResponse) ProtoMessage()    {}
func (*GetFilesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6da961f807523b7, []int{4}
}

func (m *GetFilesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFilesResponse.Unmarshal(m, b)
}
func (m *GetFilesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFilesResponse.Marshal(b, m, deterministic)
}
func (m *GetFilesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFilesResponse.Merge(m, src)
}
func (m *GetFilesResponse) XXX_Size() int {
	return xxx_messageInfo_GetFilesResponse.Size(m)
}
func (m *GetFilesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFilesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetFilesResponse proto.InternalMessageInfo

func (m *GetFilesResponse) GetFiles() []*File {
	if m != nil {
		return m.Files
	}
	return nil
}

type GetSpacesParams struct {
	PageSize             int32    `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageToken            int64    `protobuf:"varint,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetSpacesParams) Reset()         { *m = GetSpacesParams{} }
func (m *GetSpacesParams) String() string { return proto.CompactTextString(m) }
func (*GetSpacesParams) ProtoMessage()    {}
func (*GetSpacesParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6da961f807523b7, []int{5}
}

func (m *GetSpacesParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSpacesParams.Unmarshal(m, b)
}
func (m *GetSpacesParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSpacesParams.Marshal(b, m, deterministic)
}
func (m *GetSpacesParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSpacesParams.Merge(m, src)
}
func (m *GetSpacesParams) XXX_Size() int {
	return xxx_messageInfo_GetSpacesParams.Size(m)
}
func (m *GetSpacesParams) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSpacesParams.DiscardUnknown(m)
}

var xxx_messageInfo_GetSpacesParams proto.InternalMessageInfo

func (m *GetSpacesParams) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *GetSpacesParams) GetPageToken() int64 {
	if m != nil {
		return m.PageToken
	}
	return 0
}

type GetSpacesResponse struct {
	Spaces               []*Spaces `protobuf:"bytes,1,rep,name=spaces,proto3" json:"spaces,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GetSpacesResponse) Reset()         { *m = GetSpacesResponse{} }
func (m *GetSpacesResponse) String() string { return proto.CompactTextString(m) }
func (*GetSpacesResponse) ProtoMessage()    {}
func (*GetSpacesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6da961f807523b7, []int{6}
}

func (m *GetSpacesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSpacesResponse.Unmarshal(m, b)
}
func (m *GetSpacesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSpacesResponse.Marshal(b, m, deterministic)
}
func (m *GetSpacesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSpacesResponse.Merge(m, src)
}
func (m *GetSpacesResponse) XXX_Size() int {
	return xxx_messageInfo_GetSpacesResponse.Size(m)
}
func (m *GetSpacesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSpacesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetSpacesResponse proto.InternalMessageInfo

func (m *GetSpacesResponse) GetSpaces() []*Spaces {
	if m != nil {
		return m.Spaces
	}
	return nil
}

type Spaces struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Spaces) Reset()         { *m = Spaces{} }
func (m *Spaces) String() string { return proto.CompactTextString(m) }
func (*Spaces) ProtoMessage()    {}
func (*Spaces) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6da961f807523b7, []int{7}
}

func (m *Spaces) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Spaces.Unmarshal(m, b)
}
func (m *Spaces) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Spaces.Marshal(b, m, deterministic)
}
func (m *Spaces) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Spaces.Merge(m, src)
}
func (m *Spaces) XXX_Size() int {
	return xxx_messageInfo_Spaces.Size(m)
}
func (m *Spaces) XXX_DiscardUnknown() {
	xxx_messageInfo_Spaces.DiscardUnknown(m)
}

var xxx_messageInfo_Spaces proto.InternalMessageInfo

func (m *Spaces) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Spaces) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CutParams struct {
	ToSpacesName         *wrappers.StringValue `protobuf:"bytes,1,opt,name=to_spaces_name,json=toSpacesName,proto3" json:"to_spaces_name,omitempty"`
	FileId               int64                 `protobuf:"varint,2,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
	Filename             *Filename             `protobuf:"bytes,3,opt,name=filename,proto3" json:"filename,omitempty"`
	ToSpacesId           *wrappers.Int64Value  `protobuf:"bytes,4,opt,name=to_spaces_id,json=toSpacesId,proto3" json:"to_spaces_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CutParams) Reset()         { *m = CutParams{} }
func (m *CutParams) String() string { return proto.CompactTextString(m) }
func (*CutParams) ProtoMessage()    {}
func (*CutParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6da961f807523b7, []int{8}
}

func (m *CutParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CutParams.Unmarshal(m, b)
}
func (m *CutParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CutParams.Marshal(b, m, deterministic)
}
func (m *CutParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CutParams.Merge(m, src)
}
func (m *CutParams) XXX_Size() int {
	return xxx_messageInfo_CutParams.Size(m)
}
func (m *CutParams) XXX_DiscardUnknown() {
	xxx_messageInfo_CutParams.DiscardUnknown(m)
}

var xxx_messageInfo_CutParams proto.InternalMessageInfo

func (m *CutParams) GetToSpacesName() *wrappers.StringValue {
	if m != nil {
		return m.ToSpacesName
	}
	return nil
}

func (m *CutParams) GetFileId() int64 {
	if m != nil {
		return m.FileId
	}
	return 0
}

func (m *CutParams) GetFilename() *Filename {
	if m != nil {
		return m.Filename
	}
	return nil
}

func (m *CutParams) GetToSpacesId() *wrappers.Int64Value {
	if m != nil {
		return m.ToSpacesId
	}
	return nil
}

type CutResponse struct {
	NewBucketName        string   `protobuf:"bytes,1,opt,name=new_bucket_name,json=newBucketName,proto3" json:"new_bucket_name,omitempty"`
	NewBucketId          int64    `protobuf:"varint,2,opt,name=new_bucket_id,json=newBucketId,proto3" json:"new_bucket_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CutResponse) Reset()         { *m = CutResponse{} }
func (m *CutResponse) String() string { return proto.CompactTextString(m) }
func (*CutResponse) ProtoMessage()    {}
func (*CutResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6da961f807523b7, []int{9}
}

func (m *CutResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CutResponse.Unmarshal(m, b)
}
func (m *CutResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CutResponse.Marshal(b, m, deterministic)
}
func (m *CutResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CutResponse.Merge(m, src)
}
func (m *CutResponse) XXX_Size() int {
	return xxx_messageInfo_CutResponse.Size(m)
}
func (m *CutResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CutResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CutResponse proto.InternalMessageInfo

func (m *CutResponse) GetNewBucketName() string {
	if m != nil {
		return m.NewBucketName
	}
	return ""
}

func (m *CutResponse) GetNewBucketId() int64 {
	if m != nil {
		return m.NewBucketId
	}
	return 0
}

type WipeParams struct {
	BucketName           string   `protobuf:"bytes,1,opt,name=bucket_name,json=bucketName,proto3" json:"bucket_name,omitempty"`
	IsDeleteSpaces       bool     `protobuf:"varint,2,opt,name=is_delete_spaces,json=isDeleteSpaces,proto3" json:"is_delete_spaces,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WipeParams) Reset()         { *m = WipeParams{} }
func (m *WipeParams) String() string { return proto.CompactTextString(m) }
func (*WipeParams) ProtoMessage()    {}
func (*WipeParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6da961f807523b7, []int{10}
}

func (m *WipeParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WipeParams.Unmarshal(m, b)
}
func (m *WipeParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WipeParams.Marshal(b, m, deterministic)
}
func (m *WipeParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WipeParams.Merge(m, src)
}
func (m *WipeParams) XXX_Size() int {
	return xxx_messageInfo_WipeParams.Size(m)
}
func (m *WipeParams) XXX_DiscardUnknown() {
	xxx_messageInfo_WipeParams.DiscardUnknown(m)
}

var xxx_messageInfo_WipeParams proto.InternalMessageInfo

func (m *WipeParams) GetBucketName() string {
	if m != nil {
		return m.BucketName
	}
	return ""
}

func (m *WipeParams) GetIsDeleteSpaces() bool {
	if m != nil {
		return m.IsDeleteSpaces
	}
	return false
}

type WipeParamsResponse struct {
	DeletedCount         int64    `protobuf:"varint,1,opt,name=deleted_count,json=deletedCount,proto3" json:"deleted_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WipeParamsResponse) Reset()         { *m = WipeParamsResponse{} }
func (m *WipeParamsResponse) String() string { return proto.CompactTextString(m) }
func (*WipeParamsResponse) ProtoMessage()    {}
func (*WipeParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6da961f807523b7, []int{11}
}

func (m *WipeParamsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WipeParamsResponse.Unmarshal(m, b)
}
func (m *WipeParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WipeParamsResponse.Marshal(b, m, deterministic)
}
func (m *WipeParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WipeParamsResponse.Merge(m, src)
}
func (m *WipeParamsResponse) XXX_Size() int {
	return xxx_messageInfo_WipeParamsResponse.Size(m)
}
func (m *WipeParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WipeParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WipeParamsResponse proto.InternalMessageInfo

func (m *WipeParamsResponse) GetDeletedCount() int64 {
	if m != nil {
		return m.DeletedCount
	}
	return 0
}

type DeleteParams struct {
	FileId               *wrappers.Int64Value `protobuf:"bytes,1,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
	Filename             *Filename            `protobuf:"bytes,2,opt,name=filename,proto3" json:"filename,omitempty"`
	IsPermanent          bool                 `protobuf:"varint,3,opt,name=is_permanent,json=isPermanent,proto3" json:"is_permanent,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *DeleteParams) Reset()         { *m = DeleteParams{} }
func (m *DeleteParams) String() string { return proto.CompactTextString(m) }
func (*DeleteParams) ProtoMessage()    {}
func (*DeleteParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6da961f807523b7, []int{12}
}

func (m *DeleteParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteParams.Unmarshal(m, b)
}
func (m *DeleteParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteParams.Marshal(b, m, deterministic)
}
func (m *DeleteParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteParams.Merge(m, src)
}
func (m *DeleteParams) XXX_Size() int {
	return xxx_messageInfo_DeleteParams.Size(m)
}
func (m *DeleteParams) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteParams.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteParams proto.InternalMessageInfo

func (m *DeleteParams) GetFileId() *wrappers.Int64Value {
	if m != nil {
		return m.FileId
	}
	return nil
}

func (m *DeleteParams) GetFilename() *Filename {
	if m != nil {
		return m.Filename
	}
	return nil
}

func (m *DeleteParams) GetIsPermanent() bool {
	if m != nil {
		return m.IsPermanent
	}
	return false
}

type DeleteResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6da961f807523b7, []int{13}
}

func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (m *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(m, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

type Filename struct {
	Name                 string                `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	BucketId             *wrappers.Int64Value  `protobuf:"bytes,2,opt,name=bucket_id,json=bucketId,proto3" json:"bucket_id,omitempty"`
	BucketName           *wrappers.StringValue `protobuf:"bytes,3,opt,name=bucket_name,json=bucketName,proto3" json:"bucket_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Filename) Reset()         { *m = Filename{} }
func (m *Filename) String() string { return proto.CompactTextString(m) }
func (*Filename) ProtoMessage()    {}
func (*Filename) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6da961f807523b7, []int{14}
}

func (m *Filename) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Filename.Unmarshal(m, b)
}
func (m *Filename) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Filename.Marshal(b, m, deterministic)
}
func (m *Filename) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Filename.Merge(m, src)
}
func (m *Filename) XXX_Size() int {
	return xxx_messageInfo_Filename.Size(m)
}
func (m *Filename) XXX_DiscardUnknown() {
	xxx_messageInfo_Filename.DiscardUnknown(m)
}

var xxx_messageInfo_Filename proto.InternalMessageInfo

func (m *Filename) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Filename) GetBucketId() *wrappers.Int64Value {
	if m != nil {
		return m.BucketId
	}
	return nil
}

func (m *Filename) GetBucketName() *wrappers.StringValue {
	if m != nil {
		return m.BucketName
	}
	return nil
}

type File struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Filename             string   `protobuf:"bytes,2,opt,name=filename,proto3" json:"filename,omitempty"`
	CreatedAt            string   `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	MimeType             string   `protobuf:"bytes,5,opt,name=mime_type,json=mimeType,proto3" json:"mime_type,omitempty"`
	FileSize             int64    `protobuf:"varint,6,opt,name=file_size,json=fileSize,proto3" json:"file_size,omitempty"`
	UserId               int64    `protobuf:"varint,7,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	IsDraft              bool     `protobuf:"varint,8,opt,name=is_draft,json=isDraft,proto3" json:"is_draft,omitempty"`
	SpacesId             int64    `protobuf:"varint,9,opt,name=spaces_id,json=spacesId,proto3" json:"spaces_id,omitempty"`
	FileUrl              string   `protobuf:"bytes,10,opt,name=file_url,json=fileUrl,proto3" json:"file_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *File) Reset()         { *m = File{} }
func (m *File) String() string { return proto.CompactTextString(m) }
func (*File) ProtoMessage()    {}
func (*File) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6da961f807523b7, []int{15}
}

func (m *File) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_File.Unmarshal(m, b)
}
func (m *File) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_File.Marshal(b, m, deterministic)
}
func (m *File) XXX_Merge(src proto.Message) {
	xxx_messageInfo_File.Merge(m, src)
}
func (m *File) XXX_Size() int {
	return xxx_messageInfo_File.Size(m)
}
func (m *File) XXX_DiscardUnknown() {
	xxx_messageInfo_File.DiscardUnknown(m)
}

var xxx_messageInfo_File proto.InternalMessageInfo

func (m *File) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *File) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

func (m *File) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *File) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *File) GetMimeType() string {
	if m != nil {
		return m.MimeType
	}
	return ""
}

func (m *File) GetFileSize() int64 {
	if m != nil {
		return m.FileSize
	}
	return 0
}

func (m *File) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *File) GetIsDraft() bool {
	if m != nil {
		return m.IsDraft
	}
	return false
}

func (m *File) GetSpacesId() int64 {
	if m != nil {
		return m.SpacesId
	}
	return 0
}

func (m *File) GetFileUrl() string {
	if m != nil {
		return m.FileUrl
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateSpacesParams)(nil), "pepeunlimited.files.CreateSpacesParams")
	proto.RegisterType((*CreateSpacesResponse)(nil), "pepeunlimited.files.CreateSpacesResponse")
	proto.RegisterType((*GetFileParams)(nil), "pepeunlimited.files.GetFileParams")
	proto.RegisterType((*GetFilesParams)(nil), "pepeunlimited.files.GetFilesParams")
	proto.RegisterType((*GetFilesResponse)(nil), "pepeunlimited.files.GetFilesResponse")
	proto.RegisterType((*GetSpacesParams)(nil), "pepeunlimited.files.GetSpacesParams")
	proto.RegisterType((*GetSpacesResponse)(nil), "pepeunlimited.files.GetSpacesResponse")
	proto.RegisterType((*Spaces)(nil), "pepeunlimited.files.Spaces")
	proto.RegisterType((*CutParams)(nil), "pepeunlimited.files.CutParams")
	proto.RegisterType((*CutResponse)(nil), "pepeunlimited.files.CutResponse")
	proto.RegisterType((*WipeParams)(nil), "pepeunlimited.files.WipeParams")
	proto.RegisterType((*WipeParamsResponse)(nil), "pepeunlimited.files.WipeParamsResponse")
	proto.RegisterType((*DeleteParams)(nil), "pepeunlimited.files.DeleteParams")
	proto.RegisterType((*DeleteResponse)(nil), "pepeunlimited.files.DeleteResponse")
	proto.RegisterType((*Filename)(nil), "pepeunlimited.files.Filename")
	proto.RegisterType((*File)(nil), "pepeunlimited.files.File")
}

func init() { proto.RegisterFile("spaces.proto", fileDescriptor_d6da961f807523b7) }

var fileDescriptor_d6da961f807523b7 = []byte{
	// 893 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x56, 0x4d, 0x6f, 0xdc, 0x44,
	0x18, 0x96, 0xd7, 0x9b, 0x5d, 0xfb, 0xdd, 0x8f, 0x86, 0x01, 0xa9, 0x5b, 0x87, 0xb4, 0xe9, 0x04,
	0x4a, 0x90, 0xd0, 0x46, 0x4a, 0x2b, 0x44, 0x0f, 0x3d, 0x34, 0x1b, 0x9a, 0xae, 0x04, 0x28, 0x72,
	0x4a, 0xab, 0x72, 0xb1, 0xbc, 0xeb, 0x37, 0xd1, 0xa8, 0x5e, 0xdb, 0xb2, 0xc7, 0x44, 0xed, 0x89,
	0x13, 0x3f, 0x01, 0xee, 0xfc, 0x2f, 0x2e, 0xfc, 0x12, 0x34, 0x1f, 0x9e, 0xb5, 0x13, 0x27, 0x01,
	0x71, 0xe0, 0xb6, 0xf3, 0xce, 0xf3, 0x3e, 0xf3, 0x7e, 0x3e, 0x6b, 0x18, 0x16, 0x59, 0xb8, 0xc4,
	0x62, 0x9a, 0xe5, 0x29, 0x4f, 0xc9, 0xc7, 0x19, 0x66, 0x58, 0x26, 0x31, 0x5b, 0x31, 0x8e, 0xd1,
	0xf4, 0x8c, 0xc5, 0x58, 0x78, 0xf7, 0xcf, 0xd3, 0xf4, 0x3c, 0xc6, 0x7d, 0x09, 0x59, 0x94, 0x67,
	0xfb, 0x17, 0x79, 0x98, 0x65, 0x98, 0x6b, 0x27, 0x7a, 0x04, 0x64, 0x96, 0x63, 0xc8, 0xf1, 0x54,
	0x52, 0x9d, 0x84, 0x79, 0xb8, 0x2a, 0x08, 0x81, 0x6e, 0x12, 0xae, 0x70, 0x62, 0xed, 0x58, 0x7b,
	0xae, 0x2f, 0x7f, 0x13, 0x0f, 0x1c, 0x4c, 0xa2, 0x2c, 0x65, 0x09, 0x9f, 0x74, 0xa4, 0xdd, 0x9c,
	0xe9, 0xaf, 0x16, 0x7c, 0x52, 0xa7, 0xf1, 0xb1, 0xc8, 0xd2, 0xa4, 0x68, 0x3a, 0x59, 0x4d, 0x27,
	0xf2, 0x10, 0x86, 0xcb, 0x28, 0x09, 0x2e, 0x91, 0x0e, 0x96, 0x51, 0xf2, 0x6d, 0x05, 0xa9, 0xe2,
	0xb0, 0x6b, 0x71, 0x6c, 0x81, 0xab, 0xd2, 0x0e, 0x58, 0x34, 0xe9, 0xee, 0x58, 0x7b, 0xb6, 0xef,
	0x28, 0xc3, 0x3c, 0xa2, 0xbf, 0x58, 0x30, 0x3a, 0x46, 0xfe, 0x82, 0xc5, 0xa8, 0x53, 0x79, 0x02,
	0x7d, 0x51, 0x09, 0x01, 0x16, 0x01, 0x0c, 0x0e, 0xb6, 0xa6, 0xaa, 0x24, 0xd3, 0xaa, 0x24, 0xd3,
	0x79, 0xc2, 0xbf, 0x7e, 0xf2, 0x3a, 0x8c, 0x4b, 0xf4, 0x7b, 0x02, 0x3b, 0x8f, 0xc8, 0x53, 0x70,
	0xc4, 0x2f, 0xf9, 0x78, 0x47, 0xba, 0x6d, 0x4f, 0x5b, 0xca, 0x3b, 0x7d, 0xa1, 0x41, 0xbe, 0x81,
	0xd3, 0xef, 0x60, 0xac, 0x23, 0xa8, 0xaa, 0xb9, 0x05, 0x6e, 0x16, 0x9e, 0x63, 0x50, 0xb0, 0x0f,
	0xaa, 0xa4, 0x1b, 0xbe, 0x23, 0x0c, 0xa7, 0xec, 0x03, 0x92, 0x6d, 0x00, 0x79, 0xc9, 0xd3, 0x77,
	0x98, 0xc8, 0xb7, 0x6c, 0x5f, 0xc2, 0x5f, 0x09, 0x03, 0x9d, 0xc1, 0x66, 0xc5, 0x66, 0x8a, 0xba,
	0x0f, 0x1b, 0xf2, 0xf5, 0x89, 0xb5, 0x63, 0xef, 0x0d, 0x0e, 0xee, 0x5d, 0x1b, 0x99, 0xaf, 0x70,
	0xf4, 0x7b, 0xb8, 0x73, 0x8c, 0xbc, 0xd1, 0xe1, 0xff, 0x12, 0xd3, 0x4b, 0xf8, 0xc8, 0xd0, 0x99,
	0xa0, 0x1e, 0x43, 0x4f, 0x75, 0x41, 0x47, 0xb5, 0xd5, 0x1a, 0x95, 0x76, 0xd2, 0x50, 0xfa, 0x15,
	0xf4, 0x94, 0x85, 0x8c, 0xa1, 0xa3, 0x3b, 0x64, 0xfb, 0x1d, 0x16, 0x99, 0xce, 0x77, 0xd6, 0x9d,
	0xa7, 0x7f, 0x59, 0xe0, 0xce, 0x4a, 0xae, 0x33, 0x38, 0x84, 0x31, 0x4f, 0x03, 0x3d, 0x0a, 0x66,
	0x5a, 0x07, 0x07, 0x9f, 0x5e, 0xe9, 0xef, 0x29, 0xcf, 0x59, 0x72, 0xae, 0x1a, 0x3c, 0xe4, 0xa9,
	0x7a, 0xf1, 0x07, 0x31, 0x4b, 0x77, 0xd7, 0xc3, 0xa1, 0xb2, 0x6c, 0xeb, 0xbf, 0xfd, 0xaf, 0xfa,
	0x4f, 0x9e, 0xc1, 0x70, 0x1d, 0x97, 0x1e, 0xd1, 0x5b, 0xa6, 0x0e, 0xaa, 0xa0, 0xe6, 0x11, 0x7d,
	0x0b, 0x83, 0x59, 0xc9, 0x4d, 0x59, 0x1f, 0xc1, 0x9d, 0x04, 0x2f, 0x82, 0x45, 0xb9, 0x7c, 0x87,
	0x3c, 0xa8, 0x2d, 0xe5, 0x28, 0xc1, 0x8b, 0x43, 0x69, 0x95, 0x99, 0x50, 0x18, 0xd5, 0x70, 0x26,
	0x9f, 0x81, 0x41, 0xcd, 0x23, 0xfa, 0x06, 0xe0, 0x0d, 0xcb, 0xaa, 0xc5, 0x78, 0x00, 0x83, 0xab,
	0xac, 0xb0, 0x58, 0x53, 0xee, 0xc1, 0x26, 0x2b, 0x82, 0x08, 0x63, 0xe4, 0xa8, 0xf3, 0x91, 0xac,
	0x8e, 0x3f, 0x66, 0xc5, 0x91, 0x34, 0xab, 0xa8, 0xe9, 0x53, 0x20, 0x6b, 0x62, 0x13, 0xfa, 0x2e,
	0x8c, 0x94, 0x73, 0x14, 0x2c, 0xd3, 0x52, 0x0b, 0x80, 0xed, 0x0f, 0xb5, 0x71, 0x26, 0x6c, 0xf4,
	0x0f, 0x0b, 0x86, 0x8a, 0xeb, 0x7f, 0xda, 0x57, 0x21, 0x43, 0xac, 0x08, 0x32, 0xcc, 0x57, 0x61,
	0x82, 0x09, 0x97, 0xed, 0x76, 0xfc, 0x01, 0x2b, 0x4e, 0x2a, 0x13, 0xdd, 0x84, 0xb1, 0x8a, 0xb1,
	0xca, 0x8d, 0xfe, 0x6e, 0x81, 0x53, 0x71, 0xb5, 0xaa, 0xe5, 0x37, 0xe0, 0x36, 0x7b, 0x71, 0x4b,
	0x22, 0xce, 0x42, 0x77, 0x89, 0x3c, 0x6b, 0xf6, 0xc5, 0xfe, 0x07, 0x43, 0x5d, 0xeb, 0x1a, 0xfd,
	0xad, 0x03, 0x5d, 0x11, 0xd9, 0x95, 0x8d, 0xf2, 0x2e, 0x95, 0xc8, 0xad, 0xd5, 0x60, 0x1b, 0x60,
	0x29, 0xe5, 0x3b, 0x0a, 0x42, 0xae, 0xd5, 0xd6, 0xd5, 0x96, 0xe7, 0x5c, 0x5c, 0x97, 0x59, 0x54,
	0x5d, 0x77, 0xd5, 0xb5, 0xb6, 0x3c, 0xe7, 0x42, 0x4b, 0x56, 0x6c, 0x85, 0x01, 0x7f, 0x9f, 0xe1,
	0x64, 0x43, 0x51, 0x0b, 0xc3, 0xab, 0xf7, 0x99, 0x94, 0x6b, 0xd9, 0x4f, 0x29, 0x34, 0x3d, 0x25,
	0xd7, 0xc2, 0x20, 0x85, 0xe6, 0x2e, 0xf4, 0xcb, 0x02, 0x73, 0x51, 0xa3, 0xbe, 0xda, 0x3f, 0x71,
	0x9c, 0x47, 0xe4, 0x1e, 0x38, 0x62, 0xf6, 0xf2, 0xf0, 0x8c, 0x4f, 0x1c, 0xd9, 0x90, 0x3e, 0x2b,
	0x8e, 0xc4, 0xb1, 0xa9, 0xff, 0x6e, 0x53, 0xff, 0x85, 0x9f, 0x7c, 0xad, 0xcc, 0xe3, 0x09, 0xc8,
	0x48, 0xe4, 0x34, 0xfd, 0x98, 0xc7, 0x07, 0x7f, 0x76, 0x61, 0xa4, 0xe6, 0xf5, 0x14, 0xf3, 0x9f,
	0xd9, 0x12, 0xc9, 0x4b, 0xe8, 0x6b, 0x6d, 0x25, 0xb4, 0x75, 0x5a, 0x1a, 0xff, 0x24, 0xde, 0xf5,
	0x3a, 0x4b, 0x5e, 0x83, 0x53, 0xa9, 0x34, 0xd9, 0xbd, 0x89, 0x4a, 0xcb, 0xaf, 0xf7, 0xf9, 0x8d,
	0x20, 0xb3, 0x42, 0x6f, 0xc1, 0x35, 0x4a, 0x4b, 0x3e, 0xbb, 0xce, 0xa7, 0x2e, 0xec, 0xde, 0xa3,
	0x9b, 0x51, 0x86, 0xfa, 0x18, 0xec, 0x59, 0xc9, 0xc9, 0xfd, 0x56, 0xb8, 0x51, 0x59, 0x6f, 0xe7,
	0xba, 0x7b, 0x43, 0x74, 0x02, 0x3d, 0xb5, 0x1c, 0xe4, 0x61, 0x2b, 0xb6, 0xbe, 0xdd, 0xde, 0xee,
	0x0d, 0x90, 0x1a, 0x63, 0x57, 0xc8, 0x09, 0x79, 0xd0, 0x0a, 0x5e, 0x2b, 0x8d, 0xf7, 0xc5, 0x2d,
	0x00, 0xc3, 0xb8, 0x80, 0x61, 0xfd, 0xf3, 0x84, 0xb4, 0x3b, 0x5e, 0xfd, 0x10, 0xf2, 0xbe, 0xbc,
	0x15, 0x58, 0xbd, 0x71, 0xb8, 0xf1, 0x93, 0x9d, 0x67, 0xcb, 0x45, 0x4f, 0x6e, 0xe8, 0xe3, 0xbf,
	0x03, 0x00, 0x00, 0xff, 0xff, 0x3d, 0xd9, 0xba, 0x85, 0x9c, 0x09, 0x00, 0x00,
}
