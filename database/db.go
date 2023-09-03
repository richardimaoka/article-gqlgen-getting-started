package database

import "fmt"

type DbUser struct {
	Id        string
	Name      string
	Residence string
}

func GetUserFromDb(userId string) (*DbUser, error) {
	if userId == "abc123" {
		return &DbUser{
			Id:        "abc123",
			Name:      "Tarou MarsYama",
			Residence: "MARS",
		}, nil
	}

	return nil, fmt.Errorf("User not found for id = %s", userId)
}
