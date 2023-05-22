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

func InsertionSort(arr []int) []int {
	PrintIntSlice(arr)
	n := len(arr)

	// 未ソート領域(i..n-1)
	// 一番初めの値は自動的にソート済みとなる
	for i := 1; i <= n-1; i++ {
		crr := arr[i]
		j := i - 1

		// ソート済み領域(0..i-1)
		// ソート済み領域の中から、現在の未ソート値より小さいものを探す
		for j >= 0 && crr < arr[j] {
			// より大きい値が見つかれば、その値を右にずらす
			arr[j+1] = arr[j]
			j--
		}

		// 一番左までいくか、より小さいものが見つかればその右に挿入
		// (一番左までいった場合はindexが-1 となるので、+1してindexが0になる)
		arr[j+1] = crr
		PrintIntSlice(arr)
	}

	return arr
}

func main() {
	sc.Split(bufio.ScanWords)

	N := NextInt()
	A := NextIntSlice(N)

	InsertionSort(A)
}
