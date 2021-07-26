// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.2
// source: v1/message.proto

package message

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Auction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	PayloadCid       string                 `protobuf:"bytes,2,opt,name=payload_cid,json=payloadCid,proto3" json:"payload_cid,omitempty"`
	DealSize         uint64                 `protobuf:"varint,3,opt,name=deal_size,json=dealSize,proto3" json:"deal_size,omitempty"`
	DealDuration     uint64                 `protobuf:"varint,4,opt,name=deal_duration,json=dealDuration,proto3" json:"deal_duration,omitempty"`
	FilEpochDeadline uint64                 `protobuf:"varint,5,opt,name=fil_epoch_deadline,json=filEpochDeadline,proto3" json:"fil_epoch_deadline,omitempty"`
	Sources          *Sources               `protobuf:"bytes,6,opt,name=sources,proto3" json:"sources,omitempty"`
	EndsAt           *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=ends_at,json=endsAt,proto3" json:"ends_at,omitempty"`
}

func (x *Auction) Reset() {
	*x = Auction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Auction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Auction) ProtoMessage() {}

func (x *Auction) ProtoReflect() protoreflect.Message {
	mi := &file_v1_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Auction.ProtoReflect.Descriptor instead.
func (*Auction) Descriptor() ([]byte, []int) {
	return file_v1_message_proto_rawDescGZIP(), []int{0}
}

func (x *Auction) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Auction) GetPayloadCid() string {
	if x != nil {
		return x.PayloadCid
	}
	return ""
}

func (x *Auction) GetDealSize() uint64 {
	if x != nil {
		return x.DealSize
	}
	return 0
}

func (x *Auction) GetDealDuration() uint64 {
	if x != nil {
		return x.DealDuration
	}
	return 0
}

func (x *Auction) GetFilEpochDeadline() uint64 {
	if x != nil {
		return x.FilEpochDeadline
	}
	return 0
}

func (x *Auction) GetSources() *Sources {
	if x != nil {
		return x.Sources
	}
	return nil
}

func (x *Auction) GetEndsAt() *timestamppb.Timestamp {
	if x != nil {
		return x.EndsAt
	}
	return nil
}

type Bid struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuctionId        string `protobuf:"bytes,1,opt,name=auction_id,json=auctionId,proto3" json:"auction_id,omitempty"`
	MinerId          string `protobuf:"bytes,2,opt,name=miner_id,json=minerId,proto3" json:"miner_id,omitempty"`
	WalletAddrSig    []byte `protobuf:"bytes,3,opt,name=wallet_addr_sig,json=walletAddrSig,proto3" json:"wallet_addr_sig,omitempty"`
	AskPrice         int64  `protobuf:"varint,4,opt,name=ask_price,json=askPrice,proto3" json:"ask_price,omitempty"`
	VerifiedAskPrice int64  `protobuf:"varint,5,opt,name=verified_ask_price,json=verifiedAskPrice,proto3" json:"verified_ask_price,omitempty"`
	StartEpoch       uint64 `protobuf:"varint,6,opt,name=start_epoch,json=startEpoch,proto3" json:"start_epoch,omitempty"`
	FastRetrieval    bool   `protobuf:"varint,7,opt,name=fast_retrieval,json=fastRetrieval,proto3" json:"fast_retrieval,omitempty"`
}

func (x *Bid) Reset() {
	*x = Bid{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bid) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bid) ProtoMessage() {}

