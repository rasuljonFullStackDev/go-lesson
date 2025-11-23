package main

import "fmt"

func main() {
	var Name string = "Rasuljon"
	var Age int = 25
	fmt.Println("Name", Name, "Age", Age)

	var FullName = "Tursunboyev Rasuljon"
	var Age2 = 25
	fmt.Println("FullName", FullName, "Age2", Age2)

	count := 10
	fmt.Println("count", count)

	var a, b, c = 1, 2, 3
	fmt.Println("a", a, "b", b, "c", c)
	// ""=>bosh joy(string)
	// 0->flat,int
	// false->bool
	// nail->pointer,slice,map,interface

	var a1 int    //// defult 0 qiymatga ega boladi
	var b1 string /// deult "" qiymatga ega boladi
	var c1 bool   /// deulft false qiymatga ega boladi
	fmt.Println("a1", a1, "b1", b1, "c1", c1)

	// constta
	const Name1 string = "Rasuljon"
	const Age1 int = 25
	fmt.Println("name1", Name1, "age", Age1)

	// eslatma globlar ozgarivchilar ham bor bu maindan tashqairda elon qilinadi
	// agar main ichda elon qilib tashqarimga printlng qilsangiz ishlmaydi

	q1 := 6
	q2 := 7
	sum := q1 + q2
	fmt.Println("sum", sum)
	
	fmt.Println("hello world")
}
