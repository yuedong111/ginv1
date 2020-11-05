package main

import (
	"fmt"
	"unsafe"
)

func main() {
	slice_test := []int{1, 2, 3, 4, 5}
	fmt.Println(unsafe.Sizeof(slice_test))
	fmt.Printf("main:%#v,%#v,%#v\n", slice_test, len(slice_test), cap(slice_test))
	slice_value(slice_test)
	fmt.Printf("main:%#v,%#v,%#v\n", slice_test, len(slice_test), cap(slice_test))
	slice_ptr(&slice_test)
	fmt.Printf("main:%#v,%#v,%#v\n", slice_test, len(slice_test), cap(slice_test))
	fmt.Println(unsafe.Sizeof(slice_test))
}

func slice_value(slice_test []int) {
	slice_test[1] = 100                // 函数外的slice确实有被修改
	slice_test = append(slice_test, 6) // 函数外的不变
	fmt.Printf("slice_value:%#v,%#v,%#v\n", slice_test, len(slice_test), cap(slice_test))
}

func slice_ptr(slice_test *[]int) { // 这样才能修改函数外的slice
	*slice_test = append(*slice_test, 7)
	fmt.Printf("slice_ptr:%#v,%#v,%#v\n", *slice_test, len(*slice_test), cap(*slice_test))
}