func (x *Bid) ProtoReflect() protoreflect.Message {
	mi := &file_v1_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bid.ProtoReflect.Descriptor instead.
func (*Bid) Descriptor() ([]byte, []int) {
	return file_v1_message_proto_rawDescGZIP(), []int{1}
}

func (x *Bid) GetAuctionId() string {
	if x != nil {
		return x.AuctionId
	}
	return ""
}

func (x *Bid) GetMinerId() string {
	if x != nil {
		return x.MinerId
	}
	return ""
}

func (x *Bid) GetWalletAddrSig() []byte {
	if x != nil {
		return x.WalletAddrSig
	}
	return nil
}

func (x *Bid) GetAskPrice() int64 {
	if x != nil {
		return x.AskPrice
	}
	return 0
}

func (x *Bid) GetVerifiedAskPrice() int64 {
	if x != nil {
		return x.VerifiedAskPrice
	}
	return 0
}

func (x *Bid) GetStartEpoch() uint64 {
	if x != nil {
		return x.StartEpoch
	}
	return 0
}

func (x *Bid) GetFastRetrieval() bool {
	if x != nil {
		return x.FastRetrieval
	}
	return false
}

type WinningBid struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuctionId string `protobuf:"bytes,1,opt,name=auction_id,json=auctionId,proto3" json:"auction_id,omitempty"`
	BidId     string `protobuf:"bytes,2,opt,name=bid_id,json=bidId,proto3" json:"bid_id,omitempty"`
}

func (x *WinningBid) Reset() {
	*x = WinningBid{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WinningBid) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WinningBid) ProtoMessage() {}

func (x *WinningBid) ProtoReflect() protoreflect.Message {
	mi := &file_v1_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WinningBid.ProtoReflect.Descriptor instead.
func (*WinningBid) Descriptor() ([]byte, []int) {
	return file_v1_message_proto_rawDescGZIP(), []int{2}
}

func (x *WinningBid) GetAuctionId() string {
	if x != nil {
		return x.AuctionId
	}
	return ""
}

func (x *WinningBid) GetBidId() string {
	if x != nil {
		return x.BidId
	}
	return ""
}

type WinningBidProposal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuctionId   string `protobuf:"bytes,1,opt,name=auction_id,json=auctionId,proto3" json:"auction_id,omitempty"`
	BidId       string `protobuf:"bytes,2,opt,name=bid_id,json=bidId,proto3" json:"bid_id,omitempty"`
	ProposalCid string `protobuf:"bytes,3,opt,name=proposal_cid,json=proposalCid,proto3" json:"proposal_cid,omitempty"`
}

func (x *WinningBidProposal) Reset() {
	*x = WinningBidProposal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WinningBidProposal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WinningBidProposal) ProtoMessage() {}

func (x *WinningBidProposal) ProtoReflect() protoreflect.Message {
	mi := &file_v1_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WinningBidProposal.ProtoReflect.Descriptor instead.
func (*WinningBidProposal) Descriptor() ([]byte, []int) {
	return file_v1_message_proto_rawDescGZIP(), []int{3}
}

func (x *WinningBidProposal) GetAuctionId() string {
	if x != nil {
		return x.AuctionId
	}
	return ""
}

func (x *WinningBidProposal) GetBidId() string {
	if x != nil {
		return x.BidId
	}
	return ""
}

func (x *WinningBidProposal) GetProposalCid() string {
	if x != nil {
		return x.ProposalCid
	}
	return ""
}

type Sources struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CarUrl  *Sources_CARURL  `protobuf:"bytes,1,opt,name=car_url,json=carUrl,proto3" json:"car_url,omitempty"`
	CarIpfs *Sources_CARIPFS `protobuf:"bytes,2,opt,name=car_ipfs,json=carIpfs,proto3" json:"car_ipfs,omitempty"`
}

func (x *Sources) Reset() {
	*x = Sources{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sources) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sources) ProtoMessage() {}

func (x *Sources) ProtoReflect() protoreflect.Message {
	mi := &file_v1_message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sources.ProtoReflect.Descriptor instead.
func (*Sources) Descriptor() ([]byte, []int) {
	return file_v1_message_proto_rawDescGZIP(), []int{4}
}

func (x *Sources) GetCarUrl() *Sources_CARURL {
	if x != nil {
		return x.CarUrl
	}
	return nil
}

func (x *Sources) GetCarIpfs() *Sources_CARIPFS {
	if x != nil {
		return x.CarIpfs
	}
	return nil
}

type Sources_CARURL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	URL string `protobuf:"bytes,1,opt,name=URL,proto3" json:"URL,omitempty"`
}

