// Добавьте в пакет tempconv типы, константы и функции
// для работы с температурой по шкале Кельвина,
// в которой нуль градусов соответствует температуре -273,15 °C
// а разница температур в 1K имеет ту же величину,
// что и 1 °С

package tempconv

import "fmt"

type Cel float64
type Far float64
type Kel float64

const (
	AbsoluteZeroC Cel = -273.15
	FreezingC     Cel = 0
	BoilingC      Cel = 100
	KelvinZeroC   Cel = -273.15
)

func (c Cel) String() string {
	return fmt.Sprintf("%g ℃", c)
}

func (f Far) String() string {
	return fmt.Sprintf("%g ℉", f)
}

func (k Kel) String() string {
	return fmt.Sprintf("%g °K", k)
}
