package main

import (
	"bufio"
	"fmt"
	"os"
)

// Дан массив целых чисел длиной N, найти пару чисел с наибольшей разностью, вывести их индекс+1
// Если таких пар несколько, то вывести индексы с максимальной разностью

func main() {
	var n int
	r := bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &n)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &nums[i])
	}

	maxNum := nums[0]
	minNum := nums[0]
	maxIdx, minIdx := 0, 0

	for i, a := range nums {
		if a >= maxNum {
			maxNum = a
			maxIdx = i
		}
		if a < minNum {
			minNum = a
			minIdx = i
		}
	}
	fmt.Println(maxIdx+1, minIdx+1)
}
