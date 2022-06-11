package models

import (
	// "gorm.io/datatypes"
	"gorm.io/gorm"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// MODEL
type User struct {
	Id			uuid.UUID	`gorm:"type:uuid;" json:"id"`
	Username	string		`gorm:"unique;not null;" json:"username"`
	Password	string		`gorm:"not null;" json:"password"`

	Name		string		`json:"name"`
	Email		string		`json:"email"`

	// status
	IsPremium	bool		`gorm:"not null;" json:"isPremium"`
	IsAdmin		bool		`gorm:"not null;" json:"isAdmin"`
}
// constructor
func (u *User) BeforeCreate(db *gorm.DB) error {
	// uuid
	u.Id = uuid.New()
	// json update [in case its necessary]
	u.IsPremium = false
	u.IsAdmin = false

	return nil
}

// methods
func (u *User) Validate() []string {
	var errors []string
	// enum validation

	return errors
}
func (u *User) HashPassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), 4)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	return nil
}
func (u *User) ValidPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}


// SCHEMA
type UserResponse struct {
	Username	string		`json:"username"`
	
	Name		string		`json:"name"`
	Email		string		`json:"email"`
}
func (u *User) Response() *UserResponse {
	return &UserResponse {
		Username:	u.Username,

		Name:		u.Name,
		Email:		u.Email,
	}
}