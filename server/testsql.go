package main

import (
	"database/sql"
	"fmt"
)

var DB_DRIVER string

func Testsql() {
	sql.Register(DB_DRIVER, &sqlite3.SQLiteDriver{})
	database, err := sql.Open(DB_DRIVER, "mysqlite_3")
	if err != nil {
		fmt.Println("Failed to create the handle")
	}
	if err2 := database.Ping(); err2 != nil {
		fmt.Println("Failed to keep connection alive")
	}
	fmt.Println("Hello world")
}
