package main

import "fmt"

// go:generate echo Merhaba Go Dünyası
func main() {
	/**
	 * Tip ve ilk değer atayarak değişken tanımlama
	 */
	var name string = "Ahmet"
	var surname string = "Altun"
	var age int8 = 30

	fmt.Println("İsim: ", name)
	fmt.Println("Soyisim: ", surname)
	fmt.Println("Yaş: ", age)
	/**
	 * Sadece ilk değer atayarak değişken tanımlama
	 */
	var sayi = 56
	var ucret = 805.4
	var adi_soyadi, adres = "Ali Al", "Isparta"

	fmt.Println("Sayı: ", sayi)
	fmt.Println("Ücret: ", ucret)
	fmt.Println("Ad, soyad: ", adi_soyadi)
	fmt.Println("Adres: ", adres)

	/**
	 * Sadece değişken adı ve ilk değer atayarak değişken tanımlama
	 */
	number := 123
	salary := 4650.95
	fmt.Println("Number: ", number)
	fmt.Println("Salary: ", salary)

	/**
	 * Blok halinde değişken tanımlama
	 */
	var (
		universite, bolum, sinif = "Balıkesir Üniversitesi", "Bilgisayar Mühendisliği", 2
	)
	fmt.Println("Üniversite: ", universite)
	fmt.Println("Bölüm: ", bolum)
	fmt.Println("Sınıf: ", sinif)
}
