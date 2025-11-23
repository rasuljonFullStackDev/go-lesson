package request

import  "github.com/go-playground/validator"
// ProductRequest - foydalanuvchidan keladigan so‘rov structi
type ProductRequest struct {
	Name       string  `json:"name" validate:"required,min=3,max=50"`
	Price      float64 `json:"price" validate:"required,gt=0"`
	CategoryID int     `json:"category_id" validate:"required"`
	Color      string  `json:"color" validate:"omitempty,max=30"`
	Size       string  `json:"size" validate:"omitempty,max=10"`
}

// Validate funksiyasi Laravel’dagi rules() o‘rnida
func (r *ProductRequest) Validate() map[string]string {
	v := validator.New()
	err := v.Struct(r)

	if err == nil {
		return nil
	}

	errors := make(map[string]string)
	for _, e := range err.(validator.ValidationErrors) {
		switch e.Field() {
		case "Name":
			errors["name"] = "Mahsulot nomi majburiy va 3–50 belgidan iborat bo‘lishi kerak"
		case "Price":
			errors["price"] = "Narx musbat son bo‘lishi shart"
		case "CategoryID":
			errors["category_id"] = "Kategoriya majburiy tanlanishi kerak"
		default:
			errors[e.Field()] = "Noto‘g‘ri qiymat"
		}
	}

	return errors
}
