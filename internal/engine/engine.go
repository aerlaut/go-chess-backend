package engine

import (
	"errors"
	"log"

	"github.com/aerlaut/go-chess-backend/internal/message"
	"github.com/aerlaut/go-chess-backend/internal/player"
)

const CONNECT_ACTION = "join"
const LEAVE_ACTION = "leave"
const MOVE_ACTION = "move"

var actions = make(map[string]func(*message.Context, *message.Message) error)

func Execute(ctx *message.Context, m *message.Message) error {

	action, prs := actions[m.MsgType]

	if !prs {
		return errors.New("ACTION NOT FOUND")
	}

	err := action(ctx, m)

	if err != nil {
		return err
	}

	return nil
}

func InitEngine() {
	actions[CONNECT_ACTION] = player.Join
	actions[LEAVE_ACTION] = player.Leave
	actions[MOVE_ACTION] = player.MakeMove

	log.Println("[ENGINE] - Engine initialized")
}
