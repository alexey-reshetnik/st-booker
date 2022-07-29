package service

import (
	"context"
	"time"

	"st-booker/internal/model"

	"go.uber.org/zap"
)

type Storage interface {
	CreateBooking(ctx context.Context, booking model.Booking) (string, error)
	GetBookings(ctx context.Context, limit, offset int) ([]model.Booking, error)
	DeleteBooking(ctx context.Context, id string) error
}

type SpaceXClient interface {
	LaunchesForDate(launchpadID string, date time.Time) (int, error)
}

type Booking struct {
	storage Storage
	client  SpaceXClient
	logger  *zap.Logger
}

func NewBooking(storage Storage, client SpaceXClient, logger *zap.Logger) *Booking {
	return &Booking{storage: storage, client: client, logger: logger}
}
