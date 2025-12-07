package dto

type RegisterRequest struct {
	Name     string `json:"name" validate:"required,min=2,max=100"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=100"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type AuthResponse struct {
	Token string `json:"token"`
	User  struct {
		ID    uint   `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	} `json:"user"`
}

func ValidateRegister(req RegisterRequest) map[string]string {
	errs := make(map[string]string)
	if req.Name == "" {
		errs["name"] = "Name must be at least 2 characters long"
	}
	if req.Email == "" {
		errs["email"] = "Email must be a valid email address"
	}
	if req.Password == "" || len(req.Password) < 6 {
		errs["password"] = "Password must be at least 6 characters long"
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func ValidateLogin(req LoginRequest) map[string]string {
	errs := make(map[string]string)
	if req.Email == "" {
		errs["email"] = "Email must be a valid email address"
	}
	if req.Password == "" || len(req.Password) < 6 {
		errs["password"] = "Password must be at least 6 characters long"
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}
