package models

type User struct {
	ID    int
	Name  string
	Email string
}

type Users struct {
	list []User
}

func NewUsers(list []User) *Users {
	return &Users{
		list: list,
	}
}
