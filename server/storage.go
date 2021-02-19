package server

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type User struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Age     int    `json:"age"`
}

// Using inmemory slice as a storage solution. Slice of User object

// type Storage struct {
// 	users []User
// }

var storage []User

func (s *Server) GetAllUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, storage)
}

func (s *Server) CreateUser(c echo.Context) error {
	user := User{}
	user.ID = genUUID()
	err := c.Bind(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity)
	}
	storage = append(storage, user)
	return c.JSON(http.StatusCreated, user)
}

// func (us *Storage) GetUser(){
//
// }
func (s *Server) GetUser(c echo.Context) error {
	//s.conn.GetUser()
	id := c.Param("id")
	for _, user := range storage {
		if user.ID == id {
			return c.JSON(http.StatusOK, user)
		}
	}
	return c.JSON(http.StatusBadRequest, nil)
}

func (s *Server) DeleteUser(c echo.Context) error {
	id := c.Param("id")
	for item := range storage {
		if storage[item].ID == id {
			storage = append(storage[:item], storage[item+1:]...)
			return c.JSON(http.StatusOK, "User was deleted")
		}
	}
	return c.JSON(http.StatusBadRequest, nil)
}

type Storage interface {
	GetUser(id string) (*User, error)
	CreateUser(*User) error // Make sure ID is added to the struct OR return ID from this method
	DeleteUser(id string) error
}

func NewInMemoryStorage() Storage {
	// TODO: IMPLEMENT PART A
	storage = append(storage, User{genUUID(), "Mike", "Turtle", 18})
	storage = append(storage, User{genUUID(), "Leo", "Turtle", 18})
	storage = append(storage, User{genUUID(), "Splinter", "Rat", 300})
	storage = append(storage, User{genUUID(), "Shredder", "Human", 32})
	var out Storage

	// out = Storage{
	// 	users: make([]User, 0),
	// }

	return out
}
