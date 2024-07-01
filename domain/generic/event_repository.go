package generic

type EventRepository interface {
	Save(event Event) error
	FindByCorrelationID(correlationID string) (*[]Event, error)
	FindByCategory(categoryID string) (*[]Event, error)
}
