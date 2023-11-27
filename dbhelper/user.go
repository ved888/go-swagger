package dbhelper

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"go-swagger/database"
	"go-swagger/models"
)

func CreateUser(DB *sqlx.DB, user *models.Register) (*models.Register, error) {
	//SQL==language
	SQL := `INSERT INTO users (
                     name,
                     address,
                     email,
                     phone,
                     dob
                   ) 
            VALUES ($1,$2,$3,$4,$5)
            RETURNING id::uuid ;`

	arg := []interface{}{
		user.Name,
		user.Address,
		user.Email,
		user.Phone,
		user.DOB,
	}
	var userId string

	err := DB.QueryRowx(SQL, arg...).Scan(&userId)
	if err != nil {
		fmt.Println("error in movieId:", err.Error())
		return nil, err
	}
	return &models.Register{}, nil
}

// GetUserByEmailAndPhone login user by email and phone
func GetUserByEmailAndPhone(DB *sqlx.DB, phone, email *string) (*models.Register, error) {
	// SQL=language
	SQL := `Select 
                   id,
                   name,
                   address,
                   phone,
                   email,
                   dob
             From users
             where phone=$1 and email=$2
             `

	args := []interface{}{
		phone,
		email,
	}
	var user models.Register
	err := database.DB.Get(&user, SQL, args...)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
