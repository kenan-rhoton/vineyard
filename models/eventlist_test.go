package models

import (
    "testing"
)

func TestEventListContains(t *testing.T) {
    testdata := [][]string{
        []string{"20-02-2018", "Celebración"},
        []string{"30-04-2016", "Casal de Joves"},
        []string{"12-08-2030", "BATMANAAAAANANANANANANANNNAN NANANAANNAAN NAN"},
    }

    el := NewEventList()

    for _, v := range testdata {
        el.AddEvent(v)
        if !el.Contains(v) {
            t.Fatalf("Expected event %s on %s to have been saved", v[1], v[0] )
        }
    }
}

func TestEventListNotContains(t *testing.T) {
    testdata := [][]string{
        []string{"20-02-2018", "Celebración"},
        []string{"30-04-2016", "Casal de Joves"},
        []string{"12-08-2030", "BATMANAAAAANANANANANANANNNAN NANANAANNAAN NAN"},
    }

    el := NewEventList()

    for _, v := range testdata {
        if el.Contains(v) {
            t.Fatalf("Expected event %s on %s to not have been saved", v[1], v[0] )
        }
    }
}
