package auth

import (
	"backend_relawanku/model"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primaryKey"`
	Username  string         `gorm:"unique;not null" json:"username" form:"username"`
	Email     string         `gorm:"email;not null" json:"email" form:"email"`
	Password  string         `gorm:"not null" json:"password" form:"password"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func FromModel(user model.User) User {
	return User{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}
}

func (user User) ToModel() model.User {
	return model.User{
		Model: gorm.Model{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			DeletedAt: user.DeletedAt,
		},
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}
}
