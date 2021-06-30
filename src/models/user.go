package models

import "time"

//User datos del usuario
type User struct {
	Name     string    `json:"name"`
	Surname  string    `json:"surname"`
	Email    string    `json:"email"`
	CreateAt time.Time `bson:"create_at" json:"createAt"`
	UpdateAt time.Time `bson:"update_at" json:"updateAt,omitempty"`
}

//Users arreglo de usuarios
type Users []*User
