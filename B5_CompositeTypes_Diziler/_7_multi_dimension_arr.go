package main

import "fmt"

func main() {
	var coklu_dizi [2][3] string
	coklu_dizi[0][0] = "Öğrenci - 1"
	coklu_dizi[0][1] = "95"
	coklu_dizi[0][2] = "67"

	coklu_dizi[1][0] = "Öğrenci - 2"
	coklu_dizi[1][1] = "74"
	coklu_dizi[1][2] = "56"

	for i:=0; i < 2; i++ {
		fmt.Println()
		for j := 0; j < 3; j++{
			fmt.Print(" ", coklu_dizi[i][j]," ")
		}
	}
}
