// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Snapshot.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type SnapshotDescription_Type int32

const (
	SnapshotDescription_DISABLED  SnapshotDescription_Type = 0
	SnapshotDescription_FLUSH     SnapshotDescription_Type = 1
	SnapshotDescription_SKIPFLUSH SnapshotDescription_Type = 2
)

var SnapshotDescription_Type_name = map[int32]string{
	0: "DISABLED",
	1: "FLUSH",
	2: "SKIPFLUSH",
}

var SnapshotDescription_Type_value = map[string]int32{
	"DISABLED":  0,
	"FLUSH":     1,
	"SKIPFLUSH": 2,
}

func (x SnapshotDescription_Type) Enum() *SnapshotDescription_Type {
	p := new(SnapshotDescription_Type)
	*p = x
	return p
}

func (x SnapshotDescription_Type) String() string {
	return proto.EnumName(SnapshotDescription_Type_name, int32(x))
}

func (x *SnapshotDescription_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(SnapshotDescription_Type_value, data, "SnapshotDescription_Type")
	if err != nil {
		return err
	}
	*x = SnapshotDescription_Type(value)
	return nil
}

func (SnapshotDescription_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_37871977bc5f5631, []int{0, 0}
}

type SnapshotFileInfo_Type int32

const (
	SnapshotFileInfo_HFILE SnapshotFileInfo_Type = 1
	SnapshotFileInfo_WAL   SnapshotFileInfo_Type = 2
)

var SnapshotFileInfo_Type_name = map[int32]string{
	1: "HFILE",
	2: "WAL",
}

var SnapshotFileInfo_Type_value = map[string]int32{
	"HFILE": 1,
	"WAL":   2,
}

func (x SnapshotFileInfo_Type) Enum() *SnapshotFileInfo_Type {
	p := new(SnapshotFileInfo_Type)
	*p = x
	return p
}

func (x SnapshotFileInfo_Type) String() string {
	return proto.EnumName(SnapshotFileInfo_Type_name, int32(x))
}

func (x *SnapshotFileInfo_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(SnapshotFileInfo_Type_value, data, "SnapshotFileInfo_Type")
	if err != nil {
		return err
	}
	*x = SnapshotFileInfo_Type(value)
	return nil
}

func (SnapshotFileInfo_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_37871977bc5f5631, []int{1, 0}
}

//*
// Description of the snapshot to take
type SnapshotDescription struct {
	Name                 *string                   `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Table                *string                   `protobuf:"bytes,2,opt,name=table" json:"table,omitempty"`
	CreationTime         *int64                    `protobuf:"varint,3,opt,name=creation_time,json=creationTime,def=0" json:"creation_time,omitempty"`
	Type                 *SnapshotDescription_Type `protobuf:"varint,4,opt,name=type,enum=hbase.pb.SnapshotDescription_Type,def=1" json:"type,omitempty"`
	Version              *int32                    `protobuf:"varint,5,opt,name=version" json:"version,omitempty"`
	Owner                *string                   `protobuf:"bytes,6,opt,name=owner" json:"owner,omitempty"`
	UsersAndPermissions  *UsersAndPermissions      `protobuf:"bytes,7,opt,name=users_and_permissions,json=usersAndPermissions" json:"users_and_permissions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *SnapshotDescription) Reset()         { *m = SnapshotDescription{} }
func (m *SnapshotDescription) String() string { return proto.CompactTextString(m) }
func (*SnapshotDescription) ProtoMessage()    {}
func (*SnapshotDescription) Descriptor() ([]byte, []int) {
	return fileDescriptor_37871977bc5f5631, []int{0}
}

func (m *SnapshotDescription) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SnapshotDescription.Unmarshal(m, b)
}
func (m *SnapshotDescription) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SnapshotDescription.Marshal(b, m, deterministic)
}
func (m *SnapshotDescription) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SnapshotDescription.Merge(m, src)
}
func (m *SnapshotDescription) XXX_Size() int {
	return xxx_messageInfo_SnapshotDescription.Size(m)
}
func (m *SnapshotDescription) XXX_DiscardUnknown() {
	xxx_messageInfo_SnapshotDescription.DiscardUnknown(m)
}

var xxx_messageInfo_SnapshotDescription proto.InternalMessageInfo

const Default_SnapshotDescription_CreationTime int64 = 0
const Default_SnapshotDescription_Type SnapshotDescription_Type = SnapshotDescription_FLUSH

