package handler

import (
	"image"
	"image/jpeg"
	"image/png"
	"mime/multipart"
	"net/http"
	"os"

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
// @Router /api/v1/users/signup [post]
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
	token, err := h.authService.GenerateToken(newUser.Id)
	if err != nil {
		errorMessage := gin.H{
			"error": err,
		}
		response := helper.APIResponse(http.StatusInternalServerError, "Failed", "", errorMessage)
		c.JSON(http.StatusInternalServerError, response)
	}
	// Format the new user data and return success response.
	formatter := users.APIUserFormatter(newUser, token)
	response := helper.APIResponse(http.StatusCreated, "Success", "Your account has been created", formatter)
	c.JSON(http.StatusCreated, response)
}

// Login handles user login.
// It validates the input credentials and logs in the user if successful.
// @Summary Login a user
// @Description This endpoint allows an existing user to log in.
// @Tags Users
// @Accept json
// @Produce json
// @Param request body users.LoginUserInput true "User login data"
// @Success 200 {object} helper.Response "Successfully logged in"
// @Failure 422 {object} helper.Response "Validation errors"
// @Failure 500 {object} helper.Response "Server error"
// @Router /api/v1/users/login [post]
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
}

// FetchUser retrieves the currently authenticated user's data.
// @Summary Get current user data
// @Description This endpoint fetches the data of the currently logged-in user.
// @Tags Users
// @Produce json
// @Success 200 {object} helper.Response "Successfully fetch user data"
// @Failure 401 {object} helper.Response "Unauthorized"
// @Router /api/v1/users/me [get]
func (h *userHandler) FetchUser(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(users.Users)
	formatter := users.APIUserFormatter(currentUser, "")
	response := helper.APIResponse(http.StatusOK, "Success", "Successfully fetch user data", formatter)
	c.JSON(http.StatusOK, response)
}

// CheckEmail
// @Summary Check email availability
// @Description This endpoint check email when the new user signup
// @Tags Users
// @Produce json
// @Success 200 {object} helper.Response "email is available"
// @Failure 401 {object} helper.Response "email has been registered"
// @Router /api/v1/users/check-email [POST]
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

// UploadAvatar
// @Summary Upload Avatar user endpoint
// @Description this endpoint is used to upload avatar users
// @Tags Users
// @Produce json
// @Success 200 {object} helper.Response "email is available"
// @Failure 401 {object} helper.Response "email has been registered"
// @Failure 500 {object} helper.Response ""
// @Router /api/v1/users/check-email [POST]
// TODO: Need to refactor this handler after finish this course.
func (h *userHandler) UploadAvatar(c *gin.Context) {

	const maxFileSize = 2 << 20
	const maxDimension = 4000

	file, err := c.FormFile("file")
	if err != nil {
		data := gin.H{
			"is_uploaded": false,
		}
		response := helper.APIResponse(http.StatusBadRequest, "Failed", "Failed to upload", data)
		c.JSON(http.StatusBadRequest, response)
	}

	if file.Size > maxFileSize {
		data := gin.H{
			"is_uploaded": false,
		}
		response := helper.APIResponse(http.StatusBadRequest, "Failed", "image should less than 2 mb.", data)
		c.JSON(http.StatusBadRequest, response)
	}

	openedFile, err := file.Open()
	if err != nil {
		data := gin.H{
			"is_uploaded": false,
		}
		response := helper.APIResponse(http.StatusBadRequest, "Failed", "cannot open the image", data)
		c.JSON(http.StatusBadRequest, response)
	}
	defer func(openedFile multipart.File) {
		err := openedFile.Close()
		if err != nil {
			c.JSON(http.StatusInternalServerError, nil)
		}
	}(openedFile)

	buf := make([]byte, 512)
	_, err = openedFile.Read(buf)
	if err != nil {
		data := gin.H{
			"is_uploaded": false,
		}
		response := helper.APIResponse(http.StatusBadRequest, "Failed", "cannot read the image byte", data)
		c.JSON(http.StatusBadRequest, response)
	}

	fileType := http.DetectContentType(buf)
	if fileType != "image/jpeg" && fileType != "image/png" {
		data := gin.H{
			"is_uploaded": false,
		}
		response := helper.APIResponse(http.StatusBadRequest, "Failed", "Invalid file type", data)
		c.JSON(http.StatusBadRequest, response)
	}

	_, err = openedFile.Seek(0, os.SEEK_SET)
	if err != nil {
		data := gin.H{
			"is_uploaded": false,
		}
		response := helper.APIResponse(http.StatusBadRequest, "Failed", "Invalid file type", data)
		c.JSON(http.StatusBadRequest, response)
	}

	img, _, err := image.Decode(openedFile)
	if err != nil {
		data := gin.H{
			"is_uploaded": false,
		}
		response := helper.APIResponse(http.StatusBadRequest, "Failed", "Could not decode the image", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	width := img.Bounds().Max.X
	height := img.Bounds().Max.Y

	if width > maxDimension || height > maxDimension {
		data := gin.H{
			"is_uploaded": false,
		}
		response := helper.APIResponse(http.StatusBadRequest, "Failed", "Image dimensions too large. Maximum is 4000x4000 pixels", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	path := "/images/users/avatar" + file.Filename

	outFile, err := os.Create(path)
	if err != nil {
		data := gin.H{
			"is_uploaded": false,
		}
		response := helper.APIResponse(http.StatusBadRequest, "Failed", "Could not save the image", data)
		c.JSON(http.StatusBadRequest, response)
	}
	defer func(outFile *os.File) {
		err := outFile.Close()
		if err != nil {
			data := gin.H{
				"is_uploaded": false,
			}
			response := helper.APIResponse(http.StatusBadRequest, "Failed", "Could not save the image", data)
			c.JSON(http.StatusBadRequest, response)
		}
	}(outFile)

	if fileType == "image/jpeg" {
		err = jpeg.Encode(outFile, img, &jpeg.Options{Quality: 75})
	}
	if fileType == "image/png" {
		err = png.Encode(outFile, img)
	}

	if err != nil {
		data := gin.H{
			"is_uploaded": false,
		}
		response := helper.APIResponse(http.StatusInternalServerError, "Failed", "Failed to compress and save the file", data)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{
			"is_uploaded": false,
		}
		response := helper.APIResponse(http.StatusInternalServerError, "Failed", "Failed to save avatar", data)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	currentUser := c.MustGet("currentUser").(users.Users)
	userId := currentUser.Id

	_, err = h.userService.UploadAvatar(userId, path)
	if err != nil {
		data := gin.H{
			"is_uploaded": false,
		}
		response := helper.APIResponse(http.StatusInternalServerError, "Failed", "Failed to save avatar", data)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	data := gin.H{
		"is_uploaded":   true,
		"file_location": path,
	}
	response := helper.APIResponse(http.StatusOK, "Success", "successfully upload the avatar", data)
	c.JSON(http.StatusOK, response)
	return
}
