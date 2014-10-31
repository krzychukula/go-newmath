//Package newmath is a trivial example package.
package newmath

import "fmt"

func abs(x float64) float64 {
	switch {
	case x < 0:
		return -x
	case x == 0:
		return 0 // return correctly abs(-0)
	}
	return x
}

//Sqrt return an approximation to the square root of x.
func Sqrt(x float64) float64 {
	diff := 0.000000000000000001
	z := 1.0
	var newZ float64
	var count = 0
	for ; ; z = newZ {
		count += 1
		newZ = z - (z*z-x)/(2*z)
		if abs(z-newZ) < diff {
			break
		}

	}
	fmt.Println(count)
	return z
}
