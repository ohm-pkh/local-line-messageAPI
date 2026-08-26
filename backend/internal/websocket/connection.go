package ws

import (
	"errors"
	"sync"

	"github.com/gofiber/contrib/v3/websocket"
)

type Connection struct {
	conn *websocket.Conn
	mu   sync.Mutex
}

func NewConnection() *Connection {
	return &Connection{}
}

func (w *Connection) Set(conn *websocket.Conn) {
	w.mu.Lock()
	defer w.mu.Unlock()

	w.conn = conn
}

func (w *Connection) Clear() {
	w.mu.Lock()
	defer w.mu.Unlock()

	w.conn = nil
}

func (w *Connection) SendJSON(v any) error {
	w.mu.Lock()
	defer w.mu.Unlock()

	if w.conn == nil {
		return errors.New("websocket is not connected")
	}

	return w.conn.WriteJSON(v)
}
