package main

import "fmt"

func main() {
	var slice = []int{1, 3, 5, 7, 9, 11, 0, 2, 4, 6, 8}
	var tekSayilar [] int = slice[0:6]
	var ciftSayilar []int = slice[6:]

	fmt.Println("Ana Slice: ", slice)
	fmt.Println("Tek Sayılar: ", tekSayilar)
	fmt.Println("Çift Sayılar: ", ciftSayilar)
}
