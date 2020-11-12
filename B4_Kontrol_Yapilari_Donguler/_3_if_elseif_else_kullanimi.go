package main

import "fmt"

func main() {
	var not int32 = 89
	if not > 90 && not <= 100 {
		fmt.Printf("Notunuz: AA")
	} else if not > 80 && not <= 90 {
		fmt.Printf("Notunuz: BA")
	} else if not > 70 && not <= 80 {
		fmt.Printf("Notunuz: BB")
	} else if not > 60 && not <= 70 {
		fmt.Printf("Notunuz: CB")
	} else {
		fmt.Printf("Notunuz: FF")
	}
}
