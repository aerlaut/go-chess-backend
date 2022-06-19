package player

import (
	"errors"
	"log"

	"github.com/aerlaut/go-chess-backend/internal/message"
	"github.com/gorilla/websocket"
)

type player struct {
	Id         string          `json:"id"`
	Connection *websocket.Conn `json:"-"`
}

var playersList = make(map[string][]player)

func newPlayer(id string, c *websocket.Conn) player {
	return player{
		Id:         id,
		Connection: c,
	}
}

func Join(ctx *message.Context, m *message.Message) error {

	players, prs := playersList[m.GameId]

	// Initialize players list if the game doesn't have any player yet
	if !prs {
		playersList[m.GameId] = []player{}
		log.Printf("[GAME] - Game %s initialized", m.GameId)
	}

	// Check if the player exists
	playerExists := false
	for _, player := range players {
		if player.Id == m.PlayerId {
			playerExists = true
			break
		}
	}

	if !playerExists {
		playersList[m.GameId] = append(players, newPlayer(m.PlayerId, ctx.Conn))
		log.Printf("[GAME %s] - New user registered: Player %s added to the game", m.GameId, m.PlayerId)
		log.Printf("[GAME %s] - Players in game: %d", m.GameId, len(playersList[m.GameId]))
	}

	return nil
}

func Leave(ctx *message.Context, m *message.Message) error {

	players, prs := playersList[m.GameId]

	if !prs {
		return errors.New("GAME NOT FOUND")
	}

	remainingPlayers := []player{}

	for _, player := range players {
		if player.Id != m.PlayerId {
			remainingPlayers = append(remainingPlayers, player)
		}
	}

	playersList[m.GameId] = remainingPlayers

	log.Printf("[GAME %s] - Player %s left", m.GameId, m.PlayerId)
	return nil
}

func MakeMove(ctx *message.Context, m *message.Message) error {

	players, prs := playersList[m.GameId]

	if !prs {
		return errors.New("GAME NOT FOUND")
	}

	for _, player := range players {
		if player.Id != m.PlayerId {
			log.Printf("[GAME %s] - Sending move to player %s", m.GameId, player.Id)
			player.Connection.WriteMessage(ctx.MsgType, []byte(m.Data))
		}
	}

	return nil

}
