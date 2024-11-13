package main

import (
	"bufio"
	"fmt"
	"os"
)

// даны координаты точек(x, y) и их количество, необходимо вычислить максимальную площадь прямоугольника
//стороны параллельны осям координат
//находим среди всех возможных отрезков, те которые параллельны оси Х, затем сравниваем их длины и координаты

type segment struct { // структура для отрезков параллельных оси Х
	dist int
	x    int
	y1   int
	y2   int
}

func main() {
	var count int
	r := bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &count)
	x := make([]int, count)
	y := make([]int, count)
	for i := 0; i < count; i++ {
		fmt.Fscan(r, &x[i]) //считываем координаты в два слайса X и Y
		fmt.Fscan(r, &y[i])
	}

	var segments []*segment // срез структур отрезков
	for i := 0; i < count; i++ {
		for j := i + 1; j < count; j++ { // перебираем все возможные отрезки среди точек
			if x[i] == x[j] { // если х двух точек равный, то значит отрезок параллелен оси y
				dist := y[i] - y[j] // считаем расстояние между точками
				if dist < 0 {
					dist *= -1
				}
				s := new(segment)
				s.dist = dist
				s.x = x[i]
				s.y1 = y[i]
				s.y2 = y[j]
				segments = append(segments, s) //записываем параметры отрезка в структуру, а структуру в слайс
			}
		}

	}

	maxRectangle := 0
	for i := 0; i < len(segments); i++ {
		for j := i + 1; j < len(segments); j++ { // сравниваем отрезки между собой
			area := rectangleArea(segments[i], segments[j])
			if area == 0 {
				continue
			}
			if area > maxRectangle {
				maxRectangle = area //находим максимальную площадь среди возможных прямоугольников
			}

		}

	}

	fmt.Println(maxRectangle)
}

func rectangleArea(a, b *segment) int { //функция проверки отрезков являются ли они сторонами прямоугольника
	if a.dist != b.dist {
		return 0
	}
	area := 0
	if (a.y1 == b.y1 && a.y2 == b.y2) || (a.y1 == b.y2 && a.y2 == b.y1) {
		area = a.dist * (a.x - b.x) // и расчет площади
		if area < 0 {
			area *= -1
		}
	}
	return area
}
