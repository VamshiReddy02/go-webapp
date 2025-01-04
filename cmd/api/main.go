package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
)

const port = 8080

type application struct {
	Domain string
	dsn string
	DB *sql.DB
}

func main () {
	var app application


	flag.StringVar(&app.dsn, "dsn", "host=localhost port=5433 user=postgres password=postgres dbname=movies sslmode=disable timezone=UTC connect_timeout=5","postgres connection string")
	flag.Parse()

	conn, err := app.ConnectToDB()
	if err != nil {
		log.Fatal(err)
	}

	app.DB = conn
	
	app.Domain = "example.com"

	fmt.Println("serving on port", port)

	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.Routes())
	if err != nil {
		log.Fatal(err)
	}
}