func (m *SnapshotDescription) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *SnapshotDescription) GetTable() string {
	if m != nil && m.Table != nil {
		return *m.Table
	}
	return ""
}

func (m *SnapshotDescription) GetCreationTime() int64 {
	if m != nil && m.CreationTime != nil {
		return *m.CreationTime
	}
	return Default_SnapshotDescription_CreationTime
}

func (m *SnapshotDescription) GetType() SnapshotDescription_Type {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Default_SnapshotDescription_Type
}

func (m *SnapshotDescription) GetVersion() int32 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return 0
}

func (m *SnapshotDescription) GetOwner() string {
	if m != nil && m.Owner != nil {
		return *m.Owner
	}
	return ""
}

func (m *SnapshotDescription) GetUsersAndPermissions() *UsersAndPermissions {
	if m != nil {
		return m.UsersAndPermissions
	}
	return nil
}

type SnapshotFileInfo struct {
	Type                 *SnapshotFileInfo_Type `protobuf:"varint,1,req,name=type,enum=hbase.pb.SnapshotFileInfo_Type" json:"type,omitempty"`
	Hfile                *string                `protobuf:"bytes,3,opt,name=hfile" json:"hfile,omitempty"`
	WalServer            *string                `protobuf:"bytes,4,opt,name=wal_server,json=walServer" json:"wal_server,omitempty"`
	WalName              *string                `protobuf:"bytes,5,opt,name=wal_name,json=walName" json:"wal_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *SnapshotFileInfo) Reset()         { *m = SnapshotFileInfo{} }
func (m *SnapshotFileInfo) String() string { return proto.CompactTextString(m) }
func (*SnapshotFileInfo) ProtoMessage()    {}
func (*SnapshotFileInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_37871977bc5f5631, []int{1}
}

func (m *SnapshotFileInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SnapshotFileInfo.Unmarshal(m, b)
}
func (m *SnapshotFileInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SnapshotFileInfo.Marshal(b, m, deterministic)
}
func (m *SnapshotFileInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SnapshotFileInfo.Merge(m, src)
}
func (m *SnapshotFileInfo) XXX_Size() int {
	return xxx_messageInfo_SnapshotFileInfo.Size(m)
}
func (m *SnapshotFileInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_SnapshotFileInfo.DiscardUnknown(m)
}

var xxx_messageInfo_SnapshotFileInfo proto.InternalMessageInfo

func (m *SnapshotFileInfo) GetType() SnapshotFileInfo_Type {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return SnapshotFileInfo_HFILE
}

func (m *SnapshotFileInfo) GetHfile() string {
	if m != nil && m.Hfile != nil {
		return *m.Hfile
	}
	return ""
}

func (m *SnapshotFileInfo) GetWalServer() string {
	if m != nil && m.WalServer != nil {
		return *m.WalServer
	}
	return ""
}

func (m *SnapshotFileInfo) GetWalName() string {
	if m != nil && m.WalName != nil {
		return *m.WalName
	}
	return ""
}

type SnapshotRegionManifest struct {
	Version              *int32                                `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	RegionInfo           *RegionInfo                           `protobuf:"bytes,2,req,name=region_info,json=regionInfo" json:"region_info,omitempty"`
	FamilyFiles          []*SnapshotRegionManifest_FamilyFiles `protobuf:"bytes,3,rep,name=family_files,json=familyFiles" json:"family_files,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *SnapshotRegionManifest) Reset()         { *m = SnapshotRegionManifest{} }
func (m *SnapshotRegionManifest) String() string { return proto.CompactTextString(m) }
func (*SnapshotRegionManifest) ProtoMessage()    {}
func (*SnapshotRegionManifest) Descriptor() ([]byte, []int) {
	return fileDescriptor_37871977bc5f5631, []int{2}
}

func (m *SnapshotRegionManifest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SnapshotRegionManifest.Unmarshal(m, b)
}
func (m *SnapshotRegionManifest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SnapshotRegionManifest.Marshal(b, m, deterministic)
}
func (m *SnapshotRegionManifest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SnapshotRegionManifest.Merge(m, src)
}
func (m *SnapshotRegionManifest) XXX_Size() int {
	return xxx_messageInfo_SnapshotRegionManifest.Size(m)
}
func (m *SnapshotRegionManifest) XXX_DiscardUnknown() {
	xxx_messageInfo_SnapshotRegionManifest.DiscardUnknown(m)
}

var xxx_messageInfo_SnapshotRegionManifest proto.InternalMessageInfo

func (m *SnapshotRegionManifest) GetVersion() int32 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return 0
}

func (m *SnapshotRegionManifest) GetRegionInfo() *RegionInfo {
	if m != nil {
		return m.RegionInfo
	}
	return nil
}

func (m *SnapshotRegionManifest) GetFamilyFiles() []*SnapshotRegionManifest_FamilyFiles {
	if m != nil {
		return m.FamilyFiles
	}
	return nil
}

type SnapshotRegionManifest_StoreFile struct {
	Name      *string    `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Reference *Reference `protobuf:"bytes,2,opt,name=reference" json:"reference,omitempty"`
	// TODO: Add checksums or other fields to verify the file
	FileSize             *uint64  `protobuf:"varint,3,opt,name=file_size,json=fileSize" json:"file_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SnapshotRegionManifest_StoreFile) Reset()         { *m = SnapshotRegionManifest_StoreFile{} }
func (m *SnapshotRegionManifest_StoreFile) String() string { return proto.CompactTextString(m) }
func (*SnapshotRegionManifest_StoreFile) ProtoMessage()    {}
func (*SnapshotRegionManifest_StoreFile) Descriptor() ([]byte, []int) {
	return fileDescriptor_37871977bc5f5631, []int{2, 0}
}

func (m *SnapshotRegionManifest_StoreFile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SnapshotRegionManifest_StoreFile.Unmarshal(m, b)
}
func (m *SnapshotRegionManifest_StoreFile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SnapshotRegionManifest_StoreFile.Marshal(b, m, deterministic)
}
func (m *SnapshotRegionManifest_StoreFile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SnapshotRegionManifest_StoreFile.Merge(m, src)
}
func (m *SnapshotRegionManifest_StoreFile) XXX_Size() int {
	return xxx_messageInfo_SnapshotRegionManifest_StoreFile.Size(m)
}
func (m *SnapshotRegionManifest_StoreFile) XXX_DiscardUnknown() {
	xxx_messageInfo_SnapshotRegionManifest_StoreFile.DiscardUnknown(m)
}

var xxx_messageInfo_SnapshotRegionManifest_StoreFile proto.InternalMessageInfo

func (m *SnapshotRegionManifest_StoreFile) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *SnapshotRegionManifest_StoreFile) GetReference() *Reference {
	if m != nil {
		return m.Reference
	}
	return nil
}

func (m *SnapshotRegionManifest_StoreFile) GetFileSize() uint64 {
	if m != nil && m.FileSize != nil {
		return *m.FileSize
	}
	return 0
}

type SnapshotRegionManifest_FamilyFiles struct {
	FamilyName           []byte                              `protobuf:"bytes,1,req,name=family_name,json=familyName" json:"family_name,omitempty"`
	StoreFiles           []*SnapshotRegionManifest_StoreFile `protobuf:"bytes,2,rep,name=store_files,json=storeFiles" json:"store_files,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *SnapshotRegionManifest_FamilyFiles) Reset()         { *m = SnapshotRegionManifest_FamilyFiles{} }
func (m *SnapshotRegionManifest_FamilyFiles) String() string { return proto.CompactTextString(m) }
func (*SnapshotRegionManifest_FamilyFiles) ProtoMessage()    {}
func (*SnapshotRegionManifest_FamilyFiles) Descriptor() ([]byte, []int) {
	return fileDescriptor_37871977bc5f5631, []int{2, 1}
}

func (m *SnapshotRegionManifest_FamilyFiles) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SnapshotRegionManifest_FamilyFiles.Unmarshal(m, b)
}
func (m *SnapshotRegionManifest_FamilyFiles) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SnapshotRegionManifest_FamilyFiles.Marshal(b, m, deterministic)
}
func (m *SnapshotRegionManifest_FamilyFiles) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SnapshotRegionManifest_FamilyFiles.Merge(m, src)
}
func (m *SnapshotRegionManifest_FamilyFiles) XXX_Size() int {
	return xxx_messageInfo_SnapshotRegionManifest_FamilyFiles.Size(m)
}
func (m *SnapshotRegionManifest_FamilyFiles) XXX_DiscardUnknown() {
	xxx_messageInfo_SnapshotRegionManifest_FamilyFiles.DiscardUnknown(m)
}

