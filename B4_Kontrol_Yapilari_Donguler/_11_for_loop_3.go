package main

import "fmt"

func main() {
	// Range yapısı, diziler üzerinde yapılacak iteratif işlemler için kullanılır
	kelime := "Go Programlama Dili"
	for indis, character := range kelime {
		fmt.Printf("%d indisindeki karakter := %c\n", indis, character)
	}
}
