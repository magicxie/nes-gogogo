package main

import . "nes6502/nes"

func main() {

	//TODO read configuration
	mainframe := &NES{}
	mainframe.PowerOn()

	//TODO screen detection and output
	//TODO fake PnP detection
	//TODO wait for rom

}