var xxx_messageInfo_SnapshotRegionManifest_FamilyFiles proto.InternalMessageInfo

func (m *SnapshotRegionManifest_FamilyFiles) GetFamilyName() []byte {
	if m != nil {
		return m.FamilyName
	}
	return nil
}

func (m *SnapshotRegionManifest_FamilyFiles) GetStoreFiles() []*SnapshotRegionManifest_StoreFile {
	if m != nil {
		return m.StoreFiles
	}
	return nil
}

type SnapshotDataManifest struct {
	TableSchema          *TableSchema              `protobuf:"bytes,1,req,name=table_schema,json=tableSchema" json:"table_schema,omitempty"`
	RegionManifests      []*SnapshotRegionManifest `protobuf:"bytes,2,rep,name=region_manifests,json=regionManifests" json:"region_manifests,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *SnapshotDataManifest) Reset()         { *m = SnapshotDataManifest{} }
func (m *SnapshotDataManifest) String() string { return proto.CompactTextString(m) }
func (*SnapshotDataManifest) ProtoMessage()    {}
func (*SnapshotDataManifest) Descriptor() ([]byte, []int) {
	return fileDescriptor_37871977bc5f5631, []int{3}
}

func (m *SnapshotDataManifest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SnapshotDataManifest.Unmarshal(m, b)
}
func (m *SnapshotDataManifest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SnapshotDataManifest.Marshal(b, m, deterministic)
}
func (m *SnapshotDataManifest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SnapshotDataManifest.Merge(m, src)
}
func (m *SnapshotDataManifest) XXX_Size() int {
	return xxx_messageInfo_SnapshotDataManifest.Size(m)
}
func (m *SnapshotDataManifest) XXX_DiscardUnknown() {
	xxx_messageInfo_SnapshotDataManifest.DiscardUnknown(m)
}

var xxx_messageInfo_SnapshotDataManifest proto.InternalMessageInfo

func (m *SnapshotDataManifest) GetTableSchema() *TableSchema {
	if m != nil {
		return m.TableSchema
	}
	return nil
}

func (m *SnapshotDataManifest) GetRegionManifests() []*SnapshotRegionManifest {
	if m != nil {
		return m.RegionManifests
	}
	return nil
}

func init() {
	proto.RegisterEnum("hbase.pb.SnapshotDescription_Type", SnapshotDescription_Type_name, SnapshotDescription_Type_value)
	proto.RegisterEnum("hbase.pb.SnapshotFileInfo_Type", SnapshotFileInfo_Type_name, SnapshotFileInfo_Type_value)
	proto.RegisterType((*SnapshotDescription)(nil), "hbase.pb.SnapshotDescription")
	proto.RegisterType((*SnapshotFileInfo)(nil), "hbase.pb.SnapshotFileInfo")
	proto.RegisterType((*SnapshotRegionManifest)(nil), "hbase.pb.SnapshotRegionManifest")
	proto.RegisterType((*SnapshotRegionManifest_StoreFile)(nil), "hbase.pb.SnapshotRegionManifest.StoreFile")
	proto.RegisterType((*SnapshotRegionManifest_FamilyFiles)(nil), "hbase.pb.SnapshotRegionManifest.FamilyFiles")
	proto.RegisterType((*SnapshotDataManifest)(nil), "hbase.pb.SnapshotDataManifest")
}

func init() { proto.RegisterFile("Snapshot.proto", fileDescriptor_37871977bc5f5631) }

var fileDescriptor_37871977bc5f5631 = []byte{
	// 684 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0xcd, 0x6a, 0xdb, 0x40,
	0x10, 0xee, 0xca, 0x76, 0x6d, 0x8d, 0x9c, 0xd4, 0xac, 0x9d, 0xa2, 0xba, 0x84, 0x08, 0x1f, 0x8a,
	0x28, 0x45, 0x34, 0x2e, 0x85, 0x12, 0xe8, 0xc1, 0x6e, 0x62, 0x6c, 0xec, 0xb6, 0xe9, 0x2a, 0xa1,
	0x47, 0xb1, 0xb6, 0x57, 0xb1, 0x40, 0xd2, 0x8a, 0x5d, 0x25, 0x26, 0xe9, 0x0b, 0xf4, 0x25, 0x0a,
	0x3d, 0xf6, 0xde, 0xa7, 0xe9, 0xdb, 0x94, 0xd5, 0x8f, 0x9d, 0x90, 0x40, 0x6e, 0xfb, 0xcd, 0xcc,
	0xce, 0x7c, 0xdf, 0x37, 0xd2, 0xc2, 0xae, 0x1b, 0xd3, 0x44, 0xae, 0x78, 0xea, 0x24, 0x82, 0xa7,
	0x1c, 0x37, 0x56, 0x73, 0x2a, 0x99, 0x93, 0xcc, 0xbb, 0xed, 0xc1, 0x62, 0xc1, 0xa4, 0xfc, 0xc4,
	0xe3, 0x54, 0xf0, 0x30, 0x4f, 0x77, 0x1b, 0x23, 0xb7, 0x38, 0x19, 0xe3, 0x61, 0x56, 0xa8, 0x40,
	0xef, 0x9f, 0x06, 0xed, 0xb2, 0xd1, 0x31, 0x93, 0x0b, 0x11, 0x24, 0x69, 0xc0, 0x63, 0x8c, 0xa1,
	0x1a, 0xd3, 0x88, 0x99, 0xc8, 0xd2, 0x6c, 0x9d, 0x64, 0x67, 0xdc, 0x81, 0x5a, 0x4a, 0xe7, 0x21,
	0x33, 0x35, 0x0b, 0xd9, 0x3a, 0xc9, 0x01, 0x7e, 0x05, 0x3b, 0x0b, 0xc1, 0xa8, 0xba, 0xe5, 0xa5,
	0x41, 0xc4, 0xcc, 0x8a, 0x85, 0xec, 0xca, 0x11, 0x7a, 0x4b, 0x9a, 0x65, 0xfc, 0x2c, 0x88, 0x18,
	0xfe, 0x08, 0xd5, 0xf4, 0x3a, 0x61, 0x66, 0xd5, 0x42, 0xf6, 0x6e, 0xbf, 0xe7, 0x94, 0x74, 0x9d,
	0x07, 0xc6, 0x3b, 0x67, 0xd7, 0x09, 0x3b, 0xaa, 0x8d, 0x66, 0xe7, 0xee, 0x98, 0x64, 0xd7, 0xb0,
	0x09, 0xf5, 0x2b, 0x26, 0x64, 0xc0, 0x63, 0xb3, 0x66, 0x21, 0xbb, 0x46, 0x4a, 0xa8, 0x68, 0xf1,
	0x75, 0xcc, 0x84, 0xf9, 0x34, 0xa7, 0x95, 0x01, 0xfc, 0x0d, 0xf6, 0x2e, 0x25, 0x13, 0xd2, 0xa3,
	0xf1, 0xd2, 0x4b, 0x98, 0x88, 0x02, 0xa9, 0xaa, 0xa5, 0x59, 0xb7, 0x90, 0x6d, 0xf4, 0xf7, 0xb7,
	0xf3, 0xcf, 0x55, 0xd9, 0x20, 0x5e, 0x9e, 0x6e, 0x8b, 0x48, 0xfb, 0xf2, 0x7e, 0xb0, 0xe7, 0x40,
	0x55, 0xf1, 0xc2, 0x4d, 0x68, 0x1c, 0x4f, 0xdc, 0xc1, 0x70, 0x76, 0x72, 0xdc, 0x7a, 0x82, 0x75,
	0xc8, 0x79, 0xb6, 0x10, 0xde, 0x01, 0xdd, 0x9d, 0x4e, 0x4e, 0x73, 0xa8, 0xf5, 0xfe, 0x22, 0x68,
	0x95, 0xe2, 0x46, 0x41, 0xc8, 0x26, 0xb1, 0xcf, 0xf1, 0xbb, 0xc2, 0x06, 0x65, 0xec, 0x6e, 0xff,
	0xe0, 0xbe, 0x0d, 0x65, 0x65, 0xe6, 0x41, 0x21, 0xbe, 0x03, 0xb5, 0x95, 0x1f, 0x84, 0xb9, 0xb7,
	0x3a, 0xc9, 0x01, 0xde, 0x07, 0x58, 0xd3, 0xd0, 0x93, 0x4c, 0x5c, 0x31, 0x91, 0xf9, 0xaa, 0x13,
	0x7d, 0x4d, 0x43, 0x37, 0x0b, 0xe0, 0x17, 0xd0, 0x50, 0xe9, 0x6c, 0x8d, 0xb5, 0x2c, 0x59, 0x5f,
	0xd3, 0xf0, 0x0b, 0x8d, 0x58, 0xaf, 0x5b, 0x28, 0xd1, 0xa1, 0x36, 0x1e, 0x4d, 0x66, 0x27, 0x2d,
	0x84, 0xeb, 0x50, 0xf9, 0x3e, 0x98, 0xb5, 0xb4, 0xde, 0x9f, 0x0a, 0x3c, 0x2f, 0xb9, 0x10, 0x76,
	0x11, 0xf0, 0xf8, 0x33, 0x8d, 0x03, 0x9f, 0xc9, 0xf4, 0xf6, 0x0e, 0xd0, 0xdd, 0x1d, 0xbc, 0x07,
	0x43, 0x64, 0xb5, 0x5e, 0x10, 0xfb, 0xdc, 0xd4, 0x2c, 0xcd, 0x36, 0xfa, 0x9d, 0xad, 0xb8, 0xbc,
	0x91, 0x92, 0x45, 0x40, 0x6c, 0xce, 0xf8, 0x2b, 0x34, 0x7d, 0x1a, 0x05, 0xe1, 0xb5, 0xa7, 0x04,
	0x49, 0xb3, 0x62, 0x55, 0x6c, 0xa3, 0xff, 0xe6, 0xbe, 0x29, 0x77, 0x89, 0x38, 0xa3, 0xec, 0x92,
	0x72, 0x4a, 0x12, 0xc3, 0xdf, 0x82, 0x2e, 0x07, 0xdd, 0x4d, 0xb9, 0x60, 0x0a, 0x3d, 0xf8, 0x0d,
	0x1f, 0x82, 0x2e, 0x98, 0xcf, 0x04, 0x8b, 0x17, 0xf9, 0x77, 0x6c, 0xf4, 0xdb, 0xb7, 0x69, 0x16,
	0x29, 0xb2, 0xad, 0xc2, 0x2f, 0x41, 0x57, 0xec, 0x3c, 0x19, 0xdc, 0xe4, 0x0b, 0xa8, 0x92, 0x86,
	0x0a, 0xb8, 0xc1, 0x0d, 0xeb, 0xfe, 0x00, 0xe3, 0x16, 0x19, 0x7c, 0x00, 0x05, 0x1d, 0x6f, 0x33,
	0xb9, 0x49, 0x20, 0x0f, 0x29, 0xe7, 0xf1, 0x14, 0x0c, 0xa9, 0x08, 0x16, 0x82, 0xb5, 0x4c, 0xf0,
	0xeb, 0x47, 0x05, 0x6f, 0x44, 0x11, 0x90, 0xe5, 0x51, 0xf6, 0x7e, 0x21, 0xe8, 0x6c, 0xfe, 0x1e,
	0x9a, 0xd2, 0xcd, 0xa2, 0x3e, 0x40, 0x33, 0xfb, 0x39, 0x3d, 0xb9, 0x58, 0xb1, 0x88, 0x66, 0x3c,
	0x8c, 0xfe, 0xde, 0x76, 0xcc, 0x99, 0xca, 0xba, 0x59, 0x92, 0x18, 0xe9, 0x16, 0xe0, 0x29, 0xb4,
	0x8a, 0x45, 0x46, 0x45, 0xb3, 0x92, 0xa4, 0xf5, 0x18, 0x49, 0xf2, 0x4c, 0xdc, 0xc1, 0x72, 0x38,
	0x85, 0x43, 0x2e, 0x2e, 0x1c, 0x9a, 0xd0, 0xc5, 0x8a, 0x39, 0x2b, 0xba, 0xe4, 0x3c, 0x29, 0xba,
	0xc8, 0x15, 0x5d, 0xb2, 0x65, 0xfe, 0x08, 0xcd, 0x2f, 0x7d, 0xe7, 0x82, 0xc5, 0x4c, 0xd0, 0x94,
	0x2d, 0x87, 0x9b, 0x77, 0xed, 0x54, 0xe5, 0xe4, 0x18, 0xfd, 0x44, 0xe8, 0x37, 0x42, 0xff, 0x03,
	0x00, 0x00, 0xff, 0xff, 0xbd, 0x6d, 0x3b, 0x9a, 0xf0, 0x04, 0x00, 0x00,
}
