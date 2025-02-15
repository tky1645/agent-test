package entities

import "fmt"

type User struct {
	ID   int    `json:"id"`
	Name userName `json:"name"`
}


type userName string
func newUserName(name string) (userName, error) {
	if name == "" {
		return "",fmt.Errorf("name is empty")
	}

	return userName(name),nil
}

// factory method
func NewUser(id int, name string)( User, error) {
	userName, err := newUserName(name)
	if err != nil {
		return User{}, err
	}

	return User{
		ID:   id,
		Name: userName,
	},nil
}
