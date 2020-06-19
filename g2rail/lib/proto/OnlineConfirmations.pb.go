// Code generated by protoc-gen-go. DO NOT EDIT.
// source: lib/proto/OnlineConfirmations.proto

package g2rail

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ConfirmRequest struct {
	OrderId              string   `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConfirmRequest) Reset()         { *m = ConfirmRequest{} }
func (m *ConfirmRequest) String() string { return proto.CompactTextString(m) }
func (*ConfirmRequest) ProtoMessage()    {}
func (*ConfirmRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_52d952a1be3b823c, []int{0}
}

func (m *ConfirmRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfirmRequest.Unmarshal(m, b)
}
func (m *ConfirmRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfirmRequest.Marshal(b, m, deterministic)
}
func (m *ConfirmRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfirmRequest.Merge(m, src)
}
func (m *ConfirmRequest) XXX_Size() int {
	return xxx_messageInfo_ConfirmRequest.Size(m)
}
func (m *ConfirmRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfirmRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConfirmRequest proto.InternalMessageInfo

func (m *ConfirmRequest) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

type OnlineConfirmationAsyncQueryRequest struct {
	AsyncKey             string   `protobuf:"bytes,1,opt,name=async_key,json=asyncKey,proto3" json:"async_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OnlineConfirmationAsyncQueryRequest) Reset()         { *m = OnlineConfirmationAsyncQueryRequest{} }
func (m *OnlineConfirmationAsyncQueryRequest) String() string { return proto.CompactTextString(m) }
func (*OnlineConfirmationAsyncQueryRequest) ProtoMessage()    {}
func (*OnlineConfirmationAsyncQueryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_52d952a1be3b823c, []int{1}
}

func (m *OnlineConfirmationAsyncQueryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OnlineConfirmationAsyncQueryRequest.Unmarshal(m, b)
}
func (m *OnlineConfirmationAsyncQueryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OnlineConfirmationAsyncQueryRequest.Marshal(b, m, deterministic)
}
func (m *OnlineConfirmationAsyncQueryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OnlineConfirmationAsyncQueryRequest.Merge(m, src)
}
func (m *OnlineConfirmationAsyncQueryRequest) XXX_Size() int {
	return xxx_messageInfo_OnlineConfirmationAsyncQueryRequest.Size(m)
}
func (m *OnlineConfirmationAsyncQueryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OnlineConfirmationAsyncQueryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OnlineConfirmationAsyncQueryRequest proto.InternalMessageInfo

func (m *OnlineConfirmationAsyncQueryRequest) GetAsyncKey() string {
	if m != nil {
		return m.AsyncKey
	}
	return ""
}

type OnlineConfirmationResponse struct {
	Id                   string         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ConfirmAgain         string         `protobuf:"bytes,3,opt,name=confirm_again,json=confirmAgain,proto3" json:"confirm_again,omitempty"`
	Order                *Order         `protobuf:"bytes,4,opt,name=order,proto3" json:"order,omitempty"`
	TicketPrice          *Price         `protobuf:"bytes,8,opt,name=ticket_price,json=ticketPrice,proto3" json:"ticket_price,omitempty"`
	PaymentPrice         *Price         `protobuf:"bytes,9,opt,name=payment_price,json=paymentPrice,proto3" json:"payment_price,omitempty"`
	ChargingPrice        *Price         `protobuf:"bytes,10,opt,name=charging_price,json=chargingPrice,proto3" json:"charging_price,omitempty"`
	RebateAmount         *Price         `protobuf:"bytes,11,opt,name=rebate_amount,json=rebateAmount,proto3" json:"rebate_amount,omitempty"`
	Records              []*PriceDetail `protobuf:"bytes,15,rep,name=records,proto3" json:"records,omitempty"`
	Loading              bool           `protobuf:"varint,16,opt,name=loading,proto3" json:"loading,omitempty"`
	OfflineFulfillment   bool           `protobuf:"varint,17,opt,name=offline_fulfillment,json=offlineFulfillment,proto3" json:"offline_fulfillment,omitempty"`
	TrainDescription     string         `protobuf:"bytes,18,opt,name=train_description,json=trainDescription,proto3" json:"train_description,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *OnlineConfirmationResponse) Reset()         { *m = OnlineConfirmationResponse{} }
func (m *OnlineConfirmationResponse) String() string { return proto.CompactTextString(m) }
func (*OnlineConfirmationResponse) ProtoMessage()    {}
func (*OnlineConfirmationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_52d952a1be3b823c, []int{2}
}

func (m *OnlineConfirmationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OnlineConfirmationResponse.Unmarshal(m, b)
}
func (m *OnlineConfirmationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OnlineConfirmationResponse.Marshal(b, m, deterministic)
}
func (m *OnlineConfirmationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OnlineConfirmationResponse.Merge(m, src)
}
func (m *OnlineConfirmationResponse) XXX_Size() int {
	return xxx_messageInfo_OnlineConfirmationResponse.Size(m)
}
func (m *OnlineConfirmationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_OnlineConfirmationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_OnlineConfirmationResponse proto.InternalMessageInfo

func (m *OnlineConfirmationResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *OnlineConfirmationResponse) GetConfirmAgain() string {
	if m != nil {
		return m.ConfirmAgain
	}
	return ""
}

func (m *OnlineConfirmationResponse) GetOrder() *Order {
	if m != nil {
		return m.Order
	}
	return nil
}

func (m *OnlineConfirmationResponse) GetTicketPrice() *Price {
	if m != nil {
		return m.TicketPrice
	}
	return nil
}

func (m *OnlineConfirmationResponse) GetPaymentPrice() *Price {
	if m != nil {
		return m.PaymentPrice
	}
	return nil
}

func (m *OnlineConfirmationResponse) GetChargingPrice() *Price {
	if m != nil {
		return m.ChargingPrice
	}
	return nil
}

func (m *OnlineConfirmationResponse) GetRebateAmount() *Price {
	if m != nil {
		return m.RebateAmount
	}
	return nil
}

func (m *OnlineConfirmationResponse) GetRecords() []*PriceDetail {
	if m != nil {
		return m.Records
	}
	return nil
}

func (m *OnlineConfirmationResponse) GetLoading() bool {
	if m != nil {
		return m.Loading
	}
	return false
}

func (m *OnlineConfirmationResponse) GetOfflineFulfillment() bool {
	if m != nil {
		return m.OfflineFulfillment
	}
	return false
}

func (m *OnlineConfirmationResponse) GetTrainDescription() string {
	if m != nil {
		return m.TrainDescription
	}
	return ""
}

type Order struct {
	Id                   string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Railway              *Railway     `protobuf:"bytes,2,opt,name=railway,proto3" json:"railway,omitempty"`
	OfflineFulfillment   bool         `protobuf:"varint,3,opt,name=offline_fulfillment,json=offlineFulfillment,proto3" json:"offline_fulfillment,omitempty"`
	TrainDescription     string       `protobuf:"bytes,4,opt,name=train_description,json=trainDescription,proto3" json:"train_description,omitempty"`
	From                 *Station     `protobuf:"bytes,5,opt,name=from,proto3" json:"from,omitempty"`
	To                   *Station     `protobuf:"bytes,6,opt,name=to,proto3" json:"to,omitempty"`
	Departure            string       `protobuf:"bytes,7,opt,name=departure,proto3" json:"departure,omitempty"`
	Arrival              string       `protobuf:"bytes,8,opt,name=arrival,proto3" json:"arrival,omitempty"`
	PNR                  string       `protobuf:"bytes,9,opt,name=PNR,proto3" json:"PNR,omitempty"`
	Trains               []*Train     `protobuf:"bytes,12,rep,name=trains,proto3" json:"trains,omitempty"`
	Passengers           []*Passenger `protobuf:"bytes,13,rep,name=passengers,proto3" json:"passengers,omitempty"`
	Tickets              []*Ticket    `protobuf:"bytes,14,rep,name=tickets,proto3" json:"tickets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Order) Reset()         { *m = Order{} }
func (m *Order) String() string { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()    {}
func (*Order) Descriptor() ([]byte, []int) {
	return fileDescriptor_52d952a1be3b823c, []int{3}
}

func (m *Order) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Order.Unmarshal(m, b)
}
func (m *Order) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Order.Marshal(b, m, deterministic)
}
func (m *Order) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Order.Merge(m, src)
}
func (m *Order) XXX_Size() int {
	return xxx_messageInfo_Order.Size(m)
}
func (m *Order) XXX_DiscardUnknown() {
	xxx_messageInfo_Order.DiscardUnknown(m)
}

var xxx_messageInfo_Order proto.InternalMessageInfo

func (m *Order) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Order) GetRailway() *Railway {
	if m != nil {
		return m.Railway
	}
	return nil
}

