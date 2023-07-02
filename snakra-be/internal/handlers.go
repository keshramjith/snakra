package internal

import "net/http"

func handleVoicenoteRead(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("voice note for specified id"))
}

func handleVoicenotesByUserRead(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("voice notes for specified user"))
}

func handleVoicenoteCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("create voice note - anonymous or user"))
}
