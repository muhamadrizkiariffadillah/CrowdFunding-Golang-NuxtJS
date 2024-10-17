package handler

import (
	"fmt"
	"net/http"

	"github.com/muhamadrizkiariffadillah/CrowdFunding-Golang-NuxtJS/authJWT"

	"github.com/gin-gonic/gin"
	"github.com/muhamadrizkiariffadillah/CrowdFunding-Golang-NuxtJS/helper"
	"github.com/muhamadrizkiariffadillah/CrowdFunding-Golang-NuxtJS/users"
)

// userHandler handles user-related HTTP requests.
type userHandler struct {
	userService users.Service // Service layer for user operations.
	authService authJWT.Service
}

// UserHandler initializes a new userHandler with the provided user service.
func UserHandler(userService users.Service, authService authJWT.Service) *userHandler {
	return &userHandler{userService, authService}
}

func (h *userHandler) Signup(c *gin.Context) {
	var input users.RegisterUserInput

	// Bind incoming JSON request to input struct and validate it.
	err := c.ShouldBindJSON(&input)
	if err != nil {
		// Format and return validation errors as JSON response.
		errors := helper.FormatValidationError(err)
		errorMassage := gin.H{
			"error": errors,
		}
		response := helper.APIResponse(http.StatusUnprocessableEntity, "Failed", "", errorMassage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// Register the user through the service layer.
	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		// Return error response if user registration fails.
		response := helper.APIResponse(http.StatusInternalServerError, "Failed", "", nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	// Format the new user data and return success response.
	formatter := users.APIUserFormatter(newUser, "")
	response := helper.APIResponse(http.StatusCreated, "Success", "Your account has been created", formatter)
	c.JSON(http.StatusCreated, response)
}

func (h *userHandler) Login(c *gin.Context) {
	var input users.LoginUserInput

	// Bind incoming JSON request to input struct and validate it.
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{
			"Errors:": errors,
		}
		response := helper.APIResponse(http.StatusUnprocessableEntity, "Failed", "errors", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// Log in the user through the service layer.
	loggedUser, err := h.userService.LoginUser(input)
	if err != nil {
		errorMessage := gin.H{
			"errors:": err,
		}
		response := helper.APIResponse(http.StatusNotFound, "Failed", "username or password incorrect", errorMessage)
		c.JSON(http.StatusNotFound, response)
		return
	}

	token, err := h.authService.GenerateToken(loggedUser.Id)
	if err != nil {
		errorMessage := gin.H{
			"errors": "generate token error",
		}
		response := helper.APIResponse(http.StatusInternalServerError, "Failed", "errors", errorMessage)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	// Format the logged-in user data and return success response.
	formatter := users.APIUserFormatter(loggedUser, token)
	response := helper.APIResponse(http.StatusOK, "Success", "Successfully logged in", formatter)
	c.JSON(http.StatusOK, response)
	return
}

func (h *userHandler) FetchUser(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(users.Users)
	formatter := users.APIUserFormatter(currentUser, "")
	response := helper.APIResponse(http.StatusOK, "Success", "Successfully fetch user data", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) CheckEmail(c *gin.Context) {
	var input users.CheckEmailInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{
			"errors": errors,
		}
		response := helper.APIResponse(http.StatusUnprocessableEntity, "Failed", "", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	isEmailAvailable, err := h.userService.IsEmailAvailable(input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{
			"errors": errors,
		}
		response := helper.APIResponse(http.StatusInternalServerError, "Failed", "", errorMessage)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	message := gin.H{
		"is_available": isEmailAvailable,
	}

	if !isEmailAvailable {
		response := helper.APIResponse(http.StatusNotFound, "failed", "email has been registered", message)
		c.JSON(http.StatusNotFound, response)
		return
	}

	response := helper.APIResponse(http.StatusOK, "success", "email is available", message)
	c.JSON(http.StatusOK, response)

}

// TODO: Need to refactor this handler after finish this course.
func (h *userHandler) UploadAvatar(c *gin.Context) {

	// Get the current user from the context
	currentUser, exists := c.Get("currentUser")
	if !exists {
		response := helper.APIResponse(http.StatusUnauthorized, "Failed", "Unauthorized", nil)
		c.JSON(http.StatusUnauthorized, response)
		return
	}
	user := currentUser.(users.Users)
	userId := user.Id

	// Get the uploaded file from the request
	file, err := c.FormFile("avatar")
	if err != nil {
		errorMsg := gin.H{
			"error": err.Error(),
		}
		response := helper.APIResponse(http.StatusBadRequest, "Error", "Avatar not available", errorMsg)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// TODO need to update the file-name, configure the image is not allowed more then 2mb, and 2048 pixel.
	// Create the path for saving the uploaded file
	path := fmt.Sprintf("./images/user/avatar/%s", file.Filename)

	// Save the uploaded file to the specified location
	if err := c.SaveUploadedFile(file, path); err != nil {
		errorMsg := gin.H{
			"error": err.Error(),
		}
		response := helper.APIResponse(http.StatusInternalServerError, "Error", "Failed to save avatar", errorMsg)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// Save the avatar information in the database
	if _, err := h.userService.UploadAvatar(userId, path); err != nil {
		errorMsg := gin.H{
			"error": err.Error(),
		}
		response := helper.APIResponse(http.StatusInternalServerError, "Error", "Failed to save avatar", errorMsg)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	msg := gin.H{
		"is_uploaded": true,
		"file_url":    path,
	}
	response := helper.APIResponse(http.StatusOK, "success", "successfully to upload the file", msg)
	c.JSON(http.StatusOK, response)
	return
}
