// debug is false
//  $ go run ./example_without_error.go
//  panic: runtime error: invalid memory address or nil pointer dereference
//  [signal SIGSEGV: segmentation violation code=0x1 addr=0x20 pc=0x462e1d]
//
//  goroutine 1 [running]:
//  bytes.(*Buffer).Write(0x404b99?, {0xc00007cf5c?, 0x0?, 0x0?})
//        /usr/local/go/src/bytes/buffer.go:168 +0x1d
//  main.f(...)
//        /home/geek/GoApps/The_Go_Programming_Language_Exercises/ch7/7.5.1/example_without_error.go:26
//  main.main()
//        /home/geek/GoApps/The_Go_Programming_Language_Exercises/ch7/7.5.1/example_without_error.go:16 +0x55
//  exit status 2
//

package main

import (
	"bytes"
	"fmt"
	"io"
)

const debug = false

func main() {
	var buf *bytes.Buffer
	if debug {
		buf = new(bytes.Buffer) //  Накопление вывода
	}
	f(buf)
	if debug {
		// используем buf...
		fmt.Println("Debug turn on...")
	}
}

func f(out io.Writer) {
	// ... некоторые действия ...
	if out != nil {
		out.Write([]byte("выполнено:\n"))
	}
}
