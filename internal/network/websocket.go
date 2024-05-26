package network

import (
	"github.com/369guang/tg-im/internal/logs"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"log"
	"net/http"
	"time"
)

const (
	writeWait      = 10 * time.Second    // Time allowed to write a message to the peer.
	pongWait       = 60 * time.Second    // Time allowed to read the next pong message from the peer.
	pingPeriod     = (pongWait * 9) / 10 // Send pings to peer with this period. Must be less than pongWait.
	maxMessageSize = 512                 // Maximum message size allowed from peer.
)

// WebSocketServer 封装了 WebSocket 服务器
type WebSocketServer struct {
	UpGrader websocket.Upgrader
}

// WebSocketClient 封装了 WebSocket 客户端
type WebSocketClient struct {
	Conn *websocket.Conn
	Send chan []byte
}

// NewWebSocketServer 创建一个 WebSocket 服务器
func NewWebSocketServer() *WebSocketServer {
	return &WebSocketServer{
		UpGrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	}
}

// ServerHttpHandler 用于处理 WebSocket 请求
func (ws *WebSocketServer) ServerHttpHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := ws.UpGrader.Upgrade(w, r, nil)
	if err != nil {
		logs.Logger.Error("network upgrade:", zap.Error(err))
		log.Printf("network upgrade error: %v\n", err)
		return
	}

	client := &WebSocketClient{
		Conn: conn,
		Send: make(chan []byte, 256), // 缓冲区大小为 256
	}

	go client.writePump()
	go client.readPump()
}

// readPump 读取客户端发送的消息
func (c *WebSocketClient) readPump() {
	defer func() {
		c.Conn.Close()
	}()

	c.Conn.SetReadLimit(maxMessageSize)
	c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	c.Conn.SetPongHandler(func(string) error {
		c.Conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		_, _, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				logs.Logger.Error("network read:", zap.Error(err))
				log.Printf("network read error: %v\n", err)
			}
			break
		}
		logs.Logger.Info("network read:", zap.String("message", string("")))
		// 这里处理收到的消息，可以将其发送到其他客户端
	}
}

// writePump 处理向 WebSocket 连接发送消息
func (c *WebSocketClient) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// 发送通道关闭
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// Add queued chat messages to the current websocket message.
			n := len(c.Send)
			for i := 0; i < n; i++ {
				w.Write(<-c.Send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