func (x *Sources_CARURL) Reset() {
	*x = Sources_CARURL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_message_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sources_CARURL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sources_CARURL) ProtoMessage() {}

func (x *Sources_CARURL) ProtoReflect() protoreflect.Message {
	mi := &file_v1_message_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sources_CARURL.ProtoReflect.Descriptor instead.
func (*Sources_CARURL) Descriptor() ([]byte, []int) {
	return file_v1_message_proto_rawDescGZIP(), []int{4, 0}
}

func (x *Sources_CARURL) GetURL() string {
	if x != nil {
		return x.URL
	}
	return ""
}

type Sources_CARIPFS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cid        string   `protobuf:"bytes,1,opt,name=cid,proto3" json:"cid,omitempty"`
	Multiaddrs []string `protobuf:"bytes,2,rep,name=multiaddrs,proto3" json:"multiaddrs,omitempty"`
}

func (x *Sources_CARIPFS) Reset() {
	*x = Sources_CARIPFS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_message_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sources_CARIPFS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sources_CARIPFS) ProtoMessage() {}

func (x *Sources_CARIPFS) ProtoReflect() protoreflect.Message {
	mi := &file_v1_message_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sources_CARIPFS.ProtoReflect.Descriptor instead.
func (*Sources_CARIPFS) Descriptor() ([]byte, []int) {
	return file_v1_message_proto_rawDescGZIP(), []int{4, 1}
}

func (x *Sources_CARIPFS) GetCid() string {
	if x != nil {
		return x.Cid
	}
	return ""
}

func (x *Sources_CARIPFS) GetMultiaddrs() []string {
	if x != nil {
		return x.Multiaddrs
	}
	return nil
}

var File_v1_message_proto protoreflect.FileDescriptor

var file_v1_message_proto_rawDesc = []byte{
	0x0a, 0x10, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x94, 0x02, 0x0a, 0x07, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x63, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x43,
	0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x61, 0x6c, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x64, 0x65, 0x61, 0x6c, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x23, 0x0a, 0x0d, 0x64, 0x65, 0x61, 0x6c, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x64, 0x65, 0x61, 0x6c, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x12, 0x66, 0x69, 0x6c, 0x5f, 0x65, 0x70, 0x6f, 0x63,
	0x68, 0x5f, 0x64, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x10, 0x66, 0x69, 0x6c, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x44, 0x65, 0x61, 0x64, 0x6c, 0x69,
	0x6e, 0x65, 0x12, 0x33, 0x0a, 0x07, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x52, 0x07,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x33, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x73, 0x5f,
	0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x06, 0x65, 0x6e, 0x64, 0x73, 0x41, 0x74, 0x22, 0xfa, 0x01, 0x0a,
	0x03, 0x42, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x26,
	0x0a, 0x0f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x5f, 0x73, 0x69,
	0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x41,
	0x64, 0x64, 0x72, 0x53, 0x69, 0x67, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x73, 0x6b, 0x5f, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x61, 0x73, 0x6b, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x5f,
	0x61, 0x73, 0x6b, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x10, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x41, 0x73, 0x6b, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x65, 0x70, 0x6f, 0x63, 0x68,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x45, 0x70, 0x6f,
	0x63, 0x68, 0x12, 0x25, 0x0a, 0x0e, 0x66, 0x61, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x74, 0x72, 0x69,
	0x65, 0x76, 0x61, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x66, 0x61, 0x73, 0x74,
	0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x61, 0x6c, 0x22, 0x42, 0x0a, 0x0a, 0x57, 0x69, 0x6e,
	0x6e, 0x69, 0x6e, 0x67, 0x42, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x75, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x62, 0x69, 0x64, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x69, 0x64, 0x49, 0x64, 0x22, 0x6d, 0x0a,
	0x12, 0x57, 0x69, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x42, 0x69, 0x64, 0x50, 0x72, 0x6f, 0x70, 0x6f,
	0x73, 0x61, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x62, 0x69, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x62, 0x69, 0x64, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f,
	0x70, 0x6f, 0x73, 0x61, 0x6c, 0x5f, 0x63, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x43, 0x69, 0x64, 0x22, 0xdb, 0x01, 0x0a,
	0x07, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x39, 0x0a, 0x07, 0x63, 0x61, 0x72, 0x5f,
	0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x41, 0x52, 0x55, 0x52, 0x4c, 0x52, 0x06, 0x63, 0x61, 0x72,
	0x55, 0x72, 0x6c, 0x12, 0x3c, 0x0a, 0x08, 0x63, 0x61, 0x72, 0x5f, 0x69, 0x70, 0x66, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x2e, 0x43, 0x41, 0x52, 0x49, 0x50, 0x46, 0x53, 0x52, 0x07, 0x63, 0x61, 0x72, 0x49, 0x70, 0x66,
	0x73, 0x1a, 0x1a, 0x0a, 0x06, 0x43, 0x41, 0x52, 0x55, 0x52, 0x4c, 0x12, 0x10, 0x0a, 0x03, 0x55,
	0x52, 0x4c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55, 0x52, 0x4c, 0x1a, 0x3b, 0x0a,
	0x07, 0x43, 0x41, 0x52, 0x49, 0x50, 0x46, 0x53, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x75,
	0x6c, 0x74, 0x69, 0x61, 0x64, 0x64, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a,
	0x6d, 0x75, 0x6c, 0x74, 0x69, 0x61, 0x64, 0x64, 0x72, 0x73, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x78, 0x74, 0x69, 0x6c, 0x65,
	0x69, 0x6f, 0x2f, 0x62, 0x69, 0x64, 0x62, 0x6f, 0x74, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x76, 0x31,
	0x3b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_message_proto_rawDescOnce sync.Once
	file_v1_message_proto_rawDescData = file_v1_message_proto_rawDesc
)

func file_v1_message_proto_rawDescGZIP() []byte {
	file_v1_message_proto_rawDescOnce.Do(func() {
		file_v1_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_message_proto_rawDescData)
	})
	return file_v1_message_proto_rawDescData
}

