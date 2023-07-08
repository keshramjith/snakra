package server

const (
	api_version    string = "v1"
	api            string = "api"
	voicenote_base string = "/" + api + "/" + api_version + "/vn"
	voicenote_id   string = voicenote_base + "/{id}"
)

func (s *Server) registerRoutes() {
	s.router.Get(voicenote_id, s.HandleVoicenoteRead)
	s.router.Post(voicenote_base, s.HandleVoicenoteCreate)
}
