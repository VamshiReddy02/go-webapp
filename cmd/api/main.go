package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 8080

type application struct {
	Domain string
}

func main() {
	var app application

	app.Domain = "example.com"

	log.Println("stating application on port 8080")

	err := http.ListenAndServe(fmt.Sprintf(":%d",port), app.Routes())

	if err != nil {
		log.Fatal(err)
	}
}