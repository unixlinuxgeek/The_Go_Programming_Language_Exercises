### Методы с указателем в роли получателя.

```go
// Где p это указатель
func (p *Point) ScaleBy(factor float64) {
  p.X *= factor
  p.Y *= factor
}
```
