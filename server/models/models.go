package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Username    string        `json:"username" gorm:"unique;not null" binding:"required"`
	Email       string        `json:"email" gorm:"unique;not null" binding:"required,email"`
	Password    string        `json:"password" gorm:"not null" binding:"required,min=6"`
	Age         int           `json:"age" gorm:"not null" binding:"required,min=9"`
	Photo       []Photo       `gorm:"foreignKey:UserId"`
	SocialMedia []SocialMedia `gorm:"foreignKey:UserId"`
	Comment     []Comment     `gorm:"foreignKey:UserId"`
}

type Photo struct {
	gorm.Model
	Title    string    `json:"title" gorm:"not null" binding:"required"`
	Caption  string    `json:"caption"`
	PhotoUrl string    `json:"photo_url" gorm:"not null" binding:"required"`
	UserId   int       `json:"user_id"`
	Comment  []Comment `gorm:"foreignKey:PhotoId"`
}

type Comment struct {
	gorm.Model
	Message string `json:"message" gorm:"not null" binding:"required"`
	UserId  int    `json:"user_id"`
	PhotoId int    `json:"photo_id"`
}

type SocialMedia struct {
	gorm.Model
	Name           string `json:"name" gorm:"not null" binding:"required"`
	SocialMediaUrl string `json:"social_media_url" gorm:"not null" binding:"required"`
	UserId         int    `json:"user_id"`
}

type ReqUser struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}

type Photos struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserId    int       `json:"user_id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type SocialMedias struct {
	Id             int       `json:"id"`
	Name           string    `json:"title"`
	SocialMediaUrl string    `json:"caption"`
	UserId         int       `json:"user_id"`
	Username       string    `json:"username"`
	Email          string    `json:"email"`
	UserIdUser     int       `json:"user_id_user"`
	PhotoUrl       string    `json:"photo_url"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type Comments struct {
	Id           int       `json:"id"`
	Message      string    `json:"message"`
	PhotoId      int       `json:"photo_id"`
	UserId       int       `json:"user_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	UserIdUser   int       `json:"user_id_user"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	PhotoIdPhoto int       `json:"photo_id_photo"`
	Title        string    `json:"title"`
	Caption      string    `json:"caption"`
	PhotoUrl     string    `json:"photo_url"`
	UserIdPhoto  int       `json:"user_id_photo"`
}
