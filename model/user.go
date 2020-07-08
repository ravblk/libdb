package model

import (
	"time"
)

type User struct {
	ID        int       `json:"id" db:"id" gorm:"primary_key,AUTO_INCREMENT"`
	Email     string    `json:"email" db:"email"`
	Username  string    `json:"username" db:"username"`
	Password  string    `json:"password" db:"password"`
	Status    int       `json:"status" db:"status"`
	Age       int       `json:"age" db:"age"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
