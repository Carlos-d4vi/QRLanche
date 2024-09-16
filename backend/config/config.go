package config

import (
    "database/sql"
    _ "github.com/lib/pq" // Importa o driver PostgreSQL
)

var DB *sql.DB

func OpenConn() error {
    var err error
    DB, err = sql.Open("postgres", "host=localhost port=5432 user=postgres password=d4vi1234) dbname=qrLanche sslmode=disable")
    if err != nil {
        return err
    }
    return DB.Ping()
}

func CloseConn() error {
    if DB != nil {
        return DB.Close()
    }
    return nil
}
