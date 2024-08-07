package repository

import (
	"github.com/itzaddddd/ticket-reserve/entity"
	"gorm.io/gorm"
)

func (r *repository) GetEvent(eventId uint) (event entity.Event, err error) {
	if res := r.db.First(&event).Where("id = ?", eventId); res.Error != nil {
		err = res.Error
	}
	return
}

func (r *repository) CreateEvent(eventReq entity.Event) (err error) {
	res := r.db.Create(eventReq)
	err = res.Error
	return

}

func (r *repository) DecreseQuota(eventId uint, numToDecrese uint) (err error) {
	result := r.db.Model(entity.Event{}).Where("id = ?", eventId).Update("remain_quota", gorm.Expr("remain_quota - ?", numToDecrese)).Debug()
	err = result.Error
	return
}

func (r *repository) Increseuota(eventId uint, numToIncrese uint) (err error) {
	result := r.db.Model(entity.Event{}).Where("id = ?", eventId).Update("remain_quota", gorm.Expr("remain_quota + ?", numToIncrese))
	err = result.Error
	return
}

func (r *repository) InsertTicketReserveLog(userId, eventId uint, numberReserved uint) (err error) {
	log := entity.TicketReserveLog{
		UserId:         userId,
		EventId:        eventId,
		NumberReserved: numberReserved,
	}

	if res := r.db.Create(&log); res.Error != nil {
		err = res.Error
	}

	return
}
