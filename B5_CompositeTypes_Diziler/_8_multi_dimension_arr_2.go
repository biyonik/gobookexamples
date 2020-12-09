package main

import "fmt"

func main() {
	sayilar := [4][3] int {
		{1,5,6},
		{3,5,7},
		{6,8,10},
		{1,3,5},
	}

	for i := 0; i < len(sayilar); i++ {
		for j:=0; j < len(sayilar[0]); j++ {
			fmt.Printf("sayilar[%d][%d] = %d\n", i, j, sayilar[i][j])
		}
		fmt.Println()
	}
}
