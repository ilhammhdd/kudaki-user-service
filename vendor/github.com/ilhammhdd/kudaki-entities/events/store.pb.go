// Code generated by protoc-gen-go. DO NOT EDIT.
// source: events/store.proto

package events

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	store "github.com/ilhammhdd/kudaki-entities/store"
	user "github.com/ilhammhdd/kudaki-entities/user"
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

type AddStorefrontItemRequested struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Amount               int32    `protobuf:"varint,5,opt,name=amount,proto3" json:"amount,omitempty"`
	Unit                 string   `protobuf:"bytes,6,opt,name=unit,proto3" json:"unit,omitempty"`
	Price                int32    `protobuf:"varint,7,opt,name=price,proto3" json:"price,omitempty"`
	Description          string   `protobuf:"bytes,8,opt,name=description,proto3" json:"description,omitempty"`
	Photo                string   `protobuf:"bytes,9,opt,name=photo,proto3" json:"photo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddStorefrontItemRequested) Reset()         { *m = AddStorefrontItemRequested{} }
func (m *AddStorefrontItemRequested) String() string { return proto.CompactTextString(m) }
func (*AddStorefrontItemRequested) ProtoMessage()    {}
func (*AddStorefrontItemRequested) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f52bba9433e5948, []int{0}
}

func (m *AddStorefrontItemRequested) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddStorefrontItemRequested.Unmarshal(m, b)
}
func (m *AddStorefrontItemRequested) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddStorefrontItemRequested.Marshal(b, m, deterministic)
}
func (m *AddStorefrontItemRequested) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddStorefrontItemRequested.Merge(m, src)
}
func (m *AddStorefrontItemRequested) XXX_Size() int {
	return xxx_messageInfo_AddStorefrontItemRequested.Size(m)
}
func (m *AddStorefrontItemRequested) XXX_DiscardUnknown() {
	xxx_messageInfo_AddStorefrontItemRequested.DiscardUnknown(m)
}

var xxx_messageInfo_AddStorefrontItemRequested proto.InternalMessageInfo

func (m *AddStorefrontItemRequested) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *AddStorefrontItemRequested) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AddStorefrontItemRequested) GetAmount() int32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *AddStorefrontItemRequested) GetUnit() string {
	if m != nil {
		return m.Unit
	}
	return ""
}

func (m *AddStorefrontItemRequested) GetPrice() int32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *AddStorefrontItemRequested) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *AddStorefrontItemRequested) GetPhoto() string {
	if m != nil {
		return m.Photo
	}
	return ""
}

type StorefrontItemAdded struct {
	Uid                        string                      `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Storefront                 *store.Storefront           `protobuf:"bytes,2,opt,name=storefront,proto3" json:"storefront,omitempty"`
	Item                       *store.Item                 `protobuf:"bytes,3,opt,name=item,proto3" json:"item,omitempty"`
	EventStatus                *Status                     `protobuf:"bytes,4,opt,name=event_status,json=eventStatus,proto3" json:"event_status,omitempty"`
	AddStorefrontItemRequested *AddStorefrontItemRequested `protobuf:"bytes,5,opt,name=add_storefront_item_requested,json=addStorefrontItemRequested,proto3" json:"add_storefront_item_requested,omitempty"`
	XXX_NoUnkeyedLiteral       struct{}                    `json:"-"`
	XXX_unrecognized           []byte                      `json:"-"`
	XXX_sizecache              int32                       `json:"-"`
}

func (m *StorefrontItemAdded) Reset()         { *m = StorefrontItemAdded{} }
func (m *StorefrontItemAdded) String() string { return proto.CompactTextString(m) }
func (*StorefrontItemAdded) ProtoMessage()    {}
func (*StorefrontItemAdded) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f52bba9433e5948, []int{1}
}

func (m *StorefrontItemAdded) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StorefrontItemAdded.Unmarshal(m, b)
}
func (m *StorefrontItemAdded) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StorefrontItemAdded.Marshal(b, m, deterministic)
}
func (m *StorefrontItemAdded) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StorefrontItemAdded.Merge(m, src)
}
func (m *StorefrontItemAdded) XXX_Size() int {
	return xxx_messageInfo_StorefrontItemAdded.Size(m)
}
func (m *StorefrontItemAdded) XXX_DiscardUnknown() {
	xxx_messageInfo_StorefrontItemAdded.DiscardUnknown(m)
}

var xxx_messageInfo_StorefrontItemAdded proto.InternalMessageInfo

func (m *StorefrontItemAdded) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *StorefrontItemAdded) GetStorefront() *store.Storefront {
	if m != nil {
		return m.Storefront
	}
	return nil
}

