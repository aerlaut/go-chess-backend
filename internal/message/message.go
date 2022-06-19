package message

import (
	"encoding/json"
)

type Message struct {
	MsgType  string `json:"msgType"`
	GameId   string `json:"gameId"`
	PlayerId string `json:"playerId"`
	Data     string `json:"data"`
}

func ParseMessage(message string) (Message, error) {
	var msg Message
	err := json.Unmarshal([]byte(message), &msg)

	if err != nil {
		return msg, err
	}

	return msg, nil
}
