package entity

import (
	"time"

	"github.com/DedovR/events_test/server"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Event struct {
  ID         bson.ObjectID `bson:"_id"`
  Type       string        `json:"type" bson:"event_type"`
  State      byte          `json:"state" bson:"state"`             // 0 - для незавершенных событий, 1 - для завершенных
  StartedAt  time.Time     `json:"startedAt" bson:"started_at"`
  FinishedAt *time.Time    `json:"finishedAt" bson:"finished_at"`
}

type EventListParams struct {
  Type   string
  Limit  int64
  Offset int64
}

func (e *Event) ToResponse() *api.EventResponse {
  result := &api.EventResponse{
    Type: e.Type,
    StartedAt: e.StartedAt,
    FinishedAt: e.FinishedAt,
  }

  if e.FinishedAt == nil {
    result.State = api.Started
  } else {
    result.State = api.Finished
  }

  return result
}