func (m *StorefrontItemAdded) GetItem() *store.Item {
	if m != nil {
		return m.Item
	}
	return nil
}

func (m *StorefrontItemAdded) GetEventStatus() *Status {
	if m != nil {
		return m.EventStatus
	}
	return nil
}

func (m *StorefrontItemAdded) GetAddStorefrontItemRequested() *AddStorefrontItemRequested {
	if m != nil {
		return m.AddStorefrontItemRequested
	}
	return nil
}

type DeleteStorefrontItemRequested struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	ItemUuid             string   `protobuf:"bytes,2,opt,name=item_uuid,json=itemUuid,proto3" json:"item_uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteStorefrontItemRequested) Reset()         { *m = DeleteStorefrontItemRequested{} }
func (m *DeleteStorefrontItemRequested) String() string { return proto.CompactTextString(m) }
func (*DeleteStorefrontItemRequested) ProtoMessage()    {}
func (*DeleteStorefrontItemRequested) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f52bba9433e5948, []int{2}
}

func (m *DeleteStorefrontItemRequested) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteStorefrontItemRequested.Unmarshal(m, b)
}
func (m *DeleteStorefrontItemRequested) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteStorefrontItemRequested.Marshal(b, m, deterministic)
}
func (m *DeleteStorefrontItemRequested) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteStorefrontItemRequested.Merge(m, src)
}
func (m *DeleteStorefrontItemRequested) XXX_Size() int {
	return xxx_messageInfo_DeleteStorefrontItemRequested.Size(m)
}
func (m *DeleteStorefrontItemRequested) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteStorefrontItemRequested.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteStorefrontItemRequested proto.InternalMessageInfo

func (m *DeleteStorefrontItemRequested) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *DeleteStorefrontItemRequested) GetItemUuid() string {
	if m != nil {
		return m.ItemUuid
	}
	return ""
}

type StorefrontItemDeleted struct {
	Uid                           string                         `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Item                          *store.Item                    `protobuf:"bytes,2,opt,name=item,proto3" json:"item,omitempty"`
	EventStatus                   *Status                        `protobuf:"bytes,3,opt,name=event_status,json=eventStatus,proto3" json:"event_status,omitempty"`
	Storefront                    *store.Storefront              `protobuf:"bytes,4,opt,name=storefront,proto3" json:"storefront,omitempty"`
	User                          *user.User                     `protobuf:"bytes,5,opt,name=user,proto3" json:"user,omitempty"`
	DeleteStorefrontItemRequested *DeleteStorefrontItemRequested `protobuf:"bytes,6,opt,name=delete_storefront_item_requested,json=deleteStorefrontItemRequested,proto3" json:"delete_storefront_item_requested,omitempty"`
	XXX_NoUnkeyedLiteral          struct{}                       `json:"-"`
	XXX_unrecognized              []byte                         `json:"-"`
	XXX_sizecache                 int32                          `json:"-"`
}

func (m *StorefrontItemDeleted) Reset()         { *m = StorefrontItemDeleted{} }
func (m *StorefrontItemDeleted) String() string { return proto.CompactTextString(m) }
func (*StorefrontItemDeleted) ProtoMessage()    {}
func (*StorefrontItemDeleted) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f52bba9433e5948, []int{3}
}

func (m *StorefrontItemDeleted) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StorefrontItemDeleted.Unmarshal(m, b)
}
func (m *StorefrontItemDeleted) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StorefrontItemDeleted.Marshal(b, m, deterministic)
}
func (m *StorefrontItemDeleted) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StorefrontItemDeleted.Merge(m, src)
}
func (m *StorefrontItemDeleted) XXX_Size() int {
	return xxx_messageInfo_StorefrontItemDeleted.Size(m)
}
func (m *StorefrontItemDeleted) XXX_DiscardUnknown() {
	xxx_messageInfo_StorefrontItemDeleted.DiscardUnknown(m)
}

var xxx_messageInfo_StorefrontItemDeleted proto.InternalMessageInfo

func (m *StorefrontItemDeleted) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *StorefrontItemDeleted) GetItem() *store.Item {
	if m != nil {
		return m.Item
	}
	return nil
}

func (m *StorefrontItemDeleted) GetEventStatus() *Status {
	if m != nil {
		return m.EventStatus
	}
	return nil
}

func (m *StorefrontItemDeleted) GetStorefront() *store.Storefront {
	if m != nil {
		return m.Storefront
	}
	return nil
}

