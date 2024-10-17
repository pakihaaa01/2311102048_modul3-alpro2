package main

import "fmt"

func main() {
	var n, r int
	fmt.Scan(&n, &r)
	fmt.Println(permutaion(n, r))
}

func mencariFaktorial(n int, hasil *int) {
	var i int
	*hasil = 1
	for i = n; i >= 1; i-- {
		*hasil = *hasil * i
	}
}

func permutaion(n, r int) int {
	var pembilang, penyebut int
	mencariFaktorial(n, &pembilang)
	mencariFaktorial(n-r, &penyebut)
	return pembilang / penyebut
}
