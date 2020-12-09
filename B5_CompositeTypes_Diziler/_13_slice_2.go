package main

import "fmt"

func main() {
	s := make([]int, 0, 3)
	for i := 0; i < 20; i++ {
		s = append(s, i)
		fmt.Printf("Kapasite: %d, Uzunluk: %d\n", cap(s), len(s))
	}
}
