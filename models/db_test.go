package models

import (
    "testing"
    "runtime"
    "gopkg.in/mgo.v2/bson"
)

func CheckDatabaseAccessError(err error, t *testing.T) {
    if err != nil {
        _, _, line, _ := runtime.Caller(1)
        t.Fatalf("Error accessing the database on line %d: %s", line, err.Error())
    }
}

func TestInitDBHasNoError(t *testing.T) {
    _, err := InitDB("localhost")
    CheckDatabaseAccessError(err, t)
}

type testStruct struct {
    A int
    B string
    C bool
}

func TestPersistence(t *testing.T) {
    db, _ := InitDB("localhost")
    testdata := testStruct{A: 55, B: "potatoes", C: false}
    err := db.InsertInto("droptest", &testdata)
    CheckDatabaseAccessError(err, t)
    res := testStruct{}
    err = db.GetFrom("droptest", bson.M{"a": 55}, &res)
    CheckDatabaseAccessError(err, t)
    if res.A != testdata.A || res.B != testdata.B || res.C != testdata.C {
        t.Fatalf("Persistence error, retrieved object differs (%d - %d) (%s - %s) (%b - %b)", res.A, testdata.A, res.B, testdata.B, res.C, testdata.C)
    }
    err = db.DeleteFrom("droptest", bson.M{"a": 55})
    CheckDatabaseAccessError(err, t)
}