func (m *Order) GetOfflineFulfillment() bool {
	if m != nil {
		return m.OfflineFulfillment
	}
	return false
}

func (m *Order) GetTrainDescription() string {
	if m != nil {
		return m.TrainDescription
	}
	return ""
}

func (m *Order) GetFrom() *Station {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *Order) GetTo() *Station {
	if m != nil {
		return m.To
	}
	return nil
}

func (m *Order) GetDeparture() string {
	if m != nil {
		return m.Departure
	}
	return ""
}

func (m *Order) GetArrival() string {
	if m != nil {
		return m.Arrival
	}
	return ""
}

func (m *Order) GetPNR() string {
	if m != nil {
		return m.PNR
	}
	return ""
}

func (m *Order) GetTrains() []*Train {
	if m != nil {
		return m.Trains
	}
	return nil
}

func (m *Order) GetPassengers() []*Passenger {
	if m != nil {
		return m.Passengers
	}
	return nil
}

func (m *Order) GetTickets() []*Ticket {
	if m != nil {
		return m.Tickets
	}
	return nil
}

func init() {
	proto.RegisterType((*ConfirmRequest)(nil), "g2rail.ConfirmRequest")
	proto.RegisterType((*OnlineConfirmationAsyncQueryRequest)(nil), "g2rail.OnlineConfirmationAsyncQueryRequest")
	proto.RegisterType((*OnlineConfirmationResponse)(nil), "g2rail.OnlineConfirmationResponse")
	proto.RegisterType((*Order)(nil), "g2rail.Order")
}

