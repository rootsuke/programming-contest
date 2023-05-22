package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func NextString() string {
	sc.Scan()
	return sc.Text()
}

func NextInt() int {
	n, _ := strconv.Atoi(NextString())
	return n
}

func NextIntSlice(size int) []int {
	res := make([]int, size)
	for i := 0; i < size; i++ {
		res[i] = NextInt()
	}
	return res
}

func PrintIntSlice(slice []int) {
	fmt.Println(strings.Trim(fmt.Sprint(slice), "[]"))
}

// return sorted slice and swapped count.
func SelectionSort(slice []int) (sorted_slice []int, swapped_cnt int) {
	n := len(slice)
	cnt := 0

	for i := 0; i < n; i++ {
		min_value_index := i
		swapped := false

		for j := i + 1; j < n; j++ {
			if slice[j] < slice[min_value_index] {
				min_value_index = j
				swapped = true
			}
		}

		if swapped {
			slice[i], slice[min_value_index] = slice[min_value_index], slice[i]
			cnt++
		}
	}

	return slice, cnt
}

func main() {
	sc.Split(bufio.ScanWords)

	N := NextInt()
	A := NextIntSlice(N)

	A, cnt := SelectionSort(A)

	PrintIntSlice(A)
	fmt.Println(cnt)
}
