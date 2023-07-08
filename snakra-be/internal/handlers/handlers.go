package handlers

import (
	"net/http"
)

func HandleVoicenoteRead(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("voice note for specified id"))
}

func HandleVoicenoteCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("create voice note - anonymous or user"))
}