func (m *StorefrontItemDeleted) GetUser() *user.User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *StorefrontItemDeleted) GetDeleteStorefrontItemRequested() *DeleteStorefrontItemRequested {
	if m != nil {
		return m.DeleteStorefrontItemRequested
	}
	return nil
}

type UpdateStorefrontItemRequested struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	KudakiToken          string   `protobuf:"bytes,2,opt,name=kudaki_token,json=kudakiToken,proto3" json:"kudaki_token,omitempty"`
	Uuid                 string   `protobuf:"bytes,3,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Amount               int32    `protobuf:"varint,5,opt,name=amount,proto3" json:"amount,omitempty"`
	Unit                 string   `protobuf:"bytes,6,opt,name=unit,proto3" json:"unit,omitempty"`
	Price                int32    `protobuf:"varint,7,opt,name=price,proto3" json:"price,omitempty"`
	Description          string   `protobuf:"bytes,8,opt,name=description,proto3" json:"description,omitempty"`
	Photo                string   `protobuf:"bytes,9,opt,name=photo,proto3" json:"photo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateStorefrontItemRequested) Reset()         { *m = UpdateStorefrontItemRequested{} }
func (m *UpdateStorefrontItemRequested) String() string { return proto.CompactTextString(m) }
func (*UpdateStorefrontItemRequested) ProtoMessage()    {}
func (*UpdateStorefrontItemRequested) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f52bba9433e5948, []int{4}
}

func (m *UpdateStorefrontItemRequested) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateStorefrontItemRequested.Unmarshal(m, b)
}
func (m *UpdateStorefrontItemRequested) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateStorefrontItemRequested.Marshal(b, m, deterministic)
}
func (m *UpdateStorefrontItemRequested) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateStorefrontItemRequested.Merge(m, src)
}
func (m *UpdateStorefrontItemRequested) XXX_Size() int {
	return xxx_messageInfo_UpdateStorefrontItemRequested.Size(m)
}
func (m *UpdateStorefrontItemRequested) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateStorefrontItemRequested.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateStorefrontItemRequested proto.InternalMessageInfo

func (m *UpdateStorefrontItemRequested) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *UpdateStorefrontItemRequested) GetKudakiToken() string {
	if m != nil {
		return m.KudakiToken
	}
	return ""
}

func (m *UpdateStorefrontItemRequested) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *UpdateStorefrontItemRequested) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateStorefrontItemRequested) GetAmount() int32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *UpdateStorefrontItemRequested) GetUnit() string {
	if m != nil {
		return m.Unit
	}
	return ""
}

func (m *UpdateStorefrontItemRequested) GetPrice() int32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *UpdateStorefrontItemRequested) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *UpdateStorefrontItemRequested) GetPhoto() string {
	if m != nil {
		return m.Photo
	}
	return ""
}

