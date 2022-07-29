package entity

import "time"

type User struct {
	ID        int               `json:"id" gorm:"primary_key:auto_increment"`
	Name      string            `json:"name" gorm:"type: varchar(255)"`
	Email     string            `json:"email" gorm:"uniqueIndex;type:varchar(255)"`
	Password  string            `json:"-" gorm:"<-:create"`
	Status    string            `json:"status" gorm:"type: varchar(255)"`
	Profile   ProfileResponse   `json:"profile"`
	Products  []ProductResponse `json:"products"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
}

type UserResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (UserResponse) TableName() string {
	return "users"
}
