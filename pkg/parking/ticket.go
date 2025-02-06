package parking

import "time"

type Ticket struct {
	ID          string
	ParkingSpot *ParkingSpot
	VehicleNo   string
	EntryTime   time.Time
}
