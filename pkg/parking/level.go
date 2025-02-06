package parking

import "fmt"

type Level struct {
	Floor        int
	ParkingSpots []*ParkingSpot
}

func NewLevel(floor int, noSpots int) *Level {

	level := &Level{Floor: floor}
	bikeSpot := (noSpots * 40) / 100
	carSpot := (noSpots * 40) / 100

	truckSpot := noSpots - bikeSpot - carSpot
	spotNo := 1
	for i := 0; i < bikeSpot; i++ {
		level.ParkingSpots = append(level.ParkingSpots, &ParkingSpot{ID: spotNo, Type: Motorcycle})
		spotNo++
	}

	for i := 0; i < carSpot; i++ {
		level.ParkingSpots = append(level.ParkingSpots, &ParkingSpot{ID: spotNo, Type: Car})
		spotNo++
	}

	for i := 0; i < truckSpot; i++ {
		level.ParkingSpots = append(level.ParkingSpots, &ParkingSpot{ID: spotNo, Type: Truck})
		spotNo++
	}
	return level
}

func (l *Level) ParkVehicle(v *Vehicle) (*ParkingSpot, bool) {
	for _, spot := range l.ParkingSpots {
		if spot.IsAvailable() && spot.Type == v.VehicleType {
			fmt.Println("parking at", spot.ID)
			spot.IsOccupied = true

			spot.Vehicle = v
			return spot, true
		}
	}
	return nil, false
}

func (l *Level) UnParkVehicle(v *Vehicle) bool {
	for _, spot := range l.ParkingSpots {
		if spot.IsOccupied {
			spot.IsOccupied = false
			spot.Vehicle = nil

			return true
		}
	}
	return false
}

func (l *Level) CheckAvailability() []*ParkingSpot {
	var result []*ParkingSpot
	for _, spot := range l.ParkingSpots {
		if spot.IsAvailable() {
			result = append(result, spot)
			fmt.Println(" spot is available for", spot, spot.ID)
		}
	}
	return result
}
