package model

import "time"

type Booking struct {
	ID            string    `json:"id"`
	Status        string    `json:"status"`
	StatusReason  string    `json:"statusReason"`
	FirstName     string    `json:"firstName" binding:"required"`
	LastName      string    `json:"lastName" binding:"required"`
	Gender        string    `json:"gender" binding:"required"`
	Birthday      time.Time `json:"birthday" binding:"required"`
	LaunchpadID   string    `json:"launchpadId" binding:"required"`
	DestinationID string    `json:"destinationId" binding:"required"`
	LaunchDate    time.Time `json:"launchDate" binding:"required"`
}
