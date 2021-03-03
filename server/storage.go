package server

import (
	"database/sql"
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
	SaveUser(*User) error
	DeleteUser(id string) error
	GetAllUsers() ([]User, error)
}

// type UserStorage struct {
// 	storage []User
// }

// func (us *UserStorage) GetAllUsers() ([]User, error) {
// 	return us.storage, nil
// }

// func (us *UserStorage) GetUser(id string) (*User, error) {
// 	for _, user := range us.storage {
// 		if user.ID == id {
// 			return &user, nil
// 		}
// 	}
// 	return nil, fmt.Errorf("User with id %v was not found", id)
// }

// func (us *UserStorage) SaveUser(u *User) error {
// 	u.ID = genUUID()
// 	us.storage = append(us.storage, *u)
// 	return nil
// }

// func (us *UserStorage) DeleteUser(id string) error {
// 	for item := range us.storage {
// 		if us.storage[item].ID == id {
// 			us.storage = append(us.storage[:item], us.storage[item+1:]...)
// 			return fmt.Errorf("User with id %v was deleted", id)
// 		}
// 	}
// 	return nil
// }

type DBStorage struct {
	db *sql.DB
}

func (db *DBStorage) GetAllUsers() ([]User, error) {
	// db.db, err := sql.Open("mysql", "root:rock@tcp(127.0.0.1:55000)/trygodb")
	// defer db.db.Close()
	//
	// if err != nil {
	// 	fmt.Errorf("DB Connectivity error: %v", err)
	// }
	query, err := db.db.Query("SELECT * FROM users")
	if err != nil {
		fmt.Errorf("Oh no: %v", err)
	}
	defer query.Close()

	var result []User
	for query.Next() {
		user := new(User)
		err := query.Scan(&user.ID, &user.Name, &user.Surname, &user.Age)
		if err != nil {
			return nil, fmt.Errorf("Oh no, cant retreive list of users from DB: %v", err)
		}
		result = append(result, *user)
	}
	return result, nil
}

func (db *DBStorage) GetUser(id string) (*User, error) {
	user := new(User)
	query := db.db.QueryRow("SELECT id, name, surname, age FROM users WHERE id=?", id)

	err := query.Scan(&user.ID, &user.Name, &user.Surname, &user.Age)
	if err != nil {
		return nil, fmt.Errorf("Oh no, cant retreive list of users from DB: %v", err)
	}
	return user, nil
}

func (db *DBStorage) SaveUser(u *User) error {
	_, err := db.db.Exec("INSERT INTO users(id, name, surname, age) VALUES ($1,$2,$3,$4)", genUUID(), u.Name, u.Surname, u.Age)
	if err != nil {
		return fmt.Errorf("Oh no, cant insert user into DB: %v", err)
	}
	return nil
}

func (db *DBStorage) DeleteUser(id string) error {
	_, err := db.db.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("Oh no, cant delete user from DB: %v", err)
	}
	return nil
}

// func NewInMemoryStorage() Storage {
// 	 	return &UserStorage{
// 		storage: []User{
// 			{genUUID(), "Shredder", "Human", 32},
// 			{genUUID(), "Mike", "Turtle", 18},
// 			{genUUID(), "Splinter", "Rat", 300},
// 			{genUUID(), "Leo", "Turtle", 18},
// 		},
// 	}
// }

// CREATE TABLE users(id VARCHAR(255) PRIMARY KEY, name VARCHAR(255), surname VARCHAR(255), age INT);

func NewSQLStorage() (Storage, error) {
	db, err := sql.Open("mysql", "root:rock@tcp(127.0.0.1:55000)/trygodb")

	if err != nil {
		fmt.Errorf("DB Connectivity error: %v", err)
	}

	query := `
USE trygodb;
DROP TABLE IF EXISTS users;
INSERT INTO users(id, name, surname, age ) VALUES (?, 'Shredder', 'Human', 32);
INSERT INTO users(id, name, surname, age ) VALUES (?, 'Mike', 'Turtle', 18);
INSERT INTO users(id, name, surname, age ) VALUES (?, 'Leo', 'Turtle', 18);
INSERT INTO users(id, name, surname, age ) VALUES (?, 'Splinter', 'Rat', 300);
	`
	res, err := db.Query(query, genUUID(), genUUID(), genUUID(), genUUID())
	res.Close()

	if err != nil {
		return nil, err
	}
	return &DBStorage{db: db}, nil
}
