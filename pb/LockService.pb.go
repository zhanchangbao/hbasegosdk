// Code generated by protoc-gen-go. DO NOT EDIT.
// source: LockService.proto

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

type LockType int32

const (
	LockType_EXCLUSIVE LockType = 1
	LockType_SHARED    LockType = 2
)

var LockType_name = map[int32]string{
	1: "EXCLUSIVE",
	2: "SHARED",
}

var LockType_value = map[string]int32{
	"EXCLUSIVE": 1,
	"SHARED":    2,
}

func (x LockType) Enum() *LockType {
	p := new(LockType)
	*p = x
	return p
}

func (x LockType) String() string {
	return proto.EnumName(LockType_name, int32(x))
}

func (x *LockType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(LockType_value, data, "LockType")
	if err != nil {
		return err
	}
	*x = LockType(value)
	return nil
}

func (LockType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_81b19d81005d4392, []int{0}
}

type LockedResourceType int32

const (
	LockedResourceType_SERVER    LockedResourceType = 1
	LockedResourceType_NAMESPACE LockedResourceType = 2
	LockedResourceType_TABLE     LockedResourceType = 3
	LockedResourceType_REGION    LockedResourceType = 4
	LockedResourceType_PEER      LockedResourceType = 5
)

var LockedResourceType_name = map[int32]string{
	1: "SERVER",
	2: "NAMESPACE",
	3: "TABLE",
	4: "REGION",
	5: "PEER",
}

var LockedResourceType_value = map[string]int32{
	"SERVER":    1,
	"NAMESPACE": 2,
	"TABLE":     3,
	"REGION":    4,
	"PEER":      5,
}

func (x LockedResourceType) Enum() *LockedResourceType {
	p := new(LockedResourceType)
	*p = x
	return p
}

func (x LockedResourceType) String() string {
	return proto.EnumName(LockedResourceType_name, int32(x))
}

func (x *LockedResourceType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(LockedResourceType_value, data, "LockedResourceType")
	if err != nil {
		return err
	}
	*x = LockedResourceType(value)
	return nil
}

func (LockedResourceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_81b19d81005d4392, []int{1}
}

type LockHeartbeatResponse_LockStatus int32

const (
	LockHeartbeatResponse_UNLOCKED LockHeartbeatResponse_LockStatus = 1
	LockHeartbeatResponse_LOCKED   LockHeartbeatResponse_LockStatus = 2
)

var LockHeartbeatResponse_LockStatus_name = map[int32]string{
	1: "UNLOCKED",
	2: "LOCKED",
}

var LockHeartbeatResponse_LockStatus_value = map[string]int32{
	"UNLOCKED": 1,
	"LOCKED":   2,
}

func (x LockHeartbeatResponse_LockStatus) Enum() *LockHeartbeatResponse_LockStatus {
	p := new(LockHeartbeatResponse_LockStatus)
	*p = x
	return p
}

func (x LockHeartbeatResponse_LockStatus) String() string {
	return proto.EnumName(LockHeartbeatResponse_LockStatus_name, int32(x))
}

func (x *LockHeartbeatResponse_LockStatus) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(LockHeartbeatResponse_LockStatus_value, data, "LockHeartbeatResponse_LockStatus")
	if err != nil {
		return err
	}
	*x = LockHeartbeatResponse_LockStatus(value)
	return nil
}

func (LockHeartbeatResponse_LockStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_81b19d81005d4392, []int{3, 0}
}

