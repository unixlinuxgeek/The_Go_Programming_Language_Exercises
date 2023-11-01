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
	fmt.Println(w) // <nil>

	w = os.Stdout
	fmt.Println(w) // &{0xc00002a120}

	w = new(bytes.Buffer)
	fmt.Println(w) // empty output
	w = nil
	fmt.Println(w) // nil
}
