package main

import "fmt"

func main() {
	/**
	 * Sabit tanımlama -> Constant
	 * const <degisken_ismi> <degisken_tipi> = <degisken_degeri>
	 */

	const sayi int = 10
	const isim string = "Ahmet"
	const maas float32 = 4650.75
	const adres = "Balıkesir"

	fmt.Println("Sayı: ", sayi)
	fmt.Println("İsim: ", isim)
	fmt.Println("Maaş: ", maas)
	fmt.Println("Adres: ", adres)

	/**
	 * iota tanımlayıcısı artan sayıların tanımını basitleştirmek amacı ile const tanımlamalarında kullanılır
	 */
	const (
		a = iota
		b
		c
	)
	fmt.Println(a, b, c)
	const (
		d = iota + 3
		_
		e
		f
	)
	println(d, e, f)
}
