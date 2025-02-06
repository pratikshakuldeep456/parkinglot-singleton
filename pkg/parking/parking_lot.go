package parking

type ParkingLot struct {
	Levels []*Level
}

var ParkingLotInstance *ParkingLot

func GetInstance() *ParkingLot {
	if ParkingLotInstance == nil {
		ParkingLotInstance = NewParkingLot()

	}
	return ParkingLotInstance
}

func NewParkingLot() *ParkingLot {
	return &ParkingLot{}
}

func (pl *ParkingLot) AddLevel(level *Level) {
	pl.Levels = append(pl.Levels, level)
}

func (pl *ParkingLot) ParkVehicle(v *Vehicle) bool {

	for _, l := range pl.Levels {
		if l.ParkVehicle(v) {
			return true
		}
	}

	return false
}

func (pl *ParkingLot) UnParkVehicle(v *Vehicle) bool {
	for _, l := range pl.Levels {
		if l.ParkVehicle(v) {
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