func init() {
	proto.RegisterFile("lib/proto/OnlineConfirmations.proto", fileDescriptor_52d952a1be3b823c)
}

var fileDescriptor_52d952a1be3b823c = []byte{
	// 636 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xdd, 0x6e, 0xd3, 0x3c,
	0x18, 0x56, 0xdb, 0xad, 0x69, 0xdf, 0xfe, 0x6c, 0xf3, 0xa4, 0xc9, 0x5f, 0xbf, 0x49, 0x54, 0xad,
	0x90, 0x8a, 0x26, 0x3a, 0x28, 0x9c, 0x71, 0xd4, 0x6d, 0x1a, 0x42, 0x48, 0xac, 0x18, 0x8e, 0x38,
	0x89, 0xdc, 0xc4, 0x09, 0xd6, 0x52, 0x3b, 0x38, 0x2e, 0xa8, 0xb7, 0xc4, 0x31, 0x77, 0xc0, 0xf5,
	0x70, 0x0f, 0xc8, 0x4e, 0x9c, 0x52, 0x35, 0x13, 0xe2, 0x2c, 0x7e, 0x9e, 0xe7, 0xfd, 0xcb, 0xfb,
	0x03, 0xe3, 0x84, 0x2f, 0x2f, 0x53, 0x25, 0xb5, 0xbc, 0xbc, 0x13, 0x09, 0x17, 0xec, 0x5a, 0x8a,
	0x88, 0xab, 0x15, 0xd5, 0x5c, 0x8a, 0x6c, 0x6a, 0x19, 0xd4, 0x8c, 0x67, 0x8a, 0xf2, 0x64, 0x70,
	0xb6, 0x15, 0x5f, 0xcb, 0xd5, 0x4a, 0x8a, 0x9c, 0x1f, 0x5d, 0x40, 0xbf, 0x30, 0x23, 0xec, 0xcb,
	0x9a, 0x65, 0x1a, 0xfd, 0x07, 0x2d, 0xa9, 0x42, 0xa6, 0x7c, 0x1e, 0xe2, 0xda, 0xb0, 0x36, 0x69,
	0x13, 0xcf, 0xbe, 0xdf, 0x84, 0xa3, 0x2b, 0x18, 0xef, 0x47, 0x9a, 0x67, 0x1b, 0x11, 0xbc, 0x5f,
	0x33, 0xb5, 0x71, 0x1e, 0xfe, 0x87, 0x36, 0x35, 0xa0, 0x7f, 0xcf, 0x36, 0x85, 0x8b, 0x96, 0x05,
	0xde, 0xb2, 0xcd, 0xe8, 0x57, 0x03, 0x06, 0xfb, 0x4e, 0x08, 0xcb, 0x52, 0x29, 0x32, 0x86, 0xfa,
	0x50, 0x2f, 0xe3, 0xd6, 0x79, 0x88, 0xc6, 0xd0, 0x0b, 0x72, 0x9d, 0x4f, 0x63, 0xca, 0x05, 0x6e,
	0x58, 0xaa, 0x5b, 0x80, 0x73, 0x83, 0xa1, 0x31, 0x1c, 0xda, 0x14, 0xf1, 0xc1, 0xb0, 0x36, 0xe9,
	0xcc, 0x7a, 0xd3, 0xbc, 0xe8, 0xe9, 0x9d, 0x01, 0x49, 0xce, 0xa1, 0x67, 0xd0, 0xd5, 0x3c, 0xb8,
	0x67, 0xda, 0x4f, 0x15, 0x0f, 0x18, 0x6e, 0xed, 0x6a, 0x17, 0x06, 0x24, 0x9d, 0x5c, 0x62, 0x1f,
	0x68, 0x06, 0xbd, 0x94, 0x6e, 0x56, 0x4c, 0x38, 0x93, 0x76, 0x95, 0x49, 0xb7, 0xd0, 0xe4, 0x36,
	0x2f, 0xa1, 0x1f, 0x7c, 0xa6, 0x2a, 0xe6, 0x22, 0x2e, 0x8c, 0xa0, 0xca, 0xa8, 0xe7, 0x44, 0x65,
	0x24, 0xc5, 0x96, 0x54, 0x33, 0x9f, 0xae, 0xe4, 0x5a, 0x68, 0xdc, 0xa9, 0x8c, 0x94, 0x6b, 0xe6,
	0x56, 0x82, 0x9e, 0x82, 0xa7, 0x58, 0x20, 0x55, 0x98, 0xe1, 0xa3, 0x61, 0x63, 0xd2, 0x99, 0x9d,
	0xee, 0xa8, 0x6f, 0x98, 0xa6, 0x3c, 0x21, 0x4e, 0x83, 0x30, 0x78, 0x89, 0xa4, 0x21, 0x17, 0x31,
	0x3e, 0x1e, 0xd6, 0x26, 0x2d, 0xe2, 0x9e, 0xe8, 0x12, 0x4e, 0x65, 0x14, 0x99, 0x8e, 0xf8, 0xd1,
	0x3a, 0x89, 0x78, 0x92, 0x98, 0x72, 0xf0, 0x89, 0x55, 0xa1, 0x82, 0xba, 0xdd, 0x32, 0xe8, 0x02,
	0x4e, 0xb4, 0xa2, 0x5c, 0xf8, 0x21, 0xcb, 0x02, 0xc5, 0x53, 0xd3, 0x40, 0x8c, 0x6c, 0x5f, 0x8e,
	0x2d, 0x71, 0xb3, 0xc5, 0x47, 0x3f, 0x1a, 0x70, 0x68, 0xfb, 0xb0, 0xd7, 0xda, 0x27, 0xe0, 0x99,
	0x74, 0xbf, 0xd1, 0x0d, 0xae, 0xdb, 0x72, 0x8f, 0x5c, 0x01, 0x24, 0x87, 0x89, 0xe3, 0x1f, 0x4a,
	0xb1, 0xf1, 0x6f, 0x29, 0x1e, 0x54, 0xa7, 0x88, 0xc6, 0x70, 0x10, 0x29, 0xb9, 0xc2, 0x87, 0xbb,
	0x59, 0x7c, 0xd0, 0xf9, 0x68, 0x5a, 0x12, 0x3d, 0x82, 0xba, 0x96, 0xb8, 0x59, 0x2d, 0xa9, 0x6b,
	0x89, 0xce, 0xa1, 0x1d, 0xb2, 0x94, 0x2a, 0xbd, 0x56, 0x0c, 0x7b, 0x36, 0xd4, 0x16, 0x30, 0xbf,
	0x9f, 0x2a, 0xc5, 0xbf, 0xd2, 0xc4, 0x0e, 0x5e, 0x9b, 0xb8, 0x27, 0x3a, 0x86, 0xc6, 0xe2, 0x1d,
	0xb1, 0xb3, 0xd5, 0x26, 0xe6, 0x13, 0x3d, 0x86, 0xa6, 0xcd, 0x31, 0xc3, 0x5d, 0xdb, 0xd8, 0x72,
	0x0c, 0x3e, 0x1a, 0x94, 0x14, 0x24, 0x7a, 0x0e, 0x90, 0xd2, 0x2c, 0x63, 0x22, 0x66, 0x2a, 0xc3,
	0x3d, 0x2b, 0x3d, 0x29, 0x67, 0xc0, 0x31, 0xe4, 0x0f, 0x11, 0x9a, 0x80, 0x97, 0x0f, 0x78, 0x86,
	0xfb, 0x56, 0xdf, 0x2f, 0x5d, 0x5b, 0x98, 0x38, 0x7a, 0xf6, 0xb3, 0x06, 0xa7, 0x15, 0x57, 0x05,
	0xbd, 0x02, 0xaf, 0x00, 0xd0, 0x99, 0xb3, 0xdd, 0x3d, 0x20, 0x03, 0xec, 0xf0, 0x79, 0xb1, 0xf3,
	0xe5, 0x72, 0x4b, 0x38, 0xb7, 0x87, 0xc2, 0x12, 0xfb, 0xde, 0xd1, 0x45, 0xb9, 0xb8, 0x7f, 0xbf,
	0x32, 0x83, 0xd1, 0xc3, 0x62, 0x17, 0xf0, 0xea, 0x16, 0xfa, 0x5c, 0x4e, 0x63, 0x95, 0x06, 0x85,
	0xf8, 0x0a, 0x57, 0x14, 0xb5, 0x30, 0x97, 0x70, 0x51, 0xfb, 0x54, 0xdc, 0xca, 0xef, 0xf5, 0xa3,
	0xd7, 0x33, 0x92, 0x2f, 0x92, 0xd4, 0x72, 0xb9, 0x8e, 0x96, 0x4d, 0x7b, 0x2c, 0x5f, 0xfc, 0x0e,
	0x00, 0x00, 0xff, 0xff, 0x69, 0xe2, 0xf2, 0x16, 0x73, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// OnlineConfirmationsClient is the client API for OnlineConfirmations service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OnlineConfirmationsClient interface {
	// Send a confirm request
	Confirm(ctx context.Context, in *ConfirmRequest, opts ...grpc.CallOption) (*AsyncKeyResponse, error)
	// Service to query async result by using the async key
	QueryAsyncOnlineConfirmation(ctx context.Context, in *OnlineConfirmationAsyncQueryRequest, opts ...grpc.CallOption) (*OnlineConfirmationResponse, error)
}

