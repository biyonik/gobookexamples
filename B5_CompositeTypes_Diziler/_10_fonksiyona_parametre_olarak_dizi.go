package main

import "fmt"

func main() {
	dizi1 := [...]float64{10.2, 54.6, 96.5}
	x:= Topla(dizi1)
	fmt.Printf("Dizi elemanlarının toplamı = %.2f", x)
}

func Topla(paramArr [3] float64) (result float64) {
	for _, v := range paramArr {
		result += v
	}
	return
}
