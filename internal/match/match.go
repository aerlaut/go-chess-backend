package match

import (
	"fmt"
	"math/rand"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var defaultIdLength = 8
var baseUrl = "http://localhost:5000"

type match struct {
	Id   string `json:"-"`
	Link string `json:"link"`
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

func generateMatchLink(id string) string {
	return fmt.Sprintf("%s/api/match/%s", baseUrl, id)
}

func NewMatch() *match {

	id := generateMatchId(defaultIdLength)

	return &match{
		Id:   id,
		Link: generateMatchLink(id),
	}
}
