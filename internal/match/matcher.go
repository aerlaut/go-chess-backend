package match

import (
	"fmt"
	"net/http"
)

var baseUrl = "http://localhost:5000"

func GenerateMatchLink(w http.ResponseWriter, r *http.Request) {

	match := NewMatch()

	link := fmt.Sprintf("%s/match/%s", baseUrl, match.Id)
	w.Write([]byte(link))
}
