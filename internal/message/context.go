package message

import "github.com/gorilla/websocket"

type Context struct {
	Conn    *websocket.Conn
	MsgType int
}
