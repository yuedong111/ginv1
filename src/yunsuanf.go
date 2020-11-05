package main

import "fmt"

func main(){
	var a int = 4
	var ptr int
	ptr = a
	fmt.Println(ptr)//4
	a = 15
	fmt.Println(ptr)//4

	var b int = 5
	var ptr1 *int
	ptr1 = &b
	fmt.Println(*ptr1)//5
	b=15
	fmt.Println(*ptr1)//15
}
