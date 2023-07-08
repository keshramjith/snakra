package handlers

import (
	"net/http"
)

func HandleVoicenoteRead(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("voice note for specified id"))
}

func HandleVoicenotesByUserRead(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("voice notes for specified user"))
}

func HandleVoicenoteCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("create voice note - anonymous or user"))
}
