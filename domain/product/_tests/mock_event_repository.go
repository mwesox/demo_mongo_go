package _tests

import "cqrs_go/domain/generic"

type MockEventRepository struct {
	EventsByCorrelationID map[string][]generic.Event
	EventsByCategory      map[string][]generic.Event
}

func (m *MockEventRepository) Save(event generic.Event) error {
	return nil
}

func (m *MockEventRepository) FindByCorrelationID(correlationID string) (*[]generic.Event, error) {
	events := m.EventsByCorrelationID[correlationID]
	return &events, nil
}

func (m *MockEventRepository) FindByCategory(categoryID string) (*[]generic.Event, error) {
	events := m.EventsByCategory[categoryID]
	return &events, nil
}
