package main

import "fmt"

func main() {
	// var nilai8 int8 = 127
	// var nilai16 int16 = 32767
	// var nilai32 int32 = 2147483647
	// var nilai64 int64 = 9223372036854775807

	var nilai32 int32 = 32767
	var nilai16 int16 = int16(nilai32)
	var nilai64 int64 = int64(nilai32)

	fmt.Println(nilai16)
	fmt.Println(nilai32)
	fmt.Println(nilai64)

	// var name = "richie stevanus"
	// var r = name[0]
	// var rString = string(r)

	// fmt.Println(name)
	// fmt.Println(r)
	// fmt.Println(rString)

	var name = "richie stevanus"
	//	var r = name[0]
	var rString = string(name[0])

	fmt.Println(name)
	//	fmt.Println(r)
	fmt.Println(rString)
}
