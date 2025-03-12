package gateway

import (
	"encoding/json"
	"net/http"
  "github.com/DedovR/events_test/server"
)

type Server struct{}

func NewServer() Server {
	return Server{}
}

// (GET /)
func (Server) GetV1(w http.ResponseWriter, r *http.Request, params api.GetV1Params) {
	resp := api.EventsResponse{
    api.EventResponse{
      Type: "type",
    },
  }

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}

// (POST /start)
func (Server) PostV1Start(w http.ResponseWriter, r *http.Request) {
	resp := api.EventsResponse{}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}

// (POST /finish)
func (Server) PostV1Finish(w http.ResponseWriter, r *http.Request) {
	resp := api.EventsResponse{}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}
