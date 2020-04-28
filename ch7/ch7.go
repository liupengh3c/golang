package ch7

import (
	"bytes"
	"io"
)

// Ch7 第七章入口函数
func Ch7() {

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

var w io.Writer = new(bytes.Buffer)
