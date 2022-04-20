// Code generated by protoc-gen-go. DO NOT EDIT.
// source: WAL.proto

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

type ScopeType int32

const (
	ScopeType_REPLICATION_SCOPE_LOCAL  ScopeType = 0
	ScopeType_REPLICATION_SCOPE_GLOBAL ScopeType = 1
	ScopeType_REPLICATION_SCOPE_SERIAL ScopeType = 2
)

var ScopeType_name = map[int32]string{
	0: "REPLICATION_SCOPE_LOCAL",
	1: "REPLICATION_SCOPE_GLOBAL",
	2: "REPLICATION_SCOPE_SERIAL",
}

var ScopeType_value = map[string]int32{
	"REPLICATION_SCOPE_LOCAL":  0,
	"REPLICATION_SCOPE_GLOBAL": 1,
	"REPLICATION_SCOPE_SERIAL": 2,
}

func (x ScopeType) Enum() *ScopeType {
	p := new(ScopeType)
	*p = x
	return p
}

func (x ScopeType) String() string {
	return proto.EnumName(ScopeType_name, int32(x))
}

func (x *ScopeType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ScopeType_value, data, "ScopeType")
	if err != nil {
		return err
	}
	*x = ScopeType(value)
	return nil
}

func (ScopeType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_84a5c6a637a14f9b, []int{0}
}

type FlushDescriptor_FlushAction int32

const (
	FlushDescriptor_START_FLUSH  FlushDescriptor_FlushAction = 0
	FlushDescriptor_COMMIT_FLUSH FlushDescriptor_FlushAction = 1
	FlushDescriptor_ABORT_FLUSH  FlushDescriptor_FlushAction = 2
	FlushDescriptor_CANNOT_FLUSH FlushDescriptor_FlushAction = 3
)

var FlushDescriptor_FlushAction_name = map[int32]string{
	0: "START_FLUSH",
	1: "COMMIT_FLUSH",
	2: "ABORT_FLUSH",
	3: "CANNOT_FLUSH",
}

var FlushDescriptor_FlushAction_value = map[string]int32{
	"START_FLUSH":  0,
	"COMMIT_FLUSH": 1,
	"ABORT_FLUSH":  2,
	"CANNOT_FLUSH": 3,
}

func (x FlushDescriptor_FlushAction) Enum() *FlushDescriptor_FlushAction {
	p := new(FlushDescriptor_FlushAction)
	*p = x
	return p
}

func (x FlushDescriptor_FlushAction) String() string {
	return proto.EnumName(FlushDescriptor_FlushAction_name, int32(x))
}

func (x *FlushDescriptor_FlushAction) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(FlushDescriptor_FlushAction_value, data, "FlushDescriptor_FlushAction")
	if err != nil {
		return err
	}
	*x = FlushDescriptor_FlushAction(value)
	return nil
}

func (FlushDescriptor_FlushAction) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_84a5c6a637a14f9b, []int{4, 0}
}

type RegionEventDescriptor_EventType int32

const (
	RegionEventDescriptor_REGION_OPEN  RegionEventDescriptor_EventType = 0
	RegionEventDescriptor_REGION_CLOSE RegionEventDescriptor_EventType = 1
)

var RegionEventDescriptor_EventType_name = map[int32]string{
	0: "REGION_OPEN",
	1: "REGION_CLOSE",
}

var RegionEventDescriptor_EventType_value = map[string]int32{
	"REGION_OPEN":  0,
	"REGION_CLOSE": 1,
}

func (x RegionEventDescriptor_EventType) Enum() *RegionEventDescriptor_EventType {
	p := new(RegionEventDescriptor_EventType)
	*p = x
	return p
}

func (x RegionEventDescriptor_EventType) String() string {
	return proto.EnumName(RegionEventDescriptor_EventType_name, int32(x))
}

func (x *RegionEventDescriptor_EventType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(RegionEventDescriptor_EventType_value, data, "RegionEventDescriptor_EventType")
	if err != nil {
		return err
	}
	*x = RegionEventDescriptor_EventType(value)
	return nil
}

func (RegionEventDescriptor_EventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_84a5c6a637a14f9b, []int{7, 0}
}

