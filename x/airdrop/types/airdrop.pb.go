// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: quicksilver/airdrop/v1/airdrop.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Action is used as an enum to denote specific actions or tasks.
type Action int32

const (
	// Undefined action (per protobuf spec)
	ActionUndefined Action = 0
	// Initial claim action
	ActionInitialClaim Action = 1
	// Deposit tier 1 (e.g. > 5% of base_value)
	ActionDepositT1 Action = 2
	// Deposit tier 2 (e.g. > 10% of base_value)
	ActionDepositT2 Action = 3
	// Deposit tier 3 (e.g. > 15% of base_value)
	ActionDepositT3 Action = 4
	// Deposit tier 4 (e.g. > 22% of base_value)
	ActionDepositT4 Action = 5
	// Deposit tier 5 (e.g. > 30% of base_value)
	ActionDepositT5 Action = 6
	// Active QCK delegation
	ActionStakeQCK Action = 7
	// Intent is set
	ActionSignalIntent Action = 8
	// Cast governance vote on QS
	ActionQSGov Action = 9
	// Governance By Proxy (GbP): cast vote on remote zone
	ActionGbP Action = 10
	// Provide liquidity on Osmosis
	ActionOsmosis Action = 11
)

var Action_name = map[int32]string{
	0:  "ActionUndefined",
	1:  "ActionInitialClaim",
	2:  "ActionDepositT1",
	3:  "ActionDepositT2",
	4:  "ActionDepositT3",
	5:  "ActionDepositT4",
	6:  "ActionDepositT5",
	7:  "ActionStakeQCK",
	8:  "ActionSignalIntent",
	9:  "ActionQSGov",
	10: "ActionGbP",
	11: "ActionOsmosis",
}

var Action_value = map[string]int32{
	"ActionUndefined":    0,
	"ActionInitialClaim": 1,
	"ActionDepositT1":    2,
	"ActionDepositT2":    3,
	"ActionDepositT3":    4,
	"ActionDepositT4":    5,
	"ActionDepositT5":    6,
	"ActionStakeQCK":     7,
	"ActionSignalIntent": 8,
	"ActionQSGov":        9,
	"ActionGbP":          10,
	"ActionOsmosis":      11,
}

func (x Action) String() string {
	return proto.EnumName(Action_name, int32(x))
}

func (Action) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e3f0590c06bbb467, []int{0}
}

// Status is used as an enum to denote zone status.
type Status int32

const (
	StatusUndefined Status = 0
	StatusActive    Status = 1
	StatusFuture    Status = 2
	StatusExpired   Status = 3
)

var Status_name = map[int32]string{
	0: "StatusUndefined",
	1: "StatusActive",
	2: "StatusFuture",
	3: "StatusExpired",
}

var Status_value = map[string]int32{
	"StatusUndefined": 0,
	"StatusActive":    1,
	"StatusFuture":    2,
	"StatusExpired":   3,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}

func (Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e3f0590c06bbb467, []int{1}
}

// ZoneDrop represents an airdrop for a specific zone.
type ZoneDrop struct {
	ChainId     string                                   `protobuf:"bytes,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	StartTime   time.Time                                `protobuf:"bytes,2,opt,name=start_time,json=startTime,proto3,stdtime" json:"start_time" yaml:"start_time"`
	Duration    time.Duration                            `protobuf:"bytes,3,opt,name=duration,proto3,stdduration" json:"duration,omitempty" yaml:"duration"`
	Decay       time.Duration                            `protobuf:"bytes,4,opt,name=decay,proto3,stdduration" json:"decay,omitempty" yaml:"decay"`
	Allocation  uint64                                   `protobuf:"varint,5,opt,name=allocation,proto3" json:"allocation,omitempty"`
	Actions     []github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,6,rep,name=actions,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"actions"`
	IsConcluded bool                                     `protobuf:"varint,7,opt,name=is_concluded,json=isConcluded,proto3" json:"is_concluded,omitempty"`
}

func (m *ZoneDrop) Reset()         { *m = ZoneDrop{} }
func (m *ZoneDrop) String() string { return proto.CompactTextString(m) }
func (*ZoneDrop) ProtoMessage()    {}
func (*ZoneDrop) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3f0590c06bbb467, []int{0}
}
func (m *ZoneDrop) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ZoneDrop) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ZoneDrop.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ZoneDrop) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ZoneDrop.Merge(m, src)
}
func (m *ZoneDrop) XXX_Size() int {
	return m.Size()
}
func (m *ZoneDrop) XXX_DiscardUnknown() {
	xxx_messageInfo_ZoneDrop.DiscardUnknown(m)
}

