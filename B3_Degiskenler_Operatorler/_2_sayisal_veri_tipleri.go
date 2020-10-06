package main

import (
	"fmt"
	"math"
)

func main() {
	/**
	 * Sayısal Veri Tipleri
	 */
	var min8Bit int8 = math.MinInt8
	var max8Bit int8 = math.MaxInt8
	var min16Bit int16 = math.MinInt16
	var max16Bit int16 = math.MaxInt16
	var min32Bit int32 = math.MinInt32
	var max32Bit int32 = math.MaxInt32

	fmt.Printf("Min 8 Bit: \t%d\nMax 8 Bit: \t%d\nMin 16 Bit: \t%d\nMax 16 Bit: \t%d\nMin 32 Bit: \t%d\nMax 32 Bit: \t%d\n",
		min8Bit,
		max8Bit,
		min16Bit,
		max16Bit,
		min32Bit,
		max32Bit)

	/**
	 * Rune veri tipi
	 */
	var ch1 rune = '\f'
	var ch2 rune = '€'
	fmt.Println("Ch1: ", ch1)
	fmt.Println("Ch2: ", ch2)

	var byteVariable byte = 65
	fmt.Println("Byte Variable: ", byteVariable)

	/**
	 * Complex (Kompleks) Veri Tipi
	 */
	var kompleks1 complex64 = 5 + 7i
	var kompleks2 complex128 = 12 + 122i
	fmt.Println("Kompleks1: ", kompleks1)
	fmt.Println("Kompleks2: ", kompleks2)

	/**
	 * Boolean veri tipi
	 */
	var isMarried bool = false
	fmt.Println("Is Married? : ", isMarried)
}
