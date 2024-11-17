package main

import (
	"bufio"
	"fmt"
	"os"
)

// даны координаты точек(x, y) и их количество, необходимо вычислить максимальную площадь прямоугольника
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
	// x := make([]int, count)
	// y := make([]int, count)
	points := make([]point, 0, count)
	var segments []*segment // срез структур отрезков
	var maxRectangle int

	// TODO
	// points := make(map[int][]point) // map x -> point
	// segments := make(map[int][]segment) // map dist -> segment

	for i := 0; i < count; i++ {
		var p point
		fmt.Fscan(r, &p.x) //считываем координаты в два слайса X и Y
		fmt.Fscan(r, &p.y)

		for _, pointx := range points { // перебираем все возможные отрезки среди точек
			// условие, в котором раскрывается ключ к мапе
			if pointx.x == p.x { // если х двух точек равный, то значит отрезок параллелен оси y
				dist := pointx.y - p.y // считаем расстояние между точками
				if dist < 0 {
					dist *= -1
				}

				s := &segment{
					dist:   dist,
					point1: p,
					point2: pointx,
				}

				for _, seg := range segments { // сравниваем отрезки между собой
					maxRectangle = max(maxRectangle, rectangleArea(s, seg))
				}

				segments = append(segments, s) //записываем параметры отрезка в структуру, а структуру в слайс
			}
		}

		points = append(points, p)

		// TODO
		// points[p.x] = append(points[p.x], p)
	}

	fmt.Println(maxRectangle)
}

func rectangleArea(a, b *segment) int { //функция проверки отрезков являются ли они сторонами прямоугольника
	// условие, в котором раскрывается ключ к мапе
	if a.dist != b.dist {
		return 0
	}
	area := 0
	if a.point1.y == b.point1.y && a.point2.y == b.point2.y || a.point1.y == b.point2.y && a.point2.y == b.point1.y {
		area = a.dist * (a.point1.x - b.point1.x) // и расчет площади
		if area < 0 {
			area *= -1
		}
	}
	return area
}
