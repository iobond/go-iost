// Code generated by protoc-gen-go. DO NOT EDIT.
// source: core/blockcache/block_cache.proto

package blockcache

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

type BcMessageType int32

const (
	BcMessageType_LinkType         BcMessageType = 0
	BcMessageType_UpdateActiveType BcMessageType = 1
)

var BcMessageType_name = map[int32]string{
	0: "LinkType",
	1: "UpdateActiveType",
}

var BcMessageType_value = map[string]int32{
	"LinkType":         0,
	"UpdateActiveType": 1,
}

func (x BcMessageType) String() string {
	return proto.EnumName(BcMessageType_name, int32(x))
}

func (BcMessageType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_96d071d2d2082821, []int{0}
}

type BcMessage struct {
	Data                 []byte        `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Type                 BcMessageType `protobuf:"varint,2,opt,name=type,proto3,enum=blockcache.BcMessageType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *BcMessage) Reset()         { *m = BcMessage{} }
func (m *BcMessage) String() string { return proto.CompactTextString(m) }
func (*BcMessage) ProtoMessage()    {}
func (*BcMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_96d071d2d2082821, []int{0}
}

func (m *BcMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BcMessage.Unmarshal(m, b)
}
func (m *BcMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BcMessage.Marshal(b, m, deterministic)
}
func (m *BcMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BcMessage.Merge(m, src)
}
func (m *BcMessage) XXX_Size() int {
	return xxx_messageInfo_BcMessage.Size(m)
}
func (m *BcMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_BcMessage.DiscardUnknown(m)
}

var xxx_messageInfo_BcMessage proto.InternalMessageInfo

func (m *BcMessage) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *BcMessage) GetType() BcMessageType {
	if m != nil {
		return m.Type
	}
	return BcMessageType_LinkType
}

type BlockCacheRaw struct {
	BlockBytes           []byte       `protobuf:"bytes,1,opt,name=blockBytes,proto3" json:"blockBytes,omitempty"`
	WitnessList          *WitnessList `protobuf:"bytes,2,opt,name=witnessList,proto3" json:"witnessList,omitempty"`
	SerialNum            int64        `protobuf:"varint,3,opt,name=serialNum,proto3" json:"serialNum,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *BlockCacheRaw) Reset()         { *m = BlockCacheRaw{} }
func (m *BlockCacheRaw) String() string { return proto.CompactTextString(m) }
func (*BlockCacheRaw) ProtoMessage()    {}
func (*BlockCacheRaw) Descriptor() ([]byte, []int) {
	return fileDescriptor_96d071d2d2082821, []int{1}
}

func (m *BlockCacheRaw) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockCacheRaw.Unmarshal(m, b)
}
func (m *BlockCacheRaw) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockCacheRaw.Marshal(b, m, deterministic)
}
func (m *BlockCacheRaw) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockCacheRaw.Merge(m, src)
}
func (m *BlockCacheRaw) XXX_Size() int {
	return xxx_messageInfo_BlockCacheRaw.Size(m)
}
func (m *BlockCacheRaw) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockCacheRaw.DiscardUnknown(m)
}

var xxx_messageInfo_BlockCacheRaw proto.InternalMessageInfo

func (m *BlockCacheRaw) GetBlockBytes() []byte {
	if m != nil {
		return m.BlockBytes
	}
	return nil
}

func (m *BlockCacheRaw) GetWitnessList() *WitnessList {
	if m != nil {
		return m.WitnessList
	}
	return nil
}

func (m *BlockCacheRaw) GetSerialNum() int64 {
	if m != nil {
		return m.SerialNum
	}
	return 0
}

type UpdateActiveRaw struct {
	BlockHashBytes       []byte       `protobuf:"bytes,1,opt,name=blockHashBytes,proto3" json:"blockHashBytes,omitempty"`
	WitnessList          *WitnessList `protobuf:"bytes,2,opt,name=witnessList,proto3" json:"witnessList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *UpdateActiveRaw) Reset()         { *m = UpdateActiveRaw{} }
func (m *UpdateActiveRaw) String() string { return proto.CompactTextString(m) }
func (*UpdateActiveRaw) ProtoMessage()    {}
func (*UpdateActiveRaw) Descriptor() ([]byte, []int) {
	return fileDescriptor_96d071d2d2082821, []int{2}
}

func (m *UpdateActiveRaw) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateActiveRaw.Unmarshal(m, b)
}
func (m *UpdateActiveRaw) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateActiveRaw.Marshal(b, m, deterministic)
}
func (m *UpdateActiveRaw) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateActiveRaw.Merge(m, src)
}
func (m *UpdateActiveRaw) XXX_Size() int {
	return xxx_messageInfo_UpdateActiveRaw.Size(m)
}
func (m *UpdateActiveRaw) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateActiveRaw.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateActiveRaw proto.InternalMessageInfo

