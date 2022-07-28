package service

import (
	"context"

	"st-booker/internal/model"

	"go.uber.org/zap"
)

type Storage interface {
	CreateBooking(ctx context.Context, booking model.Booking) (string, error)
	GetBookings(ctx context.Context, limit, offset int) ([]model.Booking, error)
	DeleteBooking(ctx context.Context, id string) error
}

type Booking struct {
	storage Storage
	logger  *zap.Logger
}

func NewBooking(storage Storage, logger *zap.Logger) *Booking {
	return &Booking{storage: storage, logger: logger}
}
