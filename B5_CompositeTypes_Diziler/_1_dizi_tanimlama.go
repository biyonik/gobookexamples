package main

import "fmt"

func main() {
	/**
	 * Dizi; aynı veri tipine ait öğelerin belli uzunlukta ve numaradaki koleksiyonu olup,
	 * bellekte sıralı olarak saklanır
	 */
	// var <dizi_adi> [dizi_boyutu] <dizi_tipi>

	var intArr [3]int
	intArr[0] = 5
	intArr[1] = 3
	intArr[2] = 2

	fmt.Print(intArr)

	var boolArr [5]bool
	fmt.Print(boolArr)
	// [false false false false false] -> bool veri tipinin varsayılan değeri 'false' olduğu için dizi elemanları da 'false' olacaktır
}