var file_v1_message_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_v1_message_proto_goTypes = []interface{}{
	(*Auction)(nil),               // 0: proto.v1.message.Auction
	(*Bid)(nil),                   // 1: proto.v1.message.Bid
	(*WinningBid)(nil),            // 2: proto.v1.message.WinningBid
	(*WinningBidProposal)(nil),    // 3: proto.v1.message.WinningBidProposal
	(*Sources)(nil),               // 4: proto.v1.message.Sources
	(*Sources_CARURL)(nil),        // 5: proto.v1.message.Sources.CARURL
	(*Sources_CARIPFS)(nil),       // 6: proto.v1.message.Sources.CARIPFS
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
}
var file_v1_message_proto_depIdxs = []int32{
	4, // 0: proto.v1.message.Auction.sources:type_name -> proto.v1.message.Sources
	7, // 1: proto.v1.message.Auction.ends_at:type_name -> google.protobuf.Timestamp
	5, // 2: proto.v1.message.Sources.car_url:type_name -> proto.v1.message.Sources.CARURL
	6, // 3: proto.v1.message.Sources.car_ipfs:type_name -> proto.v1.message.Sources.CARIPFS
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_v1_message_proto_init() }
func file_v1_message_proto_init() {
	if File_v1_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Auction); i {
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
		file_v1_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bid); i {
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
		file_v1_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WinningBid); i {
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
		file_v1_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WinningBidProposal); i {
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
		file_v1_message_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sources); i {
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
		file_v1_message_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sources_CARURL); i {
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
		file_v1_message_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sources_CARIPFS); i {
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
			RawDescriptor: file_v1_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_message_proto_goTypes,
		DependencyIndexes: file_v1_message_proto_depIdxs,
		MessageInfos:      file_v1_message_proto_msgTypes,
	}.Build()
	File_v1_message_proto = out.File
	file_v1_message_proto_rawDesc = nil
	file_v1_message_proto_goTypes = nil
	file_v1_message_proto_depIdxs = nil
}
