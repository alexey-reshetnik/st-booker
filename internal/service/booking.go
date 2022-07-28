package service

import (
	"context"

	"st-booker/internal/model"
)

func (c *Booking) CreateBooking(ctx context.Context, booking model.Booking) (string, error) {
	id, err := c.storage.CreateBooking(ctx, booking)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (c *Booking) GetBookings(ctx context.Context, limit, offset int) ([]model.Booking, error) {
	bookings, err := c.storage.GetBookings(ctx, limit, offset)
	if err != nil {
		return []model.Booking{}, err
	}

	return bookings, nil
}

func (c *Booking) DeleteBooking(ctx context.Context, id string) error {
	return c.storage.DeleteBooking(ctx, id)
}
