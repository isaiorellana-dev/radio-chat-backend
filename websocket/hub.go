package ws

import (
	"fmt"

	m "github.com/isaiorellana-dev/radio-chat-backend/models"
)

type Hub struct {
	clients map[*Client]bool
	// broadcast  chan []byte
	Messages   chan *m.Message
	register   chan *Client
	unregister chan *Client
}

func NewHub() *Hub {
	fmt.Println("hello desde el creador del hub")
	return &Hub{
		clients: make(map[*Client]bool),
		// broadcast:  make(chan []byte),
		Messages:   make(chan *m.Message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (h *Hub) Run() {
	fmt.Println("hello desde el runner del hub")
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		// case message := <-h.broadcast:
		// 	for client := range h.clients {
		// 		select {
		// 		case client.send <- message:
		// 		default:
		// 			close(client.send)
		// 			delete(h.clients, client)
		// 		}
		// 	}
		case message := <-h.Messages:
			// messageBytes, err := json.Marshal(message)
			// if err != nil {
			// 	fmt.Println("Error marshaling message:", err)
			// 	continue
			// }
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}

// func (h *Hub) BroadcastMessage(message *m.Message) {
// 	h.messages <- message
// }
