package models

import "time"

type Pagination struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type NotificationResponse struct {
	UserID    int       `json:"user_id" gorm:"column:liked_user_id"`
	Username  string    `json:"username"`
	Profile   string    `json:"profile"`
	PostID    int       `json:"post_id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}

type Notification struct {
	UserID    int       `json:"liked_user_id" gorm:"column:liked_user_id"`
	PostID    int       `json:"post_id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}

type NotificationReq struct {
	UserID      int       `json:"user_id"`
	LikedUserID int       `json:"liked_user_id"`
	PostID      int       `json:"post_id"`
	Message     string    `json:"message"`
	CreatedAt   time.Time `json:"created_at"`
}
