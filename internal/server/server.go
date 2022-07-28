package server

import (
	"st-booker/internal/service"

	"go.uber.org/zap"
)

type Server struct {
	service service.Booking
	log     *zap.Logger
}

func NewBookingController(service service.Booking, logger *zap.Logger) *Server {
	return &Server{service: service, log: logger}
}
