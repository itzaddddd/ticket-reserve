package server

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/itzaddddd/ticket-reserve/config"
	"github.com/itzaddddd/ticket-reserve/db"
	"github.com/itzaddddd/ticket-reserve/handler"
	"github.com/itzaddddd/ticket-reserve/repository"
	"github.com/itzaddddd/ticket-reserve/service"

	"gorm.io/gorm"
)

type Server struct {
	App *gin.Engine
	Db  *gorm.DB
	Cfg *config.Config
}

func NewServer() *Server {
	return &Server{
		App: gin.Default(),
	}
}

func defaultLog(ctx *gin.Context) {
	log.Printf("calling [%s] %s", ctx.Request.Method, ctx.Request.URL)

	ctx.Next()
}

func (s *Server) SetMiddleware() {
	s.App.Use(defaultLog)
	s.App.Use(cors.Default())
}

func (s *Server) SetConfig() {
	cfg := config.NewConfig()
	s.Cfg = cfg
}

func (s *Server) SetDb() {
	db := db.NewConn(s.Cfg)
	s.Db = db
}

func (s *Server) ShutdownDb() error {
	db, err := s.Db.DB()
	if err != nil {
		return err
	}

	if err := db.Close(); err != nil {
		return err
	}

	return nil
}

func (s *Server) SetValidator() {
}

func (s *Server) SetHandler() {

	repo := repository.NewRepository(s.Db)
	service := service.NewService(repo)
	handler := handler.NewHandler(service)
	eventHandler := s.App.Group("/event")
	eventHandler.POST("/reserve", handler.ReserveTicketHandler)
}
