package entity

import "time"

type Notification interface {
	IsNotification()
}
type BaseNotification struct {
	CreateAt time.Time `json:"create_at"`
}

func (BaseNotification) IsNotification() {}

type UnreadWorkRequest struct {
	BaseNotification
	WorkID int    `json:"work_id"`
	Title  string `json:"title"`
}

type UnreadMessageNotification struct {
	BaseNotification
	Count int `json:"count"`
}
