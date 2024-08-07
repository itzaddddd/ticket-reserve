package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/itzaddddd/ticket-reserve/service"
)

type Handler struct {
	Service   *service.Service
	Validator *validator.Validate
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		Service:   service,
		Validator: validator.New(),
	}
}
