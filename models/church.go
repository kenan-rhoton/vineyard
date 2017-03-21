package models

type Church struct {
    Name string
    Address string
    events *EventList
}

func NewChurch() *Church {
    return &Church{Name: "", Address: "", events: NewEventList()}
}

func GetDefaultChurch() *Church {
    return &Church{}
}

func (c *Church) AddEvent(date, location string) {
    c.events.AddEvent([]string{date,location})
}

func (c *Church) GetEvents() *EventList {
    return c.events
}
