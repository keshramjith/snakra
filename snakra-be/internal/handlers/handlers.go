package handlers

import (
	"net/http"
)

const (
	Timed     = 0
	ViewCount = 1
)

func HandleVoicenoteRead(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("voice note for specified id"))
}

func HandleVoicenoteCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("create voice note - anonymous or user"))
}