type StorefrontItemUpdated struct {
	Uid                           string                         `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Item                          *store.Item                    `protobuf:"bytes,2,opt,name=item,proto3" json:"item,omitempty"`
	Storefront                    *store.Storefront              `protobuf:"bytes,3,opt,name=storefront,proto3" json:"storefront,omitempty"`
	EventStatus                   *Status                        `protobuf:"bytes,4,opt,name=event_status,json=eventStatus,proto3" json:"event_status,omitempty"`
	User                          *user.User                     `protobuf:"bytes,5,opt,name=user,proto3" json:"user,omitempty"`
	UpdateStorefrontItemRequested *UpdateStorefrontItemRequested `protobuf:"bytes,6,opt,name=update_storefront_item_requested,json=updateStorefrontItemRequested,proto3" json:"update_storefront_item_requested,omitempty"`
	XXX_NoUnkeyedLiteral          struct{}                       `json:"-"`
	XXX_unrecognized              []byte                         `json:"-"`
	XXX_sizecache                 int32                          `json:"-"`
}

func (m *StorefrontItemUpdated) Reset()         { *m = StorefrontItemUpdated{} }
func (m *StorefrontItemUpdated) String() string { return proto.CompactTextString(m) }
func (*StorefrontItemUpdated) ProtoMessage()    {}
func (*StorefrontItemUpdated) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f52bba9433e5948, []int{5}
}

func (m *StorefrontItemUpdated) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StorefrontItemUpdated.Unmarshal(m, b)
}
func (m *StorefrontItemUpdated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StorefrontItemUpdated.Marshal(b, m, deterministic)
}
func (m *StorefrontItemUpdated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StorefrontItemUpdated.Merge(m, src)
}
func (m *StorefrontItemUpdated) XXX_Size() int {
	return xxx_messageInfo_StorefrontItemUpdated.Size(m)
}
func (m *StorefrontItemUpdated) XXX_DiscardUnknown() {
	xxx_messageInfo_StorefrontItemUpdated.DiscardUnknown(m)
}

var xxx_messageInfo_StorefrontItemUpdated proto.InternalMessageInfo

func (m *StorefrontItemUpdated) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *StorefrontItemUpdated) GetItem() *store.Item {
	if m != nil {
		return m.Item
	}
	return nil
}

func (m *StorefrontItemUpdated) GetStorefront() *store.Storefront {
	if m != nil {
		return m.Storefront
	}
	return nil
}

func (m *StorefrontItemUpdated) GetEventStatus() *Status {
	if m != nil {
		return m.EventStatus
	}
	return nil
}

func (m *StorefrontItemUpdated) GetUser() *user.User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *StorefrontItemUpdated) GetUpdateStorefrontItemRequested() *UpdateStorefrontItemRequested {
	if m != nil {
		return m.UpdateStorefrontItemRequested
	}
	return nil
}

type RetrieveStorefrontItemsRequested struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	KudakiToken          string   `protobuf:"bytes,2,opt,name=kudaki_token,json=kudakiToken,proto3" json:"kudaki_token,omitempty"`
	Offset               int32    `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit                int32    `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RetrieveStorefrontItemsRequested) Reset()         { *m = RetrieveStorefrontItemsRequested{} }
func (m *RetrieveStorefrontItemsRequested) String() string { return proto.CompactTextString(m) }
func (*RetrieveStorefrontItemsRequested) ProtoMessage()    {}
func (*RetrieveStorefrontItemsRequested) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f52bba9433e5948, []int{6}
}

func (m *RetrieveStorefrontItemsRequested) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RetrieveStorefrontItemsRequested.Unmarshal(m, b)
}
func (m *RetrieveStorefrontItemsRequested) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RetrieveStorefrontItemsRequested.Marshal(b, m, deterministic)
}
func (m *RetrieveStorefrontItemsRequested) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RetrieveStorefrontItemsRequested.Merge(m, src)
}
func (m *RetrieveStorefrontItemsRequested) XXX_Size() int {
	return xxx_messageInfo_RetrieveStorefrontItemsRequested.Size(m)
}
func (m *RetrieveStorefrontItemsRequested) XXX_DiscardUnknown() {
	xxx_messageInfo_RetrieveStorefrontItemsRequested.DiscardUnknown(m)
}

var xxx_messageInfo_RetrieveStorefrontItemsRequested proto.InternalMessageInfo

func (m *RetrieveStorefrontItemsRequested) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *RetrieveStorefrontItemsRequested) GetKudakiToken() string {
	if m != nil {
		return m.KudakiToken
	}
	return ""
}

func (m *RetrieveStorefrontItemsRequested) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *RetrieveStorefrontItemsRequested) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type RetrieveItemsRequested struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Offset               int32    `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit                int32    `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	KudakiToken          string   `protobuf:"bytes,4,opt,name=kudaki_token,json=kudakiToken,proto3" json:"kudaki_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RetrieveItemsRequested) Reset()         { *m = RetrieveItemsRequested{} }
func (m *RetrieveItemsRequested) String() string { return proto.CompactTextString(m) }
func (*RetrieveItemsRequested) ProtoMessage()    {}
func (*RetrieveItemsRequested) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f52bba9433e5948, []int{7}
}

func (m *RetrieveItemsRequested) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RetrieveItemsRequested.Unmarshal(m, b)
}
func (m *RetrieveItemsRequested) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RetrieveItemsRequested.Marshal(b, m, deterministic)
}
func (m *RetrieveItemsRequested) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RetrieveItemsRequested.Merge(m, src)
}
func (m *RetrieveItemsRequested) XXX_Size() int {
	return xxx_messageInfo_RetrieveItemsRequested.Size(m)
}
func (m *RetrieveItemsRequested) XXX_DiscardUnknown() {
	xxx_messageInfo_RetrieveItemsRequested.DiscardUnknown(m)
}

var xxx_messageInfo_RetrieveItemsRequested proto.InternalMessageInfo

func (m *RetrieveItemsRequested) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *RetrieveItemsRequested) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *RetrieveItemsRequested) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *RetrieveItemsRequested) GetKudakiToken() string {
	if m != nil {
		return m.KudakiToken
	}
	return ""
}

