package main

import "fmt"

func main() {
	var sozcukler = [3][2] string {
		{"Google", "Yandex"},
		{"IPhone", "Samsung"},
		{"Lenovo", "Asus"},
	}

	for _, i := range sozcukler {
		for _, j := range i {
			fmt.Print(" ", j)
		}
		fmt.Println()
	}
}
