package repo

import (
	"context"
	"log"
	"time"

	"github.com/DedovR/events_test/entity"
	er "github.com/DedovR/events_test/errors"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const (
  dbName = "events"
  eventCollection = "events"
)

type EventRepo interface {
  GetList(ctx context.Context, t string, limit, offset int64) ([]entity.Event, error)
  Start(ctx context.Context, t string) error
  Finish(ctx context.Context, t string) error
}

type Event struct {
  client     *mongo.Client
  collection *mongo.Collection
}

func NewEvent(client *mongo.Client) *Event {
  return &Event{
    client:     client,
    collection: client.Database(dbName).Collection(eventCollection),
  }
}

func (e *Event) GetList(ctx context.Context, t string, limit, offset int64) ([]entity.Event, error) {
  events := make([]entity.Event, 0, limit)
  filter := bson.D{{}}

  if t != "" {
    filter = bson.D{{Key: "event_type", Value: t}}
  }
  opts := options.Find().
    SetLimit(limit).
    SetSkip(offset).
    SetSort(bson.D{{Key: "started_at", Value: -1}})
  cursor, err := e.collection.Find(ctx, filter, opts)

  if err != nil {
    return nil , err
  }

  if err = cursor.All(ctx, &events); err != nil {
    return nil, err
  }

  return events, nil
}

func (e *Event) Start(ctx context.Context, t string) error {
  event := &entity.Event{
    Type: t,
    State: 0,
    StartedAt: time.Now(),
  }

  _, err := e.collection.InsertOne(ctx, event)
  if err != nil {
    return err
  }

  return nil
}

func (e *Event) Finish(ctx context.Context, t string) error {
  event := &entity.Event{}
  filter := bson.D{
    {Key: "event_type", Value: t},
    {Key: "state", Value: 0},
  }
  err := e.collection.FindOne(ctx, filter).Decode(event)
  if err != nil {
    if err == mongo.ErrNoDocuments {
      return er.ErrNoRows
    }
    return err
  }

  log.Println(event)
  now := time.Now()
  update := bson.D{{"$set",
    bson.D{
      {Key: "state", Value: 1},
      {Key: "finished_at", Value: now},
    },
  }}

  _, err = e.collection.UpdateByID(ctx, event.ID, update)
  if err != nil {
    return err
  }

  return nil
}




