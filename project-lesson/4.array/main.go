package main

import "fmt"

type Product struct {
	Name  string
	Price int
	color string
}

func main() {
	fmt.Println("hello world")
	var nums [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(nums)
	nums2 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(nums2)
	fmt.Println(nums[0])
	fmt.Println(len(nums2)) /// uzunligi
	for index, value := range nums {
		fmt.Println("index", index, "value", value)
	}
	products := [...]Product{
		{Name: "apple", Price: 100, color: "red"},
		{Name: "apple", Price: 100, color: "red"},
	}
	for index, value := range products {
		fmt.Println("index", index, "value", value)
	}
	// eslatma ozida find,slice,filter funksiylari yuq qolda yoziladi loki kutubxona ishlatiladi
}
