package user

import (
	"encoding/json"
	"net/http"

	userDto "github.com/amr-mahmoud/mrktly/internal/interface/dto/user"
	"github.com/go-chi/chi/v5"
)

type UserController struct {
}


func NewUserController() *UserController {
	return &UserController{}
}


func (uc *UserController) Routes() chi.Router {
    r := chi.NewRouter()
    
	r.Post("/signup", uc.SignUpProfessionalHandler)
	r.Post("/login", uc.LoginProfessionalHandler)
    
    return r
}

func (uc *UserController) SignUpProfessionalHandler(w http.ResponseWriter, r *http.Request) {
	var input userDto.AuthInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	output, err := uc.SignUpProfessional(input)
	if err != nil {
		http.Error(w, "Signup failed", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}

func (uc *UserController) SignUpProfessional(input userDto.AuthInput) (userDto.AuthOutput, error) {
	return userDto.AuthOutput{
		ID:           "12345",
		Email:        input.Email,
		Phone:        input.Phone,
		Verified:     false,
	}, nil
}


func (uc *UserController) LoginProfessionalHandler(w http.ResponseWriter, r *http.Request) {
	var input userDto.AuthOutput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	output, err := uc.LoginProfessional(input)
	if err != nil {
		http.Error(w, "Login failed", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}

func (uc *UserController) LoginProfessional(input userDto.AuthOutput) (userDto.AuthOutput, error) {
	return userDto.AuthOutput{
		ID:           "12345",
		Email:        input.Email,
		Phone:        input.Phone,
		Verified:     false,
	}, nil
}