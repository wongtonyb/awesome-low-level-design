package parkinglot

func Run() {
	parkingLot := GetParkingLotInstance()
	parkingLot.AddLevel(NewLevel(1, map[VehicleType]int{
		CAR:        5,
		MOTORCYCLE: 3,
		TRUCK:      2,
	}))
	parkingLot.AddLevel(NewLevel(2, map[VehicleType]int{
		CAR:        1,
		MOTORCYCLE: 2,
		TRUCK:      2,
	}))

	car1 := NewCar("car1")
	motocycle1 := NewMotorcycle("moto1")
	truck1 := NewTruck("truck2")
	truck2 := NewTruck("truck3")
	truck3 := NewTruck("truck2")

	parkingLot.DisplayAvailability()

	parkingLot.ParkVehicle(car1)
	parkingLot.ParkVehicle(motocycle1)
	parkingLot.ParkVehicle(truck1)
	parkingLot.ParkVehicle(truck2)
	parkingLot.ParkVehicle(truck3)

	parkingLot.DisplayAvailability()

	parkingLot.UnparkVehicle(car1)

	parkingLot.DisplayAvailability()
}
