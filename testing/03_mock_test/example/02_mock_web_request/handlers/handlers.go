package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func init() {
	http.HandleFunc("/user", GetUser)
}

type user struct {
	Id        int64  `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var users = []user{
	{
		Id:        1,
		Username:  "ag",
		FirstName: "Andrew",
		LastName:  "Garfield",
	},
	{
		Id:        2,
		Username:  "th",
		FirstName: "Tom",
		LastName:  "Hiddleston",
	},
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, fmt.Sprintf(`{"error": "method %s is not allowed"}`, r.Method))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	encoder := json.NewEncoder(w)
	encoder.Encode(&users)
}
