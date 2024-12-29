package main

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

func main() {
	// Capture connection properties.
	cfg := mysql.Config{
		//User:   os.Getenv("DBUSER"),
		//Passwd: os.Getenv("DBPASS"),
		User:   "root",
		Passwd: "20230205.ken",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "recordings",
	}

	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
}
