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

//var storage []User

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

// func (s *Server) CreateUser(c echo.Context) error {
// 	user := User{}
// 	user.ID = genUUID()
// 	err := c.Bind(&user)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusUnprocessableEntity)
// 	}
// 	storage = append(storage, user)
// 	return c.JSON(http.StatusCreated, user)
// }

// func (s *Server) GetUser(c echo.Context) error {
// 	id := c.Param("id")
// 	for _, user := range storage {
// 		if user.ID == id {
// 			return c.JSON(http.StatusOK, user)
// 		}
// 	}
// 	return c.JSON(http.StatusBadRequest, nil)
// }

// func (s *Server) DeleteUser(c echo.Context) error {
// 	id := c.Param("id")
// 	for item := range storage {
// 		if storage[item].ID == id {
// 			storage = append(storage[:item], storage[item+1:]...)
// 			return c.JSON(http.StatusOK, "User was deleted")
// 		}
// 	}
// 	return c.JSON(http.StatusBadRequest, nil)
// }

func NewInMemoryStorage() Storage {
	// TODO: IMPLEMENT PART A
	// storage = append(storage, User{genUUID(), "Mike", "Turtle", 18})
	// storage = append(storage, User{genUUID(), "Leo", "Turtle", 18})
	// storage = append(storage, User{genUUID(), "Splinter", "Rat", 300})
	// storage = append(storage, User{genUUID(), "Shredder", "Human", 32})

	return &UserStorage{
		storage: []User{
			{genUUID(), "Shredder", "Human", 32},
			{genUUID(), "Mike", "Turtle", 18},
			{genUUID(), "Splinter", "Rat", 300},
			{genUUID(), "Leo", "Turtle", 18},
		},
	}
}
