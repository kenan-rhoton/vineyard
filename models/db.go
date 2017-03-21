package models

import (
    "database/sql"
    "log"
    _ "github.com/lib/pq"
)

type DB struct {
    conn *sql.DB
}

func InitDB() (*DB, error) {
    db, err := sql.Open("postgres", "user=vineyard dbname=vineyard sslmode=verify-full")
    if err != nil {
        log.Fatal(err)
    }
    err = db.Ping()
    if err != nil {
        log.Fatal(err)
    }

    return &DB{conn: db}, err
}

func (db *DB) Close() {
    db.conn.Close()
}
