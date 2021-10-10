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
	stream   ChatService_SubscribeServer
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
		time.Sleep(time.Second * 10)
	}
	r.wg.Done()

}

type Server struct {
	roomsMap map[string]*Room
}

func (s Server) ListRooms(ctx context.Context, empty *Empty) (*ListRoomResponse, error) {
	lrr := &ListRoomResponse{
		RoomNames: []string{},
	}
	for k := range s.roomsMap {
		lrr.RoomNames = append(lrr.RoomNames, k)
	}
	return lrr, nil
}

func (s Server) Subscribe(request *RoomRequest, server ChatService_SubscribeServer) error {
	roomID := request.RoomName
	var room *Room
	if _, exists := s.roomsMap[roomID]; !exists {
		room = NewRoom()
		room.connections = append(room.connections, &ClientConnection{
			clientID: request.GetInitialConnectionRequest().GetServerID(),
			stream:   server,
			active:   true,
			errChan:  make(chan error),
		})
		s.roomsMap[roomID] = room
		go room.performCleanup()
	} else {
		room = s.roomsMap[roomID]
		for _, conn := range room.connections {
			if conn.clientID == request.GetInitialConnectionRequest().GetServerID() {
				return errors.New("user with the same name is already in the room")
			}
		}
		s.roomsMap[roomID].connections = append(s.roomsMap[roomID].connections, &ClientConnection{
			clientID: request.GetInitialConnectionRequest().GetServerID(),
			stream:   server,
			active:   true,
			errChan:  make(chan error),
		})
		room.BroadcastMessage(&ChatMessage{
			Sender:    "GATEWAY",
			Recipient: "",
			Content:   []byte(fmt.Sprintf("%s joined.", request.GetInitialConnectionRequest().GetServerID())),
			Timestamp: uint64(time.Now().Unix()),
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

func (s Server) SendMessage(ctx context.Context, message *ChatMessage) (*Empty, error) {
	forClient := message.GetRecipient()
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

func (s Server) performRoomCleanup() {
	for {
		for k, room := range s.roomsMap {
			if len(room.connections) == 0 {
				delete(s.roomsMap, k)
			}
		}
		time.Sleep(time.Second * 20)
	}
}

func NewChatServer() ChatServiceServer {
	s := &Server{
		roomsMap: map[string]*Room{},
	}
	go s.performRoomCleanup()
	return s
}
