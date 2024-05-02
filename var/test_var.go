package main

import (
	"fmt"
	"sync"
)

var a float64 = 1

const (
	//iota 只能和const配合累加
	A = iota
	B
	C
)

func main() {
	//不能生成全局变量
	//e := 1
	//fmt.Println(e)
	//fmt.Println(a)
	//fmt.Println(A)
	//fmt.Println(B)
	//fmt.Println(C)
	mu := &sync.RWMutex{}
	mu.Lock()
	mu.RLock()
	fmt.Println("hello")
}
