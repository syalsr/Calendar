package api

import (
	"Calendar/interntal/model"
	"sync"
	"time"
)

type Calendar struct {
	events map[int]model.Event
	sync.RWMutex
}

func NewCalendar() *Calendar {
	return &Calendar{events: make(map[int]model.Event)}
}

func (c *Calendar) CreateEvent(event *model.Event) {
	c.Lock()
	defer c.Unlock()

	c.events[event.UUID] = *event
}
func (c *Calendar) UpdateEvent(event *model.Event) {
	c.Lock()
	defer c.Unlock()

	c.events[event.UUID] = *event
}
func (c *Calendar) DeleteEvent(UUID int) {
	c.Lock()
	defer c.Unlock()

	delete(c.events, UUID)
}
func (c *Calendar) GetEventPerDay(day time.Time) []model.Event {
	var events []model.Event

	for _, event := range c.events {
		if event.Date == day {
			events = append(events, event)
		}
	}
	return events
}

func (c *Calendar) GetEventPerWeek(day time.Time) []model.Event {
	var events []model.Event

	for _, event := range c.events {
		if event.Date.After(day) && event.Date.Before(day.AddDate(0, 0, 7)) {
			events = append(events, event)
		}
	}
	return events
}

func (c *Calendar) GetEventPerMonth(day time.Time) []model.Event {
	var events []model.Event

	for _, event := range c.events {
		if event.Date.After(day) && event.Date.Before(day.AddDate(0, 1, 0)) {
			events = append(events, event)
		}
	}
	return events
}
