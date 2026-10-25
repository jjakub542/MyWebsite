package models

import (
	"time"

	"gorm.io/gorm"

	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type User struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Email     string         `gorm:"uniqueIndex;not null" json:"email"`
	Password  string         `gorm:"not null" json:"-"`
	IsAdmin   bool           `gorm:"default:false" json:"is_admin"`
}

func (u *User) Validate() error {
	return validation.ValidateStruct(u,
		validation.Field(&u.Email, validation.Required, is.Email, validation.Length(3, 255)),
		validation.Field(&u.Password, validation.Required, validation.Length(3, 99)),
	)
}

type Message struct {
	gorm.Model
	Name    string `gorm:"size:255;not null"`
	Email   string `gorm:"size:255;not null"`
	Content string `gorm:"type:text;not null"`
}
