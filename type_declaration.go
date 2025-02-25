package main

import "fmt"

func main() {
	type kalimat string

	var ktpRichie kalimat = "11111111"

	var contoh string = "2222222"
	var contohKtp kalimat = kalimat(contoh)

	fmt.Println(ktpRichie)
	fmt.Println(contohKtp)
}
