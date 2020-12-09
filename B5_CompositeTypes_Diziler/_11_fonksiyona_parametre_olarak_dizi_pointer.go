package main

import "fmt"

func main() {
	var floatArr = [...]float64{10.2, 54.6, 96.5, 110.23}
	var x = Topla(&floatArr)
	fmt.Printf("Dizi elemanlarının toplamı : %.2f", x)
}

func Topla(paramArr *[4]float64) (result float64) {
	for _, v := range paramArr {
		result += v
	}
	return
}
