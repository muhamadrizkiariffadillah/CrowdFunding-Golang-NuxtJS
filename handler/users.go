package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muhamadrizkiariffadillah/CrowdFunding-Golang-NuxtJS/helper"
	"github.com/muhamadrizkiariffadillah/CrowdFunding-Golang-NuxtJS/users"
)

// userHandler handles user-related HTTP requests.
type userHandler struct {
	userService users.Service // Service layer for user operations.
}

// UserHandler initializes a new userHandler with the provided user service.
func UserHandler(userService users.Service) *userHandler {
	return &userHandler{userService}
}

// Signup handles the user registration process.
// It validates the input, registers the user using the service layer,
// and returns the appropriate API response.
// @Summary Register a new user
// @Description This endpoint allows you to register a new user.
// @Tags Users
// @Accept json
// @Produce json
// @Param request body users.RegisterUserInput true "User registration data"
// @Success 201 {object} helper.Response "Your account has been created"
// @Failure 422 {object} helper.Response "Validation errors"
// @Failure 500 {object} helper.Response "Server error"
// @Router /users [post]
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
	formatter := users.RegisterUserFormatter(newUser, "")
	response := helper.APIResponse(http.StatusCreated, "Success", "Your account has been created", formatter)
	c.JSON(http.StatusCreated, response)
}
