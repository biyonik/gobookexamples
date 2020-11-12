package main

import "fmt"

func main() {
	/**
	 * Tür dönüşümü T(v) şeklindedir
	 * T -> dönüştürülmek istenen tür tipi
	 * v -> dönüştürülmek istenen değerdir
	 */
	var a int = 5
	var b int64 = int64(a)
	fmt.Println(int64(a) + b)

	x := 64.5
	y := 94
	var c float64 = float64(y)
	fmt.Println(x+c)

}
