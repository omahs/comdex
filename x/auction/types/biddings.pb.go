// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: comdex/auction/v1beta1/biddings.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
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

type Biddings struct {
	Id                  uint64                                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" yaml:"id"`
	AuctionId           uint64                                  `protobuf:"varint,2,opt,name=auction_id,json=auctionId,proto3" json:"auction_id,omitempty" yaml:"auction_id"`
	AuctionStatus       string                                  `protobuf:"bytes,3,opt,name=auction_status,json=auctionStatus,proto3" json:"auction_status,omitempty" yaml:"auction_status"`
	AuctionedCollateral github_com_cosmos_cosmos_sdk_types.Coin `protobuf:"bytes,4,opt,name=auctioned_collateral,json=auctionedCollateral,proto3,casttype=github.com/cosmos/cosmos-sdk/types.Coin" json:"auctioned_collateral" yaml:"auctioned_collateral"`
	Bidder              string                                  `protobuf:"bytes,5,opt,name=bidder,proto3" json:"bidder,omitempty" yaml:"bidder"`
	Bid                 github_com_cosmos_cosmos_sdk_types.Coin `protobuf:"bytes,6,opt,name=bid,proto3,casttype=github.com/cosmos/cosmos-sdk/types.Coin" json:"bid" yaml:"bid"`
	BiddingTimestamp    time.Time                               `protobuf:"bytes,7,opt,name=bidding_timestamp,json=biddingTimestamp,proto3,stdtime" json:"bidding_timestamp" yaml:"bidding_timestamp"`
	BiddingStatus       string                                  `protobuf:"bytes,8,opt,name=bidding_status,json=biddingStatus,proto3" json:"bidding_status,omitempty" yaml:"bidding_status"`
}

func (m *Biddings) Reset()         { *m = Biddings{} }
func (m *Biddings) String() string { return proto.CompactTextString(m) }
func (*Biddings) ProtoMessage()    {}
func (*Biddings) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5a3f4b8597bafd2, []int{0}
}
func (m *Biddings) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Biddings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Biddings.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Biddings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Biddings.Merge(m, src)
}
func (m *Biddings) XXX_Size() int {
	return m.Size()
}
func (m *Biddings) XXX_DiscardUnknown() {
	xxx_messageInfo_Biddings.DiscardUnknown(m)
}

var xxx_messageInfo_Biddings proto.InternalMessageInfo

type DutchBiddings struct {
	BiddingId          uint64                                  `protobuf:"varint,1,opt,name=bidding_id,json=biddingId,proto3" json:"bidding_id,omitempty" yaml:"bidding_id"`
	AuctionId          uint64                                  `protobuf:"varint,2,opt,name=auction_id,json=auctionId,proto3" json:"auction_id,omitempty" yaml:"auction_id"`
	AuctionStatus      string                                  `protobuf:"bytes,3,opt,name=auction_status,json=auctionStatus,proto3" json:"auction_status,omitempty" yaml:"auction_status"`
	OutflowTokenAmount github_com_cosmos_cosmos_sdk_types.Coin `protobuf:"bytes,4,opt,name=outflow_token_amount,json=outflowTokenAmount,proto3,casttype=github.com/cosmos/cosmos-sdk/types.Coin" json:"outflow_token_amount" yaml:"outflow_token_amount"`
	InflowTokenAmount  github_com_cosmos_cosmos_sdk_types.Coin `protobuf:"bytes,5,opt,name=inflow_token_amount,json=inflowTokenAmount,proto3,casttype=github.com/cosmos/cosmos-sdk/types.Coin" json:"inflow_token_amount" yaml:"inflow_token_amount"`
	Bidder             string                                  `protobuf:"bytes,6,opt,name=bidder,proto3" json:"bidder,omitempty" yaml:"bidder"`
	BiddingTimestamp   time.Time                               `protobuf:"bytes,7,opt,name=bidding_timestamp,json=biddingTimestamp,proto3,stdtime" json:"bidding_timestamp" yaml:"bidding_timestamp"`
	BiddingStatus      string                                  `protobuf:"bytes,8,opt,name=bidding_status,json=biddingStatus,proto3" json:"bidding_status,omitempty" yaml:"bidding_status"`
}