type WALHeader struct {
	HasCompression       *bool    `protobuf:"varint,1,opt,name=has_compression,json=hasCompression" json:"has_compression,omitempty"`
	EncryptionKey        []byte   `protobuf:"bytes,2,opt,name=encryption_key,json=encryptionKey" json:"encryption_key,omitempty"`
	HasTagCompression    *bool    `protobuf:"varint,3,opt,name=has_tag_compression,json=hasTagCompression" json:"has_tag_compression,omitempty"`
	WriterClsName        *string  `protobuf:"bytes,4,opt,name=writer_cls_name,json=writerClsName" json:"writer_cls_name,omitempty"`
	CellCodecClsName     *string  `protobuf:"bytes,5,opt,name=cell_codec_cls_name,json=cellCodecClsName" json:"cell_codec_cls_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WALHeader) Reset()         { *m = WALHeader{} }
func (m *WALHeader) String() string { return proto.CompactTextString(m) }
func (*WALHeader) ProtoMessage()    {}
func (*WALHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_84a5c6a637a14f9b, []int{0}
}

func (m *WALHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WALHeader.Unmarshal(m, b)
}
func (m *WALHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WALHeader.Marshal(b, m, deterministic)
}
func (m *WALHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WALHeader.Merge(m, src)
}
func (m *WALHeader) XXX_Size() int {
	return xxx_messageInfo_WALHeader.Size(m)
}
func (m *WALHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_WALHeader.DiscardUnknown(m)
}

var xxx_messageInfo_WALHeader proto.InternalMessageInfo

func (m *WALHeader) GetHasCompression() bool {
	if m != nil && m.HasCompression != nil {
		return *m.HasCompression
	}
	return false
}

func (m *WALHeader) GetEncryptionKey() []byte {
	if m != nil {
		return m.EncryptionKey
	}
	return nil
}

func (m *WALHeader) GetHasTagCompression() bool {
	if m != nil && m.HasTagCompression != nil {
		return *m.HasTagCompression
	}
	return false
}

func (m *WALHeader) GetWriterClsName() string {
	if m != nil && m.WriterClsName != nil {
		return *m.WriterClsName
	}
	return ""
}

func (m *WALHeader) GetCellCodecClsName() string {
	if m != nil && m.CellCodecClsName != nil {
		return *m.CellCodecClsName
	}
	return ""
}

//
// Protocol buffer version of WALKey; see WALKey comment, not really a key but WALEdit header
// for some KVs
type WALKey struct {
	EncodedRegionName []byte  `protobuf:"bytes,1,req,name=encoded_region_name,json=encodedRegionName" json:"encoded_region_name,omitempty"`
	TableName         []byte  `protobuf:"bytes,2,req,name=table_name,json=tableName" json:"table_name,omitempty"`
	LogSequenceNumber *uint64 `protobuf:"varint,3,req,name=log_sequence_number,json=logSequenceNumber" json:"log_sequence_number,omitempty"`
	WriteTime         *uint64 `protobuf:"varint,4,req,name=write_time,json=writeTime" json:"write_time,omitempty"`
	//
	//This parameter is deprecated in favor of clusters which
	//contains the list of clusters that have consumed the change.
	//It is retained so that the log created by earlier releases (0.94)
	//can be read by the newer releases.
	ClusterId        *UUID          `protobuf:"bytes,5,opt,name=cluster_id,json=clusterId" json:"cluster_id,omitempty"` // Deprecated: Do not use.
	Scopes           []*FamilyScope `protobuf:"bytes,6,rep,name=scopes" json:"scopes,omitempty"`
	FollowingKvCount *uint32        `protobuf:"varint,7,opt,name=following_kv_count,json=followingKvCount" json:"following_kv_count,omitempty"`
	//
	//This field contains the list of clusters that have
	//consumed the change
	ClusterIds           []*UUID  `protobuf:"bytes,8,rep,name=cluster_ids,json=clusterIds" json:"cluster_ids,omitempty"`
	NonceGroup           *uint64  `protobuf:"varint,9,opt,name=nonceGroup" json:"nonceGroup,omitempty"`
	Nonce                *uint64  `protobuf:"varint,10,opt,name=nonce" json:"nonce,omitempty"`
	OrigSequenceNumber   *uint64  `protobuf:"varint,11,opt,name=orig_sequence_number,json=origSequenceNumber" json:"orig_sequence_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WALKey) Reset()         { *m = WALKey{} }
func (m *WALKey) String() string { return proto.CompactTextString(m) }
func (*WALKey) ProtoMessage()    {}
func (*WALKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_84a5c6a637a14f9b, []int{1}
}

func (m *WALKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WALKey.Unmarshal(m, b)
}
func (m *WALKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WALKey.Marshal(b, m, deterministic)
}
func (m *WALKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WALKey.Merge(m, src)
}
func (m *WALKey) XXX_Size() int {
	return xxx_messageInfo_WALKey.Size(m)
}
func (m *WALKey) XXX_DiscardUnknown() {
	xxx_messageInfo_WALKey.DiscardUnknown(m)
}

var xxx_messageInfo_WALKey proto.InternalMessageInfo

func (m *WALKey) GetEncodedRegionName() []byte {
	if m != nil {
		return m.EncodedRegionName
	}
	return nil
}

func (m *WALKey) GetTableName() []byte {
	if m != nil {
		return m.TableName
	}
	return nil
}

func (m *WALKey) GetLogSequenceNumber() uint64 {
	if m != nil && m.LogSequenceNumber != nil {
		return *m.LogSequenceNumber
	}
	return 0
}

func (m *WALKey) GetWriteTime() uint64 {
	if m != nil && m.WriteTime != nil {
		return *m.WriteTime
	}
	return 0
}

// Deprecated: Do not use.
func (m *WALKey) GetClusterId() *UUID {
	if m != nil {
		return m.ClusterId
	}
	return nil
}

func (m *WALKey) GetScopes() []*FamilyScope {
	if m != nil {
		return m.Scopes
	}
	return nil
}

func (m *WALKey) GetFollowingKvCount() uint32 {
	if m != nil && m.FollowingKvCount != nil {
		return *m.FollowingKvCount
	}
	return 0
}

func (m *WALKey) GetClusterIds() []*UUID {
	if m != nil {
		return m.ClusterIds
	}
	return nil
}

func (m *WALKey) GetNonceGroup() uint64 {
	if m != nil && m.NonceGroup != nil {
		return *m.NonceGroup
	}
	return 0
}

func (m *WALKey) GetNonce() uint64 {
	if m != nil && m.Nonce != nil {
		return *m.Nonce
	}
	return 0
}

func (m *WALKey) GetOrigSequenceNumber() uint64 {
	if m != nil && m.OrigSequenceNumber != nil {
		return *m.OrigSequenceNumber
	}
	return 0
}

type FamilyScope struct {
	Family               []byte     `protobuf:"bytes,1,req,name=family" json:"family,omitempty"`
	ScopeType            *ScopeType `protobuf:"varint,2,req,name=scope_type,json=scopeType,enum=hbase.pb.ScopeType" json:"scope_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *FamilyScope) Reset()         { *m = FamilyScope{} }
func (m *FamilyScope) String() string { return proto.CompactTextString(m) }
func (*FamilyScope) ProtoMessage()    {}
func (*FamilyScope) Descriptor() ([]byte, []int) {
	return fileDescriptor_84a5c6a637a14f9b, []int{2}
}

func (m *FamilyScope) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FamilyScope.Unmarshal(m, b)
}
func (m *FamilyScope) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FamilyScope.Marshal(b, m, deterministic)
}
func (m *FamilyScope) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FamilyScope.Merge(m, src)
}
func (m *FamilyScope) XXX_Size() int {
	return xxx_messageInfo_FamilyScope.Size(m)
}
func (m *FamilyScope) XXX_DiscardUnknown() {
	xxx_messageInfo_FamilyScope.DiscardUnknown(m)
}

var xxx_messageInfo_FamilyScope proto.InternalMessageInfo

func (m *FamilyScope) GetFamily() []byte {
	if m != nil {
		return m.Family
	}
	return nil
}

func (m *FamilyScope) GetScopeType() ScopeType {
	if m != nil && m.ScopeType != nil {
		return *m.ScopeType
	}
	return ScopeType_REPLICATION_SCOPE_LOCAL
}

//*
// Special WAL entry to hold all related to a compaction.
// Written to WAL before completing compaction.  There is
// sufficient info in the below message to complete later
// the * compaction should we fail the WAL write.
type CompactionDescriptor struct {
	TableName            []byte   `protobuf:"bytes,1,req,name=table_name,json=tableName" json:"table_name,omitempty"`
	EncodedRegionName    []byte   `protobuf:"bytes,2,req,name=encoded_region_name,json=encodedRegionName" json:"encoded_region_name,omitempty"`
	FamilyName           []byte   `protobuf:"bytes,3,req,name=family_name,json=familyName" json:"family_name,omitempty"`
	CompactionInput      []string `protobuf:"bytes,4,rep,name=compaction_input,json=compactionInput" json:"compaction_input,omitempty"`
	CompactionOutput     []string `protobuf:"bytes,5,rep,name=compaction_output,json=compactionOutput" json:"compaction_output,omitempty"`
	StoreHomeDir         *string  `protobuf:"bytes,6,req,name=store_home_dir,json=storeHomeDir" json:"store_home_dir,omitempty"`
	RegionName           []byte   `protobuf:"bytes,7,opt,name=region_name,json=regionName" json:"region_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CompactionDescriptor) Reset()         { *m = CompactionDescriptor{} }
func (m *CompactionDescriptor) String() string { return proto.CompactTextString(m) }
func (*CompactionDescriptor) ProtoMessage()    {}
func (*CompactionDescriptor) Descriptor() ([]byte, []int) {
	return fileDescriptor_84a5c6a637a14f9b, []int{3}
}

func (m *CompactionDescriptor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CompactionDescriptor.Unmarshal(m, b)
}
func (m *CompactionDescriptor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CompactionDescriptor.Marshal(b, m, deterministic)
}
func (m *CompactionDescriptor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompactionDescriptor.Merge(m, src)
}
func (m *CompactionDescriptor) XXX_Size() int {
	return xxx_messageInfo_CompactionDescriptor.Size(m)
}
func (m *CompactionDescriptor) XXX_DiscardUnknown() {
	xxx_messageInfo_CompactionDescriptor.DiscardUnknown(m)
}

var xxx_messageInfo_CompactionDescriptor proto.InternalMessageInfo

func (m *CompactionDescriptor) GetTableName() []byte {
	if m != nil {
		return m.TableName
	}
	return nil
}

func (m *CompactionDescriptor) GetEncodedRegionName() []byte {
	if m != nil {
		return m.EncodedRegionName
	}
	return nil
}

func (m *CompactionDescriptor) GetFamilyName() []byte {
	if m != nil {
		return m.FamilyName
	}
	return nil
}

func (m *CompactionDescriptor) GetCompactionInput() []string {
	if m != nil {
		return m.CompactionInput
	}
	return nil
}

func (m *CompactionDescriptor) GetCompactionOutput() []string {
	if m != nil {
		return m.CompactionOutput
	}
	return nil
}

func (m *CompactionDescriptor) GetStoreHomeDir() string {
	if m != nil && m.StoreHomeDir != nil {
		return *m.StoreHomeDir
	}
	return ""
}

func (m *CompactionDescriptor) GetRegionName() []byte {
	if m != nil {
		return m.RegionName
	}
	return nil
}

//*
// Special WAL entry to hold all related to a flush.
type FlushDescriptor struct {
	Action               *FlushDescriptor_FlushAction            `protobuf:"varint,1,req,name=action,enum=hbase.pb.FlushDescriptor_FlushAction" json:"action,omitempty"`
	TableName            []byte                                  `protobuf:"bytes,2,req,name=table_name,json=tableName" json:"table_name,omitempty"`
	EncodedRegionName    []byte                                  `protobuf:"bytes,3,req,name=encoded_region_name,json=encodedRegionName" json:"encoded_region_name,omitempty"`
	FlushSequenceNumber  *uint64                                 `protobuf:"varint,4,opt,name=flush_sequence_number,json=flushSequenceNumber" json:"flush_sequence_number,omitempty"`
	StoreFlushes         []*FlushDescriptor_StoreFlushDescriptor `protobuf:"bytes,5,rep,name=store_flushes,json=storeFlushes" json:"store_flushes,omitempty"`
	RegionName           []byte                                  `protobuf:"bytes,6,opt,name=region_name,json=regionName" json:"region_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_unrecognized     []byte                                  `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *FlushDescriptor) Reset()         { *m = FlushDescriptor{} }
func (m *FlushDescriptor) String() string { return proto.CompactTextString(m) }
func (*FlushDescriptor) ProtoMessage()    {}
func (*FlushDescriptor) Descriptor() ([]byte, []int) {
	return fileDescriptor_84a5c6a637a14f9b, []int{4}
}

func (m *FlushDescriptor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlushDescriptor.Unmarshal(m, b)
}
func (m *FlushDescriptor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlushDescriptor.Marshal(b, m, deterministic)
}
func (m *FlushDescriptor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlushDescriptor.Merge(m, src)
}
func (m *FlushDescriptor) XXX_Size() int {
	return xxx_messageInfo_FlushDescriptor.Size(m)
}
func (m *FlushDescriptor) XXX_DiscardUnknown() {
	xxx_messageInfo_FlushDescriptor.DiscardUnknown(m)
}

var xxx_messageInfo_FlushDescriptor proto.InternalMessageInfo

func (m *FlushDescriptor) GetAction() FlushDescriptor_FlushAction {
	if m != nil && m.Action != nil {
		return *m.Action
	}
	return FlushDescriptor_START_FLUSH
}

func (m *FlushDescriptor) GetTableName() []byte {
	if m != nil {
		return m.TableName
	}
	return nil
}

func (m *FlushDescriptor) GetEncodedRegionName() []byte {
	if m != nil {
		return m.EncodedRegionName
	}
	return nil
}

func (m *FlushDescriptor) GetFlushSequenceNumber() uint64 {
	if m != nil && m.FlushSequenceNumber != nil {
		return *m.FlushSequenceNumber
	}
	return 0
}

func (m *FlushDescriptor) GetStoreFlushes() []*FlushDescriptor_StoreFlushDescriptor {
	if m != nil {
		return m.StoreFlushes
	}
	return nil
}

func (m *FlushDescriptor) GetRegionName() []byte {
	if m != nil {
		return m.RegionName
	}
	return nil
}

type FlushDescriptor_StoreFlushDescriptor struct {
	FamilyName           []byte   `protobuf:"bytes,1,req,name=family_name,json=familyName" json:"family_name,omitempty"`
	StoreHomeDir         *string  `protobuf:"bytes,2,req,name=store_home_dir,json=storeHomeDir" json:"store_home_dir,omitempty"`
	FlushOutput          []string `protobuf:"bytes,3,rep,name=flush_output,json=flushOutput" json:"flush_output,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FlushDescriptor_StoreFlushDescriptor) Reset()         { *m = FlushDescriptor_StoreFlushDescriptor{} }
func (m *FlushDescriptor_StoreFlushDescriptor) String() string { return proto.CompactTextString(m) }
func (*FlushDescriptor_StoreFlushDescriptor) ProtoMessage()    {}
func (*FlushDescriptor_StoreFlushDescriptor) Descriptor() ([]byte, []int) {
	return fileDescriptor_84a5c6a637a14f9b, []int{4, 0}
}

func (m *FlushDescriptor_StoreFlushDescriptor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlushDescriptor_StoreFlushDescriptor.Unmarshal(m, b)
}
func (m *FlushDescriptor_StoreFlushDescriptor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlushDescriptor_StoreFlushDescriptor.Marshal(b, m, deterministic)
}
func (m *FlushDescriptor_StoreFlushDescriptor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlushDescriptor_StoreFlushDescriptor.Merge(m, src)
}
func (m *FlushDescriptor_StoreFlushDescriptor) XXX_Size() int {
	return xxx_messageInfo_FlushDescriptor_StoreFlushDescriptor.Size(m)
}
func (m *FlushDescriptor_StoreFlushDescriptor) XXX_DiscardUnknown() {
	xxx_messageInfo_FlushDescriptor_StoreFlushDescriptor.DiscardUnknown(m)
}

var xxx_messageInfo_FlushDescriptor_StoreFlushDescriptor proto.InternalMessageInfo

func (m *FlushDescriptor_StoreFlushDescriptor) GetFamilyName() []byte {
	if m != nil {
		return m.FamilyName
	}
	return nil
}

func (m *FlushDescriptor_StoreFlushDescriptor) GetStoreHomeDir() string {
	if m != nil && m.StoreHomeDir != nil {
		return *m.StoreHomeDir
	}
	return ""
}

func (m *FlushDescriptor_StoreFlushDescriptor) GetFlushOutput() []string {
	if m != nil {
		return m.FlushOutput
	}
	return nil
}

type StoreDescriptor struct {
	FamilyName           []byte   `protobuf:"bytes,1,req,name=family_name,json=familyName" json:"family_name,omitempty"`
	StoreHomeDir         *string  `protobuf:"bytes,2,req,name=store_home_dir,json=storeHomeDir" json:"store_home_dir,omitempty"`
	StoreFile            []string `protobuf:"bytes,3,rep,name=store_file,json=storeFile" json:"store_file,omitempty"`
	StoreFileSizeBytes   *uint64  `protobuf:"varint,4,opt,name=store_file_size_bytes,json=storeFileSizeBytes" json:"store_file_size_bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StoreDescriptor) Reset()         { *m = StoreDescriptor{} }
func (m *StoreDescriptor) String() string { return proto.CompactTextString(m) }
func (*StoreDescriptor) ProtoMessage()    {}
func (*StoreDescriptor) Descriptor() ([]byte, []int) {
	return fileDescriptor_84a5c6a637a14f9b, []int{5}
}

func (m *StoreDescriptor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoreDescriptor.Unmarshal(m, b)
}
func (m *StoreDescriptor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoreDescriptor.Marshal(b, m, deterministic)
}
func (m *StoreDescriptor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoreDescriptor.Merge(m, src)
}
func (m *StoreDescriptor) XXX_Size() int {
	return xxx_messageInfo_StoreDescriptor.Size(m)
}
func (m *StoreDescriptor) XXX_DiscardUnknown() {
	xxx_messageInfo_StoreDescriptor.DiscardUnknown(m)
}

var xxx_messageInfo_StoreDescriptor proto.InternalMessageInfo

func (m *StoreDescriptor) GetFamilyName() []byte {
	if m != nil {
		return m.FamilyName
	}
	return nil
}

func (m *StoreDescriptor) GetStoreHomeDir() string {
	if m != nil && m.StoreHomeDir != nil {
		return *m.StoreHomeDir
	}
	return ""
}

func (m *StoreDescriptor) GetStoreFile() []string {
	if m != nil {
		return m.StoreFile
	}
	return nil
}

func (m *StoreDescriptor) GetStoreFileSizeBytes() uint64 {
	if m != nil && m.StoreFileSizeBytes != nil {
		return *m.StoreFileSizeBytes
	}
	return 0
}

//*
// Special WAL entry used for writing bulk load events to WAL
type BulkLoadDescriptor struct {
	TableName            *TableName         `protobuf:"bytes,1,req,name=table_name,json=tableName" json:"table_name,omitempty"`
	EncodedRegionName    []byte             `protobuf:"bytes,2,req,name=encoded_region_name,json=encodedRegionName" json:"encoded_region_name,omitempty"`
	Stores               []*StoreDescriptor `protobuf:"bytes,3,rep,name=stores" json:"stores,omitempty"`
	BulkloadSeqNum       *int64             `protobuf:"varint,4,req,name=bulkload_seq_num,json=bulkloadSeqNum" json:"bulkload_seq_num,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *BulkLoadDescriptor) Reset()         { *m = BulkLoadDescriptor{} }
func (m *BulkLoadDescriptor) String() string { return proto.CompactTextString(m) }
func (*BulkLoadDescriptor) ProtoMessage()    {}
func (*BulkLoadDescriptor) Descriptor() ([]byte, []int) {
	return fileDescriptor_84a5c6a637a14f9b, []int{6}
}

func (m *BulkLoadDescriptor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BulkLoadDescriptor.Unmarshal(m, b)
}
func (m *BulkLoadDescriptor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BulkLoadDescriptor.Marshal(b, m, deterministic)
}
func (m *BulkLoadDescriptor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BulkLoadDescriptor.Merge(m, src)
}
func (m *BulkLoadDescriptor) XXX_Size() int {
	return xxx_messageInfo_BulkLoadDescriptor.Size(m)
}
func (m *BulkLoadDescriptor) XXX_DiscardUnknown() {
	xxx_messageInfo_BulkLoadDescriptor.DiscardUnknown(m)
}

var xxx_messageInfo_BulkLoadDescriptor proto.InternalMessageInfo

func (m *BulkLoadDescriptor) GetTableName() *TableName {
	if m != nil {
		return m.TableName
	}
	return nil
}

func (m *BulkLoadDescriptor) GetEncodedRegionName() []byte {
	if m != nil {
		return m.EncodedRegionName
	}
	return nil
}

func (m *BulkLoadDescriptor) GetStores() []*StoreDescriptor {
	if m != nil {
		return m.Stores
	}
	return nil
}

func (m *BulkLoadDescriptor) GetBulkloadSeqNum() int64 {
	if m != nil && m.BulkloadSeqNum != nil {
		return *m.BulkloadSeqNum
	}
	return 0
}

//*
// Special WAL entry to hold all related to a region event (open/close).
type RegionEventDescriptor struct {
	EventType            *RegionEventDescriptor_EventType `protobuf:"varint,1,req,name=event_type,json=eventType,enum=hbase.pb.RegionEventDescriptor_EventType" json:"event_type,omitempty"`
	TableName            []byte                           `protobuf:"bytes,2,req,name=table_name,json=tableName" json:"table_name,omitempty"`
	EncodedRegionName    []byte                           `protobuf:"bytes,3,req,name=encoded_region_name,json=encodedRegionName" json:"encoded_region_name,omitempty"`
	LogSequenceNumber    *uint64                          `protobuf:"varint,4,opt,name=log_sequence_number,json=logSequenceNumber" json:"log_sequence_number,omitempty"`
	Stores               []*StoreDescriptor               `protobuf:"bytes,5,rep,name=stores" json:"stores,omitempty"`
	Server               *ServerName                      `protobuf:"bytes,6,opt,name=server" json:"server,omitempty"`
	RegionName           []byte                           `protobuf:"bytes,7,opt,name=region_name,json=regionName" json:"region_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *RegionEventDescriptor) Reset()         { *m = RegionEventDescriptor{} }
func (m *RegionEventDescriptor) String() string { return proto.CompactTextString(m) }
func (*RegionEventDescriptor) ProtoMessage()    {}
func (*RegionEventDescriptor) Descriptor() ([]byte, []int) {
	return fileDescriptor_84a5c6a637a14f9b, []int{7}
}

func (m *RegionEventDescriptor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegionEventDescriptor.Unmarshal(m, b)
}
func (m *RegionEventDescriptor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegionEventDescriptor.Marshal(b, m, deterministic)
}
func (m *RegionEventDescriptor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegionEventDescriptor.Merge(m, src)
}
func (m *RegionEventDescriptor) XXX_Size() int {
	return xxx_messageInfo_RegionEventDescriptor.Size(m)
}
func (m *RegionEventDescriptor) XXX_DiscardUnknown() {
	xxx_messageInfo_RegionEventDescriptor.DiscardUnknown(m)
}

var xxx_messageInfo_RegionEventDescriptor proto.InternalMessageInfo

func (m *RegionEventDescriptor) GetEventType() RegionEventDescriptor_EventType {
	if m != nil && m.EventType != nil {
		return *m.EventType
	}
	return RegionEventDescriptor_REGION_OPEN
}

func (m *RegionEventDescriptor) GetTableName() []byte {
	if m != nil {
		return m.TableName
	}
	return nil
}

func (m *RegionEventDescriptor) GetEncodedRegionName() []byte {
	if m != nil {
		return m.EncodedRegionName
	}
	return nil
}

func (m *RegionEventDescriptor) GetLogSequenceNumber() uint64 {
	if m != nil && m.LogSequenceNumber != nil {
		return *m.LogSequenceNumber
	}
	return 0
}

func (m *RegionEventDescriptor) GetStores() []*StoreDescriptor {
	if m != nil {
		return m.Stores
	}
	return nil
}

func (m *RegionEventDescriptor) GetServer() *ServerName {
	if m != nil {
		return m.Server
	}
	return nil
}

func (m *RegionEventDescriptor) GetRegionName() []byte {
	if m != nil {
		return m.RegionName
	}
	return nil
}

//*
// A trailer that is appended to the end of a properly closed WAL file.
// If missing, this is either a legacy or a corrupted WAL file.
// N.B. This trailer currently doesn't contain any information and we
// purposefully don't expose it in the WAL APIs. It's for future growth.
type WALTrailer struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WALTrailer) Reset()         { *m = WALTrailer{} }
func (m *WALTrailer) String() string { return proto.CompactTextString(m) }
func (*WALTrailer) ProtoMessage()    {}
func (*WALTrailer) Descriptor() ([]byte, []int) {
	return fileDescriptor_84a5c6a637a14f9b, []int{8}
}

func (m *WALTrailer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WALTrailer.Unmarshal(m, b)
}
func (m *WALTrailer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WALTrailer.Marshal(b, m, deterministic)
}
func (m *WALTrailer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WALTrailer.Merge(m, src)
}
func (m *WALTrailer) XXX_Size() int {
	return xxx_messageInfo_WALTrailer.Size(m)
}
func (m *WALTrailer) XXX_DiscardUnknown() {
	xxx_messageInfo_WALTrailer.DiscardUnknown(m)
}

var xxx_messageInfo_WALTrailer proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("hbase.pb.ScopeType", ScopeType_name, ScopeType_value)
	proto.RegisterEnum("hbase.pb.FlushDescriptor_FlushAction", FlushDescriptor_FlushAction_name, FlushDescriptor_FlushAction_value)
	proto.RegisterEnum("hbase.pb.RegionEventDescriptor_EventType", RegionEventDescriptor_EventType_name, RegionEventDescriptor_EventType_value)
	proto.RegisterType((*WALHeader)(nil), "hbase.pb.WALHeader")
	proto.RegisterType((*WALKey)(nil), "hbase.pb.WALKey")
	proto.RegisterType((*FamilyScope)(nil), "hbase.pb.FamilyScope")
	proto.RegisterType((*CompactionDescriptor)(nil), "hbase.pb.CompactionDescriptor")
	proto.RegisterType((*FlushDescriptor)(nil), "hbase.pb.FlushDescriptor")
	proto.RegisterType((*FlushDescriptor_StoreFlushDescriptor)(nil), "hbase.pb.FlushDescriptor.StoreFlushDescriptor")
	proto.RegisterType((*StoreDescriptor)(nil), "hbase.pb.StoreDescriptor")
	proto.RegisterType((*BulkLoadDescriptor)(nil), "hbase.pb.BulkLoadDescriptor")
	proto.RegisterType((*RegionEventDescriptor)(nil), "hbase.pb.RegionEventDescriptor")
	proto.RegisterType((*WALTrailer)(nil), "hbase.pb.WALTrailer")
}

func init() { proto.RegisterFile("WAL.proto", fileDescriptor_84a5c6a637a14f9b) }

var fileDescriptor_84a5c6a637a14f9b = []byte{
	// 1116 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0x41, 0x73, 0xda, 0x46,
	0x14, 0x8e, 0x00, 0x93, 0xe8, 0x09, 0x03, 0x5e, 0xdb, 0xad, 0x9a, 0x36, 0x2d, 0x65, 0x9a, 0x96,
	0xb4, 0x89, 0xda, 0x70, 0xef, 0x01, 0x08, 0x0e, 0x4c, 0x14, 0xf0, 0x48, 0x78, 0x3c, 0x3d, 0x69,
	0x84, 0xb4, 0x06, 0x8d, 0x85, 0x56, 0xd1, 0x4a, 0xce, 0x90, 0x53, 0x8f, 0x3d, 0xf7, 0x17, 0xb4,
	0x87, 0x5e, 0xfa, 0x67, 0x7a, 0xea, 0x1f, 0xe8, 0x2f, 0xe9, 0xec, 0x93, 0x00, 0x19, 0xf0, 0x34,
	0x33, 0xcd, 0x8d, 0xfd, 0xde, 0xb7, 0xbb, 0xef, 0x7d, 0xef, 0x7b, 0x2b, 0x40, 0xbe, 0xec, 0xe8,
	0x5a, 0x18, 0xb1, 0x98, 0x91, 0x07, 0xf3, 0xa9, 0xcd, 0xa9, 0x16, 0x4e, 0x1f, 0x2a, 0x83, 0x2e,
	0xfe, 0x12, 0x70, 0xf3, 0x1f, 0x09, 0x49, 0x03, 0x6a, 0xbb, 0x34, 0x22, 0xdf, 0x40, 0x6d, 0x6e,
	0x73, 0xcb, 0x61, 0x8b, 0x30, 0xa2, 0x9c, 0x7b, 0x2c, 0x50, 0xa5, 0x86, 0xd4, 0x7a, 0x60, 0x54,
	0xe7, 0x36, 0xef, 0x6d, 0x50, 0xf2, 0x18, 0xaa, 0x34, 0x70, 0xa2, 0x65, 0x18, 0x7b, 0x2c, 0xb0,
	0xae, 0xe9, 0x52, 0x2d, 0x34, 0xa4, 0x56, 0xc5, 0x38, 0xdc, 0xa0, 0xaf, 0xe8, 0x92, 0x68, 0x70,
	0x2c, 0xce, 0x8b, 0xed, 0xd9, 0xad, 0x33, 0x8b, 0x78, 0xe6, 0xd1, 0xdc, 0xe6, 0x13, 0x7b, 0x96,
	0x3f, 0xf6, 0x6b, 0xa8, 0xbd, 0x8d, 0xbc, 0x98, 0x46, 0x96, 0xe3, 0x73, 0x2b, 0xb0, 0x17, 0x54,
	0x2d, 0x35, 0xa4, 0x96, 0x6c, 0x1c, 0xa6, 0x70, 0xcf, 0xe7, 0x23, 0x7b, 0x41, 0xc9, 0x33, 0x38,
	0x76, 0xa8, 0xef, 0x5b, 0x0e, 0x73, 0xa9, 0xb3, 0xe1, 0x1e, 0x20, 0xb7, 0x2e, 0x42, 0x3d, 0x11,
	0xc9, 0xe8, 0xcd, 0xbf, 0x8a, 0x50, 0xbe, 0xec, 0xe8, 0x59, 0x46, 0x34, 0x10, 0xdb, 0x5c, 0x2b,
	0xa2, 0x33, 0x91, 0x3c, 0xee, 0x94, 0x1a, 0x85, 0x56, 0xc5, 0x38, 0xca, 0x42, 0x06, 0x46, 0xf0,
	0xa6, 0x47, 0x00, 0xb1, 0x3d, 0xf5, 0x69, 0x4a, 0x2b, 0x20, 0x4d, 0x46, 0x04, 0xc3, 0x1a, 0x1c,
	0xfb, 0x6c, 0x66, 0x71, 0xfa, 0x26, 0xa1, 0x81, 0x43, 0xad, 0x20, 0x59, 0x4c, 0x69, 0xa4, 0x16,
	0x1b, 0x85, 0x56, 0xc9, 0x38, 0xf2, 0xd9, 0xcc, 0xcc, 0x22, 0x23, 0x0c, 0x88, 0xe3, 0xb0, 0x12,
	0x2b, 0xf6, 0xb0, 0x36, 0x41, 0x93, 0x11, 0x99, 0x78, 0x0b, 0x4a, 0x9e, 0x03, 0x38, 0x7e, 0xc2,
	0x85, 0x00, 0x9e, 0x8b, 0xe5, 0x28, 0xed, 0xaa, 0xb6, 0xea, 0x9c, 0x76, 0x71, 0x31, 0x7c, 0xd1,
	0x2d, 0xa8, 0x92, 0x21, 0x67, 0xac, 0xa1, 0x4b, 0x9e, 0x41, 0x99, 0x3b, 0x2c, 0xa4, 0x5c, 0x2d,
	0x37, 0x8a, 0x2d, 0xa5, 0x7d, 0xba, 0xa1, 0x9f, 0xd9, 0x0b, 0xcf, 0x5f, 0x9a, 0x22, 0x6a, 0x64,
	0x24, 0xf2, 0x14, 0xc8, 0x15, 0xf3, 0x7d, 0xf6, 0xd6, 0x0b, 0x66, 0xd6, 0xf5, 0x8d, 0xe5, 0xb0,
	0x24, 0x88, 0xd5, 0xfb, 0x0d, 0xa9, 0x75, 0x68, 0xd4, 0xd7, 0x91, 0x57, 0x37, 0x3d, 0x81, 0x93,
	0xef, 0x41, 0xd9, 0xe4, 0xc3, 0xd5, 0x07, 0x78, 0xc3, 0x56, 0x42, 0x06, 0xac, 0x93, 0xe1, 0xe4,
	0x73, 0x80, 0x80, 0x05, 0x0e, 0x7d, 0x19, 0xb1, 0x24, 0x54, 0xe5, 0x86, 0xd4, 0x2a, 0x19, 0x39,
	0x84, 0x9c, 0xc0, 0x01, 0xae, 0x54, 0xc0, 0x50, 0xba, 0x20, 0x3f, 0xc0, 0x09, 0x8b, 0xbc, 0x5d,
	0x19, 0x15, 0x24, 0x11, 0x11, 0xbb, 0xad, 0x63, 0xf3, 0x27, 0x50, 0x72, 0xd5, 0x91, 0x8f, 0xa0,
	0x7c, 0x85, 0xcb, 0xac, 0x91, 0xd9, 0x8a, 0xb4, 0x01, 0xb0, 0x6e, 0x2b, 0x5e, 0x86, 0x69, 0xf7,
	0xaa, 0xed, 0xe3, 0x4d, 0xfa, 0xb8, 0x79, 0xb2, 0x0c, 0xa9, 0x21, 0xf3, 0xd5, 0xcf, 0xe6, 0x1f,
	0x05, 0x38, 0x11, 0x9e, 0xb4, 0x1d, 0xe1, 0xe2, 0x17, 0x94, 0x3b, 0x91, 0x17, 0xc6, 0x2c, 0xda,
	0xb2, 0x82, 0xb4, 0xc7, 0x0a, 0xfb, 0x9c, 0x55, 0xb8, 0xcb, 0x59, 0x5f, 0x80, 0x92, 0x66, 0x99,
	0xf2, 0x8a, 0xc8, 0x83, 0x14, 0x42, 0xc2, 0x13, 0xa8, 0x3b, 0xeb, 0x3c, 0x2c, 0x2f, 0x08, 0x93,
	0x58, 0x2d, 0x35, 0x8a, 0x2d, 0xd9, 0xa8, 0x6d, 0xf0, 0xa1, 0x80, 0xc9, 0x77, 0x70, 0x94, 0xa3,
	0xb2, 0x24, 0x16, 0xdc, 0x03, 0xe4, 0xe6, 0xce, 0x18, 0x23, 0x4e, 0xbe, 0x82, 0x2a, 0x8f, 0x59,
	0x44, 0xad, 0x39, 0x5b, 0x50, 0xcb, 0xf5, 0x22, 0xb5, 0xdc, 0x28, 0xb4, 0x64, 0xa3, 0x82, 0xe8,
	0x80, 0x2d, 0xe8, 0x0b, 0x2f, 0x12, 0xe9, 0xe5, 0xcb, 0xb8, 0x8f, 0xe3, 0x0d, 0xd1, 0x3a, 0xff,
	0xe6, 0xef, 0x25, 0xa8, 0x9d, 0xf9, 0x09, 0x9f, 0xe7, 0x24, 0xfa, 0x11, 0xca, 0xe9, 0x55, 0x28,
	0x4f, 0xb5, 0xfd, 0x38, 0x67, 0xc6, 0xdb, 0xd4, 0x74, 0xdd, 0x41, 0xb2, 0x91, 0x6d, 0x7a, 0x8f,
	0x61, 0xdb, 0xa7, 0x70, 0xf1, 0x2e, 0x85, 0xdb, 0x70, 0x7a, 0x25, 0x6e, 0xd9, 0xf1, 0x55, 0x09,
	0x7d, 0x75, 0x8c, 0xc1, 0xad, 0x01, 0x35, 0xe1, 0x30, 0x15, 0x07, 0x83, 0x94, 0xa3, 0x8a, 0x4a,
	0x5b, 0xbb, 0xbb, 0x10, 0x53, 0xd0, 0xb7, 0xc0, 0x4c, 0xcb, 0xb3, 0xf4, 0x8c, 0x6d, 0x2d, 0xcb,
	0xdb, 0x5a, 0x3e, 0xfc, 0x59, 0x82, 0x93, 0x7d, 0xe7, 0x6c, 0x9b, 0x44, 0xda, 0x31, 0xc9, 0x6e,
	0x33, 0x0b, 0x7b, 0x9a, 0xf9, 0x25, 0x54, 0x52, 0x25, 0x32, 0x6b, 0x14, 0xd1, 0x1a, 0x0a, 0x62,
	0xa9, 0x2b, 0x9a, 0x26, 0x28, 0xb9, 0x96, 0x90, 0x1a, 0x28, 0xe6, 0xa4, 0x63, 0x4c, 0xac, 0x33,
	0xfd, 0xc2, 0x1c, 0xd4, 0xef, 0x91, 0x3a, 0x54, 0x7a, 0xe3, 0xd7, 0xaf, 0x87, 0x2b, 0x44, 0x12,
	0x94, 0x4e, 0x77, 0xbc, 0xa6, 0x14, 0x90, 0xd2, 0x19, 0x8d, 0xc6, 0x2b, 0xa4, 0xd8, 0xfc, 0x53,
	0x82, 0x1a, 0xd6, 0xf5, 0xe1, 0x4b, 0x7a, 0x04, 0x90, 0x35, 0xca, 0xf3, 0x69, 0x56, 0x90, 0x9c,
	0xaa, 0xee, 0xf9, 0xe2, 0x25, 0x3d, 0xdd, 0x84, 0x2d, 0xee, 0xbd, 0xa3, 0xd6, 0x74, 0x19, 0x53,
	0x9e, 0xf5, 0x9e, 0xac, 0x99, 0xa6, 0xf7, 0x8e, 0x76, 0x45, 0xa4, 0xf9, 0xb7, 0x04, 0xa4, 0x9b,
	0xf8, 0xd7, 0x3a, 0xb3, 0xdd, 0x5c, 0xbe, 0xed, 0x9d, 0xb1, 0x57, 0xf2, 0x6f, 0xc8, 0x64, 0x65,
	0xcf, 0xff, 0xf3, 0x16, 0x3c, 0x87, 0x32, 0x26, 0xc4, 0xb1, 0x10, 0xa5, 0xfd, 0x49, 0xee, 0x8d,
	0xba, 0x2d, 0x9f, 0x91, 0x11, 0x49, 0x0b, 0xea, 0xd3, 0xc4, 0xbf, 0xf6, 0x99, 0xed, 0x0a, 0x7f,
	0x0b, 0x6b, 0xe3, 0xf7, 0xa4, 0x68, 0x54, 0x57, 0xb8, 0x49, 0xdf, 0x8c, 0x92, 0x45, 0xf3, 0xd7,
	0x22, 0x9c, 0xa6, 0x77, 0xf5, 0x6f, 0x68, 0x10, 0xe7, 0x4a, 0x1b, 0x00, 0x50, 0x01, 0xa5, 0xcf,
	0x63, 0x3a, 0xb2, 0x4f, 0x36, 0x57, 0xef, 0xdd, 0xa4, 0xe1, 0x3a, 0x7d, 0x34, 0xe9, 0xea, 0xe7,
	0x87, 0x9e, 0xdc, 0x3b, 0x3e, 0xab, 0x69, 0xef, 0xf6, 0x7c, 0x56, 0x37, 0xfa, 0x1d, 0xbc, 0xaf,
	0x7e, 0x4f, 0xa1, 0xcc, 0x69, 0x74, 0x43, 0x23, 0x1c, 0x47, 0xa5, 0x7d, 0x92, 0xdb, 0x82, 0x38,
	0xf6, 0x34, 0xe3, 0xfc, 0xf7, 0x6b, 0xa8, 0x81, 0xbc, 0x16, 0x46, 0x4c, 0x86, 0xd1, 0x7f, 0x39,
	0x1c, 0x8f, 0xac, 0xf1, 0x79, 0x7f, 0x94, 0x0e, 0x4f, 0x06, 0xf4, 0xf4, 0xb1, 0xd9, 0xaf, 0x4b,
	0xcd, 0x0a, 0xc0, 0x65, 0x47, 0x9f, 0x44, 0xb6, 0xe7, 0xd3, 0xe8, 0x5b, 0x17, 0xe4, 0xf5, 0xb7,
	0x88, 0x7c, 0x0a, 0x1f, 0x1b, 0xfd, 0x73, 0x7d, 0xd8, 0xeb, 0x4c, 0xc4, 0x0e, 0xb3, 0x37, 0x3e,
	0xef, 0x5b, 0xfa, 0xb8, 0xd7, 0xd1, 0xeb, 0xf7, 0xc8, 0x67, 0xa0, 0xee, 0x06, 0x5f, 0xea, 0xe3,
	0x6e, 0x47, 0xaf, 0x4b, 0xfb, 0xa3, 0x66, 0xdf, 0x18, 0x76, 0xf4, 0x7a, 0xa1, 0x7b, 0x06, 0xcf,
	0x59, 0x34, 0xd3, 0xec, 0xd0, 0x76, 0xe6, 0x54, 0x9b, 0xdb, 0x2e, 0x63, 0x61, 0x56, 0x35, 0x9f,
	0xdb, 0x2e, 0x75, 0xd3, 0xff, 0x84, 0xd3, 0xe4, 0x4a, 0x9b, 0xd1, 0x80, 0x46, 0x76, 0x4c, 0xdd,
	0xae, 0xf8, 0x77, 0x78, 0x2e, 0x60, 0x3e, 0x90, 0x7e, 0x91, 0xee, 0xfd, 0x26, 0x49, 0xff, 0x06,
	0x00, 0x00, 0xff, 0xff, 0x50, 0x78, 0x79, 0x82, 0x56, 0x0a, 0x00, 0x00,
}
