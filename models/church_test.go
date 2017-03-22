package models

import (
    "testing"
)

func TestChurchName(t *testing.T) {
    testdata := []string{
        "asdf",
        "∫ïmbøLôs",
        "àcéntôs",
        "e s p a c i o s",
    }
    c := NewChurch()

    for _, v := range testdata {
        c.Name = v
        if c.Name != v {
            t.Fatalf("Not equal, expected %s, received %s", v, c.Name)
        }
    }
}

func TestChurchAddress(t *testing.T) {
    testdata := []string{
        "Calle normal, 123",
        "C/ Abrev. 248",
        "Avda. åÇasdfA^SDFafpqfá fadsmf oqfqm fom fmadkfm , 8749 2,4234",
    }
    c := NewChurch()

    for _, v := range testdata {
        c.Address = v
        if c.Address != v {
            t.Fatalf("Not equal, expected %s, received %s", v, c.Address)
        }
    }
}

//func TestChurchAddAndGetEventDates(t *testing.T) {
    //testdata := [][]string{
        //[]string{"20-02-2018", "Celebración"},
        //[]string{"30-04-2016", "Casal de Joves"},
        //[]string{"12-08-2030", "BATMANAAAAANANANANANANANNNAN NANANAANNAAN NAN"},
    //}
    //c := NewChurch()

    //for _, v := range testdata {
        //c.AddEvent(v[0], v[1])
        //if el := c.GetEvents(); !el.Contains(v) {
            //t.Fatalf("Expected event %s on %s to have been saved", v[1], v[0] )
        //}
    //}
//}

func TestChurchMission(t *testing.T) {
    testdata := []string{
        "Calle normal, 123",
        "C/ Abrev. 248",
        "Avda. åÇasdfA^SDFafpqfá fadsmf oqfqm fom fmadkfm , 8749 2,4234",
    }
    c := NewChurch()

    for _, v := range testdata {
        c.Mission = v
        if c.Mission != v {
            t.Fatalf("Not equal, expected %s, received %s", v, c.Address)
        }
    }
}

func TestDefaultChurch(t *testing.T) {
    c := GetDefaultChurch()

    if c.Name != "Vinya Castelldefels" {
        t.Fatalf("Expected Vinya Castelldefels, got %s", c.Name )
    }
}
