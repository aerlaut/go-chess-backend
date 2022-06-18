package match

import (
	"math/rand"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var defaultIdLength = 8

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

type match struct {
	Id string
}

func NewMatch() *match {
	return &match{
		Id: generateMatchId(defaultIdLength),
	}
}
