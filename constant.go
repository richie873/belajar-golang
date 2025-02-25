package main

import "fmt"

func main() {
	// const (
	// 	firstName = "richie"
	// 	lastName  = "stevanus"
	// )

	// jika ada nilai const yang tidak digunakan tidak akan error
	const firstName = "richie"
	const lastName = "stevanus"

	// namun nilainya tidak bisa diubah
	// firstName = "Richie New"
	// lastName = "Stevanus New"

	// fmt.Println(firstName + " " + lastName)
	fmt.Println(firstName)
}
