package parkinglot

import "fmt"

type Level struct {
	floor        int
	parkingSpots []*ParkingSpot
}

func NewLevel(floor int, spotsPerVehicleType map[VehicleType]int) *Level {
	level := &Level{floor: floor}

	for vehicleType, count := range spotsPerVehicleType {
		for i := 0; i < count; i++ {
			level.parkingSpots = append(level.parkingSpots, NewParkingSpot(fmt.Sprintf("F%d%c%d", floor, vehicleType.String()[0], i), vehicleType))
		}
	}

	return level
}

func (l *Level) DisplayAvailability() {
	for _, spot := range l.parkingSpots {
		status := "Available"
		if !spot.IsAvailable() {
			status = "Occupied"
		}
		println("Level:", l.floor, "Spot:", spot.GetSpotId(), "Status:", status, "Type:", spot.GetVehicleType().String())
	}
}

func (l *Level) ParkVehicle(vehicle Vehicle) (bool, *ParkingSpot) {
	for _, spot := range l.parkingSpots {
		isParked, parkingSpot := spot.ParkVehicle(vehicle)
		if isParked {
			return true, parkingSpot
		}
	}
	return false, nil
}

func (l *Level) UnparkVehicle(vehicle Vehicle) (bool, *ParkingSpot) {
	for _, spot := range l.parkingSpots {
		isUnparked, parkingSpot := spot.UnparkVehicle(vehicle)
		if isUnparked {
			return true, parkingSpot
		}
	}
	return false, nil
}
