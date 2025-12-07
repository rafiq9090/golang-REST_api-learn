package handlers

import (
	"encoding/json"
	"net/http"
	"task-api/internal/dto"
	"task-api/internal/models"
	"task-api/internal/service"
	"task-api/internal/utils"
)

var JWT_SECRET = []byte("your-super-secret-jwt-key-2025")

func Register(w http.ResponseWriter, r *http.Request) {
	var input dto.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.JSONError(w, "Invalid JSON", http.StatusBadRequest, nil)
		return
	}
	if errs := dto.ValidateRegister(input); errs != nil {
		utils.JSONError(w, "Validation failed", http.StatusBadRequest, errs)
		return
	}

	user := models.User{Name: input.Name, Email: input.Email, Password: input.Password}
	if err := service.Auth.Register(&user); err != nil {
		utils.JSONError(w, "Email already exists", http.StatusBadRequest, nil)
		return
	}

	utils.JSONSuccess(w, map[string]any{"message": "Registered successfully", "user_id": user.ID})
}

func Login(w http.ResponseWriter, r *http.Request) {
	var input dto.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.JSONError(w, "Invalid JSON", http.StatusBadRequest, nil)
		return
	}
	if errs := dto.ValidateLogin(input); errs != nil {
		utils.JSONError(w, "Validation failed", http.StatusBadRequest, errs)
		return
	}

	user, token, err := service.Auth.Login(input.Email, input.Password)
	if err != nil {
		utils.JSONError(w, "Invalid email or password", http.StatusUnauthorized, nil)
		return
	}

	resp := map[string]any{
		"token": token,
		"user": map[string]any{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		},
	}
	utils.JSONSuccess(w, resp)
}
