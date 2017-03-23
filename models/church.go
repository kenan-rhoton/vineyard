package models

import (
    "gopkg.in/mgo.v2/bson"
    "net/url"
    "fmt"
)

type Church struct {
    Name string
    Handle string
    Address string
    Mission string
    Default bool
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

func SaveChurch(form url.Values, _ string) error {
    c := &Church{}
    err := LoadModel(c, form)
    if err != nil {
        return err
    }
    err = c.Validate()
    if err != nil {
        return err
    }
    return c.Save()
}

func UpdateChurch(form url.Values, which string) error {
    c := NewChurch()
    LoadModel(c, form)
    return c.Update(which)
}

func (c *Church) Save() error {
    return TheDB.InsertInto("Churches", c)
}

func (c *Church) Validate() error {
    switch {
    case c.Handle == "":
        return fmt.Errorf("Empty Handle")
    }
    return nil
}

func (c *Church) Update(handle string) error {
    return TheDB.Update("Churches", bson.M{"handle": handle}, c)
}

func GetChurch(name string) *Church {
    c := NewChurch()
    TheDB.GetFrom("Churches", bson.M{"name": name}, c)
    return c
}

func GetChurches() []Church {
    var results []Church
    TheDB.GetAll("Churches", &results)
    return results
}

func DeleteChurch(handle string) error {
    return TheDB.DeleteFrom("Churches", bson.M{"handle": handle})
}

func (c *Church) Delete() {
    TheDB.DeleteFrom("Churches", bson.M{"handle": c.Handle})
}

