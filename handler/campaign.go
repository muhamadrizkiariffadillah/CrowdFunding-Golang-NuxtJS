package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/muhamadrizkiariffadillah/CrowdFunding-Golang-NuxtJS/campaigns"
	"github.com/muhamadrizkiariffadillah/CrowdFunding-Golang-NuxtJS/helper"
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
