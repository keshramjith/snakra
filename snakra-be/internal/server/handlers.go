package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	Timed     = 0
	ViewCount = 1
)

type response struct {
	Message string `json:"message"`
}

func (s *Server) HandleVoicenoteRead(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{ message: "voice note for specified id" }`))
}

func (s *Server) HandleVoicenoteCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	js, err := json.Marshal(&response{Message: "created voice note"})
	if err != nil {
		fmt.Print("Error")
	}
	w.Write(js)
}
