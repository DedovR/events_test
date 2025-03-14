package usecase

import (
	"context"
	"log"

	"github.com/DedovR/events_test/entity"
	"github.com/DedovR/events_test/repo"
)

type EventUC interface {
  GetList(ctx context.Context, params *entity.EventListParams) ([]entity.Event, error)
  Start(ctx context.Context, t string) error
  Finish(ctx context.Context, t string) error
}

type Event struct {
  repo repo.EventRepo
}

func NewEvent(repo repo.EventRepo) *Event {
  return &Event{
    repo: repo,
  }
}

func (e *Event) GetList(ctx context.Context, p *entity.EventListParams) ([]entity.Event, error) {
  events, err := e.repo.GetList(ctx, p.Type, p.Limit, p.Offset)
  if err != nil {
    return nil, err
  }

  return events, nil
}

func (e *Event) Start(ctx context.Context, t string) error {
  err := e.repo.Start(ctx, t)
  if err != nil {
    return err
  }

  log.Printf("Start event %s\n", t)
  // Здесь могла быть ваша бизнес-логика

  return nil
}

func (e *Event) Finish(ctx context.Context, t string) error {
  log.Printf("Finish event %s\n", t)
  // Здесь могла быть ваша бизнес-логика

  err := e.repo.Finish(ctx, t)
  if err != nil {
    return err
  }

  return nil
}
