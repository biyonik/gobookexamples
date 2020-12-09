package main

import "fmt"

func main() {
	/**
	 * copy metodu, 2 slice arasında kopyalama işlemi yapar
	 */

	kaynak_slice := [] string {"Go", "JavaScript", "C#", "Kotlin", "DartLang"}
	hedef_slice := make([]string, 2)

	kopyalanan_sayisi := copy(hedef_slice, kaynak_slice)
	fmt.Println("Kaynak Slice: ", kaynak_slice)
	fmt.Println("Hedef Slice: ", hedef_slice)
	fmt.Println("Kopyalanan Öğe Sayısı: ", kopyalanan_sayisi)
}
