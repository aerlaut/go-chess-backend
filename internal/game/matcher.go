package game

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/aerlaut/go-chess-backend/internal/engine"
	"github.com/aerlaut/go-chess-backend/internal/message"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func GenerateGameLink(w http.ResponseWriter, r *http.Request) {

	game := NewGame()
	jsonStr, _ := json.Marshal(game)

	w.Write(jsonStr)
}

var closeError = &websocket.CloseError{}

func ConnectToGame(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println("err:", err)
	}

	ctx := message.Context{Conn: c}

	defer c.Close()

	for {
		mt, m, err := c.ReadMessage()
		if err != nil {
			if errors.As(err, &closeError) {
				log.Println("User disconnected")
				break
			}

			log.Println("err:", err)
			break
		}

		msg, err := message.ParseMessage(string(m))
		if err != nil {
			log.Println("err:", err)
		}

		ctx.MsgType = mt
		err = engine.Execute(&ctx, &msg)

		if err != nil {
			log.Println("err:", err)
		}
	}
}
