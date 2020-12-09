package main

import "fmt"

func main() {
	pls := [] [] int {
		{5,8},
		{17},
		{6,3},
	}

	for _, deg1 := range pls {
		for _, deg2 := range deg1 {
			fmt.Printf("%d ", deg2)
		}
		fmt.Printf("\n")
	}
}
