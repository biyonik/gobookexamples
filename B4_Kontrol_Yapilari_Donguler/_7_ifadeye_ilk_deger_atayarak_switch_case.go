package main

import "fmt"

func main() {
	switch a := 5; {
		case a < 0:
			fmt.Printf("%d sayısı negatiftir", a)
		case a == 0:
			fmt.Printf("%d sayısı 0 a eşittir", a)
		case a > 0:
			fmt.Printf("%d sayısı 0'dan büyüktür", a)
	}
}
