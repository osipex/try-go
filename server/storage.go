package server

// type User struct {
// 	ID int // Think about how server will generate unique IDs
// 	Name string
// 	Surname string
// 	Age int
// }

type User struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Age     int    `json:"age"`
}

// Using inmemory slice as a storage storage solution. Slice of User object
type Users []User

var storage Users

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

	return out
}
