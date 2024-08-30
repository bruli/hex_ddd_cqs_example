package app

import (
	"github.com/google/uuid"
	"time"
)

type Event interface {
	EventID() uuid.UUID
	EventName() string
	EventAt() time.Time
	AggregateRootID() string
}
