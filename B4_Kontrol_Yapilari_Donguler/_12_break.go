package main

import "fmt"

func main() {
	// break, continue
	a := 18
	asalMi := true
	for i:=2; i < a; i++ {
		if a % i == 0 {
			asalMi = false
			break
		}
	}
	if asalMi {
		fmt.Printf("%d, sayısı asal sayıdır\n", a)
	} else {
		fmt.Printf("%d, sayısı asal sayı değildir\n", a)
	}
}
