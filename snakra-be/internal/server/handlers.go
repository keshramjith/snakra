package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gocql/gocql"
	"github.com/keshramjith/snakra/internal/dbservice"
)

const (
	Timed     = 0
	ViewCount = 1
)

type response struct {
	Id string `json:"id"`
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

	id, err := gocql.RandomUUID()
	if err != nil {
		s.errorLogger.Fatalf("failed to generated UUID: %s\n", err)
	}
	vnEntry := &dbservice.Voicenote{
		Id:         id,
		S3_key:     id.String(),
		Created_at: time.Now(),
	}

	s3err := s.s3client.AddObject(requestBody.VnString, vnEntry.Id.String())
	if s3err != nil {
		s.errorLogger.Printf("S3 error: %s\n", s3err)
	}

	s.db.InsertVoicenote(vnEntry)

	w.Header().Set("Content-Type", "application/json")
	js, err := json.Marshal(&response{Id: vnEntry.Id.String()})
	if err != nil {
		fmt.Print("Error")
	}
	w.Write(js)
}

func (s *Server) HandleVoicenoteRead(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{ message: "voice note for specified id" }`))
}
