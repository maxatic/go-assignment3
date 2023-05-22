package domain

import (
	"errors"
	"regexp"
)

type Contact struct {
	Id          int
	fullName    string
	PhoneNumber string
}

func NewContact(id int, fullName string, phoneNumber string) (*Contact, error) {
	validPhoneNum := regexp.MustCompile(`^[\\d]+$`).MatchString(phoneNumber)
	validFullName := regexp.MustCompile(`^[A-Za-z]+\\s[A-Za-z]+\\s[A-Za-z]+$`).MatchString(fullName)

	if !validPhoneNum || !validFullName {
		return nil, errors.New("invalid format")
	}
	return &Contact{Id: id, fullName: fullName, PhoneNumber: phoneNumber}, nil
}

func (c *Contact) FullName() string {
	return c.fullName
}
