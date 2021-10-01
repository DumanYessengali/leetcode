package main

type ParkingSystem struct {
	car map[int]int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{car: map[int]int{
		1: big,
		2: medium,
		3: small,
	}}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	if this.car[carType] > 0 {
		this.car[carType] -= 1
		return true
	}
	return false
}

//
//obj := Constructor(big, medium, small)
//param_1 := obj.AddCar(carType)

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */
