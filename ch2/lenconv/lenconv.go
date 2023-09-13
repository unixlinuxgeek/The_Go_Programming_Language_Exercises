package lenconv

type Foot float64
type Meter float64

const (
	FOOT = 0.3048
)

// длину в футы и метры
// 1 foot = 0.3048  meter
// 2 foot = 2*0.3048
func FToM(f float64) Meter {
	return Meter(f * FOOT)
}

// длину в футы и метры
// 1 meter = 1 / 0.3048 = 3.2 foots
// 2 meter = 2 / 0.3048 = 6.5 foots
func MToF(m float64) Foot {
	return Foot(m / FOOT)
}
