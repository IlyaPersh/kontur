// Найти в массиве последовательности элементов, сумма которых меньше К. Один элемент тоже последовательность
// перебираем массив двумя for, считаем сумму элементов
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, k int
	r := bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &n)
	fmt.Fscan(r, &k)
	nums := make([]int, n) //считываем данные
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &nums[i])
	}

	count := 0 // счетчик для количества интересных подотрезков
	for i := 0; i < n; i++ {
		if nums[i] > k { //если текущий элемент больше к, то переходим к следующему
			continue
		}
		count++
		sum := nums[i]               //начинаем считать сумму
		for j := i + 1; j < n; j++ { //второй цикл для подсчета суммы с последующими элементами
			if sum+nums[j] > k {
				break // если сумма больше к, то прерываем подсчет суммы, выходим из подцикла, начинаем считать сумму с нового элемента
			}
			count++
			sum += nums[j]
		}
	}

	fmt.Println(count)
}
