package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/itzaddddd/ticket-reserve/util/response"
)

type CreateEventReq struct {
	Name  string `json:"event" validate:"required,max=100"`
	Quota uint   `json:"quota" validate:"required,number,min=0,max=100"`
}

func (h *Handler) CreateEventHandler(ctx *gin.Context) {
	var req CreateEventReq

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

}

type GetEventReq struct {
	EventId uint `uri:"event_id" validate:"number,min=0"`
}

func (h *Handler) GetEventHandler(ctx *gin.Context) {
	var req GetEventReq

	if err := ctx.ShouldBindUri(req); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

}

type ReserveTicketReq struct {
	UserId          uint `json:"user_id" validate:"required,number,min=0"`
	EventId         uint `json:"event_id" validate:"required,number,min=0"`
	NumberToReserve uint `json:"number_to_reserve" validate:"required,number,min=0,max=10"`
}

func (h *Handler) ReserveTicketHandler(ctx *gin.Context) {
	var req ReserveTicketReq

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	if err := h.Service.ReserveTicket(req.UserId, req.EventId, req.NumberToReserve); err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(err))
	}

}
