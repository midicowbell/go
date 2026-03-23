package main

import (
	"smartHome/smarthome"
)

func main() {
	lamp := smarthome.Lamp{}
	cond := smarthome.AirConditioner{}
	smarthome.ManageDevice(lamp)
	smarthome.ManageDevice(cond)

}
