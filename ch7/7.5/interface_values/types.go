// На заметку:
// Stdin, Stdout и Stderr — это открытые файлы, указывающие на дескрипторы файлов стандартного ввода,
// стандартного вывода и стандартных ошибок.
// Обратите внимание, что среда выполнения Go записывает стандартные ошибки в случае паники и сбоев;
// закрытие Stderr может привести к тому, что эти сообщения будут отправлены в другое место, возможно, в файл, открытый позже.

package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var w io.Writer
	fmt.Printf("%T\n", w) // <nil>

	w = os.Stdout
	fmt.Printf("%T\n", w) // *os.File

	w = new(bytes.Buffer)
	fmt.Printf("%T\n", w) // *bytes.Buffer

	w = nil
	fmt.Printf("%T\n", w) // <nil>
}
