// Chaptr.2 arithmetic operators
package main

import "fmt"

func main() {
	//Type.1
	fmt.Print("My weight on the surface of Mars is ")
	fmt.Print(95.0 * 0.3783)
	fmt.Print(" Kg, and I would be ")
	fmt.Print(30 * 365 / 687)
	fmt.Print(" years old.\n")

	//Type.2 Use Oneline
	fmt.Println("My weight on the surface of Mars is", 95.0*0.3783,
		"Kg, and I would be", 30*365.2425/687, "years old.")

	//Type.3 Use Fromat
	fmt.Printf("My weight on the surface of Mars is %v lbs,", 149.0*0.3783)
	fmt.Printf(" and I would be %v years old.\n", 41*365/687)

	//Fromat with padding
	fmt.Printf("%-15v $%4v\n", "SpaceX", 94)
	fmt.Printf("%-15v $%4v\n", "Virgin Galactic", 100)

}
