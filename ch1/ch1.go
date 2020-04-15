package ch1

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"io/ioutil"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

// GetStdCmd 第一章示例代码，第1个
// 输出命令行参数
func GetStdCmd() {
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	return
}

// GetStdCmd2 第一章示例代码，第2个
// 输出命令行参数，短变量、for range使用
func GetStdCmd2() {
	s, sep := "", ""
	for _, val := range os.Args[1:] {
		s += sep + val
		sep = " "
	}
	fmt.Println(s)
	return
}

// Dup 第一章示例代码，第3个
// 统计输入中相同的行，注意scan要设置停止条件
func Dup() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		// 输入end代表结束
		if input.Text() == "end" {
			break
		}
		counts[input.Text()]++
	}
	for key, val := range counts {
		if val > 1 {
			fmt.Println(key + fmt.Sprintf("------%d", val))
		}
	}
	return
}

// Dup2 第一章示例代码，第4个
// 统计输入/文本中相同的行
func Dup2() {
	counts := make(map[string]int)
	if len(os.Args[1:]) == 0 {
		// 从标准输入读取
		countsLine(os.Stdin, counts)
	} else {
		// 从文件中读取
		for _, name := range os.Args[1:] {
			f, err := os.Open(name)
			if err != nil {
				fmt.Println(name + " open error")
				continue
			}
			countsLine(f, counts)
			f.Close()
		}
	}
	for key, val := range counts {
		if val > 1 {
			fmt.Println(key + fmt.Sprintf("------%d", val))
		}
	}
	return
}

// countsLine 基础统计函数
func countsLine(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if input.Text() == "end" {
			break
		}
		counts[input.Text()]++
	}
}

// Dup3 第一章示例代码，第5个
// 一次性文本中读取内容
func Dup3() {
	counts := make(map[string]int)
	files := os.Args[1:]
	for _, name := range files {
		data, err := ioutil.ReadFile(name)
		if err != nil {
			fmt.Println(name + " read fail")
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for key, val := range counts {
		if val > 1 {
			fmt.Println(key + fmt.Sprintf("------%d", val))
		}
	}
	return
}

// Lissajous利萨茹图形
func Lissajous() {
	rand.Seed(time.Now().UTC().UnixNano())
	handler := func(w http.ResponseWriter, r *http.Request) {
		lissajous(w)
	}
	http.HandleFunc("/", handler)
	http.ListenAndServe("localhost:8000", nil)
	return
}

func lissajous(out io.Writer) {
	const (
		cycles     = 5
		res        = 0.001
		size       = 100
		nframes    = 64
		delay      = 8
		whiteIndex = 0
		blackIndex = 1
	)
	palette := []color.Color{color.White, color.Black}
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
	return
}
