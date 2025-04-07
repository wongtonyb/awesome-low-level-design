package parkinglot

type ParkingSpot struct {
	spotId        string
	VehicleType   VehicleType
	parkedVehicle Vehicle
}

func NewParkingSpot(spotId string, vehicleType VehicleType) *ParkingSpot {
	return &ParkingSpot{spotId: spotId, VehicleType: vehicleType}
}

func (ps *ParkingSpot) IsAvailable() bool {
	return ps.parkedVehicle == nil
}

func (ps *ParkingSpot) GetSpotId() string {
	return ps.spotId
}

func (ps *ParkingSpot) GetVehicleType() VehicleType {
	return ps.VehicleType
}

func (ps *ParkingSpot) ParkVehicle(vehicle Vehicle) (bool, *ParkingSpot) {
	if ps.IsAvailable() && ps.VehicleType == vehicle.GetType() {
		ps.parkedVehicle = vehicle
		return true, ps
	}
	return false, ps
}

func (ps *ParkingSpot) UnparkVehicle(vehicle Vehicle) (bool, *ParkingSpot) {
	if !ps.IsAvailable() && ps.parkedVehicle == vehicle {
		ps.parkedVehicle = nil
		return true, ps
	}
	return false, ps
}
