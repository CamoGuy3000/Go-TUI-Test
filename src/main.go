package main

import (
	mouse "github.com/CamoGuy3000/mouse"
	screen "github.com/CamoGuy3000/screen"
)

func main() {

	thing := true
	if(thing){
		mouse.RunCords()
	} else {
		screen.RunAlt()
	}
}