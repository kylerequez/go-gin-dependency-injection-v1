package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kylerequez/go-gin-dependency-injection-v1/api/repository"
	"github.com/kylerequez/go-gin-dependency-injection-v1/types"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	ur *repository.UserRepository
	js *JwtService
}

func NewAuthService(ur *repository.UserRepository, js *JwtService) *AuthService {
	return &AuthService{
		ur: ur,
		js: js,
	}
}

func (as *AuthService) RegistrationHandler(c *gin.Context) {
	type Body struct {
		FirstName       string `json:"first_name"`
		MiddleName      string `json:"middle_name"`
		LastName        string `json:"last_name"`
		Email           string `json:"email"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirm_password"`
	}

	body := new(Body)

	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Error",
			"message": "There was an error in binding the response body",
		})
		return
	}

	firstName := body.FirstName
	middleName := body.MiddleName
	lastName := body.LastName
	email := body.Email
	password := body.Password
	confirmPassword := body.ConfirmPassword

	if firstName == "" ||
		middleName == "" ||
		lastName == "" ||
		email == "" ||
		password == "" ||
		confirmPassword == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Error",
			"message": "Inputs must not be empty strings",
		})
		return
	}

	if password != confirmPassword {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Error",
			"message": "Password/s must be the same",
		})
		return
	}

	result, err := as.ur.GetUserByEmail(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Error",
			"message": err,
		})
		return
	}

	if result != nil {
		c.JSON(http.StatusConflict, gin.H{
			"status":  "Error",
			"message": "User already exists",
		})
		return
	}

	hashedPassword, err := HashPassword(password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Error",
			"message": err,
		})
		return
	}

	var newUser types.User = types.User{
		FirstName:  firstName,
		MiddleName: middleName,
		LastName:   lastName,
		Email:      email,
		Password:   hashedPassword,
	}

	result, err = as.ur.CreateUser(&newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Error",
			"message": err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "Success",
		"message": "Successfully registered the user",
		"user":    result,
	})
}

func (as *AuthService) LoginHandler(c *gin.Context) {
	type Body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	body := new(Body)

	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Error",
			"message": "There was an error in binding the response body",
		})
		return
	}

	email := body.Email
	password := body.Password

	if email == "" ||
		password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Error",
			"message": "Inputs must not be empty strings",
		})
		return
	}

	result, err := as.ur.GetUserByEmail(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Error",
			"message": err,
		})
		return
	}

	if result == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "Error",
			"message": "User does not exists",
		})
		return
	}

	var user types.User = *result
	hashedPassword := user.Password

	isVerified := VerifyPassword(password, hashedPassword)
	if !isVerified {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Error",
			"message": "Password does not match",
		})
		return
	}

	token, err := as.js.GenerateJWT(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Error",
			"message": err,
		})
		return
	}

	c.SetCookie(
		"go-gin-dependency-injection-v1-token",
		token,
		3600,
		"/",
		"localhost",
		false,
		false,
	)

	c.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Successfully logged in",
	})
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func VerifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
