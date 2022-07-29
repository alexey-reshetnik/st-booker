package service

import (
	"context"
	"errors"

	"st-booker/internal/model"
)

var ErrNoAvailableLaunchTime = errors.New("can't launch the rocket at this launchpad at given date")

func (b *Booking) CreateBooking(ctx context.Context, booking model.Booking) (string, error) {
	launchesCount, err := b.client.LaunchesForDate(booking.LaunchpadID, booking.LaunchDate)
	if err != nil {
		return "", err
	}

	if launchesCount != 0 {
		return "", ErrNoAvailableLaunchTime
	}

	id, err := b.storage.CreateBooking(ctx, booking)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (b *Booking) GetBookings(ctx context.Context, limit, offset int) ([]model.Booking, error) {
	bookings, err := b.storage.GetBookings(ctx, limit, offset)
	if err != nil {
		return []model.Booking{}, err
	}

	return bookings, nil
}

func (b *Booking) DeleteBooking(ctx context.Context, id string) error {
	return b.storage.DeleteBooking(ctx, id)
}
