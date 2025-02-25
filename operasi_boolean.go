package main

import "fmt"

func main() {
	var nilaiAkhir = 90
	var nilaiAbsen = 80

	var lulusNilaiAkhir bool = nilaiAkhir > 80
	var lulusnilaiAbsen bool = nilaiAbsen >= 80

	var lulus bool = lulusNilaiAkhir && lulusnilaiAbsen

	if lulus {
		fmt.Println("Selamat Anda Lulus")
	} else {
		fmt.Println("Maaf Anda Tidak Lulus")
	}
}
