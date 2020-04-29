package ch7

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"time"
)

// Ch7 第七章入口函数
func Ch7() {
	// sleep()
	newParam()
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
