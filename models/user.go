package models

// User type represents a person using the system.
type User struct {
	Name   string `json:"name" datastore:"name,noindex"`
	Gender string `json:"gender" datastore:"gender,noindex"`
	Age    int    `json:"age" datastore:"age,noindex"`
	ID     string `json:"id" datastore:"id"`
}

// IsEmpty returns a boolean value representing if the object is empty.
func (u User) IsEmpty() bool {
	return u.Name == "" && u.Gender == "" && u.Age == 0 && u.ID == ""
}

// UserNotFoundError identifies when a user is not found
type UserNotFoundError struct {
	Message string
}

func (u UserNotFoundError) Error() string {
	return u.Message
}
