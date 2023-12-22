// Упражнение 7.17
// Расширьте возможности ```xmlselect``` так, чтобы элементы могли быть выбраны не только
// по имени, но и по аттрибутам, в духе CSS, так что, например, элемент наподобие
// "<div id="page" class="wide">" может быть выбран как по соответствию аттрибутов id
// или class, так и по его имени.

package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	dec := xml.NewDecoder(os.Stdin)
	var stack []string // Стек имен элементов
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			for _, n := range tok.Attr {
				//for _, el := range args {
				//	if el == n.Name.Local {
				//		stack = append(stack, tok.Name.Local) // Внесение в стек
				//		fmt.Print("Element: ", n.Name.Local, " | Value:", n.Value)
				//		fmt.Print("\n")
				//	}
				//}
				for _, el := range args {
					if el == n.Value {
						stack = append(stack, tok.Name.Local) // Внесение в стек
						fmt.Print("Element: ", n.Name.Local, " | Value:", n.Value)
						fmt.Print("\n")
					}
				}
			}
			//stack = append(stack, tok.Name.Local) // Внесение в стек
		case xml.EndElement:
			if (len(stack) - 1) >= 0 {
				stack = stack[:len(stack)-1] // Снятие со стека
			}
		case xml.CharData:
			if containsAll2(stack, os.Args[1:]) {
				fmt.Printf("%s: %s\n", strings.Join(stack, " "), tok)
			}
		}
	}
}

func containsAll2(x, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if x[0] == y[0] {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}
