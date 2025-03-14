package gateway

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"regexp"

	"github.com/DedovR/events_test/entity"
	er "github.com/DedovR/events_test/errors"
	"github.com/DedovR/events_test/server"
	"github.com/DedovR/events_test/usecase"
)

const (
  defaultLimit = 25
  maxLimit     = 100
)

var typeRegexp = regexp.MustCompile("^[a-z0-9]+$")

type Server struct{
  uc usecase.EventUC
}

func NewServer(uc usecase.EventUC) Server {
  return Server{
    uc: uc,
  }
}

// (GET v1/)
func (s Server) GetV1(w http.ResponseWriter, r *http.Request, params api.GetV1Params) {
  listParams, err := parseParams(params)
  if err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
  }

  ctx := context.Background()
  events, err := s.uc.GetList(ctx, listParams)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  resp := make(api.EventsResponse, 0, listParams.Limit)
  for _, e := range events {
    resp = append(resp, *e.ToResponse())
  }

  w.WriteHeader(http.StatusOK)
  _ = json.NewEncoder(w).Encode(resp)
}

// (POST v1/start)
func (s Server) PostV1Start(w http.ResponseWriter, r *http.Request) {
  var payload api.EventRequest
  err := json.NewDecoder(r.Body).Decode(&payload)
  if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
  }
  if err = validateType(payload.Type); err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
  }

  ctx := context.Background()
  err = s.uc.Start(ctx, payload.Type)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.WriteHeader(http.StatusOK)
}

// (POST v1/finish)
func (s Server) PostV1Finish(w http.ResponseWriter, r *http.Request) {
  var payload api.EventRequest
  err := json.NewDecoder(r.Body).Decode(&payload)
  if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
  }
  if err = validateType(payload.Type); err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
  }

  ctx := context.Background()
  err = s.uc.Finish(ctx, payload.Type)
  if errors.Is(err, er.ErrNoRows) {
    http.Error(w, err.Error(), http.StatusNotFound)
    return
  }
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.WriteHeader(http.StatusOK)
}

func validateType(t string) error {
  if !typeRegexp.MatchString(t) {
    return er.ErrInvalidType
  }

  return nil
}

func parseParams(p api.GetV1Params) (*entity.EventListParams, error) {
  params := &entity.EventListParams{
    Limit: defaultLimit,
  }

  if p.Limit != nil {
    if *p.Limit > maxLimit {
      return nil, er.ErrInvalidLimit
    }
    params.Limit = int64(*p.Limit)
  }

  if p.Offset != nil {
    params.Offset = int64(*p.Offset)
  }

  if p.Type != nil {
    params.Type = *p.Type
    if err := validateType(params.Type); err != nil {
      return nil, err
    }
  }

  return params, nil
}
