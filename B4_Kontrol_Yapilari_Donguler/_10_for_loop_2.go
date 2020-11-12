package main

import "fmt"

func main() {
	i := 4
	faktoriyel := 1
	for i > 0 {
		faktoriyel *= i
		i--
	}
	fmt.Printf("%d\n", faktoriyel)
	fmt.Printf("----------------------------")

	sayi := 0
	for {
		if sayi > 20 {
			break
		}
		fmt.Printf("%d\n", sayi)
		sayi += 5
	}
}