var xxx_messageInfo_ZoneDrop proto.InternalMessageInfo

// ClaimRecord represents a users' claim (including completed claims) for a
// given zone.
type ClaimRecord struct {
	ChainId string `protobuf:"bytes,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	// Protobuf3 does not allow enum as map key
	ActionsCompleted map[int32]*CompletedAction `protobuf:"bytes,3,rep,name=actions_completed,json=actionsCompleted,proto3" json:"actions_completed,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	MaxAllocation    uint64                     `protobuf:"varint,4,opt,name=max_allocation,json=maxAllocation,proto3" json:"max_allocation,omitempty"`
	BaseValue        uint64                     `protobuf:"varint,5,opt,name=base_value,json=baseValue,proto3" json:"base_value,omitempty"`
}

func (m *ClaimRecord) Reset()         { *m = ClaimRecord{} }
func (m *ClaimRecord) String() string { return proto.CompactTextString(m) }
func (*ClaimRecord) ProtoMessage()    {}
func (*ClaimRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3f0590c06bbb467, []int{1}
}
func (m *ClaimRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClaimRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClaimRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClaimRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClaimRecord.Merge(m, src)
}
func (m *ClaimRecord) XXX_Size() int {
	return m.Size()
}
func (m *ClaimRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_ClaimRecord.DiscardUnknown(m)
}

var xxx_messageInfo_ClaimRecord proto.InternalMessageInfo

// CompletedAction represents a claim action completed by the user.
type CompletedAction struct {
	CompleteTime time.Time `protobuf:"bytes,1,opt,name=complete_time,json=completeTime,proto3,stdtime" json:"complete_time" yaml:"complete_time"`
	ClaimAmount  uint64    `protobuf:"varint,2,opt,name=claim_amount,json=claimAmount,proto3" json:"claim_amount,omitempty"`
}

func (m *CompletedAction) Reset()         { *m = CompletedAction{} }
func (m *CompletedAction) String() string { return proto.CompactTextString(m) }
func (*CompletedAction) ProtoMessage()    {}
func (*CompletedAction) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3f0590c06bbb467, []int{2}
}
func (m *CompletedAction) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CompletedAction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CompletedAction.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CompletedAction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompletedAction.Merge(m, src)
}
func (m *CompletedAction) XXX_Size() int {
	return m.Size()
}
func (m *CompletedAction) XXX_DiscardUnknown() {
	xxx_messageInfo_CompletedAction.DiscardUnknown(m)
}

var xxx_messageInfo_CompletedAction proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("quicksilver.airdrop.v1.Action", Action_name, Action_value)
	proto.RegisterEnum("quicksilver.airdrop.v1.Status", Status_name, Status_value)
	proto.RegisterType((*ZoneDrop)(nil), "quicksilver.airdrop.v1.ZoneDrop")
	proto.RegisterType((*ClaimRecord)(nil), "quicksilver.airdrop.v1.ClaimRecord")
	proto.RegisterMapType((map[int32]*CompletedAction)(nil), "quicksilver.airdrop.v1.ClaimRecord.ActionsCompletedEntry")
	proto.RegisterType((*CompletedAction)(nil), "quicksilver.airdrop.v1.CompletedAction")
}

func init() {
	proto.RegisterFile("quicksilver/airdrop/v1/airdrop.proto", fileDescriptor_e3f0590c06bbb467)
}

