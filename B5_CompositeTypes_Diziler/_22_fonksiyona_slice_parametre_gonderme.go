package main

import "fmt"

func main() {
	var sayiDizisi = [3]float64 {10.2, 54.6, 96.5}
	var toplam float64 = Topla(sayiDizisi[:])
	fmt.Println("Dizinin elemanları toplamı: ", toplam)
}

func Topla(a []float64) float64  {
	s:= 0.0
	for i := 0; i<len(a); i++ {
		s += a[i]
	}
	return s
}
