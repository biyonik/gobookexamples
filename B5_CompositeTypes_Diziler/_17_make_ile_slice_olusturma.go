package main

import "fmt"

func main() {
	/**
	 * make( [] slice_tipi, uzunluk, kapasite)
	 */

	s := make([]int, 3, 10)
	fmt.Println("Elemanlar = ", s, ", Uzunluk = ", len(s), ", Kapasite = ", cap(s))
}
