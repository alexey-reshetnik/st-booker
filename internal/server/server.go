package server

import (
	"st-booker/internal/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Server struct {
	router  *gin.Engine
	service *service.Booking
	log     *zap.Logger
}

func (s *Server) InitRoutes() {
	s.router.GET("/api/v1/booking/all", s.CreateBooking)
	s.router.POST("/api/v1/booking", s.CreateBooking)
	s.router.DELETE("/api/v1/booking", s.CreateBooking)
}

func NewServer(gin *gin.Engine, service *service.Booking, logger *zap.Logger) *Server {
	return &Server{
		router:  gin,
		service: service,
		log:     logger,
	}
}
