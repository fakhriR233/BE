package handlers

import (
	authdto "_waysbook/dto/auth"
	dto "_waysbook/dto/result"
	"_waysbook/models"
	"_waysbook/pkg/bcrypt"
	jwtToken "_waysbook/pkg/jwt"
	"_waysbook/repositories"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
)

type handlerAuth struct {
  AuthRepository repositories.AuthRepository
}

func HandlerAuth(AuthRepository repositories.AuthRepository) *handlerAuth {
  return &handlerAuth{AuthRepository}
}

func (h *handlerAuth) Register(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

  request := new(authdto.RegisterRequest)
  if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
    w.WriteHeader(http.StatusBadRequest)
    response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
    json.NewEncoder(w).Encode(response)
    return
  }

  validation := validator.New()
  err := validation.Struct(request)
  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
    json.NewEncoder(w).Encode(response)
    return
  }

  password, err := bcrypt.HashingPassword(request.Password)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
    json.NewEncoder(w).Encode(response)
  }

  user := models.User{
    FullName:     	request.FullName,
    Email:    		request.Email,
    Password: 		password,
    Status: 		"user",
  }

  claims := jwt.MapClaims{}
  claims["id"] = user.ID
  claims["status"] = user.Status
  claims["exp"] = time.Now().Add(time.Hour * 2).Unix() // 2 hours expired

  token, errGenerateToken := jwtToken.GenerateToken(&claims)
  if errGenerateToken != nil {
	log.Println(errGenerateToken)
	fmt.Println("Unauthorize")
	return
  }

  

  data, err := h.AuthRepository.Register(user)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
    json.NewEncoder(w).Encode(response)
  }

  loginResponse := authdto.RegisterResponse{
	Email:    		data.Email,
	Token:			token,
  }

  w.WriteHeader(http.StatusOK)
  response := dto.SuccessResult{Code: http.StatusOK, Data: loginResponse}
  json.NewEncoder(w).Encode(response)
}

func (h *handlerAuth) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
  
	request := new(authdto.LoginRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
	  w.WriteHeader(http.StatusBadRequest)
	  response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
	  json.NewEncoder(w).Encode(response)
	  return
	}
  
	user := models.User{
	  Email:    request.Email,
	  Password: request.Password,
	}
  
	// Check email
	user, err := h.AuthRepository.Login(user.Email)
	if err != nil {
	  w.WriteHeader(http.StatusBadRequest)
	  response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
	  json.NewEncoder(w).Encode(response)
	  return
	}
  
	// Check password
	isValid := bcrypt.CheckPasswordHash(request.Password, user.Password)
	if !isValid {
	  w.WriteHeader(http.StatusBadRequest)
	  response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "wrong email or password"}
	  json.NewEncoder(w).Encode(response)
	  return
	}
  
	claims := jwt.MapClaims{}
	claims["id"] = user.ID
	claims["fullname"] = user.FullName
	claims["status"] = user.Status
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix() // 2 hours expired
  
	token, errGenerateToken := jwtToken.GenerateToken(&claims)
	if errGenerateToken != nil {
	  log.Println(errGenerateToken)
	  fmt.Println("Unauthorize")
	  return
	}
  
	loginResponse := authdto.LoginResponse{
	  FullName:     user.FullName,
	  Email:    	user.Email,
	  Status:    	user.Status,
	  Token:    	token,
	}
  
	w.Header().Set("Content-Type", "application/json")
	response := dto.SuccessResult{Code: http.StatusOK, Data: loginResponse}
	json.NewEncoder(w).Encode(response)
  
  }

  func (h *handlerAuth) CheckAuth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	
	// Check User by Id
	user, err := h.AuthRepository.Getuser(userId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	CheckAuthResponse := authdto.CheckAuthResponse{
		ID:       		user.ID,
		FullName:     	user.FullName,
		Email:    		user.Email,
		Status:   		user.Status,
		Gender:   		user.Gender,
		Phone:   		user.Phone,
		Avatar:   		user.Avatar,
	}

	w.Header().Set("Content-Type", "application/json")
	response := dto.SuccessResult{Code: http.StatusOK, Data: CheckAuthResponse}
	json.NewEncoder(w).Encode(response)
}

func convertResponseRegister(u models.User) authdto.RegisterResponse {
	return authdto.RegisterResponse{
	  Email:    	u.Email,
	}
}