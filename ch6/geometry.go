package geometry

import "math"

type Point struct {
	X, Y float64
}

// Distance То же, но как метод типа Point
// p это получатель метода приемника (отправка сообщения объекту)
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}
