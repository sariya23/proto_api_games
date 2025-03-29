// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.21.12
// source: game/game.proto

package gamev1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	date "google.golang.org/genproto/googleapis/type/date"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Game struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Title         string                 `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Genres        []string               `protobuf:"bytes,2,rep,name=genres,proto3" json:"genres,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	ReleaseYear   *date.Date             `protobuf:"bytes,4,opt,name=release_year,json=releaseYear,proto3" json:"release_year,omitempty"`
	CoverImage    []byte                 `protobuf:"bytes,5,opt,name=cover_image,json=coverImage,proto3" json:"cover_image,omitempty"`
	Tags          []string               `protobuf:"bytes,6,rep,name=tags,proto3" json:"tags,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Game) Reset() {
	*x = Game{}
	mi := &file_game_game_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Game) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Game) ProtoMessage() {}

func (x *Game) ProtoReflect() protoreflect.Message {
	mi := &file_game_game_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Game.ProtoReflect.Descriptor instead.
func (*Game) Descriptor() ([]byte, []int) {
	return file_game_game_proto_rawDescGZIP(), []int{0}
}

func (x *Game) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Game) GetGenres() []string {
	if x != nil {
		return x.Genres
	}
	return nil
}

func (x *Game) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Game) GetReleaseYear() *date.Date {
	if x != nil {
		return x.ReleaseYear
	}
	return nil
}

func (x *Game) GetCoverImage() []byte {
	if x != nil {
		return x.CoverImage
	}
	return nil
}

func (x *Game) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

type GameWithRating struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Game          *Game                  `protobuf:"bytes,1,opt,name=game,proto3" json:"game,omitempty"`
	Rating        float32                `protobuf:"fixed32,2,opt,name=rating,proto3" json:"rating,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GameWithRating) Reset() {
	*x = GameWithRating{}
	mi := &file_game_game_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GameWithRating) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameWithRating) ProtoMessage() {}

func (x *GameWithRating) ProtoReflect() protoreflect.Message {
	mi := &file_game_game_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameWithRating.ProtoReflect.Descriptor instead.
func (*GameWithRating) Descriptor() ([]byte, []int) {
	return file_game_game_proto_rawDescGZIP(), []int{1}
}

func (x *GameWithRating) GetGame() *Game {
	if x != nil {
		return x.Game
	}
	return nil
}

func (x *GameWithRating) GetRating() float32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

type AddGameRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Game          *Game                  `protobuf:"bytes,1,opt,name=game,proto3" json:"game,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddGameRequest) Reset() {
	*x = AddGameRequest{}
	mi := &file_game_game_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddGameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddGameRequest) ProtoMessage() {}

func (x *AddGameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_game_game_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddGameRequest.ProtoReflect.Descriptor instead.
func (*AddGameRequest) Descriptor() ([]byte, []int) {
	return file_game_game_proto_rawDescGZIP(), []int{2}
}

func (x *AddGameRequest) GetGame() *Game {
	if x != nil {
		return x.Game
	}
	return nil
}

type AddGameResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Game          *Game                  `protobuf:"bytes,1,opt,name=game,proto3" json:"game,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddGameResponse) Reset() {
	*x = AddGameResponse{}
	mi := &file_game_game_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddGameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddGameResponse) ProtoMessage() {}

