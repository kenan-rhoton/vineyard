package models

import (
    "testing"
)

func TestInitDBHasNoError(t *testing.T) {
    _, err := InitDB()
    if err != nil {
        t.Fatalf("Error accessing the database: %s", err.Error())
    }
}
