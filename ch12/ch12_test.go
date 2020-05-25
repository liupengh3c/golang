package ch12

import "reflect"

func Example_slice() {
	display("slice", reflect.ValueOf([]*int{new(int), nil}))
}