type RetrieveItemRequested struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	ItemUuid             string   `protobuf:"bytes,2,opt,name=item_uuid,json=itemUuid,proto3" json:"item_uuid,omitempty"`
	KudakiToken          string   `protobuf:"bytes,3,opt,name=kudaki_token,json=kudakiToken,proto3" json:"kudaki_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RetrieveItemRequested) Reset()         { *m = RetrieveItemRequested{} }
func (m *RetrieveItemRequested) String() string { return proto.CompactTextString(m) }
func (*RetrieveItemRequested) ProtoMessage()    {}
func (*RetrieveItemRequested) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f52bba9433e5948, []int{8}
}

func (m *RetrieveItemRequested) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RetrieveItemRequested.Unmarshal(m, b)
}
func (m *RetrieveItemRequested) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RetrieveItemRequested.Marshal(b, m, deterministic)
}
func (m *RetrieveItemRequested) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RetrieveItemRequested.Merge(m, src)
}
func (m *RetrieveItemRequested) XXX_Size() int {
	return xxx_messageInfo_RetrieveItemRequested.Size(m)
}
func (m *RetrieveItemRequested) XXX_DiscardUnknown() {
	xxx_messageInfo_RetrieveItemRequested.DiscardUnknown(m)
}

var xxx_messageInfo_RetrieveItemRequested proto.InternalMessageInfo

func (m *RetrieveItemRequested) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *RetrieveItemRequested) GetItemUuid() string {
	if m != nil {
		return m.ItemUuid
	}
	return ""
}

func (m *RetrieveItemRequested) GetKudakiToken() string {
	if m != nil {
		return m.KudakiToken
	}
	return ""
}

type SearchItemsRequested struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	KudakiToken          string   `protobuf:"bytes,2,opt,name=kudaki_token,json=kudakiToken,proto3" json:"kudaki_token,omitempty"`
	Keyword              string   `protobuf:"bytes,3,opt,name=keyword,proto3" json:"keyword,omitempty"`
	Offset               int32    `protobuf:"varint,4,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit                int32    `protobuf:"varint,5,opt,name=limit,proto3" json:"limit,omitempty"`
	Amount               int32    `protobuf:"varint,6,opt,name=amount,proto3" json:"amount,omitempty"`
	PriceFrom            int32    `protobuf:"varint,7,opt,name=price_from,json=priceFrom,proto3" json:"price_from,omitempty"`
	PriceTo              int32    `protobuf:"varint,8,opt,name=price_to,json=priceTo,proto3" json:"price_to,omitempty"`
	RatingFrom           float32  `protobuf:"fixed32,9,opt,name=rating_from,json=ratingFrom,proto3" json:"rating_from,omitempty"`
	RatingTo             float32  `protobuf:"fixed32,10,opt,name=rating_to,json=ratingTo,proto3" json:"rating_to,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchItemsRequested) Reset()         { *m = SearchItemsRequested{} }
func (m *SearchItemsRequested) String() string { return proto.CompactTextString(m) }
func (*SearchItemsRequested) ProtoMessage()    {}
func (*SearchItemsRequested) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f52bba9433e5948, []int{9}
}

func (m *SearchItemsRequested) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchItemsRequested.Unmarshal(m, b)
}
func (m *SearchItemsRequested) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchItemsRequested.Marshal(b, m, deterministic)
}
func (m *SearchItemsRequested) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchItemsRequested.Merge(m, src)
}
func (m *SearchItemsRequested) XXX_Size() int {
	return xxx_messageInfo_SearchItemsRequested.Size(m)
}
func (m *SearchItemsRequested) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchItemsRequested.DiscardUnknown(m)
}

var xxx_messageInfo_SearchItemsRequested proto.InternalMessageInfo

func (m *SearchItemsRequested) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *SearchItemsRequested) GetKudakiToken() string {
	if m != nil {
		return m.KudakiToken
	}
	return ""
}

func (m *SearchItemsRequested) GetKeyword() string {
	if m != nil {
		return m.Keyword
	}
	return ""
}

func (m *SearchItemsRequested) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *SearchItemsRequested) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *SearchItemsRequested) GetAmount() int32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *SearchItemsRequested) GetPriceFrom() int32 {
	if m != nil {
		return m.PriceFrom
	}
	return 0
}

func (m *SearchItemsRequested) GetPriceTo() int32 {
	if m != nil {
		return m.PriceTo
	}
	return 0
}

func (m *SearchItemsRequested) GetRatingFrom() float32 {
	if m != nil {
		return m.RatingFrom
	}
	return 0
}

func (m *SearchItemsRequested) GetRatingTo() float32 {
	if m != nil {
		return m.RatingTo
	}
	return 0
}

type ItemsSearched struct {
	Uid                  string                `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	User                 *user.User            `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Items                []*store.Item         `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	EventStatus          *Status               `protobuf:"bytes,4,opt,name=event_status,json=eventStatus,proto3" json:"event_status,omitempty"`
	Offset               int32                 `protobuf:"varint,5,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit                int32                 `protobuf:"varint,6,opt,name=limit,proto3" json:"limit,omitempty"`
	SearchItemsRequested *SearchItemsRequested `protobuf:"bytes,7,opt,name=search_items_requested,json=searchItemsRequested,proto3" json:"search_items_requested,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ItemsSearched) Reset()         { *m = ItemsSearched{} }
