package handler

import (
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
	var input campaigns.GetCampaignDetailInput

	err := c.ShouldBindQuery(&input)
	if err != nil {
		response := helper.APIResponse(http.StatusNotFound, "failed", "fail to parser the request", err)
		c.JSON(http.StatusNotFound, response)
		return
	}

	campaign, err := h.service.GetCampaignById(input)
	if err != nil {
		response := helper.APIResponse(http.StatusInternalServerError, "failed", "fail to get campaign detail", err)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// formatter
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
		response := helper.APIResponse(http.StatusInternalServerError, "failed", "fail to get entity", errorMessage)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	formatter := campaigns.CampaignFormatter(newCampaign)
	response := helper.APIResponse(http.StatusCreated, "success", "successfully to create a campaign", formatter)
	c.JSON(http.StatusCreated, response)
	return
}
