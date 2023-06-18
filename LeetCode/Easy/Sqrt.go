package main

import (
	"fmt"
	"math"
)

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}

	left := 1
	right := x

	for left <= right {
		mid := (left + right) / 2

		if mid*mid == x {
			return mid
		} else if mid*mid > x {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return right
}

func main() {
	res := mySqrt(0)
	fmt.Printf("0: %v\n", res)
	res = mySqrt(1)
	fmt.Printf("1: %v\n", res)
	res = mySqrt(2)
	fmt.Printf("2: %v\n", res)
	res = mySqrt(3)
	fmt.Printf("3: %v\n", res)
	res = mySqrt(4)
	fmt.Printf("4: %v\n", res)
	res = mySqrt(5)
	fmt.Printf("5: %v\n", res)
	res = mySqrt(8)
	fmt.Printf("8: %v\n", res)
	res = mySqrt(9)
	fmt.Printf("9: %v\n", res)
	res = mySqrt(10)
	fmt.Printf("10: %v\n", res)
	res = mySqrt(15)
	fmt.Printf("15: %v\n", res)
	res = mySqrt(16)
	fmt.Printf("16: %v\n", res)
	res = mySqrt(17)
	fmt.Printf("17: %v\n", res)

	x := math.Pow(2.0, 8.0) - 1
	fmt.Println(x)
	res = mySqrt(int(x))
	fmt.Printf("x: %v\n", res)
}