func (m *ItemsSearched) String() string { return proto.CompactTextString(m) }
func (*ItemsSearched) ProtoMessage()    {}
func (*ItemsSearched) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f52bba9433e5948, []int{10}
}

func (m *ItemsSearched) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ItemsSearched.Unmarshal(m, b)
}
func (m *ItemsSearched) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ItemsSearched.Marshal(b, m, deterministic)
}
func (m *ItemsSearched) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ItemsSearched.Merge(m, src)
}
func (m *ItemsSearched) XXX_Size() int {
	return xxx_messageInfo_ItemsSearched.Size(m)
}
func (m *ItemsSearched) XXX_DiscardUnknown() {
	xxx_messageInfo_ItemsSearched.DiscardUnknown(m)
}

var xxx_messageInfo_ItemsSearched proto.InternalMessageInfo

func (m *ItemsSearched) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ItemsSearched) GetUser() *user.User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *ItemsSearched) GetItems() []*store.Item {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *ItemsSearched) GetEventStatus() *Status {
	if m != nil {
		return m.EventStatus
	}
	return nil
}

func (m *ItemsSearched) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *ItemsSearched) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ItemsSearched) GetSearchItemsRequested() *SearchItemsRequested {
	if m != nil {
		return m.SearchItemsRequested
	}
	return nil
}

func init() {
	proto.RegisterType((*AddStorefrontItemRequested)(nil), "event.AddStorefrontItemRequested")
	proto.RegisterType((*StorefrontItemAdded)(nil), "event.StorefrontItemAdded")
	proto.RegisterType((*DeleteStorefrontItemRequested)(nil), "event.DeleteStorefrontItemRequested")
	proto.RegisterType((*StorefrontItemDeleted)(nil), "event.StorefrontItemDeleted")
	proto.RegisterType((*UpdateStorefrontItemRequested)(nil), "event.UpdateStorefrontItemRequested")
	proto.RegisterType((*StorefrontItemUpdated)(nil), "event.StorefrontItemUpdated")
	proto.RegisterType((*RetrieveStorefrontItemsRequested)(nil), "event.RetrieveStorefrontItemsRequested")
	proto.RegisterType((*RetrieveItemsRequested)(nil), "event.RetrieveItemsRequested")
	proto.RegisterType((*RetrieveItemRequested)(nil), "event.RetrieveItemRequested")
	proto.RegisterType((*SearchItemsRequested)(nil), "event.SearchItemsRequested")
	proto.RegisterType((*ItemsSearched)(nil), "event.ItemsSearched")
}

func init() { proto.RegisterFile("events/store.proto", fileDescriptor_4f52bba9433e5948) }