func (x *AddGameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_game_game_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddGameResponse.ProtoReflect.Descriptor instead.
func (*AddGameResponse) Descriptor() ([]byte, []int) {
	return file_game_game_proto_rawDescGZIP(), []int{3}
}

func (x *AddGameResponse) GetGame() *Game {
	if x != nil {
		return x.Game
	}
	return nil
}

func (x *AddGameResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetGameRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	GameId        uint64                 `protobuf:"varint,1,opt,name=game_id,json=gameId,proto3" json:"game_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetGameRequest) Reset() {
	*x = GetGameRequest{}
	mi := &file_game_game_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetGameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGameRequest) ProtoMessage() {}

func (x *GetGameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_game_game_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGameRequest.ProtoReflect.Descriptor instead.
func (*GetGameRequest) Descriptor() ([]byte, []int) {
	return file_game_game_proto_rawDescGZIP(), []int{4}
}

func (x *GetGameRequest) GetGameId() uint64 {
	if x != nil {
		return x.GameId
	}
	return 0
}

type GetGameResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Game          *GameWithRating        `protobuf:"bytes,1,opt,name=game,proto3" json:"game,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetGameResponse) Reset() {
	*x = GetGameResponse{}
	mi := &file_game_game_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetGameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGameResponse) ProtoMessage() {}

func (x *GetGameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_game_game_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGameResponse.ProtoReflect.Descriptor instead.
func (*GetGameResponse) Descriptor() ([]byte, []int) {
	return file_game_game_proto_rawDescGZIP(), []int{5}
}

func (x *GetGameResponse) GetGame() *GameWithRating {
	if x != nil {
		return x.Game
	}
	return nil
}

type GetTopGamesRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Year          string                 `protobuf:"bytes,1,opt,name=year,proto3" json:"year,omitempty"`
	Limit         uint32                 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Genres        []string               `protobuf:"bytes,3,rep,name=genres,proto3" json:"genres,omitempty"`
	Tags          []string               `protobuf:"bytes,4,rep,name=tags,proto3" json:"tags,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetTopGamesRequest) Reset() {
	*x = GetTopGamesRequest{}
	mi := &file_game_game_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTopGamesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTopGamesRequest) ProtoMessage() {}

func (x *GetTopGamesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_game_game_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTopGamesRequest.ProtoReflect.Descriptor instead.
func (*GetTopGamesRequest) Descriptor() ([]byte, []int) {
	return file_game_game_proto_rawDescGZIP(), []int{6}
}

func (x *GetTopGamesRequest) GetYear() string {
	if x != nil {
		return x.Year
	}
	return ""
}

func (x *GetTopGamesRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetTopGamesRequest) GetGenres() []string {
	if x != nil {
		return x.Genres
	}
	return nil
}

func (x *GetTopGamesRequest) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

type GetTopGamesResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Games         []*GameWithRating      `protobuf:"bytes,1,rep,name=games,proto3" json:"games,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetTopGamesResponse) Reset() {
	*x = GetTopGamesResponse{}
	mi := &file_game_game_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTopGamesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTopGamesResponse) ProtoMessage() {}

func (x *GetTopGamesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_game_game_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTopGamesResponse.ProtoReflect.Descriptor instead.
func (*GetTopGamesResponse) Descriptor() ([]byte, []int) {
	return file_game_game_proto_rawDescGZIP(), []int{7}
}

func (x *GetTopGamesResponse) GetGames() []*GameWithRating {
	if x != nil {
		return x.Games
	}
	return nil
}

type DeleteGameRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	GameId        uint64                 `protobuf:"varint,1,opt,name=game_id,json=gameId,proto3" json:"game_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteGameRequest) Reset() {
	*x = DeleteGameRequest{}
	mi := &file_game_game_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteGameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteGameRequest) ProtoMessage() {}

func (x *DeleteGameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_game_game_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteGameRequest.ProtoReflect.Descriptor instead.
func (*DeleteGameRequest) Descriptor() ([]byte, []int) {
	return file_game_game_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteGameRequest) GetGameId() uint64 {
	if x != nil {
		return x.GameId
	}
	return 0
}

type DeleteGameResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Game          *Game                  `protobuf:"bytes,1,opt,name=game,proto3" json:"game,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteGameResponse) Reset() {
	*x = DeleteGameResponse{}
	mi := &file_game_game_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteGameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteGameResponse) ProtoMessage() {}

func (x *DeleteGameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_game_game_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteGameResponse.ProtoReflect.Descriptor instead.
func (*DeleteGameResponse) Descriptor() ([]byte, []int) {
	return file_game_game_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteGameResponse) GetGame() *Game {
	if x != nil {
		return x.Game
	}
	return nil
}

var File_game_game_proto protoreflect.FileDescriptor

var file_game_game_proto_rawDesc = string([]byte{
	0x0a, 0x0f, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x61, 0x75, 0x74, 0x68, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc1, 0x01, 0x0a,
	0x04, 0x47, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x67,
	0x65, 0x6e, 0x72, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6e,
	0x72, 0x65, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x0c, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x5f, 0x79, 0x65, 0x61, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x52, 0x0b,
	0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x59, 0x65, 0x61, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x0a, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x61, 0x67, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73,
	0x22, 0x48, 0x0a, 0x0e, 0x47, 0x61, 0x6d, 0x65, 0x57, 0x69, 0x74, 0x68, 0x52, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x12, 0x1e, 0x0a, 0x04, 0x67, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0a, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x04, 0x67, 0x61,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x30, 0x0a, 0x0e, 0x41, 0x64,
	0x64, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x04,
	0x67, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x04, 0x67, 0x61, 0x6d, 0x65, 0x22, 0x4b, 0x0a, 0x0f,
	0x41, 0x64, 0x64, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1e, 0x0a, 0x04, 0x67, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x04, 0x67, 0x61, 0x6d, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x29, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x67,
	0x61, 0x6d, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x67, 0x61,
	0x6d, 0x65, 0x49, 0x64, 0x22, 0x3b, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x67, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x47, 0x61, 0x6d,
	0x65, 0x57, 0x69, 0x74, 0x68, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x04, 0x67, 0x61, 0x6d,
	0x65, 0x22, 0x6a, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x47, 0x61, 0x6d, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x22, 0x41, 0x0a,
	0x13, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x05, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x57,
	0x69, 0x74, 0x68, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x05, 0x67, 0x61, 0x6d, 0x65, 0x73,
	0x22, 0x2c, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x67, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x22, 0x34,
	0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x04, 0x67, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x04,
	0x67, 0x61, 0x6d, 0x65, 0x32, 0xf4, 0x02, 0x0a, 0x0b, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x4f, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x47, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x41, 0x64, 0x64, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x41, 0x64, 0x64,
	0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x11, 0x3a, 0x01, 0x2a, 0x22, 0x0c, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x61, 0x6d,
	0x65, 0x2f, 0x61, 0x64, 0x64, 0x12, 0x56, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x47, 0x61, 0x6d, 0x65,
	0x12, 0x14, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x47, 0x65,
	0x74, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x18, 0x12, 0x16, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2f,
	0x67, 0x65, 0x74, 0x2f, 0x7b, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x58, 0x0a,
	0x0b, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x18, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x47, 0x65,
	0x74, 0x54, 0x6f, 0x70, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x12, 0x0c, 0x2f, 0x76, 0x31, 0x2f, 0x67,
	0x61, 0x6d, 0x65, 0x2f, 0x74, 0x6f, 0x70, 0x12, 0x62, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x47, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x61, 0x6d, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b,
	0x2a, 0x19, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x2f, 0x7b, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x35, 0x5a, 0x33, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x61, 0x72, 0x69, 0x79, 0x61,
	0x32, 0x33, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x67, 0x61, 0x6d,
	0x65, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x3b, 0x67, 0x61, 0x6d, 0x65,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_game_game_proto_rawDescOnce sync.Once
	file_game_game_proto_rawDescData []byte
)

func file_game_game_proto_rawDescGZIP() []byte {
	file_game_game_proto_rawDescOnce.Do(func() {
		file_game_game_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_game_game_proto_rawDesc), len(file_game_game_proto_rawDesc)))
	})
	return file_game_game_proto_rawDescData
}

