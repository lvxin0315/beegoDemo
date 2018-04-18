package controllers

import (
	"github.com/gorilla/websocket"
	"github.com/astaxie/beego/logs"
	"beegoDemo/models"
)

type WSController struct {
	BaseController
}

func (c *WSController) Get()  {
	username := c.GetUsername(c.Ctx.ResponseWriter, c.Ctx.Request)
	logs.Info("username:",username)
	// Upgrade from http request to WebSocket.
	ws, err := websocket.Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, nil, 1024, 1024)
	if _, ok := err.(websocket.HandshakeError); ok {
		logs.Error(c.Ctx.ResponseWriter, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		logs.Error("Cannot setup WebSocket connection:", err)
		return
	}

	// Join chat room.
	Join(username, ws)
	defer Leave(username)

	// Message receive loop.
	for {
		_, p, err := ws.ReadMessage()
		logs.Info("p:",p)
		if err != nil {
			return
		}
		publish <- newEvent(models.EVENT_MESSAGE, username, string(p))
	}
}

// broadcastWebSocket broadcasts messages to WebSocket users.
func broadcastWebSocket(event models.Event) {
	//data, err := json.Marshal(event)
	//if err != nil {
	//	beego.Error("Fail to marshal event:", err)
	//	return
	//}

	for sub := subscribers.Front(); sub != nil; sub = sub.Next() {
		// Immediately send event to WebSocket users.
		ws := sub.Value.(Subscriber).Conn
		name := sub.Value.(Subscriber).Name
		if name == event.User {
			continue
		}
		if ws != nil {
			logs.Info("broadcastWebSocket:",string(event.Content))
			if ws.WriteMessage(websocket.TextMessage, []byte(event.Content)) != nil {
				// User disconnected.
				logs.Error("WriteMessage error")
				unsubscribe <- sub.Value.(Subscriber).Name
			}
		}
	}
}
