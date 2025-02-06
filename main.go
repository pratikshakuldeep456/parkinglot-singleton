package main

import "pratikshakuldeep456/parkinglot-singleton/pkg/parking"

func main() {
	parkingLot := parking.GetInstance()
	parkingLot.AddLevel(parking.NewLevel(1, 10))
	parkingLot.AddLevel(parking.NewLevel(2, 10))
	parkingLot.AddLevel(parking.NewLevel(3, 10))

	v1 := &parking.Vehicle{LicencePlate: "ABC123", VehicleType: parking.Car}
	parkingLot.Availablility()
	parkingLot.ParkVehicle(v1)
	parkingLot.Availablility()
}
