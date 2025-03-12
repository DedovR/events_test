package entity

import "time"

type Event struct {
  Type       string     `json:"type" bson:"type"`
  State      string     `json:"state" bson:"state"`
  StartedAt  time.Time  `json:"started_at" bson:"started_at"`
  FinishedAt *time.Time `json:"finished_at" bson:"finished_at"`
}
