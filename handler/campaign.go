package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/muhamadrizkiariffadillah/CrowdFunding-Golang-NuxtJS/campaigns"
	"github.com/muhamadrizkiariffadillah/CrowdFunding-Golang-NuxtJS/helper"
	"github.com/muhamadrizkiariffadillah/CrowdFunding-Golang-NuxtJS/users"
)

type campaignHandler struct {
	service campaigns.Service
}

func CampaignHandler(service campaigns.Service) *campaignHandler {
	return &campaignHandler{service: service}
}

func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Query("user_id"))

	campaign, err := h.service.GetCampaigns(userId)
	if err != nil {
		errMsg := gin.H{
			"error": err,
		}
		response := helper.APIResponse(http.StatusInternalServerError, "failed", "fail to load campaigns", errMsg)
		c.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}

	formatter := campaigns.CampaignsFormatter(campaign)
	response := helper.APIResponse(http.StatusOK, "success", "success to load campaign", formatter)
	c.JSON(http.StatusOK, response)
	return
}

func (h *campaignHandler) GetCampaign(c *gin.Context) {
	campaignId, _ := strconv.Atoi(c.Query("id"))

	campaign, err := h.service.GetCampaignById(campaignId)
	if err != nil {
		response := helper.APIResponse(http.StatusInternalServerError, "failed", "fail to get campaign detail", err)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	formatter := campaigns.GetCampaignDetailFormatter(campaign)
	response := helper.APIResponse(http.StatusOK, "success", "successfully to get campaign detail", formatter)
	c.JSON(http.StatusOK, response)
	return
}

func (h *campaignHandler) CreateCampaign(c *gin.Context) {

	var input campaigns.CreateCampaignInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse(http.StatusUnprocessableEntity, "failed", "fail to get entity", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(users.Users)
	input.User = currentUser

	newCampaign, err := h.service.CreateCampaign(input)
	if err != nil {
		errorMessage := gin.H{"error": err}
		response := helper.APIResponse(http.StatusInternalServerError, "failed", "fail to create a campaign", errorMessage)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	formatter := campaigns.CampaignFormatter(newCampaign)
	response := helper.APIResponse(http.StatusCreated, "success", "successfully to create a campaign", formatter)
	c.JSON(http.StatusCreated, response)
	return
}

func (h *campaignHandler) UpdateCampaign(c *gin.Context) {
	// url: http://server/api/v1/campaign/update?id=
	campaignId, _ := strconv.Atoi(c.Query("id"))

	// body request JSON
	var inputData campaigns.CreateCampaignInput
	err := c.ShouldBindJSON(&inputData)
	if err != nil {
		formatter := helper.FormatValidationError(err)
		errorMsg := gin.H{"error": formatter}
		response := helper.APIResponse(http.StatusUnprocessableEntity, "failed", "failed to update the campaign", errorMsg)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(users.Users)
	inputData.User = currentUser

	campaign, err := h.service.GetCampaignById(campaignId)
	if err != nil {
		errorMsg := gin.H{"error": err}
		response := helper.APIResponse(http.StatusInternalServerError, "failed", "failed to get campaign", errorMsg)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	// Security Layer for IDOR.
	if campaign.User.Id != currentUser.Id {
		response := helper.APIResponse(http.StatusUnauthorized, "failed", "Unathorized action", nil)
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	updateCampaign, err := h.service.UpdateCampaign(campaignId, inputData)
	if err != nil {
		formatter := helper.FormatValidationError(err)
		errorMsg := gin.H{"error": formatter}
		response := helper.APIResponse(http.StatusInternalServerError, "failed", "failed to update the campaign", errorMsg)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	//
	formatter := campaigns.CampaignFormatter(updateCampaign)
	responseMsg := helper.APIResponse(http.StatusOK, "success", "successfully update campaign", formatter)
	c.JSON(http.StatusOK, responseMsg)
}

func (h *campaignHandler) SaveCampaignImage(c *gin.Context) {
	var input campaigns.CreateCampaignImageInput
	err := c.ShouldBind(&input)
	if err != nil {
		formatter := helper.FormatValidationError(err)
		errorMsg := gin.H{"error": formatter}
		response := helper.APIResponse(http.StatusUnprocessableEntity, "failed", "failed to save campaign images", errorMsg)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		errorMsg := gin.H{"is_uploaded": false}
		response := helper.APIResponse(http.StatusUnprocessableEntity, "failed", "failed to upload the image", errorMsg)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(users.Users)
	userId := currentUser.Id

	path := fmt.Sprintf("images/campaigns_images/%d-%s", userId, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		errorMsg := gin.H{"is_uploaded": false, "error": err}
		response := helper.APIResponse(http.StatusUnprocessableEntity, "failed", "failed to get campaign by id", errorMsg)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	campaign, err := h.service.GetCampaignById(input.CampaignId)
	if userId != campaign.User.Id {
		errorMsg := gin.H{"is_uploaded": false, "error": err}
		response := helper.APIResponse(http.StatusUnprocessableEntity, "failed", "unauthorized action", errorMsg)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	_, err = h.service.SaveCampaignImage(input, path)
	if err != nil {
		errorMsg := gin.H{"is_uploaded": false, "error": err}
		response := helper.APIResponse(http.StatusUnprocessableEntity, "failed", "failed to get campaign by id", errorMsg)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	responseMsg := gin.H{"is_uploaded": true}
	response := helper.APIResponse(http.StatusOK, "Success", "Successfully upload campaign image", responseMsg)
	c.JSON(http.StatusOK, response)
}
