package main

import "fmt"

func main() {
	/**
	 * İlk değer atama ile dizi tanımlama
	 * var <dizi_adi> = [dizi_boyutu] <dizi_tipi> { dizi_elemanlari }
	 */

	var firstArr = [5]string{"Java", "Go", "C", "C#", "Python"}
		secondArr := [3]float64{15.6, 20.5, 36.8}
	var thirthArr = [...] int {6, 8, 5, 9}
	var fourthArr = [...] bool {3: true, 4: false, 5: true}

	fmt.Println(firstArr)
	fmt.Println(secondArr)
	fmt.Println(thirthArr)
	fmt.Println(fourthArr)
	fmt.Printf("fourthArr dizisinin eleman sayısı := %d", len(fourthArr))
}