type LockRequest struct {
	LockType             *LockType     `protobuf:"varint,1,req,name=lock_type,json=lockType,enum=hbase.pb.LockType" json:"lock_type,omitempty"`
	Namespace            *string       `protobuf:"bytes,2,opt,name=namespace" json:"namespace,omitempty"`
	TableName            *TableName    `protobuf:"bytes,3,opt,name=table_name,json=tableName" json:"table_name,omitempty"`
	RegionInfo           []*RegionInfo `protobuf:"bytes,4,rep,name=region_info,json=regionInfo" json:"region_info,omitempty"`
	Description          *string       `protobuf:"bytes,5,opt,name=description" json:"description,omitempty"`
	NonceGroup           *uint64       `protobuf:"varint,6,opt,name=nonce_group,json=nonceGroup,def=0" json:"nonce_group,omitempty"`
	Nonce                *uint64       `protobuf:"varint,7,opt,name=nonce,def=0" json:"nonce,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *LockRequest) Reset()         { *m = LockRequest{} }
func (m *LockRequest) String() string { return proto.CompactTextString(m) }
func (*LockRequest) ProtoMessage()    {}
func (*LockRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_81b19d81005d4392, []int{0}
}

func (m *LockRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LockRequest.Unmarshal(m, b)
}
func (m *LockRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LockRequest.Marshal(b, m, deterministic)
}
func (m *LockRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LockRequest.Merge(m, src)
}
func (m *LockRequest) XXX_Size() int {
	return xxx_messageInfo_LockRequest.Size(m)
}
func (m *LockRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LockRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LockRequest proto.InternalMessageInfo

const Default_LockRequest_NonceGroup uint64 = 0
const Default_LockRequest_Nonce uint64 = 0

func (m *LockRequest) GetLockType() LockType {
	if m != nil && m.LockType != nil {
		return *m.LockType
	}
	return LockType_EXCLUSIVE
}

func (m *LockRequest) GetNamespace() string {
	if m != nil && m.Namespace != nil {
		return *m.Namespace
	}
	return ""
}

func (m *LockRequest) GetTableName() *TableName {
	if m != nil {
		return m.TableName
	}
	return nil
}

func (m *LockRequest) GetRegionInfo() []*RegionInfo {
	if m != nil {
		return m.RegionInfo
	}
	return nil
}

func (m *LockRequest) GetDescription() string {
	if m != nil && m.Description != nil {
		return *m.Description
	}
	return ""
}

func (m *LockRequest) GetNonceGroup() uint64 {
	if m != nil && m.NonceGroup != nil {
		return *m.NonceGroup
	}
	return Default_LockRequest_NonceGroup
}

func (m *LockRequest) GetNonce() uint64 {
	if m != nil && m.Nonce != nil {
		return *m.Nonce
	}
	return Default_LockRequest_Nonce
}

type LockResponse struct {
	ProcId               *uint64  `protobuf:"varint,1,req,name=proc_id,json=procId" json:"proc_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LockResponse) Reset()         { *m = LockResponse{} }
func (m *LockResponse) String() string { return proto.CompactTextString(m) }
func (*LockResponse) ProtoMessage()    {}
func (*LockResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_81b19d81005d4392, []int{1}
}

func (m *LockResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LockResponse.Unmarshal(m, b)
}
func (m *LockResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LockResponse.Marshal(b, m, deterministic)
}
func (m *LockResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LockResponse.Merge(m, src)
}
func (m *LockResponse) XXX_Size() int {
	return xxx_messageInfo_LockResponse.Size(m)
}
func (m *LockResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LockResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LockResponse proto.InternalMessageInfo

func (m *LockResponse) GetProcId() uint64 {
	if m != nil && m.ProcId != nil {
		return *m.ProcId
	}
	return 0
}

type LockHeartbeatRequest struct {
	ProcId               *uint64  `protobuf:"varint,1,req,name=proc_id,json=procId" json:"proc_id,omitempty"`
	KeepAlive            *bool    `protobuf:"varint,2,opt,name=keep_alive,json=keepAlive,def=1" json:"keep_alive,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LockHeartbeatRequest) Reset()         { *m = LockHeartbeatRequest{} }
func (m *LockHeartbeatRequest) String() string { return proto.CompactTextString(m) }
func (*LockHeartbeatRequest) ProtoMessage()    {}
func (*LockHeartbeatRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_81b19d81005d4392, []int{2}
}

func (m *LockHeartbeatRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LockHeartbeatRequest.Unmarshal(m, b)
}
func (m *LockHeartbeatRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LockHeartbeatRequest.Marshal(b, m, deterministic)
}
func (m *LockHeartbeatRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LockHeartbeatRequest.Merge(m, src)
}
func (m *LockHeartbeatRequest) XXX_Size() int {
	return xxx_messageInfo_LockHeartbeatRequest.Size(m)
}
func (m *LockHeartbeatRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LockHeartbeatRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LockHeartbeatRequest proto.InternalMessageInfo

const Default_LockHeartbeatRequest_KeepAlive bool = true

func (m *LockHeartbeatRequest) GetProcId() uint64 {
	if m != nil && m.ProcId != nil {
		return *m.ProcId
	}
	return 0
}

func (m *LockHeartbeatRequest) GetKeepAlive() bool {
	if m != nil && m.KeepAlive != nil {
		return *m.KeepAlive
	}
	return Default_LockHeartbeatRequest_KeepAlive
}

type LockHeartbeatResponse struct {
	LockStatus *LockHeartbeatResponse_LockStatus `protobuf:"varint,1,req,name=lock_status,json=lockStatus,enum=hbase.pb.LockHeartbeatResponse_LockStatus" json:"lock_status,omitempty"`
	// Timeout of lock (if locked).
	TimeoutMs            *uint32  `protobuf:"varint,2,opt,name=timeout_ms,json=timeoutMs" json:"timeout_ms,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LockHeartbeatResponse) Reset()         { *m = LockHeartbeatResponse{} }
func (m *LockHeartbeatResponse) String() string { return proto.CompactTextString(m) }
func (*LockHeartbeatResponse) ProtoMessage()    {}
func (*LockHeartbeatResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_81b19d81005d4392, []int{3}
}

func (m *LockHeartbeatResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LockHeartbeatResponse.Unmarshal(m, b)
}
func (m *LockHeartbeatResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LockHeartbeatResponse.Marshal(b, m, deterministic)
}
func (m *LockHeartbeatResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LockHeartbeatResponse.Merge(m, src)
}
func (m *LockHeartbeatResponse) XXX_Size() int {
	return xxx_messageInfo_LockHeartbeatResponse.Size(m)
}
func (m *LockHeartbeatResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LockHeartbeatResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LockHeartbeatResponse proto.InternalMessageInfo

func (m *LockHeartbeatResponse) GetLockStatus() LockHeartbeatResponse_LockStatus {
	if m != nil && m.LockStatus != nil {
		return *m.LockStatus
	}
	return LockHeartbeatResponse_UNLOCKED
}

func (m *LockHeartbeatResponse) GetTimeoutMs() uint32 {
	if m != nil && m.TimeoutMs != nil {
		return *m.TimeoutMs
	}
	return 0
}

type LockProcedureData struct {
	LockType             *LockType     `protobuf:"varint,1,req,name=lock_type,json=lockType,enum=hbase.pb.LockType" json:"lock_type,omitempty"`
	Namespace            *string       `protobuf:"bytes,2,opt,name=namespace" json:"namespace,omitempty"`
	TableName            *TableName    `protobuf:"bytes,3,opt,name=table_name,json=tableName" json:"table_name,omitempty"`
	RegionInfo           []*RegionInfo `protobuf:"bytes,4,rep,name=region_info,json=regionInfo" json:"region_info,omitempty"`
	Description          *string       `protobuf:"bytes,5,opt,name=description" json:"description,omitempty"`
	IsMasterLock         *bool         `protobuf:"varint,6,opt,name=is_master_lock,json=isMasterLock,def=0" json:"is_master_lock,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *LockProcedureData) Reset()         { *m = LockProcedureData{} }
func (m *LockProcedureData) String() string { return proto.CompactTextString(m) }
func (*LockProcedureData) ProtoMessage()    {}
func (*LockProcedureData) Descriptor() ([]byte, []int) {
	return fileDescriptor_81b19d81005d4392, []int{4}
}

func (m *LockProcedureData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LockProcedureData.Unmarshal(m, b)
}
func (m *LockProcedureData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LockProcedureData.Marshal(b, m, deterministic)
}
func (m *LockProcedureData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LockProcedureData.Merge(m, src)
}
func (m *LockProcedureData) XXX_Size() int {
	return xxx_messageInfo_LockProcedureData.Size(m)
}
func (m *LockProcedureData) XXX_DiscardUnknown() {
	xxx_messageInfo_LockProcedureData.DiscardUnknown(m)
}

var xxx_messageInfo_LockProcedureData proto.InternalMessageInfo

const Default_LockProcedureData_IsMasterLock bool = false

func (m *LockProcedureData) GetLockType() LockType {
	if m != nil && m.LockType != nil {
		return *m.LockType
	}
	return LockType_EXCLUSIVE
}

func (m *LockProcedureData) GetNamespace() string {
	if m != nil && m.Namespace != nil {
		return *m.Namespace
	}
	return ""
}

func (m *LockProcedureData) GetTableName() *TableName {
	if m != nil {
		return m.TableName
	}
	return nil
}

func (m *LockProcedureData) GetRegionInfo() []*RegionInfo {
	if m != nil {
		return m.RegionInfo
	}
	return nil
}

func (m *LockProcedureData) GetDescription() string {
	if m != nil && m.Description != nil {
		return *m.Description
	}
	return ""
}

func (m *LockProcedureData) GetIsMasterLock() bool {
	if m != nil && m.IsMasterLock != nil {
		return *m.IsMasterLock
	}
	return Default_LockProcedureData_IsMasterLock
}

type LockedResource struct {
	ResourceType                *LockedResourceType `protobuf:"varint,1,req,name=resource_type,json=resourceType,enum=hbase.pb.LockedResourceType" json:"resource_type,omitempty"`
	ResourceName                *string             `protobuf:"bytes,2,opt,name=resource_name,json=resourceName" json:"resource_name,omitempty"`
	LockType                    *LockType           `protobuf:"varint,3,req,name=lock_type,json=lockType,enum=hbase.pb.LockType" json:"lock_type,omitempty"`
	ExclusiveLockOwnerProcedure *Procedure          `protobuf:"bytes,4,opt,name=exclusive_lock_owner_procedure,json=exclusiveLockOwnerProcedure" json:"exclusive_lock_owner_procedure,omitempty"`
	SharedLockCount             *int32              `protobuf:"varint,5,opt,name=shared_lock_count,json=sharedLockCount" json:"shared_lock_count,omitempty"`
	WaitingProcedures           []*Procedure        `protobuf:"bytes,6,rep,name=waitingProcedures" json:"waitingProcedures,omitempty"`
	XXX_NoUnkeyedLiteral        struct{}            `json:"-"`
	XXX_unrecognized            []byte              `json:"-"`
	XXX_sizecache               int32               `json:"-"`
}

func (m *LockedResource) Reset()         { *m = LockedResource{} }
func (m *LockedResource) String() string { return proto.CompactTextString(m) }
func (*LockedResource) ProtoMessage()    {}
func (*LockedResource) Descriptor() ([]byte, []int) {
	return fileDescriptor_81b19d81005d4392, []int{5}
}

func (m *LockedResource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LockedResource.Unmarshal(m, b)
}
func (m *LockedResource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LockedResource.Marshal(b, m, deterministic)
}
func (m *LockedResource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LockedResource.Merge(m, src)
}
func (m *LockedResource) XXX_Size() int {
	return xxx_messageInfo_LockedResource.Size(m)
}
func (m *LockedResource) XXX_DiscardUnknown() {
	xxx_messageInfo_LockedResource.DiscardUnknown(m)
}

var xxx_messageInfo_LockedResource proto.InternalMessageInfo

func (m *LockedResource) GetResourceType() LockedResourceType {
	if m != nil && m.ResourceType != nil {
		return *m.ResourceType
	}
	return LockedResourceType_SERVER
}

func (m *LockedResource) GetResourceName() string {
	if m != nil && m.ResourceName != nil {
		return *m.ResourceName
	}
	return ""
}

func (m *LockedResource) GetLockType() LockType {
	if m != nil && m.LockType != nil {
		return *m.LockType
	}
	return LockType_EXCLUSIVE
}

func (m *LockedResource) GetExclusiveLockOwnerProcedure() *Procedure {
	if m != nil {
		return m.ExclusiveLockOwnerProcedure
	}
	return nil
}

func (m *LockedResource) GetSharedLockCount() int32 {
	if m != nil && m.SharedLockCount != nil {
		return *m.SharedLockCount
	}
	return 0
}

func (m *LockedResource) GetWaitingProcedures() []*Procedure {
	if m != nil {
		return m.WaitingProcedures
	}
	return nil
}

func init() {
	proto.RegisterEnum("hbase.pb.LockType", LockType_name, LockType_value)
	proto.RegisterEnum("hbase.pb.LockedResourceType", LockedResourceType_name, LockedResourceType_value)
	proto.RegisterEnum("hbase.pb.LockHeartbeatResponse_LockStatus", LockHeartbeatResponse_LockStatus_name, LockHeartbeatResponse_LockStatus_value)
	proto.RegisterType((*LockRequest)(nil), "hbase.pb.LockRequest")
	proto.RegisterType((*LockResponse)(nil), "hbase.pb.LockResponse")
	proto.RegisterType((*LockHeartbeatRequest)(nil), "hbase.pb.LockHeartbeatRequest")
	proto.RegisterType((*LockHeartbeatResponse)(nil), "hbase.pb.LockHeartbeatResponse")
	proto.RegisterType((*LockProcedureData)(nil), "hbase.pb.LockProcedureData")
	proto.RegisterType((*LockedResource)(nil), "hbase.pb.LockedResource")
}

func init() { proto.RegisterFile("LockService.proto", fileDescriptor_81b19d81005d4392) }

var fileDescriptor_81b19d81005d4392 = []byte{
	// 763 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x55, 0xd1, 0x6e, 0xdb, 0x36,
	0x14, 0x85, 0x1c, 0x3b, 0xb5, 0xae, 0x92, 0xd4, 0xe1, 0xda, 0x55, 0xc8, 0xba, 0xce, 0x50, 0xb1,
	0xcd, 0xc8, 0x00, 0x6f, 0x33, 0xb0, 0x97, 0x60, 0x2f, 0x8e, 0x2b, 0xd4, 0x46, 0x1d, 0xc7, 0xa0,
	0xdd, 0xa2, 0x6f, 0x02, 0x2d, 0x5d, 0xdb, 0x42, 0x6c, 0x51, 0x23, 0xa9, 0x74, 0xfd, 0x83, 0x7d,
	0xc4, 0x06, 0xec, 0x69, 0x2f, 0xfb, 0xbb, 0x7d, 0xc1, 0x40, 0x4a, 0x96, 0xed, 0x64, 0xc1, 0xf6,
	0xbc, 0xb7, 0xcb, 0x73, 0xef, 0xb9, 0x22, 0xcf, 0xb9, 0xa4, 0xe0, 0x74, 0xc8, 0xc3, 0x9b, 0x09,
	0x8a, 0xdb, 0x38, 0xc4, 0x76, 0x2a, 0xb8, 0xe2, 0xa4, 0xbe, 0x9c, 0x31, 0x89, 0xed, 0x74, 0x76,
	0xe6, 0xf4, 0x2f, 0x4d, 0xa4, 0xe1, 0xb3, 0xc7, 0x63, 0xc1, 0x43, 0x8c, 0x32, 0x51, 0x00, 0xde,
	0x1f, 0x15, 0x70, 0x34, 0x9b, 0xe2, 0x4f, 0x19, 0x4a, 0x45, 0xbe, 0x05, 0x7b, 0xc5, 0xc3, 0x9b,
	0x40, 0x7d, 0x4c, 0xd1, 0xb5, 0x9a, 0x95, 0xd6, 0x49, 0x87, 0xb4, 0x37, 0xbd, 0xda, 0xba, 0x72,
	0xfa, 0x31, 0x45, 0x5a, 0x5f, 0x15, 0x11, 0x79, 0x0e, 0x76, 0xc2, 0xd6, 0x28, 0x53, 0x16, 0xa2,
	0x5b, 0x69, 0x5a, 0x2d, 0x9b, 0x6e, 0x01, 0xd2, 0x01, 0x50, 0x6c, 0xb6, 0xc2, 0x40, 0x43, 0xee,
	0x41, 0xd3, 0x6a, 0x39, 0x9d, 0x4f, 0xb6, 0xfd, 0xa6, 0x3a, 0x37, 0x62, 0x6b, 0xa4, 0xb6, 0xda,
	0x84, 0xe4, 0x07, 0x70, 0x04, 0x2e, 0x62, 0x9e, 0x04, 0x71, 0x32, 0xe7, 0x6e, 0xb5, 0x79, 0xd0,
	0x72, 0x3a, 0x4f, 0xb6, 0x24, 0x6a, 0x92, 0x83, 0x64, 0xce, 0x29, 0x88, 0x32, 0x26, 0x4d, 0x70,
	0x22, 0x94, 0xa1, 0x88, 0x53, 0x15, 0xf3, 0xc4, 0xad, 0x99, 0xad, 0xec, 0x42, 0xc4, 0x03, 0x27,
	0xe1, 0x49, 0x88, 0xc1, 0x42, 0xf0, 0x2c, 0x75, 0x0f, 0x9b, 0x56, 0xab, 0x7a, 0x61, 0x7d, 0x47,
	0xc1, 0xa0, 0xaf, 0x35, 0x48, 0x9e, 0x41, 0xcd, 0xac, 0xdc, 0x47, 0x9b, 0x6c, 0xbe, 0xf6, 0xbe,
	0x86, 0xa3, 0x5c, 0x27, 0x99, 0xf2, 0x44, 0x22, 0x79, 0x06, 0x8f, 0x52, 0xc1, 0xc3, 0x20, 0x8e,
	0x8c, 0x4c, 0x55, 0x7a, 0xa8, 0x97, 0x83, 0xc8, 0x9b, 0xc2, 0x13, 0x5d, 0xd8, 0x47, 0x26, 0xd4,
	0x0c, 0x99, 0xda, 0x28, 0xfb, 0x10, 0x81, 0xbc, 0x04, 0xb8, 0x41, 0x4c, 0x03, 0xb6, 0x8a, 0x6f,
	0x73, 0x09, 0xeb, 0x17, 0x55, 0x25, 0x32, 0xa4, 0xb6, 0xc6, 0xbb, 0x1a, 0xf6, 0xfe, 0xb4, 0xe0,
	0xe9, 0x9d, 0xb6, 0xc5, 0x46, 0xde, 0x80, 0x63, 0x1c, 0x93, 0x8a, 0xa9, 0x4c, 0x16, 0x9e, 0x9d,
	0xef, 0x7b, 0x76, 0x8f, 0x65, 0xd0, 0x89, 0x61, 0x50, 0x58, 0x95, 0x31, 0xf9, 0x1c, 0x40, 0xc5,
	0x6b, 0xe4, 0x99, 0x0a, 0xd6, 0xd2, 0xec, 0xe5, 0x98, 0xda, 0x05, 0x72, 0x25, 0xbd, 0xaf, 0x00,
	0xb6, 0x44, 0x72, 0x04, 0xf5, 0xb7, 0xa3, 0xe1, 0x75, 0xef, 0x8d, 0xff, 0xaa, 0x61, 0x11, 0x80,
	0xc3, 0x22, 0xae, 0x78, 0xbf, 0x56, 0xf2, 0x99, 0x2c, 0xa7, 0xed, 0x15, 0x53, 0xec, 0x7f, 0x3d,
	0x5b, 0xdf, 0xc0, 0x49, 0x2c, 0x83, 0x35, 0x93, 0x0a, 0x45, 0xa0, 0x0f, 0x60, 0xc6, 0xab, 0x7e,
	0x51, 0x9b, 0xb3, 0x95, 0x44, 0x7a, 0x14, 0xcb, 0x2b, 0x93, 0xd3, 0xa7, 0xf4, 0xfe, 0xaa, 0xc0,
	0x89, 0x0e, 0x30, 0xa2, 0x28, 0x79, 0x26, 0x42, 0x24, 0x5d, 0x38, 0x16, 0x45, 0xbc, 0xab, 0xcf,
	0xf3, 0x7d, 0x7d, 0xb6, 0x04, 0xa3, 0xd4, 0x91, 0xd8, 0x59, 0x91, 0x97, 0x3b, 0x2d, 0x8c, 0x24,
	0xb9, 0x62, 0x65, 0x91, 0x11, 0x60, 0xcf, 0x83, 0x83, 0xff, 0xe0, 0xc1, 0x7b, 0x78, 0x81, 0x3f,
	0x87, 0xab, 0x4c, 0xc6, 0xb7, 0x68, 0x0e, 0x16, 0xf0, 0x0f, 0x09, 0x8a, 0x20, 0xdd, 0x58, 0xeb,
	0x56, 0xef, 0x2a, 0x5f, 0xba, 0x4e, 0x3f, 0x2b, 0xa9, 0xba, 0xf3, 0xb5, 0x26, 0x96, 0x49, 0x72,
	0x0e, 0xa7, 0x72, 0xc9, 0x04, 0x46, 0x79, 0xdb, 0x90, 0x67, 0x89, 0x32, 0xd2, 0xd6, 0xe8, 0xe3,
	0x3c, 0xa1, 0x49, 0x3d, 0x0d, 0x93, 0x2e, 0x9c, 0x7e, 0x60, 0xb1, 0x8a, 0x93, 0x45, 0xc9, 0x97,
	0xee, 0xa1, 0x71, 0xef, 0x1f, 0x3f, 0x7c, 0xbf, 0xfa, 0xfc, 0x4b, 0xa8, 0x6f, 0x8e, 0x47, 0x8e,
	0xc1, 0xf6, 0xdf, 0xf7, 0x86, 0x6f, 0x27, 0x83, 0x77, 0x7e, 0x3e, 0xba, 0x93, 0x7e, 0x97, 0xea,
	0xd1, 0x3d, 0x1f, 0x03, 0xb9, 0xaf, 0xb4, 0xa9, 0xf0, 0xe9, 0x3b, 0x9f, 0x36, 0x2c, 0x4d, 0x1e,
	0x75, 0xaf, 0xfc, 0xc9, 0xb8, 0xdb, 0xf3, 0x1b, 0x15, 0x62, 0x43, 0x6d, 0xda, 0xbd, 0x1c, 0xfa,
	0x8d, 0x03, 0x5d, 0x45, 0xfd, 0xd7, 0x83, 0xeb, 0x51, 0xa3, 0x4a, 0xea, 0x50, 0x1d, 0xfb, 0x3e,
	0x6d, 0xd4, 0x3a, 0xbf, 0x59, 0xf9, 0x13, 0x5b, 0x3c, 0xd0, 0xe4, 0x47, 0x70, 0x8a, 0x37, 0x41,
	0xa3, 0xe4, 0xe9, 0xbe, 0xfc, 0x45, 0xea, 0xec, 0xd3, 0xbb, 0x70, 0x71, 0xdd, 0xc7, 0x70, 0xbc,
	0x77, 0xa3, 0xc9, 0x8b, 0x07, 0xaf, 0x7a, 0xde, 0xe8, 0x8b, 0x7f, 0x79, 0x0a, 0x2e, 0x47, 0xf0,
	0x3d, 0x17, 0x8b, 0x36, 0x4b, 0x59, 0xb8, 0xc4, 0xf6, 0x92, 0x45, 0x9c, 0xa7, 0x05, 0x47, 0x2e,
	0x59, 0x84, 0x51, 0xfe, 0xab, 0x98, 0x65, 0xf3, 0xf6, 0x02, 0x13, 0x14, 0x4c, 0x61, 0x74, 0xb9,
	0xfb, 0xcb, 0x19, 0xeb, 0xb4, 0xec, 0x5b, 0xbf, 0x58, 0xd6, 0xef, 0x96, 0xf5, 0x77, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x0c, 0xad, 0x86, 0x91, 0x8e, 0x06, 0x00, 0x00,
}
