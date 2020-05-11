package ch7

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

// Ch7 第七章入口函数
func Ch7() {
	// sleep()
	// newParam()
	chSort()
	ch7Interface()
}

// 空接口，空接口可以接收任何类型
type empty interface {
}
type reader interface {
	Read(p []byte) (n int, err error)
}
type writer interface {
	Write(p []byte) (n int, err error)
}
type closer interface {
	Closer() error
}
type readerWriter interface {
	reader
	writer
}

func sleep() {
	period := flag.Duration("period", 1*time.Second, "sleep period")
	flag.Parse()
	fmt.Printf("sleeping for %v...", *period)
	time.Sleep(*period)
	fmt.Println()
}

type celsius float64

func (c celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

type clesiusFlag struct {
	celsius
}

func (f *clesiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C", "°C":
		f.celsius = celsius(value)
		fmt.Println("intput is 摄氏度")
		return nil
	case "F", "°F":
		f.celsius = celsius(value)
		fmt.Println("intput is 华氏度")
		return nil
	}

	return fmt.Errorf("invilid input %s", s)
}

func newParam() {
	var s = new(clesiusFlag)
	s.celsius = 20.0
	flag.CommandLine.Var(s, "temp", "the temprature")
	flag.Parse()
	fmt.Println(s.celsius)
	var a bytes.Buffer
	fmt.Printf("%T\n", a)
	var b io.Writer
	fmt.Printf("%T\n", b)
}

// 7.6使用sort.interface来排序

// Track 播放列表
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
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, val := range tracks {
		fmt.Fprintf(tw, format, val.Title, val.Artist, val.Album, val.Year, val.Length)
	}
	tw.Flush()
}

// 按艺术家排序
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

// 多维度排序

type customSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (x customSort) Len() int {
	return len(x.t)
}

func (x customSort) Less(i, j int) bool {
	return x.less(x.t[i], x.t[j])
}

func (x customSort) Swap(i, j int) {
	x.t[i], x.t[j] = x.t[j], x.t[i]
}

func chSort() {
	sort.Sort(byArtist(tracks))
	printTracks(tracks)

	sort.Sort(customSort{tracks, func(x, y *Track) bool {
		if x.Title != y.Title {
			return x.Title < y.Title
		}
		if x.Year != y.Year {
			return x.Year < y.Year
		}
		if x.Length != y.Length {
			return x.Length < y.Length
		}
		return false
	}})
	printTracks(tracks)
}

type database map[string]float32

var db = database{"shoes": 50.0, "socks": 4.5}

func (db database) list(w http.ResponseWriter, r *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s:%v\n", item, price)
	}
}
func (db database) price(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	if price, ok := db[item]; ok {
		fmt.Fprintf(w, "%v\n", price)
		return
	}
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "no such item:%s", item)
	return
}
func chHander() {
	mux := http.NewServeMux()
	// http.HandlerFunc 为函数类型，其实现了http.Handler接口，serveHttp函数是调用自己
	// mux.Handle("/list", http.HandlerFunc(db.list))
	// mux.Handle("/price", http.HandlerFunc(db.price))

	// 简化版
	mux.HandleFunc("/list", db.list)
	mux.HandleFunc("price", db.price)
	http.ListenAndServe(":8000", mux)
}

type human interface {
	SayHi()
}

type student struct {
	name  string
	age   int32
	money int32
}

type employee struct {
	name string
	age  int32
	work string
}

func info(person human) {
	person.SayHi()
}

func (s student) SayHi() {
	fmt.Println("i am a student,my name is:", s.name)
}

func (e employee) SayHi() {
	fmt.Println("i am a employee,my name is:", e.name)
}

func ch7Interface() {
	stu := student{
		name:  "tom",
		age:   10,
		money: 10,
	}
	emp := employee{
		name: "lucy",
		age:  18,
		work: "engneer",
	}
	info(stu)
	info(emp)
}
