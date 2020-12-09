package main

import "fmt"

func main() {
	/**
	 * append(slice_adi, eklenecek_eleman/elemanlar)
	 */

	s := make([] int, 0 , 3)
	a := make([] int, 2, 5)

	for i := 0; i < 5; i++ {
		s = append(s, i)
		fmt.Printf("Kapasite: %d, Uzunluk: %d, Pointer: %p\n", cap(s), len(s), s)
	}
	fmt.Println("-----------------------------------------")
	for i := 0; i < 5; i++ {
		a = append(a, i)
		fmt.Printf("Kapasite: %d, Uzunluk: %d, Pointer: %p\n", cap(a), len(a), a)
	}
}
