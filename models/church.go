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

func GetDefaultChurch() *Church {
    c := &Church{}
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
    c := &Church{}
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

func GetChurch(handle string) *Church {
    c := &Church{}
    TheDB.GetFrom("Churches", bson.M{"handle": handle}, c)
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
