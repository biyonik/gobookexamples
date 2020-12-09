package main

import "fmt"

func main() {
	var kelime = [3] string {"Go", "Java", "Python"}
	var sayilar = [4] float64 {12.5, 10.0, 58.9, 56.7}
	var d_y = [5] bool {true, false, true, true, false}

	for indis := range kelime {
		fmt.Print("İndis =", indis, " ")
	}
	fmt.Println()
	for index, value := range sayilar {
		fmt.Print("İndis = ", index," Değer = ", value, " - ")
	}
	fmt.Println()

	for _, val := range d_y {
		fmt.Print("Değer = ", val, " ")
	}
}
