package handler

import (
	"net/http"
	"strconv"

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

func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.service.GetCampaigns(userId)
	if err != nil {
		errMsg := gin.H{
			"error": err,
		}
		response := helper.APIResponse(http.StatusInternalServerError, "failed", "fail to load campaigns", errMsg)
		c.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}

	response := helper.APIResponse(http.StatusOK, "success", "success to load campaign", campaigns)
	c.JSON(http.StatusOK, response)
	return
}
