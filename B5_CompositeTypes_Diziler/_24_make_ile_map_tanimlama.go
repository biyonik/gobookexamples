package main

import "fmt"

func main() {
	var makeMap = make(map [string] string)

	makeMap["Kalkülüs"] = "Matematik"
	makeMap["Algoritma"] = "Bilgisayar Mühendisliği"
	makeMap["Termodinamik"] = "Makine Mühendisliği"
	makeMap["Kalite Yönetimi"] = "Endüstri Mühendisliği"

	var makeMap2 = make(map [int]string)
	makeMap2[1] = "Bir"
	makeMap2[2] = "İki"
	makeMap2[3] = "Üç"
	makeMap2[4] = "Dört"

	for key, value := range makeMap2 {
		fmt.Println(key," => ", value)
	}
}
