package auth

import (
	"backend_relawanku/model"
	"time"

	"gorm.io/gorm"
)

type Admin struct {
	ID        uint           `gorm:"primaryKey"`
	Username  string         `gorm:"unique;not null" json:"username" form:"username"`
	Email     string         `gorm:"email;not null" json:"email" form:"email"`
	Password  string         `gorm:"not null" json:"password" form:"password"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func FromModelAdmin(admin model.Admin) Admin {
	return Admin{
		ID:        admin.ID,
		Username:  admin.Username,
		Email:     admin.Email,
		Password:  admin.Password,
		CreatedAt: admin.CreatedAt,
		UpdatedAt: admin.UpdatedAt,
		DeletedAt: admin.DeletedAt,
	}
}

func (admin Admin) ToModelAdmin() model.Admin {
	return model.Admin{
		Model: gorm.Model{
			ID:        admin.ID,
			CreatedAt: admin.CreatedAt,
			UpdatedAt: admin.UpdatedAt,
			DeletedAt: admin.DeletedAt,
		},
		Username: admin.Username,
		Email:    admin.Email,
		Password: admin.Password,
	}
}