func (m *DutchBiddings) Reset()         { *m = DutchBiddings{} }
func (m *DutchBiddings) String() string { return proto.CompactTextString(m) }
func (*DutchBiddings) ProtoMessage()    {}
func (*DutchBiddings) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5a3f4b8597bafd2, []int{1}
}
func (m *DutchBiddings) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DutchBiddings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DutchBiddings.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DutchBiddings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DutchBiddings.Merge(m, src)
}
func (m *DutchBiddings) XXX_Size() int {
	return m.Size()
}
func (m *DutchBiddings) XXX_DiscardUnknown() {
	xxx_messageInfo_DutchBiddings.DiscardUnknown(m)
}

var xxx_messageInfo_DutchBiddings proto.InternalMessageInfo

type UserBiddings struct {
	Id         uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" yaml:"id"`
	Bidder     string   `protobuf:"bytes,2,opt,name=bidder,proto3" json:"bidder,omitempty" yaml:"bidder"`
	BiddingIds []uint64 `protobuf:"varint,3,rep,packed,name=bidding_ids,json=biddingIds,proto3" json:"bidding_ids,omitempty" yaml:"bidding_ids"`
}

func (m *UserBiddings) Reset()         { *m = UserBiddings{} }
func (m *UserBiddings) String() string { return proto.CompactTextString(m) }
func (*UserBiddings) ProtoMessage()    {}
func (*UserBiddings) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5a3f4b8597bafd2, []int{2}
}
func (m *UserBiddings) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserBiddings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UserBiddings.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UserBiddings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserBiddings.Merge(m, src)
}
func (m *UserBiddings) XXX_Size() int {
	return m.Size()
}
func (m *UserBiddings) XXX_DiscardUnknown() {
	xxx_messageInfo_UserBiddings.DiscardUnknown(m)
}

var xxx_messageInfo_UserBiddings proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Biddings)(nil), "comdex.auction.v1beta1.Biddings")
	proto.RegisterType((*DutchBiddings)(nil), "comdex.auction.v1beta1.DutchBiddings")
	proto.RegisterType((*UserBiddings)(nil), "comdex.auction.v1beta1.UserBiddings")
}

func init() {
	proto.RegisterFile("comdex/auction/v1beta1/biddings.proto", fileDescriptor_a5a3f4b8597bafd2)
}

