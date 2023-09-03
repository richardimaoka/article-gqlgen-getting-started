// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Date struct {
	Year  *int `json:"year,omitempty"`
	Month *int `json:"month,omitempty"`
	Day   *int `json:"day,omitempty"`
}

type NewTodo struct {
	Text   string `json:"text"`
	UserID string `json:"userId"`
}

type Todo struct {
	ID    string  `json:"id"`
	Title *string `json:"title,omitempty"`
	Text  string  `json:"text"`
	Done  bool    `json:"done"`
	User  *User   `json:"user"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
