package dao

import (
	"context"
	"fmt"
	"log"

	"github.com/akscodebay/databaseapi/model"
	"github.com/jackc/pgx/v5"
)

const connectionstring = "postgres://aks:password@localhost:5432/godb"

func getConnection() *pgx.Conn {

	conn, err := pgx.Connect(context.Background(), connectionstring)
	if err != nil {
		log.Println("Unable to connect to database: ", err)
	}
	return conn
}

func CreateTable() {
	conn := getConnection()
	defer conn.Close(context.Background())
	fmt.Println("Inside Create table...")
	_, err := conn.Exec(context.Background(), `
		CREATE TABLE IF NOT EXISTS netflix (
			id SERIAL PRIMARY KEY,
			title TEXT,
			watched BOOLEAN
		)
	`)
	if err != nil {
		log.Println("Unable to create table: ", err)
	}
}

func InsertRecord(title string, watched bool) model.Netflix {
	conn := getConnection()
	defer conn.Close(context.Background())
	movie := model.Netflix{Title: title, Watched: watched}
	err := conn.QueryRow(context.Background(), `
		INSERT INTO netflix (title, watched) VALUES ($1, $2) RETURNING id
	`, movie.Title, movie.Watched).Scan(&movie.ID)
	if err != nil {
		log.Println("Unable to insert record: ", err)
	}
	return movie
}

func CheckTitleExists(id int, title string) bool {
	conn := getConnection()
	defer conn.Close(context.Background())

	var exists bool
	err := conn.QueryRow(context.Background(), `
		SELECT EXISTS(SELECT 1 FROM netflix WHERE title = $1 or id = $2)
	`, title, id).Scan(&exists)
	if err != nil {
		log.Println("Unable to check if title exists: ", err)
	}
	return exists
}

func GetAllRecords() []model.Netflix {
	conn := getConnection()
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), `
		SELECT id, title, watched FROM netflix
	`)
	if err != nil {
		log.Println("Unable to retrieve records: ", err)
	}
	defer rows.Close()

	var movies []model.Netflix = []model.Netflix{}
	for rows.Next() {
		var movie model.Netflix
		if err := rows.Scan(&movie.ID, &movie.Title, &movie.Watched); err != nil {
			log.Println("Unable to scan row: ", err)
		}
		movies = append(movies, movie)
	}
	return movies
}

func GetOneMovie(id int) (model.Netflix, error) {
	conn := getConnection()
	defer conn.Close(context.Background())

	var movie model.Netflix = model.Netflix{}
	err := conn.QueryRow(context.Background(), `
		SELECT id, title, watched FROM netflix WHERE id = $1
	`, id).Scan(&movie.ID, &movie.Title, &movie.Watched)
	if err != nil {
		if err == pgx.ErrNoRows {
			return movie, nil
		}
		log.Println("Unable to retrieve record: ", err)
	}
	return movie, nil
}

func DeleteOneMovie(id int) bool {
	conn := getConnection()
	defer conn.Close(context.Background())

	_, err := conn.Exec(context.Background(), `
		DELETE FROM netflix WHERE id = $1
	`, id)
	if err != nil {
		log.Println("Unable to delete record: ", err)
		return false
	}
	return true
}

func DeleteAllMovies() bool {
	conn := getConnection()
	defer conn.Close(context.Background())

	_, err := conn.Exec(context.Background(), `
		DELETE FROM netflix
	`)
	if err != nil {
		log.Println("Unable to delete records: ", err)
		return false
	}
	return true
}

func UpdateOneMovie(movie model.Netflix) (model.Netflix, error) {
	conn := getConnection()
	defer conn.Close(context.Background())

	var updatedMovie model.Netflix
	err := conn.QueryRow(context.Background(), `
		UPDATE netflix SET title = $1, watched = $2 WHERE id = $3 RETURNING id, title, watched
	`, movie.Title, movie.Watched, movie.ID).Scan(&updatedMovie.ID, &updatedMovie.Title, &updatedMovie.Watched)
	if err != nil {
		if err == pgx.ErrNoRows {
			return movie, nil
		}
		log.Println("Unable to update record: ", err)
	}
	return updatedMovie, nil
}
