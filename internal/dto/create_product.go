package dto

type CreateProductRequest struct {
	Name        string `json:"name" validate:"required,min=2,max=100"`
	Description string `json:"description" validate:"required,min=2,max=100"`
	Price       uint   `json:"price" validate:"required,gt=0"`
}

func ValidateCreateProduct(req CreateProductRequest) map[string]string {
	err := validate.Struct(req)
	if err == nil {
		return nil
	}
	return map[string]string{"error": err.Error()}
}
