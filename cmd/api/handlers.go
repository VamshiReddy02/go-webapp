package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"example.com/webapp/internals/models"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	var payload = struct {
		Status string `json: "status"`
		Message string `json: "message"`
		Version string `json: "version"`
	}{
		Status: "Active",
		Message: "Go Movies up and running",
		Version: "1.0.0",
	}

	out, err := json.Marshal(payload)
	if err != nil {
		fmt.Print(err)
	}


	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}

func (app *application) AllMovies(w http.ResponseWriter, r *http.Request) {
	var Movies []models.Movie

	rd, _ := time.Parse("2006-01-02", "1981-06-12")

	Avengers := models.Movie {
		ID: 1,
		Title: "Avengers",
		ReleaseDate: rd,
		MPAARating: "R",
		RunTime: 116,
		Description: "earth mightiest heroes",
		CreatedAt: time.Now(),
		UpdatedField: time.Now(),
	}

	Movies = append(Movies, Avengers)

	rd, _ = time.Parse("2019-04-26", "1980-04-26")

	EndGame := models.Movie {
		ID: 2,
		Title: "EndGame",
		ReleaseDate: rd,
		MPAARating: "R",
		RunTime: 116,
		Description: "earth mightiest heroes",
		CreatedAt: time.Now(),
		UpdatedField: time.Now(),
	}

	Movies = append(Movies, EndGame)

	out, err := json.Marshal(Movies)
	if err != nil {
		fmt.Print(err)
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)

}