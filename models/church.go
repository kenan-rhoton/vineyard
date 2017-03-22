package models

import (
    "gopkg.in/mgo.v2/bson"
)

type Church struct {
    Name string
    Address string
    Mission string
    Default bool
    //events *EventList
}

var TheDB, _ = InitDB("mongo")

func NewChurch() *Church {
    return &Church{Name: "", Address: "", Mission: "", Default: false}
}

func GetDefaultChurch() *Church {
    c := NewChurch()
    TheDB.GetFrom("Churches", bson.M{"default": true}, c)
    return c
}

//func (c *Church) AddEvent(date, location string) {
    //c.events.AddEvent([]string{date,location})
//}

//func (c *Church) GetEvents() *EventList {
    //return c.events
//}

func (c *Church) Save() {
    TheDB.InsertInto("Churches", c)
}

func GetChurch(name String) *Church {
    c := NewChurch()
    TheDB.GetFrom("Churches", bson.M{"name": name}, c)
    return c
}

func GetChurches(name String) []Church {
    var results []Church
    TheDB.GetAll("Churches", results)
    return results
}

func (c *Church) Delete() {
    TheDB.DeleteFrom("Church", bson.M{"name": c.Name})
}

