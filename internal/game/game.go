package game

import (
	"math/rand"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var defaultIdLength = 8

type game struct {
	Id string `json:"id"`
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func generateGameId(n int) string {

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func NewGame() *game {

	id := generateGameId(defaultIdLength)

	return &game{
		Id: id,
	}
}