var file_game_game_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_game_game_proto_goTypes = []any{
	(*Game)(nil),                // 0: auth.Game
	(*GameWithRating)(nil),      // 1: auth.GameWithRating
	(*AddGameRequest)(nil),      // 2: auth.AddGameRequest
	(*AddGameResponse)(nil),     // 3: auth.AddGameResponse
	(*GetGameRequest)(nil),      // 4: auth.GetGameRequest
	(*GetGameResponse)(nil),     // 5: auth.GetGameResponse
	(*GetTopGamesRequest)(nil),  // 6: auth.GetTopGamesRequest
	(*GetTopGamesResponse)(nil), // 7: auth.GetTopGamesResponse
	(*DeleteGameRequest)(nil),   // 8: auth.DeleteGameRequest
	(*DeleteGameResponse)(nil),  // 9: auth.DeleteGameResponse
	(*date.Date)(nil),           // 10: google.type.Date
}
var file_game_game_proto_depIdxs = []int32{
	10, // 0: auth.Game.release_year:type_name -> google.type.Date
	0,  // 1: auth.GameWithRating.game:type_name -> auth.Game
	0,  // 2: auth.AddGameRequest.game:type_name -> auth.Game
	0,  // 3: auth.AddGameResponse.game:type_name -> auth.Game
	1,  // 4: auth.GetGameResponse.game:type_name -> auth.GameWithRating
	1,  // 5: auth.GetTopGamesResponse.games:type_name -> auth.GameWithRating
	0,  // 6: auth.DeleteGameResponse.game:type_name -> auth.Game
	2,  // 7: auth.GameService.AddGame:input_type -> auth.AddGameRequest
	4,  // 8: auth.GameService.GetGame:input_type -> auth.GetGameRequest
	6,  // 9: auth.GameService.GetTopGames:input_type -> auth.GetTopGamesRequest
	8,  // 10: auth.GameService.DeleteGame:input_type -> auth.DeleteGameRequest
	3,  // 11: auth.GameService.AddGame:output_type -> auth.AddGameResponse
	5,  // 12: auth.GameService.GetGame:output_type -> auth.GetGameResponse
	7,  // 13: auth.GameService.GetTopGames:output_type -> auth.GetTopGamesResponse
	9,  // 14: auth.GameService.DeleteGame:output_type -> auth.DeleteGameResponse
	11, // [11:15] is the sub-list for method output_type
	7,  // [7:11] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_game_game_proto_init() }
func file_game_game_proto_init() {
	if File_game_game_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_game_game_proto_rawDesc), len(file_game_game_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_game_game_proto_goTypes,
		DependencyIndexes: file_game_game_proto_depIdxs,
		MessageInfos:      file_game_game_proto_msgTypes,
	}.Build()
	File_game_game_proto = out.File
	file_game_game_proto_goTypes = nil
	file_game_game_proto_depIdxs = nil
}
