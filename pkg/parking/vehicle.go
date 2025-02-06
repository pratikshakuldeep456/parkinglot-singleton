package parking

type VehicleType int

const (
	Car VehicleType = 0
	Motorcycle
	Truck
)

type Vehicle struct {
	LicencePlate string
	VehicleType  VehicleType
}