var fileDescriptor_4f52bba9433e5948 = []byte{
	// 763 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x56, 0xcd, 0x6e, 0xd3, 0x4a,
	0x14, 0x56, 0xec, 0x38, 0x3f, 0x27, 0xad, 0x6e, 0xe5, 0xa6, 0x91, 0x6f, 0xaa, 0xe8, 0xa6, 0xd6,
	0x95, 0x6e, 0x74, 0x51, 0x13, 0x54, 0x76, 0xec, 0x8a, 0x10, 0x12, 0x1b, 0x24, 0xa6, 0xed, 0x86,
	0x4d, 0xe4, 0x66, 0x26, 0xcd, 0x28, 0xb5, 0x27, 0x8c, 0xc7, 0x45, 0x88, 0x35, 0x4f, 0x82, 0x58,
	0xf0, 0x0e, 0xbc, 0x08, 0x2f, 0xc2, 0x1a, 0xcd, 0x19, 0xbb, 0x76, 0x1a, 0x37, 0x6d, 0x5a, 0x16,
	0x6c, 0x22, 0x9f, 0x6f, 0xce, 0xdf, 0x7c, 0xe7, 0x7c, 0x8e, 0xc1, 0x65, 0x57, 0x2c, 0x52, 0xf1,
	0x28, 0x56, 0x42, 0xb2, 0xe1, 0x42, 0x0a, 0x25, 0x5c, 0x07, 0xb1, 0xee, 0x0e, 0x62, 0x23, 0xae,
	0x58, 0x68, 0x0e, 0xba, 0xbb, 0xd7, 0xce, 0x81, 0x4a, 0xe2, 0x14, 0xfc, 0x2b, 0x89, 0x99, 0x1c,
	0xe9, 0x9f, 0x14, 0xe8, 0x98, 0x38, 0xfc, 0x9d, 0x4a, 0x11, 0x29, 0x83, 0xfb, 0xdf, 0x2b, 0xd0,
	0x3d, 0xa6, 0xf4, 0xe4, 0x1a, 0x7f, 0xad, 0x58, 0x48, 0xd8, 0xfb, 0x84, 0xc5, 0x8a, 0x51, 0x77,
	0x07, 0xec, 0x84, 0x53, 0xaf, 0xd2, 0xaf, 0x0c, 0x9a, 0x44, 0x3f, 0xba, 0x2e, 0x54, 0xa3, 0x20,
	0x64, 0x5e, 0x15, 0x21, 0x7c, 0x76, 0x3b, 0x50, 0x0b, 0x42, 0x91, 0x44, 0xca, 0x73, 0xfa, 0x95,
	0x81, 0x43, 0x52, 0x4b, 0xfb, 0x26, 0x11, 0x57, 0x5e, 0xcd, 0xf8, 0xea, 0x67, 0xb7, 0x0d, 0xce,
	0x42, 0xf2, 0x09, 0xf3, 0xea, 0xe8, 0x6a, 0x0c, 0xb7, 0x0f, 0x2d, 0xca, 0xe2, 0x89, 0xe4, 0x0b,
	0xc5, 0x45, 0xe4, 0x35, 0x30, 0xa0, 0x08, 0x61, 0xdc, 0x4c, 0x28, 0xe1, 0x35, 0xf1, 0xcc, 0x18,
	0xfe, 0x57, 0x0b, 0x76, 0x97, 0x7b, 0x3f, 0xa6, 0xb4, 0xb4, 0xef, 0xe7, 0x00, 0xf9, 0xe5, 0x3d,
	0xab, 0x5f, 0x19, 0xb4, 0x8e, 0xba, 0x43, 0x16, 0x29, 0xae, 0x38, 0x8b, 0x87, 0x86, 0xea, 0x3c,
	0x15, 0x29, 0x78, 0xbb, 0x03, 0xa8, 0x6a, 0xc2, 0x3d, 0x1b, 0xa3, 0xda, 0x37, 0xa3, 0x90, 0x32,
	0xf4, 0x70, 0x9f, 0xc2, 0x16, 0x8e, 0x63, 0x6c, 0xa6, 0x81, 0x2c, 0xb5, 0x8e, 0xb6, 0x87, 0x08,
	0x0e, 0x4f, 0x10, 0x24, 0x2d, 0xb4, 0x8c, 0xe1, 0x52, 0xe8, 0x05, 0x94, 0x8e, 0xf3, 0x6a, 0x63,
	0x9d, 0x68, 0x2c, 0xb3, 0x11, 0x20, 0xa5, 0xad, 0xa3, 0x83, 0x34, 0xc5, 0xed, 0xb3, 0x22, 0xdd,
	0xe0, 0xd6, 0x33, 0xff, 0x0d, 0xf4, 0x5e, 0xb2, 0x4b, 0xa6, 0xd8, 0xfd, 0x07, 0xbd, 0x0f, 0x4d,
	0xec, 0x24, 0xd1, 0xb8, 0x85, 0x78, 0x43, 0x03, 0x67, 0x09, 0xa7, 0xfe, 0x0f, 0x0b, 0xf6, 0x96,
	0x53, 0x99, 0xf4, 0x65, 0x89, 0x32, 0xf6, 0xac, 0x8d, 0xd9, 0xb3, 0xef, 0x64, 0x6f, 0x79, 0xaa,
	0xd5, 0x8d, 0xa6, 0xfa, 0x1f, 0x54, 0xb5, 0x40, 0x52, 0x82, 0x77, 0xf3, 0x28, 0x94, 0xcd, 0x59,
	0xcc, 0x24, 0x41, 0x07, 0x37, 0x84, 0x3e, 0xc5, 0xdb, 0xad, 0x99, 0x52, 0x0d, 0x93, 0xfc, 0x9b,
	0xb6, 0xba, 0x96, 0x6b, 0xd2, 0xa3, 0xeb, 0x8e, 0xfd, 0x9f, 0x15, 0xe8, 0x9d, 0x2d, 0x68, 0xb0,
	0xc9, 0xb0, 0x0e, 0x60, 0x6b, 0x9e, 0xd0, 0x60, 0xce, 0xc7, 0x4a, 0xcc, 0x59, 0x94, 0xce, 0xab,
	0x65, 0xb0, 0x53, 0x0d, 0xa1, 0x18, 0x75, 0x94, 0x9d, 0x8a, 0xf1, 0xcf, 0x13, 0xf3, 0xea, 0x52,
	0x19, 0x1a, 0x1e, 0xb7, 0x54, 0xcb, 0x2b, 0x62, 0x6f, 0xb4, 0x22, 0x9b, 0xcb, 0x79, 0x93, 0xa5,
	0x4a, 0xf0, 0x76, 0xf7, 0x5e, 0xaa, 0xb5, 0x3b, 0x41, 0x7a, 0xc9, 0xba, 0x63, 0xff, 0x73, 0x05,
	0xfa, 0x84, 0x29, 0xc9, 0xd9, 0xd5, 0x0d, 0x9f, 0xf8, 0x91, 0x7b, 0xd5, 0x81, 0x9a, 0x98, 0x4e,
	0x63, 0x66, 0xb8, 0x75, 0x48, 0x6a, 0xe9, 0x19, 0x5f, 0xf2, 0x90, 0x1b, 0x55, 0x3a, 0xc4, 0x18,
	0xfe, 0x27, 0xe8, 0x64, 0x6d, 0xdc, 0x59, 0x3c, 0xcf, 0x6c, 0x95, 0x67, 0xb6, 0x0b, 0x99, 0x57,
	0x5a, 0xad, 0xae, 0xb4, 0xea, 0x73, 0xd8, 0x2b, 0x16, 0x7f, 0xe8, 0xdb, 0x6f, 0xa5, 0x94, 0xbd,
	0x5a, 0xea, 0x8b, 0x05, 0xed, 0x13, 0x16, 0xc8, 0xc9, 0xec, 0x77, 0x70, 0xec, 0x41, 0x7d, 0xce,
	0x3e, 0x7e, 0x10, 0x32, 0x93, 0x6f, 0x66, 0x16, 0x38, 0xaa, 0x96, 0x73, 0xe4, 0x14, 0x39, 0xca,
	0xb5, 0x5d, 0x5b, 0xd2, 0x76, 0x0f, 0x00, 0xa5, 0x3b, 0x9e, 0x4a, 0x11, 0xa6, 0x62, 0x6e, 0x22,
	0xf2, 0x4a, 0x8a, 0xd0, 0xfd, 0x1b, 0x1a, 0xe6, 0x58, 0x09, 0x54, 0xb3, 0x43, 0xea, 0x68, 0x9f,
	0x0a, 0xf7, 0x1f, 0x68, 0xc9, 0x40, 0xf1, 0xe8, 0xc2, 0x84, 0x6a, 0x3d, 0x5b, 0x04, 0x0c, 0x84,
	0xb1, 0xfb, 0xd0, 0x4c, 0x1d, 0x94, 0xf0, 0x00, 0x8f, 0x1b, 0x06, 0x38, 0x15, 0xfe, 0x37, 0x0b,
	0xb6, 0x91, 0x1f, 0x43, 0x55, 0x29, 0x3d, 0x99, 0xa2, 0xac, 0xbb, 0x14, 0xf5, 0x3f, 0x38, 0x7a,
	0x42, 0xfa, 0x6f, 0xc3, 0xbe, 0xf5, 0x9d, 0x60, 0x5c, 0x1e, 0x20, 0xec, 0x9c, 0x68, 0xa7, 0x9c,
	0xe8, 0x5a, 0x91, 0xe8, 0xb7, 0xd0, 0x89, 0xf1, 0x4a, 0x28, 0xe9, 0xb8, 0xa0, 0xe9, 0x3a, 0x56,
	0xda, 0xcf, 0x2a, 0x95, 0xac, 0x08, 0x69, 0xc7, 0x25, 0xe8, 0x8b, 0xc3, 0x77, 0x4f, 0x2e, 0xb8,
	0x9a, 0x25, 0xe7, 0xc3, 0x89, 0x08, 0x47, 0xfc, 0x72, 0x16, 0x84, 0xe1, 0x8c, 0xd2, 0x91, 0xd9,
	0x93, 0xc3, 0xec, 0xb2, 0x23, 0xf3, 0x39, 0x78, 0x5e, 0xc3, 0xef, 0xbb, 0x67, 0xbf, 0x02, 0x00,
	0x00, 0xff, 0xff, 0xcd, 0xfd, 0x17, 0xbf, 0x4c, 0x0a, 0x00, 0x00,
}
