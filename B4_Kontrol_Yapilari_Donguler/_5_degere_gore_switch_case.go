package main

import "fmt"

func main() {
	a := 5
	switch a {
		case 4, 5:
			fmt.Printf("a sayısı 5 veya 4 tür")
		case 2:
			fmt.Printf("a sayısı 2'dir")
		case 3:
			fmt.Printf("a sayısı 3'tür")
	}
}
