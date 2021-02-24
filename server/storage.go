package server

import (
	"fmt"
)

type User struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Age     int    `json:"age"`
}

type Storage interface {
	GetUser(id string) (*User, error)
	CreateUser(*User) error // Make sure ID is added to the struct OR return ID from this method
	DeleteUser(id string) error
	GetAllUsers() (*[]User, error)
}

type UserStorage struct {
	storage []User
}

func (us *UserStorage) GetAllUsers() (*[]User, error) {
	return &us.storage, nil
}

func (us *UserStorage) GetUser(id string) (*User, error) {
	for _, user := range us.storage {
		if user.ID == id {
			return &user, nil
		}
	}
	return nil, fmt.Errorf("User with id %v was not found", id)
}

func (us *UserStorage) CreateUser(*User) error {
	user := User{}
	user.ID = genUUID()
	us.storage = append(us.storage, user)
	return nil
}

func (us *UserStorage) DeleteUser(id string) error {
	for item := range us.storage {
		if us.storage[item].ID == id {
			us.storage = append(us.storage[:item], us.storage[item+1:]...)
			return fmt.Errorf("User with id %v was deleted", id)
		}
	}
	return nil
}

func NewInMemoryStorage() Storage {
	// TODO: IMPLEMENT PART A
	return &UserStorage{
		storage: []User{
			{genUUID(), "Shredder", "Human", 32},
			{genUUID(), "Mike", "Turtle", 18},
			{genUUID(), "Splinter", "Rat", 300},
			{genUUID(), "Leo", "Turtle", 18},
		},
	}
}
