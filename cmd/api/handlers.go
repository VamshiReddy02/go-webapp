package main

import (
	"fmt"
	"net/http"
)


func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, world from %s", app.Domain)
}