package models

import (
    "container/list"
)

type EventList struct {
    events *list.List
}

func NewEventList() *EventList {
    return &EventList{events: list.New()}
}

func (e *EventList) AddEvent(event []string) {
    e.events.PushBack(event)
}

func (e *EventList) Contains(search []string) bool {
    for it := e.events.Front(); it != nil; it = it.Next() {
        if testEq(it.Value.([]string), search) {
            return true
        }
    }
    return false
}

func testEq(a, b []string) bool {

    if a == nil && b == nil {
        return true;
    }

    if a == nil || b == nil {
        return false;
    }

    if len(a) != len(b) {
        return false
    }

    for i := range a {
        if a[i] != b[i] {
            return false
        }
    }

    return true
}
