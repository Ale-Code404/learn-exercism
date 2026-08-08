package jedlik

import (
	"fmt"
	"math"
)

// TODO: define the 'Drive()' method
func (car *Car) Drive() {
	if car.battery >= car.batteryDrain {
		car.distance += car.speed
		car.battery -= car.batteryDrain
	}
}

// DisplayBattery shows the current distance driven by the car
func (car *Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", car.distance)
}

// DisplayBattery shows the current battery percentage
func (car *Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", car.battery)
}

// CanFinish checks if the car can finish a specific race distance
func (car *Car) CanFinish(distance int) bool {
	needed := math.Ceil(float64(distance)/float64(car.speed)) * float64(car.batteryDrain)
	return car.battery >= int(needed)
}
