package main

import "fmt"

func main() {
	a := 1
	b := 0
	goto Etiket

	Etiket:
		b += a
		fmt.Printf("a = %d ve b = %d", a, b)
}
