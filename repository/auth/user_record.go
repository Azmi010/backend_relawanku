package auth

import (
	"backend_relawanku/model"
	"time"

	"gorm.io/gorm"
)

type UserRole string

const (
	RoleUser  UserRole = "user"
	RoleAdmin UserRole = "admin"
)

type User struct {
	ID        uint           `gorm:"primaryKey"`
	Username  string         `gorm:"unique;not null" json:"username" form:"username"`
	Email     string         `gorm:"email;not null" json:"email" form:"email"`
	Password  string         `gorm:"not null" json:"password" form:"password"`
	Gender    string         `json:"gender" form:"gender"`
	Address   string         `json:"address" form:"address"`
	Role         UserRole            `gorm:"not null" json:"role"`
	ImageUrl  string         `json:"image_url" form:"image_url"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	Transactions []model.Transaction `gorm:"foreignKey:UserID" json:"user_id"`
}

func FromModelUser(user model.User) User {
	return User{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Password:  user.Password,
		Gender:    user.Gender,
		Address:   user.Address,
		ImageUrl:  user.ImageUrl,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}
}

func (user User) ToModelUser() model.User {
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
		Gender:   user.Gender,
		Address:  user.Address,
		ImageUrl: user.ImageUrl,
	}
}
