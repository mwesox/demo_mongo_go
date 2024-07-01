package generic

import (
	"time"
)

type Event struct {
	Id            string
	EventType     string
	Payload       []byte `json:"payload"`
	Timestamp     time.Time
	CorrelationId string
	Category      string
}
