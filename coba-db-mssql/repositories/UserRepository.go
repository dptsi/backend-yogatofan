package repositories

import (
	"coba-db-mssql/models"
	"database/sql"
	"fmt"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u UserRepository) AllUser() ([]models.User, error) {
	q := fmt.Sprintf("SELECT id, name, email FROM users")
	rows, err := u.db.Query(q)

	if err != nil {
		fmt.Println("Error reading rows: " + err.Error())
		return nil, err
	}

	defer rows.Close()

	var (
		user  models.User
		users []models.User
	)
	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Name, &user.Email)

		if err != nil {
			fmt.Println("Error reading rows: " + err.Error())
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
