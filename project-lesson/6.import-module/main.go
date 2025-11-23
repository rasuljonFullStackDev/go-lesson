package main

import (
	"6.import-module/utils"
	"fmt"
)

// eslatma loyiha nomi va folder korsatilar ekan ichadgi barcha funksiylarni oqiydi
// eslatma  funksiya bosh harfi katta bolsa import qilib ishlatsa boladi kichiki bolsa ozida ishlatiladi
// funsiyalar returun holatiga doim type berish kerak
func main() {
	fmt.Println("loyiha ishga tushdi")
	fmt.Println(utils.GetAge())
	fmt.Println("calculate", utils.Calculate(1, 2))
	// untils.test() ishlatib bolmaydi
}
