// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: proto/protobuf/Chat.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_protobuf_Chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_protobuf_Chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_proto_protobuf_Chat_proto_rawDescGZIP(), []int{0}
}

type RoomRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomName                 string             `protobuf:"bytes,1,opt,name=roomName,proto3" json:"roomName,omitempty"`
	InitialConnectionRequest *ConnectionRequest `protobuf:"bytes,2,opt,name=initialConnectionRequest,proto3" json:"initialConnectionRequest,omitempty"`
}

func (x *RoomRequest) Reset() {
	*x = RoomRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_protobuf_Chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomRequest) ProtoMessage() {}

func (x *RoomRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_protobuf_Chat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomRequest.ProtoReflect.Descriptor instead.
func (*RoomRequest) Descriptor() ([]byte, []int) {
	return file_proto_protobuf_Chat_proto_rawDescGZIP(), []int{1}
}

func (x *RoomRequest) GetRoomName() string {
	if x != nil {
		return x.RoomName
	}
	return ""
}

func (x *RoomRequest) GetInitialConnectionRequest() *ConnectionRequest {
	if x != nil {
		return x.InitialConnectionRequest
	}
	return nil
}

type ListRoomResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomNames []string `protobuf:"bytes,1,rep,name=roomNames,proto3" json:"roomNames,omitempty"`
}

func (x *ListRoomResponse) Reset() {
	*x = ListRoomResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_protobuf_Chat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRoomResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRoomResponse) ProtoMessage() {}

func (x *ListRoomResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_protobuf_Chat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRoomResponse.ProtoReflect.Descriptor instead.
func (*ListRoomResponse) Descriptor() ([]byte, []int) {
	return file_proto_protobuf_Chat_proto_rawDescGZIP(), []int{2}
}

func (x *ListRoomResponse) GetRoomNames() []string {
	if x != nil {
		return x.RoomNames
	}
	return nil
}

type ConnectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerID string `protobuf:"bytes,1,opt,name=serverID,proto3" json:"serverID,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *ConnectionRequest) Reset() {
	*x = ConnectionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_protobuf_Chat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectionRequest) ProtoMessage() {}

func (x *ConnectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_protobuf_Chat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectionRequest.ProtoReflect.Descriptor instead.
func (*ConnectionRequest) Descriptor() ([]byte, []int) {
	return file_proto_protobuf_Chat_proto_rawDescGZIP(), []int{3}
}

func (x *ConnectionRequest) GetServerID() string {
	if x != nil {
		return x.ServerID
	}
	return ""
}

func (x *ConnectionRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

// ChatMessage represents a message sent from user A to user B
type ChatMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sender    string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`       // the ID of the sender. Typically the username at the specific gateway
	Recipient string `protobuf:"bytes,2,opt,name=recipient,proto3" json:"recipient,omitempty"` // the ID of the recipient
	Content   []byte `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Timestamp uint64 `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *ChatMessage) Reset() {
	*x = ChatMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_protobuf_Chat_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatMessage) ProtoMessage() {}

func (x *ChatMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_protobuf_Chat_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatMessage.ProtoReflect.Descriptor instead.
func (*ChatMessage) Descriptor() ([]byte, []int) {
	return file_proto_protobuf_Chat_proto_rawDescGZIP(), []int{4}
}

func (x *ChatMessage) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

func (x *ChatMessage) GetRecipient() string {
	if x != nil {
		return x.Recipient
	}
	return ""
}

func (x *ChatMessage) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *ChatMessage) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

var File_proto_protobuf_Chat_proto protoreflect.FileDescriptor

var file_proto_protobuf_Chat_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x43, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x07, 0x0a, 0x05, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x79, 0x0a, 0x0b, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x6f, 0x6f, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x6f, 0x6f, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x4e, 0x0a, 0x18, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x18, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x30, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x6f, 0x6f, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x72, 0x6f, 0x6f, 0x6d, 0x4e, 0x61, 0x6d, 0x65,
	0x73, 0x22, 0x4b, 0x0a, 0x11, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x7b,
	0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69,
	0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x32, 0xb3, 0x01, 0x0a, 0x0b,
	0x43, 0x68, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x23, 0x0a, 0x0b, 0x53,
	0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0c, 0x2e, 0x43, 0x68, 0x61,
	0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x12, 0x29, 0x0a, 0x09, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x0c, 0x2e,
	0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x43, 0x68,
	0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x30, 0x01, 0x12, 0x2c, 0x0a, 0x0e, 0x55,
	0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x41, 0x6c, 0x6c, 0x12, 0x12, 0x2e,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x26, 0x0a, 0x09, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x73, 0x12, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x11,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_protobuf_Chat_proto_rawDescOnce sync.Once
	file_proto_protobuf_Chat_proto_rawDescData = file_proto_protobuf_Chat_proto_rawDesc
)

func file_proto_protobuf_Chat_proto_rawDescGZIP() []byte {
	file_proto_protobuf_Chat_proto_rawDescOnce.Do(func() {
		file_proto_protobuf_Chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_protobuf_Chat_proto_rawDescData)
	})
	return file_proto_protobuf_Chat_proto_rawDescData
}

var file_proto_protobuf_Chat_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_protobuf_Chat_proto_goTypes = []interface{}{
	(*Empty)(nil),             // 0: Empty
	(*RoomRequest)(nil),       // 1: RoomRequest
	(*ListRoomResponse)(nil),  // 2: ListRoomResponse
	(*ConnectionRequest)(nil), // 3: ConnectionRequest
	(*ChatMessage)(nil),       // 4: ChatMessage
}
var file_proto_protobuf_Chat_proto_depIdxs = []int32{
	3, // 0: RoomRequest.initialConnectionRequest:type_name -> ConnectionRequest
	4, // 1: ChatService.SendMessage:input_type -> ChatMessage
	1, // 2: ChatService.Subscribe:input_type -> RoomRequest
	3, // 3: ChatService.UnsubscribeAll:input_type -> ConnectionRequest
	0, // 4: ChatService.ListRooms:input_type -> Empty
	0, // 5: ChatService.SendMessage:output_type -> Empty
	4, // 6: ChatService.Subscribe:output_type -> ChatMessage
	0, // 7: ChatService.UnsubscribeAll:output_type -> Empty
	2, // 8: ChatService.ListRooms:output_type -> ListRoomResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_protobuf_Chat_proto_init() }
func file_proto_protobuf_Chat_proto_init() {
	if File_proto_protobuf_Chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_protobuf_Chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_proto_protobuf_Chat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomRequest); i {
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
		file_proto_protobuf_Chat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRoomResponse); i {
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
		file_proto_protobuf_Chat_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectionRequest); i {
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
		file_proto_protobuf_Chat_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatMessage); i {
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
			RawDescriptor: file_proto_protobuf_Chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_protobuf_Chat_proto_goTypes,
		DependencyIndexes: file_proto_protobuf_Chat_proto_depIdxs,
		MessageInfos:      file_proto_protobuf_Chat_proto_msgTypes,
	}.Build()
	File_proto_protobuf_Chat_proto = out.File
	file_proto_protobuf_Chat_proto_rawDesc = nil
	file_proto_protobuf_Chat_proto_goTypes = nil
	file_proto_protobuf_Chat_proto_depIdxs = nil
}
