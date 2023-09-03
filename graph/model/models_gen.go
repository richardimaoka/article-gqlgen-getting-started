// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

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
	ID        string  `json:"id"`
	Title     *string `json:"title,omitempty"`
	Text      string  `json:"text"`
	Done      bool    `json:"done"`
	User      *User   `json:"user"`
	CreatedAt *Date   `json:"createdAt,omitempty"`
	UpdatedAt *Date   `json:"updatedAt,omitempty"`
}

type User struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Residence *Residence `json:"residence,omitempty"`
}

type Residence string

const (
	ResidenceEarth Residence = "EARTH"
	ResidenceMoon  Residence = "MOON"
	ResidenceMars  Residence = "MARS"
)

var AllResidence = []Residence{
	ResidenceEarth,
	ResidenceMoon,
	ResidenceMars,
}

func (e Residence) IsValid() bool {
	switch e {
	case ResidenceEarth, ResidenceMoon, ResidenceMars:
		return true
	}
	return false
}

func (e Residence) String() string {
	return string(e)
}

func (e *Residence) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Residence(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Residence", str)
	}
	return nil
}

func (e Residence) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
