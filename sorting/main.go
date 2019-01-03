/*
* this package is to demonstrate sorting by custom field implementation
* In order for it to work, our custom type must implement
* 3 methods to satisfy sorting interface 1. Len; 2. Less; and 3. Swap
 */
package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

func main() {
	tracksByArtist := getTracks()
	sort.Sort(byArtist(tracksByArtist))
	printTracks(tracksByArtist)

	tracksByYear := getTracks()
	sort.Sort(byYear(tracksByYear))
	printTracks(tracksByYear)

	sortByYearUsingSlice := getTracks()
	sortUsingSlice(sortByYearUsingSlice)
	printTracks(sortByYearUsingSlice)
}

func sortByYear(list []*track) func(int, int) bool {
	return func(i, j int) bool {
		return list[i].Year > list[j].Year
	}
}

func sortUsingSlice(list []*track) {
	sort.Slice(list, sortByYear(list))
}

func getTracks() []*track {
	return []*track{
		{"Nasty", "The Prodigy", "The Day Is My Enemy", 2015, length("4m03s")},
		{"Diesel Pover", "The Prodigy", "The fat of the Land", 1997, length("4m18s")},
		{"Galvanize", "Chemical Brothers", "Push The Buttons", 2005, length("6m34s")},
		{"Attack", "System of a Down", "Hypnotize", 2005, length("3m06s")},
	}
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(err)
	}
	return d
}

func printTracks(tracks []*track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")

	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}

	tw.Flush()
}

type byArtist []*track

func (x byArtist) Len() int {
	return len(x)
}

func (x byArtist) Less(i, j int) bool {
	return x[i].Artist < x[j].Artist
}

func (x byArtist) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

type byYear []*track

func (x byYear) Len() int {
	return len(x)
}

func (x byYear) Less(i, j int) bool {
	return x[i].Year < x[j].Year
}

func (x byYear) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}
