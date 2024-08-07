package entity

import "gorm.io/gorm"

type Event struct {
	gorm.Model
	Name        string `json:"event"`
	Quota       uint   `json:"quota"`
	RemainQuota uint   `json:"remain_quota"`
}
