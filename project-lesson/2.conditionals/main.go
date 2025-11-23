package main

import "fmt"
type Product struct {
	Name string
	Color string 
	Price int
	Body int
}
func main() {
	fmt.Println("loyiha ishga tushdi")
	age := 20
	if age > 18 {
		fmt.Println("Yoshi 18 dan katta")
	} else {
		fmt.Println("Yoshi 18 dan kichik")
	}
	if age > 18 || age == 20 {
		fmt.Println("20 ga teng yoki yoshi 18 dan katta")
	}
	// a := 1
	// b := "1"
	// fmt.Println(a==b)
	// eslatma string va numberlarni solshtrib bolmas ekan har xil typelarni
	// == va === bormi ?
	// string ichaga qanday ozgaruvchi qoshib chiqarsa boladi?
	name1 := "Ali"
	age1 := 20

	msg := fmt.Sprintf("Salom, %s! Siz %d yoshdasiz.", name1, age1)
	fmt.Println(msg)
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

	day := "yakshanba"

	switch day {
	case "dushanba":
		fmt.Println("Hafta boshlandi")
	case "yakshanba":
		fmt.Println("Dam olish kuni!")
	default:
		fmt.Println("Oddiy ish kuni")
	}

	switch {
	case age < 18:
		fmt.Println("Voyaga yetmagan")
	case age >= 18 && age < 30:
		fmt.Println("Yosh avlod")
	default:
		fmt.Println("Katta yoshli")
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
