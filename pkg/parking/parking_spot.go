package parking

type ParkingSpot struct {
	ID         int
	Type       VehicleType
	IsOccupied bool
	Vehicle    *Vehicle
}

func NewParkingSpot(id int, vt VehicleType) *ParkingSpot {

	return &ParkingSpot{
		ID:         id,
		Type:       vt,
		IsOccupied: false,
		Vehicle:    &Vehicle{},
	}

}

func (ps *ParkingSpot) IsAvailable() bool {

	return ps.Vehicle == nil

}

func (ps *ParkingSpot) Park(v Vehicle) {
	if ps.IsAvailable() && ps.Type == v.VehicleType {
		ps.Vehicle = &v
		ps.IsOccupied = true
	}

}

func (ps *ParkingSpot) UnPark(v Vehicle) {
	if ps.IsOccupied {
		ps.Vehicle = nil
		ps.IsOccupied = false
	}

}
