package wtconv

import "fmt"

type Lb float64
type Kg float64

const (
	LB = 0.453592
)

func (lb Lb) String() string {
	return fmt.Sprintf("%.2g lb", lb)
}
func (kg Kg) String() string {
	return fmt.Sprintf("%.2g kg", kg)
}

// Lb to Kg
func LbToKg(p Lb) Kg {
	return Kg(p * LB)
}

// Kg to Lb
func KgToLb(k Kg) Lb {
	return Lb(k / LB)
}
