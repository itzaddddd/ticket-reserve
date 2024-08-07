package entity

import "gorm.io/gorm"

type TicketReserveLog struct {
	gorm.Model
	UserId         uint `json:"user_id"`
	EventId        uint `json:"event_id"`
	Event          Event
	NumberReserved uint `json:"number_reserved"`
}
