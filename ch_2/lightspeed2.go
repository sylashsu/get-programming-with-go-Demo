// The SpaceX Interplanetary Transport System lacks a warp drive,
// but it will coast to Mars at a respectable 100,800 km/h.
// An ambitious launch date of January 2025 would place Mars and Earth 96,300,000 km apart.
// How many days would it take to reach Mars?
package main

import "fmt"

func main() {
	const perHoursDay = 24
	/*
		const lightSpeed = 100800 // km/h
		var distance = 96300000   // km
	*/

	// Declare multiple variables type.1
	var (
		distance   = 56000000
		lightSpeed = 100800
	)

	// Declare multiple variables type.2
	/*
		var distance, lightSpeed = 56000000, 100800
	*/

	fmt.Println(distance/lightSpeed/perHoursDay, "Days")
}
