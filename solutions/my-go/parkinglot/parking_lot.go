package parkinglot

type ParkingLot struct {
	levels []*Level
}

// Singleton implmentation
// defined globally as a package-level variable
var instance *ParkingLot

// also defined globally
func GetParkingLotInstance() *ParkingLot {
	if instance == nil {
		instance = &ParkingLot{levels: []*Level{}}
	}
	return instance
}

func (p *ParkingLot) AddLevel(level *Level) {
	p.levels = append(p.levels, level)
}

func (p *ParkingLot) DisplayAvailability() {
	for _, level := range p.levels {
		level.DisplayAvailability()
	}
}

func (p *ParkingLot) ParkVehicle(vehicle Vehicle) (bool, *ParkingSpot) {
	for _, level := range p.levels {
		isParked, parkingSpot := level.ParkVehicle(vehicle)
		if isParked {
			println("Successfully parked at", parkingSpot.GetSpotId())
			return true, parkingSpot
		}
	}
	println("Unable to park")
	return false, nil
}

func (p *ParkingLot) UnparkVehicle(vehicle Vehicle) (bool, *ParkingSpot) {
	for _, level := range p.levels {
		isUnparked, parkingSpot := level.UnparkVehicle(vehicle)
		if isUnparked {
			println("Successfully unparked at", parkingSpot.GetSpotId())
			return true, parkingSpot
		}
	}
	println("Unable to find car in parking lot")
	return false, nil
}
