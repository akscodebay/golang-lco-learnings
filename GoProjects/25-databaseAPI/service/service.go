package service

import (
	"errors"
	"log"

	"github.com/akscodebay/databaseapi/dao"
	"github.com/akscodebay/databaseapi/model"
)

func CreateMovie(movie model.Netflix) model.Netflix {
	if dao.CheckTitleExists(movie.ID, movie.Title) {
		log.Println("Movie already exists:", movie.Title)
		return movie
	}
	if movie.Title == "" {
		log.Println("Title cannot be empty")
		return model.Netflix{}
	}
	return dao.InsertRecord(movie.Title, movie.Watched)
}

func GetMovies() []model.Netflix {
	return dao.GetAllRecords()
}

func GetMovie(id int) (model.Netflix, error) {
	return dao.GetOneMovie(id)
}

func DeleteMovie(id int) bool {
	return dao.DeleteOneMovie(id)
}

func DeleteAll() bool {
	return dao.DeleteAllMovies()
}

func UpdateMovie(movie model.Netflix) (model.Netflix, error) {
	if !dao.CheckTitleExists(movie.ID, movie.Title) {
		log.Println("Movie not found:", movie.ID)
		return model.Netflix{}, errors.New("movie not found")
	}
	return dao.UpdateOneMovie(movie)
}
