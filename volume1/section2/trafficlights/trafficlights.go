package main

import (
	"fmt"
)

const (
	_ = iota
	TrafficLightStateRedLight
	TrafficLightStateGreenLight
	TrafficLightStateYellowLight
)

func main() {

	fmt.Println("Red Light State Code: ", TrafficLightStateRedLight)
	fmt.Println("Green Light State Code: ", TrafficLightStateGreenLight)
	fmt.Println("Yellow Light State Code: ", TrafficLightStateYellowLight)
}