var fileDescriptor_a5a3f4b8597bafd2 = []byte{
	// 643 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x55, 0xbd, 0x6e, 0xd4, 0x4c,
	0x14, 0xf5, 0x6c, 0xb2, 0xfb, 0x65, 0x27, 0xdf, 0x22, 0xd6, 0xf9, 0x91, 0xb3, 0x08, 0x3b, 0x1a,
	0x81, 0x58, 0x8a, 0xd8, 0x0a, 0x44, 0x42, 0x42, 0x42, 0x02, 0x87, 0x26, 0x05, 0x14, 0x4e, 0x68,
	0x68, 0x56, 0xfe, 0x5b, 0x67, 0x88, 0xed, 0x89, 0x76, 0xc6, 0x40, 0xde, 0x80, 0x32, 0x74, 0x48,
	0xf0, 0x00, 0x3c, 0x4a, 0xca, 0x88, 0x8a, 0xca, 0xc0, 0xe6, 0x0d, 0x5c, 0x52, 0x21, 0x7b, 0xc6,
	0x76, 0x1c, 0x90, 0x96, 0x88, 0x02, 0x51, 0xed, 0xce, 0xdc, 0x33, 0xe7, 0x9e, 0xb9, 0xf7, 0x5c,
	0x0f, 0xbc, 0xe9, 0x92, 0xc8, 0xf3, 0x5f, 0x1b, 0x76, 0xe2, 0x32, 0x4c, 0x62, 0xe3, 0xe5, 0xa6,
	0xe3, 0x33, 0x7b, 0xd3, 0x70, 0xb0, 0xe7, 0xe1, 0x38, 0xa0, 0xfa, 0xe1, 0x84, 0x30, 0x22, 0xaf,
	0x72, 0x98, 0x2e, 0x60, 0xba, 0x80, 0x0d, 0x96, 0x03, 0x12, 0x90, 0x02, 0x62, 0xe4, 0xff, 0x38,
	0x7a, 0xa0, 0x05, 0x84, 0x04, 0xa1, 0x6f, 0x14, 0x2b, 0x27, 0x19, 0x1b, 0x0c, 0x47, 0x3e, 0x65,
	0x76, 0x74, 0x28, 0x00, 0xaa, 0x4b, 0x68, 0x44, 0xa8, 0xe1, 0xd8, 0xd4, 0xaf, 0x52, 0xba, 0x04,
	0xc7, 0x3c, 0x8e, 0xde, 0xb4, 0xe1, 0x82, 0x29, 0x14, 0xc8, 0xd7, 0x61, 0x0b, 0x7b, 0x0a, 0x58,
	0x07, 0xc3, 0x79, 0xb3, 0x97, 0xa5, 0x5a, 0xf7, 0xc8, 0x8e, 0xc2, 0xfb, 0x08, 0x7b, 0xc8, 0x6a,
	0x61, 0x4f, 0xde, 0x82, 0x50, 0xa8, 0x1a, 0x61, 0x4f, 0x69, 0x15, 0xb0, 0x95, 0x2c, 0xd5, 0xfa,
	0x1c, 0x56, 0xc7, 0x90, 0xd5, 0x15, 0x8b, 0x1d, 0x4f, 0x7e, 0x08, 0xaf, 0x94, 0x11, 0xca, 0x6c,
	0x96, 0x50, 0x65, 0x6e, 0x1d, 0x0c, 0xbb, 0xe6, 0x5a, 0x96, 0x6a, 0x2b, 0xcd, 0x93, 0x3c, 0x8e,
	0xac, 0x9e, 0xd8, 0xd8, 0x2d, 0xd6, 0xf2, 0x07, 0x00, 0x97, 0xc5, 0x8e, 0xef, 0x8d, 0x5c, 0x12,
	0x86, 0x36, 0xf3, 0x27, 0x76, 0xa8, 0xcc, 0xaf, 0x83, 0xe1, 0xe2, 0x9d, 0x35, 0x9d, 0xdf, 0x51,
	0xcf, 0xef, 0x58, 0xd6, 0x4b, 0xdf, 0x26, 0x38, 0x36, 0x9f, 0x9e, 0xa4, 0x9a, 0x94, 0xa5, 0xda,
	0xb5, 0x46, 0x9e, 0x06, 0x09, 0xfa, 0x9e, 0x6a, 0xb7, 0x02, 0xcc, 0xf6, 0x13, 0x47, 0x77, 0x49,
	0x64, 0x88, 0x7a, 0xf1, 0x9f, 0x0d, 0xea, 0x1d, 0x18, 0xec, 0xe8, 0xd0, 0xa7, 0x05, 0x9f, 0xb5,
	0x54, 0x31, 0x6c, 0x57, 0x04, 0xf2, 0x6d, 0xd8, 0xc9, 0x7b, 0xe8, 0x4f, 0x94, 0x76, 0x71, 0xb1,
	0x7e, 0x96, 0x6a, 0x3d, 0x9e, 0x90, 0xef, 0x23, 0x4b, 0x00, 0xe4, 0x17, 0x70, 0xce, 0xc1, 0x9e,
	0xd2, 0x99, 0xa5, 0xfb, 0x81, 0xd0, 0x0d, 0x2b, 0x9a, 0x4b, 0xc9, 0xcc, 0x93, 0xc8, 0x11, 0xec,
	0x0b, 0x6b, 0x8d, 0x2a, 0x53, 0x28, 0xff, 0x15, 0x99, 0x07, 0x3a, 0xb7, 0x8d, 0x5e, 0xda, 0x46,
	0xdf, 0x2b, 0x11, 0xe6, 0x0d, 0x91, 0x5a, 0xa9, 0x6f, 0xd0, 0xa0, 0x40, 0xc7, 0x5f, 0x34, 0x60,
	0x5d, 0x15, 0xfb, 0xd5, 0xb9, 0xbc, 0xcd, 0x25, 0x56, 0xb4, 0x79, 0xe1, 0x62, 0x9b, 0x9b, 0x71,
	0x64, 0xf5, 0xc4, 0x06, 0x6f, 0x33, 0xfa, 0xd4, 0x86, 0xbd, 0xc7, 0x09, 0x73, 0xf7, 0x2b, 0x3f,
	0x6e, 0x41, 0x58, 0x9e, 0xa9, 0x7c, 0x79, 0xce, 0x70, 0x75, 0x0c, 0x59, 0x5d, 0xb1, 0xd8, 0xf9,
	0x7b, 0x36, 0x7d, 0x0f, 0xe0, 0x32, 0x49, 0xd8, 0x38, 0x24, 0xaf, 0x46, 0x8c, 0x1c, 0xf8, 0xf1,
	0xc8, 0x8e, 0x48, 0x12, 0xb3, 0x4b, 0xdb, 0xf4, 0x57, 0x24, 0x97, 0xea, 0xbf, 0x2c, 0x18, 0xf6,
	0x72, 0x82, 0x47, 0xc5, 0x79, 0xf9, 0x1d, 0x80, 0x4b, 0x38, 0xfe, 0x59, 0x5c, 0x7b, 0x96, 0xb8,
	0x27, 0x42, 0xdc, 0x40, 0x7c, 0x0c, 0xe2, 0x3f, 0xd3, 0xd6, 0xe7, 0x04, 0xe7, 0xa5, 0xd5, 0x03,
	0xd4, 0x99, 0x35, 0x40, 0xff, 0x9c, 0xa9, 0xdf, 0x02, 0xf8, 0xff, 0x33, 0xea, 0x4f, 0x7e, 0xf7,
	0x1b, 0x5b, 0xd7, 0xa2, 0x35, 0xab, 0x16, 0xf7, 0xe0, 0x62, 0x3d, 0x01, 0xb9, 0x5d, 0xe7, 0x86,
	0xf3, 0xe6, 0x6a, 0x96, 0x6a, 0xf2, 0xc5, 0xf1, 0xa0, 0xc8, 0x82, 0xd5, 0x7c, 0x50, 0x73, 0xf7,
	0xe4, 0x9b, 0x2a, 0x7d, 0x9c, 0xaa, 0xd2, 0xc9, 0x54, 0x05, 0xa7, 0x53, 0x15, 0x7c, 0x9d, 0xaa,
	0xe0, 0xf8, 0x4c, 0x95, 0x4e, 0xcf, 0x54, 0xe9, 0xf3, 0x99, 0x2a, 0x3d, 0xdf, 0x6c, 0x74, 0x33,
	0x7f, 0x8f, 0x36, 0xc8, 0x78, 0x8c, 0x5d, 0x6c, 0x87, 0x62, 0x6d, 0xd4, 0x0f, 0x59, 0xd1, 0x5c,
	0xa7, 0x53, 0x94, 0xfd, 0xee, 0x8f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbc, 0x2c, 0xcd, 0x22, 0xe7,
	0x06, 0x00, 0x00,
}

