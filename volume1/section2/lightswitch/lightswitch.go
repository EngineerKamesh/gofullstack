package main

import "fmt"

var lightSwitchIsOn bool = false

func main() {

	printLightSwitchState()
	toggleLightSwitch()
	printLightSwitchState()
	toggleLightSwitch()
	printLightSwitchState()

}

func printLightSwitchState() {
	fmt.Println("The light switch is off:", lightSwitchIsOn)

}

func toggleLightSwitch() {

	lightSwitchIsOn = !lightSwitchIsOn

}
