package main

import "fmt"

func test() {
	fmt.Println("test")
}
func test2(a string) {
	fmt.Println(a)
}

func sum(a int, b int) int {
	return a + b
}
 

func main() {
	data := [...]any{}
	fmt.Println(data)

	fmt.Println("loyiha ishga tushdi")
	test()
	test2("salom")
	fmt.Println(sum(1, 2))
	// ananim funksiya
	ananim := func(a int, b int) int {
		return a / b
	}
	fmt.Println(ananim(1, 2))
}
