package parkinglot

// enumerated type represented as integers
type VehicleType int

const (
	CAR        VehicleType = iota //VehicleType.CAR = 0
	MOTORCYCLE                    // VehicleType.MOTORCYCLE = 1
	TRUCK
)

func (v VehicleType) String() string {
	switch v {
	case CAR:
		return "CAR"
	case MOTORCYCLE:
		return "MOTORCYCLE"
	case TRUCK:
		return "TRUCK"
	default:
		return "UNKNOWN"
	}
}

// interface (behavior) for vehicles
type Vehicle interface {
	GetLicensePlate() string
	GetType() VehicleType
}

// struct (data) for vehicles
type BaseVehicle struct {
	licensePlate string
	vehicleType  VehicleType
}

func (v BaseVehicle) GetLicensePlate() string {
	return v.licensePlate
}

func (v BaseVehicle) GetType() VehicleType {
	return v.vehicleType
}
