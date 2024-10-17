package main

import (
	"fmt"
	"math"
)

// untuk menghitung jarak antara dua titik (a, b) dan (c, d)
func jarak(a, b, c, d float64) float64 {
	return math.Sqrt((a-c)*(a-c) + (b-d)*(b-d))
}

// untuk memeriksa apakah sebuah titik (x, y) berada di
// dalam lingkaran dengan pusat (cx, cy) dan jari-jari r
func diDalamLingkaran(x, y, cx, cy, r float64) bool {
	return jarak(x, y, cx, cy) <= r
}

func main() {
	var cx1, cy1, r1, cx2, cy2, r2, x, y float64

	//lingkaran 1 : pusat (cx1, cy1) dan jari-jari r1
	fmt.Print("Masukkan pusat (cx1, cy1) dan jari-jari (r1) lingkaran 1 : ")
	fmt.Scan(&cx1, &cy1, &r1)

	//lingkaran 2 : pusat (cx2, cy2) dan jari-jari r2
	fmt.Print("Masukkan pusat (cx2, cy2) dan jari-jari (r2) lingkaran 2 : ")
	fmt.Scan(&cx2, &cy2, &r2)

	//titik (x, y)
	fmt.Print("Masukkan titik (x, y) untuk diperiksa : ")
	fmt.Scan(&x, &y)

	//memeriksa apakah titik berada di dalam lingkaran 1
	diLingkaran1 := diDalamLingkaran(x, y, cx1, cy1, r1)

	//memeriksa apakah titik berada di dalam lingkaran 2
	diLingkaran2 := diDalamLingkaran(x, y, cx2, cy2, r2)

	//tentukan posisi titik
	if diLingkaran1 && diLingkaran2 {
		fmt.Println("Titik berada di dalam lingkaran 1 dan 2")
	} else if diLingkaran1 {
		fmt.Println("Titik berada di dalam lingkaran 1")
	} else if diLingkaran2 {
		fmt.Println("Titik berada di dalam lingkaran 2")
	} else {
		fmt.Println("Titik berada di luar lingkaran 1 dan 2")
	}
}
