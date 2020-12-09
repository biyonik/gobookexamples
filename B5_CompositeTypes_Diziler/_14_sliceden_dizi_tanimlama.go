package main

import "fmt"

func main() {
	var strArr = [5] string {"Go", "Java", "Kotlin", "C#", "JavaScript"}
	var slice1 []string = strArr[2:4]
	var slice2 []string = strArr[1:3]

	fmt.Println("Dizi = ", strArr)
	fmt.Println("Slice 1 = ", slice1)
	fmt.Println("Slice 2 = ", slice2)
}
