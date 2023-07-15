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

type post_request_body struct {
	VnString string `json:"vnString"`
}

func (s *Server) HandleVoicenoteCreate(w http.ResponseWriter, r *http.Request) {
	var requestBody post_request_body
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		s.errorLogger.Println("failed to unmarshal request object")
		return
	}
	// s.s3client.AddObject("cat", "suidhawiud")
	// s.db.Session.
	w.Header().Set("Content-Type", "application/json")
	js, err := json.Marshal(&response{Message: requestBody.VnString})
	if err != nil {
		fmt.Print("Error")
	}
	w.Write(js)
}

func (s *Server) HandleVoicenoteRead(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{ message: "voice note for specified id" }`))
}