var fileDescriptor_e3f0590c06bbb467 = []byte{
	// 857 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x95, 0xcf, 0x6f, 0x1b, 0x45,
	0x14, 0xc7, 0xbd, 0xfe, 0xed, 0x67, 0xa7, 0x9e, 0x0c, 0xa5, 0x72, 0x2c, 0x75, 0xed, 0x5a, 0xfc,
	0xb0, 0x0a, 0x59, 0x2b, 0x29, 0x20, 0x88, 0xe0, 0x90, 0xc4, 0xa5, 0x8a, 0x38, 0xd0, 0x6e, 0x4a,
	0x85, 0x22, 0x24, 0x6b, 0xbc, 0x3b, 0x71, 0x46, 0xd9, 0xdd, 0x59, 0x76, 0x67, 0x2d, 0xfb, 0xcc,
	0xa5, 0xc7, 0x1e, 0xb9, 0x20, 0x21, 0xf1, 0x2f, 0xf0, 0x47, 0xf4, 0x58, 0x71, 0x42, 0x1c, 0x0c,
	0x4a, 0x38, 0x20, 0x8e, 0xfd, 0x0b, 0xd0, 0xcc, 0xec, 0xd6, 0x6e, 0x5c, 0xda, 0xd3, 0xbe, 0xf9,
	0xcc, 0x7b, 0xdf, 0xf7, 0xe6, 0xcd, 0x3c, 0x2d, 0xbc, 0xf3, 0x7d, 0xc2, 0x9c, 0xf3, 0x98, 0x79,
	0x53, 0x1a, 0x0d, 0x08, 0x8b, 0xdc, 0x88, 0x87, 0x83, 0xe9, 0x4e, 0x66, 0x5a, 0x61, 0xc4, 0x05,
	0xc7, 0x37, 0x56, 0xbc, 0xac, 0x6c, 0x6b, 0xba, 0xd3, 0xde, 0x72, 0x78, 0xec, 0xf3, 0x78, 0xa4,
	0xbc, 0x06, 0x7a, 0xa1, 0x43, 0xda, 0xd7, 0x27, 0x7c, 0xc2, 0x35, 0x97, 0x56, 0x4a, 0xcd, 0x09,
	0xe7, 0x13, 0x8f, 0x0e, 0xd4, 0x6a, 0x9c, 0x9c, 0x0e, 0xdc, 0x24, 0x22, 0x82, 0xf1, 0x20, 0xdd,
	0xef, 0x5c, 0xdd, 0x17, 0xcc, 0xa7, 0xb1, 0x20, 0x7e, 0x5a, 0x49, 0xef, 0x9f, 0x02, 0x54, 0x4f,
	0x78, 0x40, 0x87, 0x11, 0x0f, 0xf1, 0x16, 0x54, 0x9d, 0x33, 0xc2, 0x82, 0x11, 0x73, 0x5b, 0x46,
	0xd7, 0xe8, 0xd7, 0xec, 0x8a, 0x5a, 0x1f, 0xb9, 0xf8, 0x5b, 0x80, 0x58, 0x90, 0x48, 0x8c, 0xa4,
	0x40, 0x2b, 0xdf, 0x35, 0xfa, 0xf5, 0xdd, 0xb6, 0xa5, 0xd5, 0xad, 0x4c, 0xdd, 0x7a, 0x98, 0xa9,
	0x1f, 0xdc, 0x7c, 0xba, 0xe8, 0xe4, 0x9e, 0x2f, 0x3a, 0x9b, 0x73, 0xe2, 0x7b, 0x7b, 0xbd, 0x65,
	0x6c, 0xef, 0xc9, 0x9f, 0x1d, 0xc3, 0xae, 0x29, 0x20, 0xdd, 0xf1, 0x19, 0x54, 0xb3, 0xa2, 0x5b,
	0x05, 0xa5, 0xbb, 0xb5, 0xa6, 0x3b, 0x4c, 0x1d, 0x0e, 0x76, 0xa4, 0xec, 0xbf, 0x8b, 0x0e, 0xce,
	0x42, 0x3e, 0xe4, 0x3e, 0x13, 0xd4, 0x0f, 0xc5, 0xfc, 0xf9, 0xa2, 0xd3, 0xd4, 0xc9, 0xb2, 0xbd,
	0xde, 0x8f, 0x32, 0xd5, 0x0b, 0x75, 0xfc, 0x1d, 0x94, 0x5c, 0xea, 0x90, 0x79, 0xab, 0xf8, 0xa6,
	0x34, 0x1f, 0xa4, 0x69, 0x9a, 0xca, 0xff, 0xa5, 0x1c, 0x8d, 0x34, 0x87, 0xdc, 0xd0, 0x09, 0xb4,
	0x28, 0x36, 0x01, 0x88, 0xe7, 0x71, 0x47, 0x9f, 0xa4, 0xd4, 0x35, 0xfa, 0x45, 0x7b, 0x85, 0xe0,
	0x47, 0x50, 0x21, 0x8e, 0xb4, 0xe2, 0x56, 0xb9, 0x5b, 0xe8, 0xd7, 0x0e, 0x3e, 0x97, 0x49, 0xfe,
	0x58, 0x74, 0xde, 0x9b, 0x30, 0x71, 0x96, 0x8c, 0x2d, 0x87, 0xfb, 0xe9, 0x95, 0xa7, 0x9f, 0xed,
	0xd8, 0x3d, 0x1f, 0x88, 0x79, 0x48, 0x63, 0x6b, 0x48, 0x9d, 0xdf, 0x7e, 0xdd, 0x86, 0xf4, 0x45,
	0x0c, 0xa9, 0x63, 0x67, 0x62, 0xf8, 0x16, 0x34, 0x58, 0x3c, 0x72, 0x78, 0xe0, 0x78, 0x89, 0x4b,
	0xdd, 0x56, 0xa5, 0x6b, 0xf4, 0xab, 0x76, 0x9d, 0xc5, 0x87, 0x19, 0xda, 0x2b, 0x3e, 0xfe, 0xb9,
	0x93, 0xeb, 0xfd, 0x9d, 0x87, 0xfa, 0xa1, 0x47, 0x98, 0x6f, 0x53, 0x87, 0x47, 0xee, 0xeb, 0x6e,
	0xbb, 0x05, 0x15, 0xe2, 0xba, 0x11, 0x8d, 0x63, 0x75, 0xd5, 0x35, 0x3b, 0x5b, 0xe2, 0x53, 0xd8,
	0x4c, 0x13, 0x8f, 0x1c, 0xee, 0x87, 0x1e, 0x15, 0xd4, 0x6d, 0x15, 0xba, 0x85, 0x7e, 0x7d, 0xf7,
	0x33, 0xeb, 0xd5, 0xaf, 0xda, 0x5a, 0x49, 0x6a, 0xed, 0xeb, 0xe0, 0xc3, 0x2c, 0xf6, 0x6e, 0x20,
	0xa2, 0xb9, 0x8d, 0xc8, 0x15, 0x8c, 0xdf, 0x85, 0x6b, 0x3e, 0x99, 0x8d, 0x56, 0x3a, 0x5a, 0x54,
	0x1d, 0xdd, 0xf0, 0xc9, 0x6c, 0x7f, 0xd9, 0xd4, 0x9b, 0x00, 0x63, 0x12, 0xd3, 0xd1, 0x94, 0x78,
	0x09, 0x4d, 0x9b, 0x5e, 0x93, 0xe4, 0x91, 0x04, 0x6d, 0x0f, 0xde, 0x7e, 0x65, 0x42, 0x8c, 0xa0,
	0x70, 0x4e, 0xe7, 0xea, 0xd8, 0x25, 0x5b, 0x9a, 0xf8, 0x0b, 0x28, 0x69, 0x11, 0xfd, 0xb6, 0xdf,
	0xff, 0xdf, 0xc3, 0x64, 0x42, 0x5a, 0xd8, 0xd6, 0x51, 0x7b, 0xf9, 0x4f, 0x8d, 0xb4, 0xcd, 0x3f,
	0x19, 0xd0, 0xbc, 0xe2, 0x84, 0x09, 0x6c, 0x64, 0xdd, 0xd2, 0x03, 0x64, 0xbc, 0x71, 0x80, 0xba,
	0xe9, 0x00, 0x5d, 0xd7, 0xef, 0xed, 0xa5, 0x70, 0x3d, 0x43, 0x8d, 0x8c, 0xa9, 0x31, 0xba, 0x05,
	0x0d, 0x47, 0xf6, 0x79, 0x44, 0x7c, 0x9e, 0x04, 0x42, 0x1d, 0xa3, 0x68, 0xd7, 0x15, 0xdb, 0x57,
	0x48, 0xd7, 0x77, 0xfb, 0x87, 0x3c, 0x94, 0xd3, 0xb2, 0xde, 0x82, 0xa6, 0xb6, 0xbe, 0x09, 0x5c,
	0x7a, 0xca, 0x02, 0xea, 0xa2, 0x1c, 0xbe, 0x01, 0x58, 0xc3, 0xa3, 0x80, 0x09, 0x46, 0x3c, 0x75,
	0x7b, 0xc8, 0x58, 0x3a, 0x0f, 0x69, 0xc8, 0x63, 0x26, 0x1e, 0xee, 0xa0, 0xfc, 0x3a, 0xdc, 0x45,
	0x85, 0x75, 0x78, 0x07, 0x15, 0xd7, 0xe1, 0x47, 0xa8, 0xb4, 0x0e, 0x3f, 0x46, 0x65, 0x8c, 0xe1,
	0x9a, 0x86, 0xc7, 0x82, 0x9c, 0xd3, 0x07, 0x87, 0x5f, 0xa1, 0xca, 0xb2, 0xa8, 0x63, 0x36, 0x09,
	0x88, 0x77, 0x14, 0x08, 0x1a, 0x08, 0x54, 0xc5, 0x4d, 0xa8, 0x6b, 0xfe, 0xe0, 0xf8, 0x1e, 0x9f,
	0xa2, 0x1a, 0xde, 0x80, 0x9a, 0x06, 0xf7, 0xc6, 0xf7, 0x11, 0xe0, 0x4d, 0xd8, 0xd0, 0xcb, 0xaf,
	0xe5, 0xe0, 0xb0, 0x18, 0xd5, 0xdb, 0xc5, 0xc7, 0xbf, 0x98, 0xb9, 0xdb, 0x27, 0x50, 0x3e, 0x16,
	0x44, 0x24, 0xb1, 0xac, 0x41, 0x5b, 0xab, 0x4d, 0x40, 0xd0, 0xd0, 0x50, 0x46, 0x4f, 0x29, 0x32,
	0x96, 0xe4, 0xcb, 0x44, 0x24, 0x11, 0x45, 0x79, 0xa9, 0xad, 0xc9, 0xdd, 0x59, 0xc8, 0x22, 0xea,
	0xa2, 0x82, 0xd6, 0x3e, 0xb8, 0xff, 0xf4, 0xc2, 0x34, 0x9e, 0x5d, 0x98, 0xc6, 0x5f, 0x17, 0xa6,
	0xf1, 0xe4, 0xd2, 0xcc, 0x3d, 0xbb, 0x34, 0x73, 0xbf, 0x5f, 0x9a, 0xb9, 0x93, 0x4f, 0x56, 0x46,
	0x9d, 0x05, 0x13, 0x1a, 0x24, 0x4c, 0xcc, 0xb7, 0xc7, 0x09, 0xf3, 0xdc, 0xc1, 0xea, 0x8f, 0x63,
	0xf6, 0xe2, 0xd7, 0xa1, 0xc6, 0x7f, 0x5c, 0x56, 0x0f, 0xe4, 0xce, 0x7f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x1f, 0x20, 0x1e, 0x88, 0x5e, 0x06, 0x00, 0x00,
}

