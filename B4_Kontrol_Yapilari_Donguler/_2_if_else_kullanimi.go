package main

import "fmt"

func main() {
	s := "Go"
	if s == "Go" {
		fmt.Printf("If bloğu çalışıyor\n")
	} else {
		fmt.Printf("Else bloğu çalışıyor\n")
	}

	n := 9
	if n%2 == 0 { // {if false} deyimide kullanılabilir
		fmt.Printf("%d sayısı çifttir.\n", n)
 	} else {
 		fmt.Printf("%d sayısı tektir.\n", n)
	}
}
