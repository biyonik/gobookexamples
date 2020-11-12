package main

import "fmt"

func main() {
	switch {
	case 3 < 6:
		fmt.Printf("3 sayısı 6'dan küçüktür")
	case (2 == 5):
		fmt.Printf("2 sayısı 5'e eşittir")
	case (5%2 == 0):
		fmt.Printf("5 sayısıs çifttir")
	default:
		fmt.Printf("Hiç bir işlem yapılmadı")
	}
}
