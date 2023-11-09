package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

// Сортируем по артистам
type byArtist []*Track

func (x byArtist) Len() int           { return len(x) }
func (x byArtist) Less(i, j int) bool { return x[i].Artist < x[j].Artist }
func (x byArtist) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type byYear []*Track

func (x byYear) Len() int           { return len(x) }
func (x byYear) Less(i, j int) bool { return x[i].Year < x[j].Year }
func (x byYear) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 0, 6, ' ', 0)
	fmt.Fprintf(os.Stdout, format, "Title        ", "Artist             ", "Album            ", "Year       ", "Length    ")
	fmt.Fprintf(os.Stdout, format, "-------------", "-------------------", "-----------------", "-----------", "----------")

	//sort.Sort(byYear(tracks))
	//sort.Sort(sort.Reverse(byYear(tracks)))

	//sort.Sort(byArtist(tracks))
	sort.Sort(sort.Reverse(byArtist(tracks)))

	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // Вычисление размеров столбцов и вывод таблицы
}

func main() {
	printTracks(tracks)
}
