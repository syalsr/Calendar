package api

import (
	"Calendar/interntal/model"
	"sync"
	"time"
)

var Calendar = NewCalendar()

type calendar struct {
	Events   map[int]model.Event
	LastUUID int
	sync.RWMutex
}

func NewCalendar() *calendar {
	return &calendar{Events: make(map[int]model.Event)}
}

func (c *calendar) CreateEvent(event *model.Event) {
	c.Lock()
	defer c.Unlock()

	c.Events[event.UUID] = *event
}
func (c *calendar) UpdateEvent(event *model.Event) {
	c.Lock()
	defer c.Unlock()

	c.Events[event.UUID] = *event
}
func (c *calendar) DeleteEvent(UUID int) {
	c.Lock()
	defer c.Unlock()

	delete(c.Events, UUID)
}
func (c *calendar) GetEventPerDay(day time.Time, userID int) []model.Event {
	var events []model.Event

	for _, event := range c.Events {
		if event.Date == day && event.UserID == userID {
			events = append(events, event)
		}
	}
	return events
}

func (c *calendar) GetEventPerWeek(day time.Time, userID int) []model.Event {
	var events []model.Event

	for _, event := range c.Events {
		if event.Date.After(day) &&
			event.Date.Before(day.AddDate(0, 0, 7)) &&
			event.UserID == userID {
			events = append(events, event)
		}
	}
	return events
}

func (c *calendar) GetEventPerMonth(day time.Time, userID int) []model.Event {
	var events []model.Event

	for _, event := range c.Events {
		if event.Date.After(day) &&
			event.Date.Before(day.AddDate(0, 1, 0)) &&
			event.UserID == userID {
			events = append(events, event)
		}
	}
	return events
}
