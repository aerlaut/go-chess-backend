package match

import (
	"math/rand"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var defaultIdLength = 8

type match struct {
	Id string `json:"id"`
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func generateMatchId(n int) string {

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func NewMatch() *match {

	id := generateMatchId(defaultIdLength)

	return &match{
		Id: id,
	}
}
