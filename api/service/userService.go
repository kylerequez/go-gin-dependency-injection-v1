package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kylerequez/go-gin-dependency-injection-v1/api/repository"
	"github.com/kylerequez/go-gin-dependency-injection-v1/types"
)

type UserService struct {
	ur *repository.UserRepository
}

func NewUserService(ur *repository.UserRepository) *UserService {
	return &UserService{ur: ur}
}

func (us *UserService) GetAllUsers(c *gin.Context) {
	users, err := us.ur.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusFound, gin.H{
		"message": "Success",
		"users":   users,
	})
}

func (us *UserService) GetUserById(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Error",
			"message": "id is not in the url path",
		})
		return
	}

	uuid, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Error",
			"message": "id was not parsed into a valid uuid",
		})
		return
	}

	result, err := us.ur.GetUserById(uuid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Error",
			"message": err,
		})
		return
	}

	c.JSON(http.StatusFound, gin.H{
		"status": "Success",
		"user":   result,
	})
}

func (us *UserService) CreateUser(c *gin.Context) {
	type Body struct {
		FirstName  string `json:"first_name"`
		MiddleName string `json:"middle_name"`
		LastName   string `json:"last_name"`
		Email      string `json:"email"`
		Password   string `json:"password"`
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

	if firstName == "" ||
		middleName == "" ||
		lastName == "" ||
		email == "" ||
		password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Error",
			"message": "Inputs must not be empty strings",
		})
		return
	}

	var newUser types.User = types.User{
		FirstName:  firstName,
		MiddleName: middleName,
		LastName:   lastName,
		Email:      email,
		Password:   password,
	}

	result, err := us.ur.CreateUser(&newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Error",
			"message": err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "Success",
		"message": "You have successfully created a new user",
		"user":    result,
	})
}

func (us *UserService) PatchUpdateUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Error",
			"message": "id is not in the url path",
		})
		return
	}

	uuid, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Error",
			"message": "id was not parsed into a valid uuid",
		})
		return
	}

	type Body struct {
		FirstName  string `json:"first_name"`
		MiddleName string `json:"middle_name"`
		LastName   string `json:"last_name"`
		Authority  string `json:"authority"`
		Email      string `json:"email"`
		Password   string `json:"password"`
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
	authority := body.Authority
	email := body.Email
	password := body.Password

	if firstName == "" ||
		middleName == "" ||
		lastName == "" ||
		authority == "" ||
		email == "" ||
		password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Error",
			"message": "Inputs must not be empty strings",
		})
		return
	}

	result, err := us.ur.GetUserById(uuid)
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
			"message": "User was not found",
		})
		return
	}

	var user *types.User = result
	user.FirstName = firstName
	user.MiddleName = middleName
	user.LastName = lastName
	user.Email = email
	user.Password = password

	result, err = us.ur.UpdateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Error",
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "The user has been updated",
		"user":    result,
	})
}

func (us *UserService) PutUpdateUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Error",
			"message": "id is not in the url path",
		})
		return
	}

	uuid, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Error",
			"message": "id was not parsed into a valid uuid",
		})
		return
	}

	type Body struct {
		FirstName  string `json:"first_name"`
		MiddleName string `json:"middle_name"`
		LastName   string `json:"last_name"`
		Authority  string `json:"authority"`
		Email      string `json:"email"`
		Password   string `json:"password"`
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
	authority := body.Authority
	email := body.Email
	password := body.Password

	if firstName == "" ||
		middleName == "" ||
		lastName == "" ||
		email == "" ||
		password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Error",
			"message": "Inputs must not be empty strings",
		})
		return
	}

	result, err := us.ur.GetUserById(uuid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Error",
			"message": err,
		})
		return
	}

	var user *types.User

	if result != nil {
		user = result
		user.FirstName = firstName
		user.MiddleName = middleName
		user.LastName = lastName
		user.Authority = authority
		user.Email = email
		user.Password = password
	} else {
		user = &types.User{
			FirstName:  firstName,
			MiddleName: middleName,
			LastName:   lastName,
			Authority:  authority,
			Email:      email,
			Password:   password,
		}
	}

	if result == nil {
		result, err := us.ur.CreateUser(user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  "Error",
				"message": err,
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"status":  "Success",
			"message": "The user has been created",
			"user":    result,
		})
		return
	} else {
		result, err = us.ur.UpdateUser(user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "Error",
				"message": err,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  "Success",
			"message": "The user has been updated",
			"user":    result,
		})
		return
	}
}

func (us *UserService) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Error",
			"message": "id is not in the url path",
		})
		return
	}

	uuid, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Error",
			"message": "id was not parsed into a valid uuid",
		})
		return
	}

	result, err := us.ur.DeleteUser(uuid)
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
			"message": "User does not exists or is already deleted",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "You have deleted the user",
		"result":  result,
	})
}
