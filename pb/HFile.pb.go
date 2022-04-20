// Code generated by protoc-gen-go. DO NOT EDIT.
// source: HFile.proto

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

type CompactionEventTracker struct {
	CompactedStoreFile   [][]byte `protobuf:"bytes,1,rep,name=compacted_store_file,json=compactedStoreFile" json:"compacted_store_file,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CompactionEventTracker) Reset()         { *m = CompactionEventTracker{} }
func (m *CompactionEventTracker) String() string { return proto.CompactTextString(m) }
func (*CompactionEventTracker) ProtoMessage()    {}
func (*CompactionEventTracker) Descriptor() ([]byte, []int) {
	return fileDescriptor_4146a58545c8feae, []int{0}
}

func (m *CompactionEventTracker) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CompactionEventTracker.Unmarshal(m, b)
}
func (m *CompactionEventTracker) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CompactionEventTracker.Marshal(b, m, deterministic)
}
func (m *CompactionEventTracker) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompactionEventTracker.Merge(m, src)
}
func (m *CompactionEventTracker) XXX_Size() int {
	return xxx_messageInfo_CompactionEventTracker.Size(m)
}
func (m *CompactionEventTracker) XXX_DiscardUnknown() {
	xxx_messageInfo_CompactionEventTracker.DiscardUnknown(m)
}

var xxx_messageInfo_CompactionEventTracker proto.InternalMessageInfo

func (m *CompactionEventTracker) GetCompactedStoreFile() [][]byte {
	if m != nil {
		return m.CompactedStoreFile
	}
	return nil
}

// Map of name/values
type FileInfoProto struct {
	MapEntry             []*BytesBytesPair `protobuf:"bytes,1,rep,name=map_entry,json=mapEntry" json:"map_entry,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *FileInfoProto) Reset()         { *m = FileInfoProto{} }
func (m *FileInfoProto) String() string { return proto.CompactTextString(m) }
func (*FileInfoProto) ProtoMessage()    {}
func (*FileInfoProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_4146a58545c8feae, []int{1}
}

func (m *FileInfoProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileInfoProto.Unmarshal(m, b)
}
func (m *FileInfoProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileInfoProto.Marshal(b, m, deterministic)
}
func (m *FileInfoProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileInfoProto.Merge(m, src)
}
func (m *FileInfoProto) XXX_Size() int {
	return xxx_messageInfo_FileInfoProto.Size(m)
}
func (m *FileInfoProto) XXX_DiscardUnknown() {
	xxx_messageInfo_FileInfoProto.DiscardUnknown(m)
}

var xxx_messageInfo_FileInfoProto proto.InternalMessageInfo

func (m *FileInfoProto) GetMapEntry() []*BytesBytesPair {
	if m != nil {
		return m.MapEntry
	}
	return nil
}

// HFile file trailer
type FileTrailerProto struct {
	FileInfoOffset            *uint64  `protobuf:"varint,1,opt,name=file_info_offset,json=fileInfoOffset" json:"file_info_offset,omitempty"`
	LoadOnOpenDataOffset      *uint64  `protobuf:"varint,2,opt,name=load_on_open_data_offset,json=loadOnOpenDataOffset" json:"load_on_open_data_offset,omitempty"`
	UncompressedDataIndexSize *uint64  `protobuf:"varint,3,opt,name=uncompressed_data_index_size,json=uncompressedDataIndexSize" json:"uncompressed_data_index_size,omitempty"`
	TotalUncompressedBytes    *uint64  `protobuf:"varint,4,opt,name=total_uncompressed_bytes,json=totalUncompressedBytes" json:"total_uncompressed_bytes,omitempty"`
	DataIndexCount            *uint32  `protobuf:"varint,5,opt,name=data_index_count,json=dataIndexCount" json:"data_index_count,omitempty"`
	MetaIndexCount            *uint32  `protobuf:"varint,6,opt,name=meta_index_count,json=metaIndexCount" json:"meta_index_count,omitempty"`
	EntryCount                *uint64  `protobuf:"varint,7,opt,name=entry_count,json=entryCount" json:"entry_count,omitempty"`
	NumDataIndexLevels        *uint32  `protobuf:"varint,8,opt,name=num_data_index_levels,json=numDataIndexLevels" json:"num_data_index_levels,omitempty"`
	FirstDataBlockOffset      *uint64  `protobuf:"varint,9,opt,name=first_data_block_offset,json=firstDataBlockOffset" json:"first_data_block_offset,omitempty"`
	LastDataBlockOffset       *uint64  `protobuf:"varint,10,opt,name=last_data_block_offset,json=lastDataBlockOffset" json:"last_data_block_offset,omitempty"`
	ComparatorClassName       *string  `protobuf:"bytes,11,opt,name=comparator_class_name,json=comparatorClassName" json:"comparator_class_name,omitempty"`
	CompressionCodec          *uint32  `protobuf:"varint,12,opt,name=compression_codec,json=compressionCodec" json:"compression_codec,omitempty"`
	EncryptionKey             []byte   `protobuf:"bytes,13,opt,name=encryption_key,json=encryptionKey" json:"encryption_key,omitempty"`
	XXX_NoUnkeyedLiteral      struct{} `json:"-"`
	XXX_unrecognized          []byte   `json:"-"`
	XXX_sizecache             int32    `json:"-"`
}

func (m *FileTrailerProto) Reset()         { *m = FileTrailerProto{} }
func (m *FileTrailerProto) String() string { return proto.CompactTextString(m) }
func (*FileTrailerProto) ProtoMessage()    {}
func (*FileTrailerProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_4146a58545c8feae, []int{2}
}

func (m *FileTrailerProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileTrailerProto.Unmarshal(m, b)
}
func (m *FileTrailerProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileTrailerProto.Marshal(b, m, deterministic)
}
func (m *FileTrailerProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileTrailerProto.Merge(m, src)
}
func (m *FileTrailerProto) XXX_Size() int {
	return xxx_messageInfo_FileTrailerProto.Size(m)
}
func (m *FileTrailerProto) XXX_DiscardUnknown() {
	xxx_messageInfo_FileTrailerProto.DiscardUnknown(m)
}

var xxx_messageInfo_FileTrailerProto proto.InternalMessageInfo

func (m *FileTrailerProto) GetFileInfoOffset() uint64 {
	if m != nil && m.FileInfoOffset != nil {
		return *m.FileInfoOffset
	}
	return 0
}

func (m *FileTrailerProto) GetLoadOnOpenDataOffset() uint64 {
	if m != nil && m.LoadOnOpenDataOffset != nil {
		return *m.LoadOnOpenDataOffset
	}
	return 0
}

func (m *FileTrailerProto) GetUncompressedDataIndexSize() uint64 {
	if m != nil && m.UncompressedDataIndexSize != nil {
		return *m.UncompressedDataIndexSize
	}
	return 0
}

func (m *FileTrailerProto) GetTotalUncompressedBytes() uint64 {
	if m != nil && m.TotalUncompressedBytes != nil {
		return *m.TotalUncompressedBytes
	}
	return 0
}

func (m *FileTrailerProto) GetDataIndexCount() uint32 {
	if m != nil && m.DataIndexCount != nil {
		return *m.DataIndexCount
	}
	return 0
}

func (m *FileTrailerProto) GetMetaIndexCount() uint32 {
	if m != nil && m.MetaIndexCount != nil {
		return *m.MetaIndexCount
	}
	return 0
}

func (m *FileTrailerProto) GetEntryCount() uint64 {
	if m != nil && m.EntryCount != nil {
		return *m.EntryCount
	}
	return 0
}

func (m *FileTrailerProto) GetNumDataIndexLevels() uint32 {
	if m != nil && m.NumDataIndexLevels != nil {
		return *m.NumDataIndexLevels
	}
	return 0
}

func (m *FileTrailerProto) GetFirstDataBlockOffset() uint64 {
	if m != nil && m.FirstDataBlockOffset != nil {
		return *m.FirstDataBlockOffset
	}
	return 0
}

func (m *FileTrailerProto) GetLastDataBlockOffset() uint64 {
	if m != nil && m.LastDataBlockOffset != nil {
		return *m.LastDataBlockOffset
	}
	return 0
}

func (m *FileTrailerProto) GetComparatorClassName() string {
	if m != nil && m.ComparatorClassName != nil {
		return *m.ComparatorClassName
	}
	return ""
}

func (m *FileTrailerProto) GetCompressionCodec() uint32 {
	if m != nil && m.CompressionCodec != nil {
		return *m.CompressionCodec
	}
	return 0
}

func (m *FileTrailerProto) GetEncryptionKey() []byte {
	if m != nil {
		return m.EncryptionKey
	}
	return nil
}

func init() {
	proto.RegisterType((*CompactionEventTracker)(nil), "hbase.pb.CompactionEventTracker")
	proto.RegisterType((*FileInfoProto)(nil), "hbase.pb.FileInfoProto")
	proto.RegisterType((*FileTrailerProto)(nil), "hbase.pb.FileTrailerProto")
}

func init() { proto.RegisterFile("HFile.proto", fileDescriptor_4146a58545c8feae) }

var fileDescriptor_4146a58545c8feae = []byte{
	// 535 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x53, 0x6d, 0x6b, 0xdb, 0x3e,
	0x10, 0xc7, 0xff, 0xb6, 0xff, 0x35, 0xca, 0x03, 0x99, 0xdb, 0x66, 0xda, 0x18, 0x2c, 0x04, 0x06,
	0x86, 0x81, 0x59, 0x3b, 0x3a, 0xf6, 0x6e, 0x90, 0xac, 0x25, 0xdd, 0xc6, 0x52, 0xdc, 0xee, 0xb5,
	0xb8, 0xd8, 0xe7, 0x46, 0xc4, 0x96, 0x8c, 0xa4, 0x94, 0xa5, 0x9f, 0x60, 0xb0, 0x2f, 0xb1, 0x8f,
	0x3a, 0x4e, 0x4e, 0x5a, 0xaf, 0xec, 0x8d, 0x2d, 0x7e, 0x0f, 0x77, 0xfa, 0x1d, 0x27, 0xd6, 0x9e,
	0x9e, 0xcb, 0x02, 0xe3, 0xca, 0x68, 0xa7, 0xc3, 0xfd, 0xc5, 0x1c, 0x2c, 0xc6, 0xd5, 0xfc, 0x45,
	0x7b, 0x3a, 0xf6, 0x27, 0x82, 0x47, 0x9f, 0xd9, 0x60, 0xa2, 0xcb, 0x0a, 0x52, 0x27, 0xb5, 0x3a,
	0xbb, 0x45, 0xe5, 0xae, 0x0d, 0xa4, 0x4b, 0x34, 0xe1, 0x5b, 0x76, 0x98, 0xd6, 0x0c, 0x66, 0xc2,
	0x3a, 0x6d, 0x50, 0xe4, 0xb2, 0x40, 0x1e, 0x0c, 0x77, 0xa2, 0x4e, 0x12, 0xde, 0x73, 0x57, 0x44,
	0x51, 0xa3, 0xd1, 0x39, 0xeb, 0xd2, 0xff, 0x42, 0xe5, 0xfa, 0xd2, 0xf7, 0x3c, 0x65, 0xad, 0x12,
	0x2a, 0x81, 0xca, 0x99, 0xb5, 0xf7, 0xb5, 0x4f, 0x78, 0xbc, 0xbd, 0x47, 0x3c, 0x5e, 0x3b, 0xb4,
	0xfe, 0x73, 0x09, 0xd2, 0x24, 0xfb, 0x25, 0x54, 0x67, 0xa4, 0x1c, 0xfd, 0xda, 0x63, 0x7d, 0x2a,
	0x74, 0x6d, 0x40, 0x16, 0x68, 0xea, 0x5a, 0x11, 0xeb, 0x53, 0x7b, 0x21, 0x55, 0xae, 0x85, 0xce,
	0x73, 0x8b, 0x8e, 0x07, 0xc3, 0x20, 0xda, 0x4d, 0x7a, 0xf9, 0xa6, 0xe9, 0xcc, 0xa3, 0xe1, 0x7b,
	0xc6, 0x0b, 0x0d, 0x99, 0xd0, 0x4a, 0xe8, 0x0a, 0x95, 0xc8, 0xc0, 0xc1, 0xd6, 0xf1, 0x9f, 0x77,
	0x1c, 0x12, 0x3f, 0x53, 0xb3, 0x0a, 0xd5, 0x27, 0x70, 0xb0, 0xf1, 0x7d, 0x64, 0x2f, 0x57, 0x8a,
	0x62, 0x19, 0xb4, 0x16, 0xb3, 0xda, 0x27, 0x55, 0x86, 0x3f, 0x84, 0x95, 0x77, 0xc8, 0x77, 0xbc,
	0xf7, 0x79, 0x53, 0x43, 0xee, 0x0b, 0x52, 0x5c, 0xc9, 0x3b, 0x0c, 0x3f, 0x30, 0xee, 0xb4, 0x83,
	0x42, 0xfc, 0x55, 0x66, 0x4e, 0x09, 0xf9, 0xae, 0x37, 0x0f, 0x3c, 0xff, 0xbd, 0x41, 0xfb, 0xfc,
	0x14, 0xae, 0xd1, 0x2d, 0xd5, 0x2b, 0xe5, 0xf8, 0xde, 0x30, 0x88, 0xba, 0x49, 0x2f, 0xdb, 0xb6,
	0x98, 0x10, 0x4a, 0xca, 0x12, 0x1f, 0x29, 0xff, 0xaf, 0x95, 0x84, 0x37, 0x94, 0xaf, 0x58, 0xdb,
	0x0f, 0x7e, 0x23, 0x7a, 0xe2, 0x2f, 0xc0, 0x3c, 0x54, 0x0b, 0x8e, 0xd9, 0x91, 0x5a, 0x95, 0xcd,
	0x98, 0x05, 0xde, 0x62, 0x61, 0xf9, 0xbe, 0xaf, 0x17, 0xaa, 0x55, 0x79, 0x9f, 0xef, 0xab, 0x67,
	0xc2, 0x53, 0xf6, 0x2c, 0x97, 0xc6, 0xba, 0xda, 0x34, 0x2f, 0x74, 0xba, 0xdc, 0x4e, 0xb6, 0x55,
	0x4f, 0xd6, 0xd3, 0x64, 0x1b, 0x13, 0xb9, 0x99, 0xec, 0x3b, 0x36, 0x28, 0xe0, 0x9f, 0x2e, 0xe6,
	0x5d, 0x07, 0xc4, 0x3e, 0x36, 0x9d, 0xb0, 0x23, 0xbf, 0x63, 0x06, 0x9c, 0x36, 0x22, 0x2d, 0xc0,
	0x5a, 0xa1, 0xa0, 0x44, 0xde, 0x1e, 0x06, 0x51, 0x2b, 0x39, 0x78, 0x20, 0x27, 0xc4, 0x7d, 0x83,
	0x12, 0xc3, 0x37, 0xec, 0xe9, 0x76, 0xb4, 0x52, 0x2b, 0x91, 0xea, 0x0c, 0x53, 0xde, 0xf1, 0x71,
	0xfa, 0x0d, 0x62, 0x42, 0x78, 0xf8, 0x9a, 0xf5, 0x50, 0xa5, 0x66, 0x5d, 0xd1, 0xea, 0x8b, 0x25,
	0xae, 0x79, 0x77, 0x18, 0x44, 0x9d, 0xa4, 0xfb, 0x80, 0x7e, 0xc1, 0xf5, 0x78, 0xca, 0x8e, 0xb5,
	0xb9, 0x89, 0xa1, 0x82, 0x74, 0x81, 0xf1, 0x02, 0x32, 0xad, 0xab, 0xcd, 0x12, 0xdb, 0x05, 0x64,
	0x98, 0xd5, 0x2f, 0x69, 0xbe, 0xca, 0xe3, 0x1b, 0x54, 0x68, 0xc0, 0x61, 0x36, 0xae, 0x9f, 0x9e,
	0xdf, 0x5c, 0x3b, 0x0d, 0x7e, 0x06, 0xc1, 0xef, 0x20, 0xf8, 0x13, 0x00, 0x00, 0xff, 0xff, 0x3f,
	0x1a, 0x37, 0x65, 0x90, 0x03, 0x00, 0x00,
}
