package model

import (
	"errors"
	"time"

	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

//role
//1=admin
//2=sppv
//3=user

type User struct {
	ID        uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Email     string `gorm:"size:255;not null;unique" json:"email"`
	Password  string `gorm:"size:255" json:"password,omitempty"`
	Role      int64  `gorm:"not null;default:3" json:"role"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (s *Server) ValidateEmail(email string) error {
	if email == "" {
		return errors.New("required email")
	}
	if email != "" {
		if err := checkmail.ValidateFormat(email); err != nil {
			return errors.New("invalid email")
		}
	}
	return nil
}

func (s *Server) CreateUser(user *User) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	emailErr := s.ValidateEmail(user.Email)
	if emailErr != nil {
		return emailErr
	}
	newUser := &User{
		Email:    user.Email,
		Password: string(passwordHash),
		Role:     user.Role,
	}
	err1 := s.DB.Debug().Create(&newUser).Error
	if err1 != nil {
		return err1
	}
	return nil
}

func (s *Server) UpdateUser(user *User) (*User, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	emailErr := s.ValidateEmail(user.Email)
	if emailErr != nil {
		return nil, emailErr
	}
	newUser := &User{
		Email:    user.Email,
		Password: string(passwordHash),
		Role:     user.Role,
	}
	err1 := s.DB.Debug().Model(&user).Where("id = ?", user.ID).Update(&newUser)
	if err1 != nil {
		if gorm.IsRecordNotFoundError(err) {
			errors.New("User not found")
		}
		return &User{}, err
	}
	return user, nil
}

func (s *Server) GetUserByEmail(email string) (*User, error) {
	user := &User{}
	err := s.DB.Debug().Where("email = ?", email).Take(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
