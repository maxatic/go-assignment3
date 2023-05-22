package domain

import (
	"errors"
	"unicode/utf8"
)

type Group struct {
	Id   int
	Name string
}

func NewGroup(name string) (*Group, error) {
	nameLength := utf8.RuneCountInString(name)
	if nameLength > 250 {
		return nil, errors.New("invalid format of group name")
	}
	return &Group{Name: name}, nil
}
