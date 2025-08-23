package entities

import "time"

type User struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Name  string `json:"name" gorm:"not null"`
	Email string `json:"email" gorm:"unique;index;not null"`

	CreatedAt time.Time `json:"-" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"-" gorm:"autoUpdateTime"`
}

func (User) TableName() string {
	return "users"
}
