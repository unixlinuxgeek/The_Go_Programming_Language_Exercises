// Упражнение 5.1
// Измените программу ```findlinks``` так,
// чтобы она обходила связанный список ```n.FirstChild``` с помощью рекурсивных вызовов visit,
// а не с помощью цикла.

// Findlinks1 выводит ссылки в HTML-документе
// прочитанным со стандартного входа.

package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

// visit добавляет в links все ссылки,
// найденные в n, и возвращает результат.
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	// Измените программу ```findlinks``` так,
	// чтобы она обходила связанный список ```n.FirstChild``` с помощью рекурсивных вызовов visit,
	// а не с помощью цикла.

	//for c := n.FirstChild; c != nil; c = c.NextSibling {
	//	links = visit(links, c)
	//}

	// n.FirstChild
	//type Node struct {
	//	Parent, FirstChild, LastChild, PrevSibling, NextSibling *Node
	//	Type                                                    NodeType
	//	DataAtom                                                atom.Atom
	//	Data                                                    string
	//	Namespace                                               string
	//	Attr                                                    []Attribute
	//}

	return links
}
