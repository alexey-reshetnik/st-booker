package storage

import (
	"context"

	"st-booker/internal/model"
)

func (c *Client) CreateBooking(ctx context.Context, booking model.Booking) (string, error) {
	return "", nil
}

func (c *Client) GetBookings(ctx context.Context, limit, offset int) ([]model.Booking, error) {
	return []model.Booking{}, nil
}

func (c *Client) DeleteBooking(ctx context.Context, id string) error {
	return nil
}
