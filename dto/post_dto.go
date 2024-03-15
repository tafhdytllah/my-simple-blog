package dto

import "mime/multipart"

type PostResponse struct {
	ID         int    `json:"id"`
	UserID     int    `json:"-"`
	User       User   `gorm:"foreignKey:UserId" json:"user"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	PictureUrl string `json:"picture_url"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

type PostRequest struct {
	UserID  int                   `form:"user_id"`
	Title   string                `form:"title"`
	Content string                `form:"content"`
	Picture *multipart.FileHeader `form:"picture"`
}

type User struct {
	ID    int `json:"id"`
	Email int `json:"email"`
	Name  int `json:"name"`
}
