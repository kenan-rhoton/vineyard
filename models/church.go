package models

import (
    "fmt"
)

type Church struct {
    Name string
    Handle string
    Address string
    Mission string
    Default bool
}

func (c *Church) class() string {
    return "Churches"
}

func (c *Church) getKey() string {
    return "handle"
}

func (c *Church) getValue() interface{} {
    return c.Handle
}

func (c *Church) validate() error {
    switch {
    case c.Handle == "":
        return fmt.Errorf("Empty Handle")
    }
    return nil
}