func (m *Biddings) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Biddings) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Biddings) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BiddingStatus) > 0 {
		i -= len(m.BiddingStatus)
		copy(dAtA[i:], m.BiddingStatus)
		i = encodeVarintBiddings(dAtA, i, uint64(len(m.BiddingStatus)))
		i--
		dAtA[i] = 0x42
	}
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.BiddingTimestamp, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.BiddingTimestamp):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintBiddings(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x3a
	{
		size, err := m.Bid.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintBiddings(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	if len(m.Bidder) > 0 {
		i -= len(m.Bidder)
		copy(dAtA[i:], m.Bidder)
		i = encodeVarintBiddings(dAtA, i, uint64(len(m.Bidder)))
		i--
		dAtA[i] = 0x2a
	}
	{
		size, err := m.AuctionedCollateral.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintBiddings(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.AuctionStatus) > 0 {
		i -= len(m.AuctionStatus)
		copy(dAtA[i:], m.AuctionStatus)
		i = encodeVarintBiddings(dAtA, i, uint64(len(m.AuctionStatus)))
		i--
		dAtA[i] = 0x1a
	}
	if m.AuctionId != 0 {
		i = encodeVarintBiddings(dAtA, i, uint64(m.AuctionId))
		i--
		dAtA[i] = 0x10
	}
	if m.Id != 0 {
		i = encodeVarintBiddings(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *DutchBiddings) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DutchBiddings) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DutchBiddings) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BiddingStatus) > 0 {
		i -= len(m.BiddingStatus)
		copy(dAtA[i:], m.BiddingStatus)
		i = encodeVarintBiddings(dAtA, i, uint64(len(m.BiddingStatus)))
		i--
		dAtA[i] = 0x42
	}
	n4, err4 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.BiddingTimestamp, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.BiddingTimestamp):])
	if err4 != nil {
		return 0, err4
	}
	i -= n4
	i = encodeVarintBiddings(dAtA, i, uint64(n4))
	i--
	dAtA[i] = 0x3a
	if len(m.Bidder) > 0 {
		i -= len(m.Bidder)
		copy(dAtA[i:], m.Bidder)
		i = encodeVarintBiddings(dAtA, i, uint64(len(m.Bidder)))
		i--
		dAtA[i] = 0x32
	}
	{
		size, err := m.InflowTokenAmount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintBiddings(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size, err := m.OutflowTokenAmount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintBiddings(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.AuctionStatus) > 0 {
		i -= len(m.AuctionStatus)
		copy(dAtA[i:], m.AuctionStatus)
		i = encodeVarintBiddings(dAtA, i, uint64(len(m.AuctionStatus)))
		i--
		dAtA[i] = 0x1a
	}
	if m.AuctionId != 0 {
		i = encodeVarintBiddings(dAtA, i, uint64(m.AuctionId))
		i--
		dAtA[i] = 0x10
	}
	if m.BiddingId != 0 {
		i = encodeVarintBiddings(dAtA, i, uint64(m.BiddingId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *UserBiddings) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserBiddings) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UserBiddings) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BiddingIds) > 0 {
		dAtA8 := make([]byte, len(m.BiddingIds)*10)
		var j7 int
		for _, num := range m.BiddingIds {
			for num >= 1<<7 {
				dAtA8[j7] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j7++
			}
			dAtA8[j7] = uint8(num)
			j7++
		}
		i -= j7
		copy(dAtA[i:], dAtA8[:j7])
		i = encodeVarintBiddings(dAtA, i, uint64(j7))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Bidder) > 0 {
		i -= len(m.Bidder)
		copy(dAtA[i:], m.Bidder)
		i = encodeVarintBiddings(dAtA, i, uint64(len(m.Bidder)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintBiddings(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintBiddings(dAtA []byte, offset int, v uint64) int {
	offset -= sovBiddings(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Biddings) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovBiddings(uint64(m.Id))
	}
	if m.AuctionId != 0 {
		n += 1 + sovBiddings(uint64(m.AuctionId))
	}
	l = len(m.AuctionStatus)
	if l > 0 {
		n += 1 + l + sovBiddings(uint64(l))
	}
	l = m.AuctionedCollateral.Size()
	n += 1 + l + sovBiddings(uint64(l))
	l = len(m.Bidder)
	if l > 0 {
		n += 1 + l + sovBiddings(uint64(l))
	}
	l = m.Bid.Size()
	n += 1 + l + sovBiddings(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.BiddingTimestamp)
	n += 1 + l + sovBiddings(uint64(l))
	l = len(m.BiddingStatus)
	if l > 0 {
		n += 1 + l + sovBiddings(uint64(l))
	}
	return n
}

func (m *DutchBiddings) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BiddingId != 0 {
		n += 1 + sovBiddings(uint64(m.BiddingId))
	}
	if m.AuctionId != 0 {
		n += 1 + sovBiddings(uint64(m.AuctionId))
	}
	l = len(m.AuctionStatus)
	if l > 0 {
		n += 1 + l + sovBiddings(uint64(l))
	}
	l = m.OutflowTokenAmount.Size()
	n += 1 + l + sovBiddings(uint64(l))
	l = m.InflowTokenAmount.Size()
	n += 1 + l + sovBiddings(uint64(l))
	l = len(m.Bidder)
	if l > 0 {
		n += 1 + l + sovBiddings(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.BiddingTimestamp)
	n += 1 + l + sovBiddings(uint64(l))
	l = len(m.BiddingStatus)
	if l > 0 {
		n += 1 + l + sovBiddings(uint64(l))
	}
	return n
}

func (m *UserBiddings) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovBiddings(uint64(m.Id))
	}
	l = len(m.Bidder)
	if l > 0 {
		n += 1 + l + sovBiddings(uint64(l))
	}
	if len(m.BiddingIds) > 0 {
		l = 0
		for _, e := range m.BiddingIds {
			l += sovBiddings(uint64(e))
		}
		n += 1 + sovBiddings(uint64(l)) + l
	}
	return n
}

func sovBiddings(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBiddings(x uint64) (n int) {
	return sovBiddings(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Biddings) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBiddings
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
			return fmt.Errorf("proto: Biddings: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Biddings: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBiddings
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuctionId", wireType)
			}
			m.AuctionId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBiddings
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AuctionId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuctionStatus", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBiddings
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
				return ErrInvalidLengthBiddings
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBiddings
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AuctionStatus = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuctionedCollateral", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBiddings
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
				return ErrInvalidLengthBiddings
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBiddings
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AuctionedCollateral.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bidder", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBiddings
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
				return ErrInvalidLengthBiddings
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBiddings
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bidder = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bid", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBiddings
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
				return ErrInvalidLengthBiddings
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBiddings
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Bid.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BiddingTimestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBiddings
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
				return ErrInvalidLengthBiddings
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBiddings
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.BiddingTimestamp, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BiddingStatus", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBiddings
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
				return ErrInvalidLengthBiddings
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBiddings
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BiddingStatus = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBiddings(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBiddings
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
func (m *DutchBiddings) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBiddings
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
			return fmt.Errorf("proto: DutchBiddings: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DutchBiddings: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BiddingId", wireType)
			}
			m.BiddingId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBiddings
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BiddingId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuctionId", wireType)
			}
			m.AuctionId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBiddings
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AuctionId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuctionStatus", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBiddings
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
				return ErrInvalidLengthBiddings
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBiddings
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AuctionStatus = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OutflowTokenAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBiddings
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
				return ErrInvalidLengthBiddings
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBiddings
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.OutflowTokenAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InflowTokenAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBiddings
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
				return ErrInvalidLengthBiddings
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBiddings
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.InflowTokenAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bidder", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBiddings
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
				return ErrInvalidLengthBiddings
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBiddings
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bidder = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BiddingTimestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBiddings
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
				return ErrInvalidLengthBiddings
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBiddings
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.BiddingTimestamp, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BiddingStatus", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBiddings
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
				return ErrInvalidLengthBiddings
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBiddings
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BiddingStatus = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBiddings(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBiddings
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
func (m *UserBiddings) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBiddings
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
			return fmt.Errorf("proto: UserBiddings: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserBiddings: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBiddings
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bidder", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBiddings
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
				return ErrInvalidLengthBiddings
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBiddings
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bidder = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowBiddings
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.BiddingIds = append(m.BiddingIds, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowBiddings
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthBiddings
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthBiddings
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.BiddingIds) == 0 {
					m.BiddingIds = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowBiddings
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.BiddingIds = append(m.BiddingIds, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field BiddingIds", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBiddings(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBiddings
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
func skipBiddings(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBiddings
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
					return 0, ErrIntOverflowBiddings
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
					return 0, ErrIntOverflowBiddings
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
				return 0, ErrInvalidLengthBiddings
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBiddings
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBiddings
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBiddings        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBiddings          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBiddings = fmt.Errorf("proto: unexpected end of group")
)