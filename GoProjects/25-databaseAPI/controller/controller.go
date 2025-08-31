package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/akscodebay/databaseapi/model"
	"github.com/akscodebay/databaseapi/service"
	"github.com/gorilla/mux"
)

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	var movie model.Netflix
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		log.Println("Error binding JSON:", err)
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	createdMovie := service.CreateMovie(movie)
	if createdMovie.ID == 0 {
		http.Error(w, "Failed to create movie", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdMovie)
}

func GetMovies(w http.ResponseWriter, r *http.Request) {
	movies := service.GetMovies()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(movies)
}

func GetMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid movie ID", http.StatusBadRequest)
		return
	}
	movie, err := service.GetMovie(id)
	if err != nil {
		http.Error(w, "Movie not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(movie)
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid movie ID", http.StatusBadRequest)
		return
	}
	if err := service.DeleteMovie(id); !err {
		http.Error(w, "Movie not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func DeleteAll(w http.ResponseWriter, r *http.Request) {
	if err := service.DeleteAll(); !err {
		http.Error(w, "Failed to delete movies", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	var movie model.Netflix
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	if movie.Title == "" {
		http.Error(w, "Movie title is required", http.StatusBadRequest)
		return
	}
	updatedMovie, err := service.UpdateMovie(movie)
	if err != nil {
		http.Error(w, "Failed to update movie", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedMovie)
}
