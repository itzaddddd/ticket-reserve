package repository

import (
	"context"

	"github.com/itzaddddd/ticket-reserve/entity"
	"gorm.io/gorm"
)

type repository struct {
	ctx *context.Context
	db  *gorm.DB
}

type Repository interface {
	Transaction(fn func(repo Repository) error) error
	WithTx(tx *gorm.DB) Repository
	GetEvent(eventId uint) (event entity.Event, err error)
	CreateEvent(eventReq entity.Event) (err error)
	Increseuota(eventId uint, numToIncrese uint) (err error)
	DecreseQuota(eventId uint, numToDecrese uint) (err error)
	InsertTicketReserveLog(userId, eventId uint, numberReserved uint) (err error)
}

func NewRepository(db *gorm.DB) *repository {
	ctx := context.Background()

	return &repository{
		ctx: &ctx,
		db:  db,
	}
}

func (r *repository) WithTx(tx *gorm.DB) Repository {
	return &repository{
		db: tx,
	}
}

func (r *repository) Transaction(fn func(repo Repository) error) error {
	tx := r.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	repo := r.WithTx(tx)
	err := fn(repo)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
