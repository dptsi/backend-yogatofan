package repositories

import (
	"coba-db-mssql/models"
)

type UserRepositoryInterface interface {
	AllUser() ([]models.User, error)
}
