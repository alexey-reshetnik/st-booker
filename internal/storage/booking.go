package storage

import (
	"context"

	"st-booker/internal/model"
)

const (
	InsertBookingQuery = `INSERT INTO bookings (first_name, last_name, gender, birthday, launchpad_id, destination_id, launch_date) VALUES ($1, $2, $3, $4, $5, $6, $7) returning id;`
	SelectBookingQuery = "SELECT * FROM bookings WHERE id = $1;"
	DeleteBookingQuery = "DELETE * FROM bookings WHERE id = $1;"
)

func (c *Client) CreateBooking(ctx context.Context, booking model.Booking) (string, error) {
	var id string

	row := c.conn.QueryRow(InsertBookingQuery, booking)
	if err := row.Scan(&id); err != nil {
		return "", err
	}

	return id, nil
}

func (c *Client) GetBookings(ctx context.Context, limit, offset int) ([]model.Booking, error) {
	var booking []model.Booking

	err := c.conn.Get(&booking, SelectBookingQuery, booking)

	return booking, err
}

func (c *Client) DeleteBooking(ctx context.Context, id string) error {
	_, err := c.conn.ExecContext(ctx, DeleteBookingQuery, id)
	if err != nil {
		return err
	}

	return err
}