func (m *ZoneDrop) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ZoneDrop) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ZoneDrop) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.IsConcluded {
		i--
		if m.IsConcluded {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x38
	}
	if len(m.Actions) > 0 {
		for iNdEx := len(m.Actions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size := m.Actions[iNdEx].Size()
				i -= size
				if _, err := m.Actions[iNdEx].MarshalTo(dAtA[i:]); err != nil {
					return 0, err
				}
				i = encodeVarintAirdrop(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if m.Allocation != 0 {
		i = encodeVarintAirdrop(dAtA, i, uint64(m.Allocation))
		i--
		dAtA[i] = 0x28
	}
	n1, err1 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.Decay, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.Decay):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintAirdrop(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x22
	n2, err2 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.Duration, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.Duration):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintAirdrop(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x1a
	n3, err3 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.StartTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTime):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintAirdrop(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0x12
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintAirdrop(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ClaimRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClaimRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClaimRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BaseValue != 0 {
		i = encodeVarintAirdrop(dAtA, i, uint64(m.BaseValue))
		i--
		dAtA[i] = 0x28
	}
	if m.MaxAllocation != 0 {
		i = encodeVarintAirdrop(dAtA, i, uint64(m.MaxAllocation))
		i--
		dAtA[i] = 0x20
	}
	if len(m.ActionsCompleted) > 0 {
		for k := range m.ActionsCompleted {
			v := m.ActionsCompleted[k]
			baseI := i
			if v != nil {
				{
					size, err := v.MarshalToSizedBuffer(dAtA[:i])
					if err != nil {
						return 0, err
					}
					i -= size
					i = encodeVarintAirdrop(dAtA, i, uint64(size))
				}
				i--
				dAtA[i] = 0x12
			}
			i = encodeVarintAirdrop(dAtA, i, uint64(k))
			i--
			dAtA[i] = 0x8
			i = encodeVarintAirdrop(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintAirdrop(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintAirdrop(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CompletedAction) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CompletedAction) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CompletedAction) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ClaimAmount != 0 {
		i = encodeVarintAirdrop(dAtA, i, uint64(m.ClaimAmount))
		i--
		dAtA[i] = 0x10
	}
	n5, err5 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.CompleteTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.CompleteTime):])
	if err5 != nil {
		return 0, err5
	}
	i -= n5
	i = encodeVarintAirdrop(dAtA, i, uint64(n5))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintAirdrop(dAtA []byte, offset int, v uint64) int {
	offset -= sovAirdrop(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ZoneDrop) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovAirdrop(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTime)
	n += 1 + l + sovAirdrop(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.Duration)
	n += 1 + l + sovAirdrop(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.Decay)
	n += 1 + l + sovAirdrop(uint64(l))
	if m.Allocation != 0 {
		n += 1 + sovAirdrop(uint64(m.Allocation))
	}
	if len(m.Actions) > 0 {
		for _, e := range m.Actions {
			l = e.Size()
			n += 1 + l + sovAirdrop(uint64(l))
		}
	}
	if m.IsConcluded {
		n += 2
	}
	return n
}

func (m *ClaimRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovAirdrop(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovAirdrop(uint64(l))
	}
	if len(m.ActionsCompleted) > 0 {
		for k, v := range m.ActionsCompleted {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
				l += 1 + sovAirdrop(uint64(l))
			}
			mapEntrySize := 1 + sovAirdrop(uint64(k)) + l
			n += mapEntrySize + 1 + sovAirdrop(uint64(mapEntrySize))
		}
	}
	if m.MaxAllocation != 0 {
		n += 1 + sovAirdrop(uint64(m.MaxAllocation))
	}
	if m.BaseValue != 0 {
		n += 1 + sovAirdrop(uint64(m.BaseValue))
	}
	return n
}

func (m *CompletedAction) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.CompleteTime)
	n += 1 + l + sovAirdrop(uint64(l))
	if m.ClaimAmount != 0 {
		n += 1 + sovAirdrop(uint64(m.ClaimAmount))
	}
	return n
}

func sovAirdrop(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAirdrop(x uint64) (n int) {
	return sovAirdrop(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ZoneDrop) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAirdrop
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ZoneDrop: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ZoneDrop: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAirdrop
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAirdrop
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAirdrop
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAirdrop
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthAirdrop
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAirdrop
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.StartTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAirdrop
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthAirdrop
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAirdrop
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.Duration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Decay", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAirdrop
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthAirdrop
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAirdrop
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.Decay, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Allocation", wireType)
			}
			m.Allocation = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAirdrop
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Allocation |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Actions", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAirdrop
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAirdrop
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAirdrop
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cosmos_cosmos_sdk_types.Dec
			m.Actions = append(m.Actions, v)
			if err := m.Actions[len(m.Actions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsConcluded", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAirdrop
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsConcluded = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipAirdrop(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAirdrop
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ClaimRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAirdrop
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ClaimRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClaimRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAirdrop
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAirdrop
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAirdrop
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAirdrop
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAirdrop
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAirdrop
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ActionsCompleted", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAirdrop
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthAirdrop
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAirdrop
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ActionsCompleted == nil {
				m.ActionsCompleted = make(map[int32]*CompletedAction)
			}
			var mapkey int32
			var mapvalue *CompletedAction
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowAirdrop
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowAirdrop
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapkey |= int32(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowAirdrop
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthAirdrop
					}
					postmsgIndex := iNdEx + mapmsglen
					if postmsgIndex < 0 {
						return ErrInvalidLengthAirdrop
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &CompletedAction{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipAirdrop(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthAirdrop
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.ActionsCompleted[mapkey] = mapvalue
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxAllocation", wireType)
			}
			m.MaxAllocation = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAirdrop
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxAllocation |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseValue", wireType)
			}
			m.BaseValue = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAirdrop
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BaseValue |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAirdrop(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAirdrop
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CompletedAction) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAirdrop
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CompletedAction: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CompletedAction: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CompleteTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAirdrop
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthAirdrop
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAirdrop
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.CompleteTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimAmount", wireType)
			}
			m.ClaimAmount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAirdrop
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ClaimAmount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAirdrop(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAirdrop
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipAirdrop(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAirdrop
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowAirdrop
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowAirdrop
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthAirdrop
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAirdrop
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAirdrop
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAirdrop        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAirdrop          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAirdrop = fmt.Errorf("proto: unexpected end of group")
)
