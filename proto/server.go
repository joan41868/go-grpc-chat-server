package proto

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

type ClientConnection struct {
	clientID string
	stream   ChatService_ConnectServer
	active   bool
	errChan  chan error
}

type Room struct {
	connections []*ClientConnection
	errChan     chan error
	wg          *sync.WaitGroup
}

func NewRoom() *Room {
	return &Room{
		connections: []*ClientConnection{},
		errChan:     make(chan error),
		wg:          new(sync.WaitGroup),
	}
}

func (r *Room) BroadcastMessage(msg *ChatMessage) {
	for _, connection := range r.connections {
		intermediary := connection
		go func() {
			err := intermediary.stream.Send(msg)
			if err != nil {
				intermediary.active = false
				intermediary.errChan <- errors.New("disconnected")
			}
		}()
	}
}

func (r *Room) performCleanup() {
	r.wg.Add(1)
	for {
		for i, connection := range r.connections {
			if !connection.active {
				r.connections = append(r.connections[:i], r.connections[i:]...)
			}
		}
		if len(r.connections) == 0 {
			r.errChan <- errors.New("room empty")
			break
		}
		time.Sleep(time.Second * 10)
	}
	r.wg.Done()
}

type Server struct {
	roomsMap map[string]*Room
}

func (s Server) ListRooms(ctx context.Context, empty *Empty) (*ListRoomResponse, error) {
	lrr := &ListRoomResponse{
		RoomIDs: []string{},
	}
	for k := range s.roomsMap {
		lrr.RoomIDs = append(lrr.RoomIDs, k)
	}
	return lrr, nil
}

func (s Server) Subscribe(request *RoomRequest, server ChatService_SubscribeServer) error {
	roomID := request.RoomID
	var room *Room
	if _, exists := s.roomsMap[roomID]; !exists {
		room = NewRoom()
		room.connections = append(room.connections, &ClientConnection{
			clientID: request.InitialConnectionRequest.ServerId,
			stream:   server,
			active:   true,
			errChan:  make(chan error),
		})
		s.roomsMap[roomID] = room
		go room.performCleanup()
	} else {
		room = s.roomsMap[roomID]
		for _, conn := range room.connections {
			if conn.clientID == request.InitialConnectionRequest.ServerId {
				return errors.New("user with the same name is already in the room")
			}
		}
		s.roomsMap[roomID].connections = append(s.roomsMap[roomID].connections, &ClientConnection{
			clientID: request.InitialConnectionRequest.ServerId,
			stream:   server,
			active:   true,
			errChan:  make(chan error),
		})
		room.BroadcastMessage(&ChatMessage{
			SenderID:       "GATEWAY",
			RecipientID:    "",
			Content:        []byte(fmt.Sprintf("%s joined.", request.InitialConnectionRequest.ServerId)),
			Timestamp:      uint64(time.Now().Unix()),
			SenderUsername: "GATEWAY",
		})
	}
	return <-room.errChan
}

func (s Server) UnsubscribeAll(ctx context.Context, request *ConnectionRequest) (*Empty, error) {
	for _, v := range s.roomsMap {
		for i, conn := range v.connections {
			conn.errChan <- errors.New("disconnected")
			v.connections = append(v.connections[:i], v.connections[i+1:]...)
		}
	}
	return &Empty{}, nil
}

func (s Server) Disconnect(ctx context.Context, request *ConnectionRequest) (*Empty, error) {
	panic("Deprecated")
}

func (s Server) Connect(request *ConnectionRequest, server ChatService_ConnectServer) error {
	panic("Deprecated")
}

func (s Server) SendMessage(ctx context.Context, message *ChatMessage) (*Empty, error) {
	forClient := message.GetRecipientID()
	for k, room := range s.roomsMap {
		if k == forClient {
			room.BroadcastMessage(message)
		}
	}
	return &Empty{}, nil
}

func (s Server) mustEmbedUnimplementedChatServiceServer() {
	panic("implement me")
}

var EMPTY_MESSAGE = ChatMessage{
	Timestamp: 0,
}

func NewChatServer() ChatServiceServer {
	s := &Server{
		roomsMap: map[string]*Room{},
	}
	return s
}
