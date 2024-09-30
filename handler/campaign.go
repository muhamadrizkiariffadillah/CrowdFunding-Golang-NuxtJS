package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muhamadrizkiariffadillah/CrowdFunding-Golang-NuxtJS/campaign"
	"github.com/muhamadrizkiariffadillah/CrowdFunding-Golang-NuxtJS/helper"
)

type campaignHandler struct {
	service campaign.Service
}

func CampaignHandler(service campaign.Service) *campaignHandler {
	return &campaignHandler{service: service}
}

func (h *campaignHandler) SaveCampaign(c *gin.Context) {
	// currentUser := c.MustGet("currentUser").(users.Users)

	var input campaign.CreateCampaignInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		message := gin.H{
			"error": errors,
		}
		response := helper.APIResponse(http.StatusUnprocessableEntity, "failed", "create campaign fail", message)
		c.JSON(http.StatusUnsupportedMediaType, response)
		return
	}

	response := helper.APIResponse(http.StatusOK, "success", "succesfully to create campaign", nil)
	c.JSON(http.StatusOK, response)
	return
}
