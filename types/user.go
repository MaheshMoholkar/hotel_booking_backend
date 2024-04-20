package types

import (
	"fmt"
	"regexp"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

const (
	bcryptCost   = 12
	minFirstName = 2
	minLastName  = 2
	minPassword  = 6
)

type PostUserParams struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type User struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName         string             `bson:"firstName" json:"firstName"`
	LastName          string             `bson:"lastName" json:"lastName"`
	Email             string             `bson:"email" json:"email"`
	EncryptedPassword string             `bson:"EncryptedPassword" json:"-"`
}

func (params PostUserParams) Validate() []string {
	errors := []string{}
	if len(params.FirstName) < minFirstName {
		errors = append(errors, fmt.Sprintf("firstName length should be at least %d characters", minFirstName))
	}
	if len(params.LastName) < minLastName {
		errors = append(errors, fmt.Sprintf("lastName length should be at least %d characters", minLastName))
	}
	if len(params.Password) < minPassword {
		errors = append(errors, fmt.Sprintf("password length should be at least %d characters", minPassword))
	}
	if !isEmailValid(params.Email) {
		errors = append(errors, "email is invalid")
	}
	return errors
}

func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}

func NewUserFromParams(params PostUserParams) (*User, error) {
	encpw, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcryptCost)
	if err != nil {
		return nil, err
	}
	return &User{
		FirstName:         params.FirstName,
		LastName:          params.LastName,
		Email:             params.Email,
		EncryptedPassword: string(encpw),
	}, nil
}
