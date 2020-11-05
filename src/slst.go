package main

import (
	"bytes"
	"fmt"
	"sync"
)

func main() {
	//	for i := 0; i < 5; i++ {
	//		fmt.Printf("i= %d\n", i)
	//		if i == 3 {
	//			goto label
	//
	//
	//		}
	//	}
	//label:
	//	fmt.Println("the i is 3")
	printData := func(wg *sync.WaitGroup, data []byte) {
		defer wg.Done()
		var buff bytes.Buffer
		for _, b := range data {
			fmt.Fprintf(&buff, "%c", b)
		}
		fmt.Println(buff.String())
	}
	var wg sync.WaitGroup
	wg.Add(2)
	data := []byte("golang")
	go printData(&wg, data[:3])
	go printData(&wg, data[3:])
	var s string
	fmt.Printf("%q\n", s)
	wg.Wait()
}
