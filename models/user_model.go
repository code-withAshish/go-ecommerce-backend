package models

import (
	"github.com/google/uuid"
	"time"
)

//type User struct {
//	ID   uuid.UUID `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4();"`
//	Name string    `json:"name"`
//}

type User struct {
	ID        uuid.UUID `json:"id" gorm:"primary_key;type:uuid;default:gen_random_uuid();"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
