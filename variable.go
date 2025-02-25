package main

import "fmt"

func main() {
	// var name string

	// name = "richie stevanus"
	// fmt.Println(name)

	// name = "richie new"
	// fmt.Println(name)

	// var name = "richie stevanus"
	// fmt.Println(name)

	// name = "richie new"
	// fmt.Println(name)

	name := "richie Stevanus"
	fmt.Println(name)

	name = "richie baru"
	fmt.Println(name)

	//jika ada nilai var yang tidak digunakan maka akan error
	var (
		firstName = "Richie"
		lastName  = "Stevanus"
	)

	fmt.Println(firstName + " " + lastName)
	fmt.Println(lastName)
}
