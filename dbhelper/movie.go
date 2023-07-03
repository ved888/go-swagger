package dbhelper

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"go-swagger/models"
)

func MovieCreate(DB *sqlx.DB, movie *models.Movie) (*models.Movie, error) {
	//SQL==language
	SQL := `INSERT INTO movie (
                   title,
                   genre
                   ) 
            VALUES ($1, $2)
            RETURNING id::uuid ;`

	arg := []interface{}{
		movie.Title,
		movie.Genre,
	}
	var movieId string

	err := DB.QueryRowx(SQL, arg...).Scan(&movieId)
	if err != nil {
		fmt.Println("error in movieId:", err.Error())
		return nil, err
	}

	return &models.Movie{Title: movie.Title, Genre: movie.Genre}, nil
}

// GetMovieById get movie by given id
func GetMovieById(DB *sqlx.DB, movieId string) (*models.Movie, error) {

	// SQL=language
	SQL := `SELECT 
                title,
                genre
          FROM movie
          WHERE id=$1::uuid
                     `

	var movie models.Movie
	err := DB.Get(&movie, SQL, &movieId)
	if err != nil {
		fmt.Println("error in movieId:", err.Error())
		return &models.Movie{}, err
	}
	return &movie, nil
}
