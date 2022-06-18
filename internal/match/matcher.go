package match

import (
	"encoding/json"
	"net/http"
)

func GenerateMatchLink(w http.ResponseWriter, r *http.Request) {

	match := NewMatch()
	str, _ := json.Marshal(match)

	w.Write(str)
}
