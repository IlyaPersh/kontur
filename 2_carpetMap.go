package main

import (
	"bufio"
	"fmt"
	"os"
)

// Даны координаты точек(x, y) и их количество, необходимо вычислить максимальную площадь прямоугольника
// стороны параллельны осям координат
// находим среди всех возможных отрезков, те которые параллельны оси Х, затем сравниваем их длины и координаты
type point struct {
	x int
	y int
}

type segment struct { // структура для отрезков параллельных оси Х
	point1 point
	point2 point
	dist   int
}

func main() {
	var count int
	r := bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &count)
	var maxRectangle int
	points := make(map[int][]point)     // мапа для точек
	segments := make(map[int][]segment) // мапа для отрезков

	for i := 0; i < count; i++ {
		var p point
		fmt.Fscan(r, &p.x) //считываем координаты в структуру
		fmt.Fscan(r, &p.y)
		if _, ok := points[p.x]; ok { // если х двух точек равный, то значит отрезок параллелен оси y
			for _, xMap := range points[p.x] { // смотрим в мапе точки с одинаковым х
				dist := xMap.y - p.y // считаем расстояние между точками
				if dist < 0 {
					dist *= -1
				}
				s := &segment{ // отрезок
					dist:   dist,
					point1: p,
					point2: xMap,
				}
				if _, ok := segments[dist]; ok { //находим отрезки с одинаковой длиной

					for _, segMap := range segments[dist] { // сравниваем отрезки между собой
						maxRectangle = max(maxRectangle, rectangleArea(s, &segMap))
					}
				}
				segments[dist] = append(segments[dist], *s) //отрезок в мапу
			}
		}
		points[p.x] = append(points[p.x], p)
	}

	fmt.Println(maxRectangle)
}

func rectangleArea(a, b *segment) int { //функция проверки отрезков являются ли они сторонами прямоугольника
	var area int
	if a.point1.y == b.point1.y && a.point2.y == b.point2.y || a.point1.y == b.point2.y && a.point2.y == b.point1.y {
		area = a.dist * (a.point1.x - b.point1.x) // и расчет площади
		if area < 0 {
			area *= -1
		}
	}
	return area
}
