package lenconv

import "fmt"

type Foot float64
type Meter float64

const (
	FOOT = 0.3048
)

func (c Meter) String() string {
	return fmt.Sprintf("%.2g m", c)
}

func (f Foot) String() string {
	return fmt.Sprintf("%.2g ft", f)
}

// футы в метры
func FToM(f Foot) Meter {
	return Meter(f * FOOT)
}

// метры в футы
func MToF(m Meter) Foot {
	return Foot(m / FOOT)
}