func (m *UpdateActiveRaw) GetBlockHashBytes() []byte {
	if m != nil {
		return m.BlockHashBytes
	}
	return nil
}

func (m *UpdateActiveRaw) GetWitnessList() *WitnessList {
	if m != nil {
		return m.WitnessList
	}
	return nil
}

type WitnessList struct {
	ActiveWitnessList    []string                `protobuf:"bytes,1,rep,name=activeWitnessList,proto3" json:"activeWitnessList,omitempty"`
	PendingWitnessList   []string                `protobuf:"bytes,2,rep,name=pendingWitnessList,proto3" json:"pendingWitnessList,omitempty"`
	PendingWitnessNumber int64                   `protobuf:"varint,3,opt,name=pendingWitnessNumber,proto3" json:"pendingWitnessNumber,omitempty"`
	WitnessInfo          map[string]*WitnessInfo `protobuf:"bytes,4,rep,name=witnessInfo,proto3" json:"witnessInfo,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *WitnessList) Reset()         { *m = WitnessList{} }
func (m *WitnessList) String() string { return proto.CompactTextString(m) }
func (*WitnessList) ProtoMessage()    {}
func (*WitnessList) Descriptor() ([]byte, []int) {
	return fileDescriptor_96d071d2d2082821, []int{3}
}

func (m *WitnessList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WitnessList.Unmarshal(m, b)
}
func (m *WitnessList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WitnessList.Marshal(b, m, deterministic)
}
func (m *WitnessList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WitnessList.Merge(m, src)
}
func (m *WitnessList) XXX_Size() int {
	return xxx_messageInfo_WitnessList.Size(m)
}
func (m *WitnessList) XXX_DiscardUnknown() {
	xxx_messageInfo_WitnessList.DiscardUnknown(m)
}

var xxx_messageInfo_WitnessList proto.InternalMessageInfo

func (m *WitnessList) GetActiveWitnessList() []string {
	if m != nil {
		return m.ActiveWitnessList
	}
	return nil
}

func (m *WitnessList) GetPendingWitnessList() []string {
	if m != nil {
		return m.PendingWitnessList
	}
	return nil
}

func (m *WitnessList) GetPendingWitnessNumber() int64 {
	if m != nil {
		return m.PendingWitnessNumber
	}
	return 0
}

func (m *WitnessList) GetWitnessInfo() map[string]*WitnessInfo {
	if m != nil {
		return m.WitnessInfo
	}
	return nil
}

type WitnessInfo struct {
	Loc                  string   `protobuf:"bytes,1,opt,name=Loc,proto3" json:"Loc,omitempty"`
	URL                  string   `protobuf:"bytes,2,opt,name=URL,proto3" json:"URL,omitempty"`
	NetID                string   `protobuf:"bytes,3,opt,name=NetID,proto3" json:"NetID,omitempty"`
	Online               bool     `protobuf:"varint,4,opt,name=Online,proto3" json:"Online,omitempty"`
	Score                int64    `protobuf:"varint,5,opt,name=Score,proto3" json:"Score,omitempty"`
	Votes                int64    `protobuf:"varint,6,opt,name=Votes,proto3" json:"Votes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WitnessInfo) Reset()         { *m = WitnessInfo{} }
func (m *WitnessInfo) String() string { return proto.CompactTextString(m) }
func (*WitnessInfo) ProtoMessage()    {}
func (*WitnessInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_96d071d2d2082821, []int{4}
}

func (m *WitnessInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WitnessInfo.Unmarshal(m, b)
}
func (m *WitnessInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WitnessInfo.Marshal(b, m, deterministic)
}
func (m *WitnessInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WitnessInfo.Merge(m, src)
}
func (m *WitnessInfo) XXX_Size() int {
	return xxx_messageInfo_WitnessInfo.Size(m)
}
func (m *WitnessInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_WitnessInfo.DiscardUnknown(m)
}

var xxx_messageInfo_WitnessInfo proto.InternalMessageInfo

func (m *WitnessInfo) GetLoc() string {
	if m != nil {
		return m.Loc
	}
	return ""
}

func (m *WitnessInfo) GetURL() string {
	if m != nil {
		return m.URL
	}
	return ""
}

func (m *WitnessInfo) GetNetID() string {
	if m != nil {
		return m.NetID
	}
	return ""
}

func (m *WitnessInfo) GetOnline() bool {
	if m != nil {
		return m.Online
	}
	return false
}

func (m *WitnessInfo) GetScore() int64 {
	if m != nil {
		return m.Score
	}
	return 0
}

func (m *WitnessInfo) GetVotes() int64 {
	if m != nil {
		return m.Votes
	}
	return 0
}

func init() {
	proto.RegisterEnum("blockcache.BcMessageType", BcMessageType_name, BcMessageType_value)
	proto.RegisterType((*BcMessage)(nil), "blockcache.bcMessage")
	proto.RegisterType((*BlockCacheRaw)(nil), "blockcache.BlockCacheRaw")
	proto.RegisterType((*UpdateActiveRaw)(nil), "blockcache.UpdateActiveRaw")
	proto.RegisterType((*WitnessList)(nil), "blockcache.WitnessList")
	proto.RegisterMapType((map[string]*WitnessInfo)(nil), "blockcache.WitnessList.WitnessInfoEntry")
	proto.RegisterType((*WitnessInfo)(nil), "blockcache.WitnessInfo")
}

func init() { proto.RegisterFile("core/blockcache/block_cache.proto", fileDescriptor_96d071d2d2082821) }

var fileDescriptor_96d071d2d2082821 = []byte{
	// 447 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0x51, 0x6f, 0xd3, 0x30,
	0x10, 0xc6, 0x4d, 0x5b, 0x2d, 0xd7, 0x6d, 0x94, 0x53, 0x05, 0x01, 0x21, 0x14, 0xf2, 0x80, 0x22,
	0xc4, 0x8a, 0xd4, 0xbd, 0x00, 0x6f, 0x0c, 0x90, 0x18, 0x0a, 0x45, 0x32, 0x8c, 0x3e, 0x22, 0xd7,
	0x35, 0x5b, 0xd4, 0xce, 0x8e, 0x62, 0x77, 0x53, 0xfe, 0x01, 0x12, 0x7f, 0x8b, 0x1f, 0x86, 0xec,
	0x64, 0x8d, 0x3b, 0x06, 0x2f, 0xbc, 0x7d, 0xf7, 0x7d, 0xdf, 0xdd, 0xf9, 0x4e, 0x67, 0x78, 0xcc,
	0x55, 0x29, 0x9e, 0xcf, 0x57, 0x8a, 0x2f, 0x39, 0xe3, 0x67, 0x0d, 0xfc, 0xe6, 0xf0, 0xb8, 0x28,
	0x95, 0x51, 0x08, 0xad, 0x9a, 0x4c, 0x21, 0x9c, 0xf3, 0x8f, 0x42, 0x6b, 0x76, 0x2a, 0x10, 0xa1,
	0xbb, 0x60, 0x86, 0x45, 0x24, 0x26, 0xe9, 0x2e, 0x75, 0x18, 0x0f, 0xa0, 0x6b, 0xaa, 0x42, 0x44,
	0x9d, 0x98, 0xa4, 0xfb, 0x93, 0xfb, 0xe3, 0x36, 0x77, 0xbc, 0x49, 0xfc, 0x52, 0x15, 0x82, 0x3a,
	0x5b, 0xf2, 0x83, 0xc0, 0xde, 0x91, 0xb5, 0xbc, 0xb1, 0x16, 0xca, 0x2e, 0xf1, 0x11, 0xd4, 0xfd,
	0x8e, 0x2a, 0x23, 0x74, 0x53, 0xda, 0x63, 0xf0, 0x25, 0x0c, 0x2e, 0x73, 0x23, 0x85, 0xd6, 0x59,
	0xae, 0x8d, 0xeb, 0x33, 0x98, 0xdc, 0xf3, 0xfb, 0xcc, 0x5a, 0x99, 0xfa, 0x5e, 0x7c, 0x08, 0xa1,
	0x16, 0x65, 0xce, 0x56, 0xd3, 0xf5, 0x79, 0x14, 0xc4, 0x24, 0x0d, 0x68, 0x4b, 0x24, 0x06, 0x6e,
	0x9f, 0x14, 0x0b, 0x66, 0xc4, 0x6b, 0x6e, 0xf2, 0x0b, 0xf7, 0x96, 0x27, 0xb0, 0xef, 0xea, 0xbe,
	0x67, 0xfa, 0xcc, 0x7f, 0xcf, 0x35, 0xf6, 0x3f, 0xde, 0x94, 0xfc, 0xea, 0xc0, 0xc0, 0x13, 0xf1,
	0x19, 0xdc, 0x61, 0xae, 0xbf, 0x47, 0x46, 0x24, 0x0e, 0xd2, 0x90, 0xfe, 0x29, 0xe0, 0x18, 0xb0,
	0x10, 0x72, 0x91, 0xcb, 0xd3, 0xd9, 0x56, 0x7f, 0x6b, 0xbf, 0x41, 0xc1, 0x09, 0x8c, 0xb6, 0xd9,
	0xe9, 0xfa, 0x7c, 0x2e, 0xca, 0x66, 0x19, 0x37, 0x6a, 0xf8, 0x61, 0x33, 0xdc, 0xb1, 0xfc, 0xae,
	0xa2, 0x6e, 0x1c, 0xa4, 0x83, 0x49, 0xfa, 0x97, 0xe1, 0xae, 0xb0, 0xb5, 0xbe, 0x93, 0xa6, 0xac,
	0xa8, 0x9f, 0xfc, 0x60, 0x06, 0xc3, 0xeb, 0x06, 0x1c, 0x42, 0xb0, 0x14, 0x95, 0xdb, 0x6c, 0x48,
	0x2d, 0xc4, 0x03, 0xe8, 0x5d, 0xb0, 0xd5, 0x5a, 0xfc, 0x63, 0x91, 0x36, 0x9d, 0xd6, 0xae, 0x57,
	0x9d, 0x17, 0x24, 0xf9, 0x49, 0x36, 0x6b, 0xb4, 0x92, 0x2d, 0x9a, 0x29, 0x7e, 0x55, 0x34, 0x53,
	0xdc, 0x32, 0x27, 0x34, 0x73, 0x25, 0x43, 0x6a, 0x21, 0x8e, 0xa0, 0x37, 0x15, 0xe6, 0xf8, 0xad,
	0x9b, 0x3e, 0xa4, 0x75, 0x80, 0x77, 0xa1, 0xff, 0x49, 0xae, 0x72, 0x29, 0xa2, 0x6e, 0x4c, 0xd2,
	0x1d, 0xda, 0x44, 0xd6, 0xfd, 0xd9, 0xfe, 0x95, 0xa8, 0xe7, 0x76, 0x55, 0x07, 0x96, 0xfd, 0xaa,
	0xec, 0x61, 0xf4, 0x6b, 0xd6, 0x05, 0x4f, 0x0f, 0x61, 0x6f, 0xeb, 0xd8, 0x71, 0x17, 0x76, 0xb2,
	0x5c, 0x2e, 0x2d, 0x1e, 0xde, 0xc2, 0x11, 0x0c, 0xfd, 0x4b, 0x73, 0x2c, 0x99, 0xf7, 0xdd, 0x6f,
	0x3b, 0xfc, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xa1, 0x6b, 0x3e, 0x5b, 0x92, 0x03, 0x00, 0x00,
}
