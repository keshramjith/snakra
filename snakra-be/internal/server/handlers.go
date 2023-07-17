package server

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"
	"time"

	"github.com/gocql/gocql"
	"github.com/keshramjith/snakra/internal/dbservice"
)

const (
	Timed     = 0
	ViewCount = 1
)

type get_response_body struct {
	AudioStr string `json:"audio"`
}

type post_request_body struct {
	VnString string `json:"vnString"`
}

type post_response_body struct {
	Id string `json:"id"`
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
	js, err := json.Marshal(&post_response_body{Id: vnEntry.Id.String()})
	if err != nil {
		fmt.Print("Error")
	}
	w.Write(js)
}

func (s *Server) HandleVoicenoteRead(w http.ResponseWriter, r *http.Request) {
	pathId := chi.URLParam(r, "id")

	id, err := gocql.ParseUUID(pathId)
	if err != nil {
		s.errorLogger.Fatalf("Failed to parse uuid from client")
		return
	}
	fetchedVoicenote := &dbservice.Voicenote{Id: id}
	s.db.GetVoicenote(fetchedVoicenote)

	s3obj := s.s3client.GetObject(fetchedVoicenote.S3_key)

	audio, err := io.ReadAll(s3obj)
	js, err := json.Marshal(&get_response_body{AudioStr: string(audio)})
	if err != nil {
		s.errorLogger.Print("Failed to serialize audio string")
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
