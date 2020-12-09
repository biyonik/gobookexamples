package main

import "fmt"

func main() {
	ogrenci_no := [3]int{101, 102, 103}
	ad_soyad := [3]string{"Ali Al", "Ahmet Demir", "Ayşe Tok"}
	not := [4]int{56, 95, 42, 100}

	for i := 0; i < len(ogrenci_no); i++ {
		fmt.Println(ogrenci_no[i], " numaralı ", ad_soyad[i],"'in notu := ", not[i]," dir.")
	}
}
