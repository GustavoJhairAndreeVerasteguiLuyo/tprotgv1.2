package ws

import (
    "log"
    "net/http"

    "github.com/gorilla/websocket"
)

type Hub struct {
    clients   map[*websocket.Conn]bool
    broadcast chan []byte
}

func NewHub() *Hub { return &Hub{clients: make(map[*websocket.Conn]bool), broadcast: make(chan []byte)} }

func (h *Hub) Run() {
    for msg := range h.broadcast {
        for c := range h.clients {
            if err := c.WriteMessage(websocket.TextMessage, msg); err != nil {
                c.Close()
                delete(h.clients, c)
            }
        }
    }
}

var upgrader = websocket.Upgrader{ CheckOrigin: func(r *http.Request) bool { return true } }

func (h *Hub) ServeWS(w http.ResponseWriter, r *http.Request) {
    c, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Print("ws upgrade: ", err)
        return
    }
    h.clients[c] = true
}

func (h *Hub) Publish(msg []byte) { h.broadcast <- msg }
