// https://pkg.go.dev/fmt
//
// %s	не-интерпретированные байты строки или среза
// %q	строка в двойных кавычках, безопасно экранированная с помощью синтаксиса Go
// %x	основание 16, строчные буквы, два символа в байте
// %X	основание 16, верхний регистр, два символа в байте

package custom

import "fmt"

type Point struct {
	X, Y float64
}

func (p Point) Add(q Point) Point {
	return Point{p.X + q.X, p.Y + q.Y}
}
func (p Point) Sub(q Point) Point {
	return Point{p.X - q.X, p.Y - q.Y}
}

type Path []Point

func (path Path) TranslateBy(offset Point, add bool) {
	var op func(p, q Point) Point
	if add {
		op = Point.Add
	} else {
		op = Point.Sub
	}
	for i := range path {
		// Вызов либо path[i].Add(offset), либо path[i].Sub(offset)
		path[i] = op(path[i], offset)
		fmt.Println(path[i])
	}
}
