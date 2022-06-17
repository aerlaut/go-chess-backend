package matcher

import "net/http"

func GenerateMatchLink(w http.ResponseWriter, r *http.Request) {
	link := "http://localhost:5000/match/123123"
	w.Write([]byte(link))
}
