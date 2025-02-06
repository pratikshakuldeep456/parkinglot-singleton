package parking

import (
	"fmt"
	"sync"
	"time"
)

type ParkingLot struct {
	Levels  []*Level
	Tickets map[string]*Ticket
	Gate    *Gate
	Lock    sync.Mutex
}

var ParkingLotInstance *ParkingLot

func GetInstance() *ParkingLot {
	if ParkingLotInstance == nil {
		ParkingLotInstance = NewParkingLot()

	}
	return ParkingLotInstance
}

func NewParkingLot() *ParkingLot {
	return &ParkingLot{
		Levels:  []*Level{},
		Tickets: make(map[string]*Ticket),
	}
}

func (pl *ParkingLot) AddLevel(level *Level) {
	pl.Levels = append(pl.Levels, level)
}

func (pl *ParkingLot) ParkVehicle(v *Vehicle) bool {
	pl.Lock.Lock()
	defer pl.Lock.Unlock()
	for _, l := range pl.Levels {
		spot, parked := l.ParkVehicle(v)
		if parked {
			ticket := &Ticket{
				ID:          fmt.Sprintf("TICKET-%s", v.LicencePlate),
				ParkingSpot: spot,
				VehicleNo:   v.LicencePlate,
				EntryTime:   time.Now(),
			}
			pl.Tickets[v.LicencePlate] = ticket
			fmt.Println("Ticket Generated:", ticket.ID)
			return true
		}

		// if l.ParkVehicle(v) {
		// 	return true
		// }
	}

	return false
}

func (pl *ParkingLot) UnParkVehicle(v *Vehicle) bool {
	pl.Lock.Lock()
	defer pl.Lock.Unlock()

	ticket, exists := pl.Tickets[v.LicencePlate]
	if !exists {
		fmt.Println("no vehicle found for ", v.LicencePlate)
	}

	for _, l := range pl.Levels {
		if l.UnParkVehicle(v) {
			delete(pl.Tickets, v.LicencePlate)
			fee := pl.CalculateFee(ticket)
			fmt.Printf("Vehicle %s exited. Fee: $%.2f\n", v.LicencePlate, fee)
			return true
		}
	}
	return false
}
func (pl *ParkingLot) Availablility() {
	for _, l := range pl.Levels {
		l.CheckAvailability()
	}
}

func (pl *ParkingLot) CalculateFee(t *Ticket) int {
	duration := time.Since(t.EntryTime).Hours()
	return int(duration) * 4.0
}
