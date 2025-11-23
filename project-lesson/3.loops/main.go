package main

import "fmt"

type Product struct {
	Name  string
	Color string
	Price int
	Body  int
}

func main() {
	fmt.Println("loyiha ishga tushdi")
	for i := 1; i < 5; i++ {
		fmt.Println("Natija ", i)
	}

	k := 0
	for k < 3 {
		fmt.Println("yangi natija", k)
		k++
	}
	// for orqali cheksiz sikl qilish imkoni ekan
	count := 1
	for {
		fmt.Print("Natija 2 ", count)
		count++
		if count < 3 {
			break
		}
	}

	// arraylar bilan ishlash
	nums := []int{1, 2, 3, 4, 4}

	for index, value := range nums {
		fmt.Println("index", index, "value", value)
	}
	// assositv massiv
	data := map[string]string{
		"name":  "loyiha",
		"email": "loyiha@.com",
	}
	for key, value := range data {
		fmt.Println("key", key, "value", value)
	}
	fmt.Println("data", data["name"])
	fmt.Println("data", data)
	//  product datasini tuzib korish
	//
	for i := 1; i <= 5; i++ {
		if i == 3 {
			continue // 3 ni o‘tkazib ketadi
		}
		if i == 5 {
			break // 5 da to‘xtaydi
		}
		fmt.Println(i)
	}

	for i := 1; i <= 5; i++ {
		switch {
		case i%2 == 0:
			fmt.Println(i, "— juft son")
		case i%2 != 0:
			fmt.Println(i, "— toq son")
		}
	}

}