type onlineConfirmationsClient struct {
	cc *grpc.ClientConn
}

func NewOnlineConfirmationsClient(cc *grpc.ClientConn) OnlineConfirmationsClient {
	return &onlineConfirmationsClient{cc}
}

func (c *onlineConfirmationsClient) Confirm(ctx context.Context, in *ConfirmRequest, opts ...grpc.CallOption) (*AsyncKeyResponse, error) {
	out := new(AsyncKeyResponse)
	err := c.cc.Invoke(ctx, "/g2rail.OnlineConfirmations/Confirm", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *onlineConfirmationsClient) QueryAsyncOnlineConfirmation(ctx context.Context, in *OnlineConfirmationAsyncQueryRequest, opts ...grpc.CallOption) (*OnlineConfirmationResponse, error) {
	out := new(OnlineConfirmationResponse)
	err := c.cc.Invoke(ctx, "/g2rail.OnlineConfirmations/QueryAsyncOnlineConfirmation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OnlineConfirmationsServer is the server API for OnlineConfirmations service.
type OnlineConfirmationsServer interface {
	// Send a confirm request
	Confirm(context.Context, *ConfirmRequest) (*AsyncKeyResponse, error)
	// Service to query async result by using the async key
	QueryAsyncOnlineConfirmation(context.Context, *OnlineConfirmationAsyncQueryRequest) (*OnlineConfirmationResponse, error)
}

func RegisterOnlineConfirmationsServer(s *grpc.Server, srv OnlineConfirmationsServer) {
	s.RegisterService(&_OnlineConfirmations_serviceDesc, srv)
}

func _OnlineConfirmations_Confirm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfirmRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnlineConfirmationsServer).Confirm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/g2rail.OnlineConfirmations/Confirm",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnlineConfirmationsServer).Confirm(ctx, req.(*ConfirmRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OnlineConfirmations_QueryAsyncOnlineConfirmation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OnlineConfirmationAsyncQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnlineConfirmationsServer).QueryAsyncOnlineConfirmation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/g2rail.OnlineConfirmations/QueryAsyncOnlineConfirmation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnlineConfirmationsServer).QueryAsyncOnlineConfirmation(ctx, req.(*OnlineConfirmationAsyncQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _OnlineConfirmations_serviceDesc = grpc.ServiceDesc{
	ServiceName: "g2rail.OnlineConfirmations",
	HandlerType: (*OnlineConfirmationsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Confirm",
			Handler:    _OnlineConfirmations_Confirm_Handler,
		},
		{
			MethodName: "QueryAsyncOnlineConfirmation",
			Handler:    _OnlineConfirmations_QueryAsyncOnlineConfirmation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lib/proto/OnlineConfirmations.proto",
}
