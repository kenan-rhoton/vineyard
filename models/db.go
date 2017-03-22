package models

import (
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)

type DB struct {
    conn *mgo.Session
}

func InitDB(addr string) (*DB, error) {
    db, err := mgo.Dial(addr)
    if err != nil {
        return nil, err
    }
    db.SetMode(mgo.Monotonic, true)
    return &DB{conn: db}, err
}

func (db *DB) Close() {
    db.conn.Close()
}

func (db *DB) InsertInto(table string, object interface{}) error {
    c := db.conn.DB("vineyard").C(table)
    err := c.Insert(object)
    return err
}

func (db *DB) GetFrom(table string, query bson.M, result interface{}) (error){
    c := db.conn.DB("vineyard").C(table)
    err := c.Find(query).One(result)
    return err
}

func (db *DB) GetAll(table string, result interface{}) (error){
    c := db.conn.DB("vineyard").C(table)
    err := c.Find(nil).All(result)
    return err
}

func (db *DB) DeleteFrom(table string, query bson.M) error {
    c := db.conn.DB("vineyard").C(table)
    err := c.Remove(query)
    return err
}
