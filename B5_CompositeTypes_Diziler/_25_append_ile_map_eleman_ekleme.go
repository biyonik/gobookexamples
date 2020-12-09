package main

import "fmt"

func main() {
	mp1 := make(map[string] []string)

	mp1["key"] = append(mp1["key"], "Value-1", "Value-2", "Value-3")
	fmt.Println(mp1["key"])
}
