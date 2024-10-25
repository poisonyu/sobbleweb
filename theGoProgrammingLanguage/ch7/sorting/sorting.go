// 给自定义的[]*Track实现sort.Interface接口，
// 从而用sort.Sort()指定Track的字段给[]*Track排序

package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots up", 2012, length("3m38s")},
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
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush()
}

// type Interface interface {
// 	Len() int
// 	Less(i, j int) bool
// 	Swap(i, j int)
// }

// type reverse struct {
// 	Interface
// }

// func (r reverse) Less(i, j int) bool {
// 	return r.Interface.Less(j, i)
// }

// func Reverse(data Interface) Interface {
// 	return reverse{data}
// }

// 根据Artist排序
type byArtist []*Track

func (x byArtist) Len() int {
	return len(x)
}
func (x byArtist) Less(i, j int) bool {
	return x[i].Artist < x[j].Artist
}

func (x byArtist) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

// 根据Year排序
type byYear []*Track

func (x byYear) Len() int {
	return len(x)
}
func (x byYear) Less(i, j int) bool {
	return x[i].Year < x[j].Year
}

func (x byYear) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

type customSort struct {
	t []*Track
	// 定义排序规则
	less func(x, y *Track) bool
}

// 实现sort.Interface
func (x customSort) Len() int {
	return len(x.t)
}

func (x customSort) Less(i, j int) bool {
	return x.less(x.t[i], x.t[j])
}

func (x customSort) Swap(i, j int) {
	x.t[i], x.t[j] = x.t[j], x.t[i]
}

type palindrome []int

func (p palindrome) Len() int {
	return len(p)
}

func (p palindrome) Less(i, j int) bool {
	return p[i] == p[j]
}

func (p palindrome) Swap(i, j int) {

}

func IsPalindrome(s sort.Interface) bool {
	// var a = s
	// sort.Sort(a)
	// sort.Sort(sort.Reverse(s))
	// return a.([]int) == s.([]int)
	// // if a == s {
	// // 	return true
	// // }

	i := 0
	j := s.Len() - 1
	for i < j {
		if !s.Less(i, j) {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	// printTracks(tracks)
	// fmt.Println()
	// sort.Sort(byArtist(tracks))
	// printTracks(tracks)
	// fmt.Println()
	// artist := sort.Reverse(byArtist(tracks))
	// sort.Sort(artist)
	// printTracks(tracks)

	// sort.Sort(byYear(tracks))               // 升序
	// sort.Sort(sort.Reverse(byYear(tracks))) // 降序
	// printTracks(tracks)

	// sort.Sort(customSort{tracks, func(x, y *Track) bool {
	// 	// 优先用Title去升序排列，如果Title相同，再根据Year排列，如果Year相同，再根据Length排序
	// 	if x.Title != y.Title {
	// 		return x.Title < y.Title
	// 	}
	// 	if x.Year != y.Year {
	// 		return x.Year < y.Year
	// 	}
	// 	if x.Length != y.Length {
	// 		return x.Length < y.Length
	// 	}
	// 	return false
	// }})
	// printTracks(tracks)

	// fmt.Println(sort.IsSorted(byArtist(tracks)))

	// a := []int{1, 2, 3}
	// b := []string{"sgas", "agsg", "aka", "fsg"}
	// // slices.Sort(b)
	// sort.Sort(sort.Reverse(sort.IntSlice(a)))
	// sort.Ints(a)
	// sort.IsSorted(sort.IntSlice(a))
	// sort.IntsAreSorted(a)
	// sort.Strings(b)
	// sort.Sort(sort.IntSlice(a))
	// sort.Sort(sort.Reverse(sort.IntSlice(a)))

	pal := []int{1, 2, 3, 4, 3, 2, 1}
	fmt.Println(IsPalindrome(palindrome(pal)))

}
