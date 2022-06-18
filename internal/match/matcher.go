package match

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func GenerateMatchLink(w http.ResponseWriter, r *http.Request) {

	match := NewMatch()
	jsonStr, _ := json.Marshal(match)

	w.Write(jsonStr)
}

func ConnectToMatch(w http.ResponseWriter, r *http.Request) {

	matchId := chi.URLParam(r, "matchId")

	log.Println("Receiving connection on match:", matchId)

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	defer c.Close()

	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
		}
		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}
