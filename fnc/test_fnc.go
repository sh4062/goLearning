package main

import "fmt"

func foo1(a string, b string) (int, int) {
	return 123, 456
}
func foo2(a string, b string) (r1 int, r2 int) {
	r1 = 123
	r2 = 456
	return
}

func food1() {
	fmt.Println("food1")
}

func food2() {
	fmt.Println("food2")
}
func food3() {
	fmt.Println("food3")

}

func deferFunc() {
	fmt.Println("deferFunc")
}
func returnFunc() int {
	fmt.Println("returnFunc")
	return 1
}
func drFunc() int {
	defer deferFunc()
	return returnFunc()
}
func main() {
	fmt.Println(foo1("a", "b"))
	fmt.Println(foo2("a", "b"))
	defer food1()
	defer food2()
	defer food3()
	drFunc()

}
