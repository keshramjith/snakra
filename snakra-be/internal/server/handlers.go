package server

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid/v5"
	"io"
	"net/http"
	"time"

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
		s.logger.Errorln("failed to unmarshal request object")
		return
	}

	var id, generate_uuid_err = uuid.NewV4()
	if generate_uuid_err != nil {
		s.logger.Fatalf("failed to generated UUID: %s\n", err)
	}
	vnEntry := &dbservice.Voicenote{
		Id:         id,
		S3_key:     id.String(),
		Created_at: time.Now(),
	}

	s3err := s.s3client.AddObject(requestBody.VnString, vnEntry.Id.String())
	if s3err != nil {
		s.logger.Errorf("S3 error: %s\n", s3err)
	}

	dbErr := s.db.InsertVoicenote(vnEntry)
	if dbErr != nil {
		s.logger.Errorf("Db Error occured: %s", dbErr)
	}

	w.Header().Set("Content-Type", "application/json")
	js, err := json.Marshal(&post_response_body{Id: vnEntry.Id.String()})
	if err != nil {
		s.logger.Errorf("Error marshalling json: %s", err)
	}
	w.Write(js)
}

func (s *Server) HandleVoicenoteRead(w http.ResponseWriter, r *http.Request) {
	pathId := chi.URLParam(r, "id")

	id, err := uuid.FromString(pathId)
	if err != nil {
		s.logger.Errorln("Failed to parse uuid from client")
		w.WriteHeader(404)
		return
	}
	fetchedVoicenote := &dbservice.Voicenote{Id: id}
	dbErr := s.db.GetVoicenote(fetchedVoicenote)
	if dbErr != nil {
		s.logger.Errorf("Db Error Occured: %s", dbErr)
	}

	if fetchedVoicenote.S3_key == "" {
		w.WriteHeader(404)
		return
	}

	s3obj := s.s3client.GetObject(fetchedVoicenote.S3_key)

	audio, err := io.ReadAll(s3obj)
	js, err := json.Marshal(&get_response_body{AudioStr: string(audio)})
	if err != nil {
		s.logger.Infoln("Failed to serialize audio string")
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
