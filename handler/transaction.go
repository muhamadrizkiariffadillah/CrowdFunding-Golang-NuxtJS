package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/muhamadrizkiariffadillah/CrowdFunding-Golang-NuxtJS/helper"
	"github.com/muhamadrizkiariffadillah/CrowdFunding-Golang-NuxtJS/transaction"
	"github.com/muhamadrizkiariffadillah/CrowdFunding-Golang-NuxtJS/users"
)

type handler struct {
	service transaction.Service
}

func TransactionsHandler(service transaction.Service) *handler {
	return &handler{service}
}

func (h *handler) GetCampaignTransactions(c *gin.Context) {
	campaignId, _ := strconv.Atoi(c.Query("campaign_id"))

	transactionsList, err := h.service.GetTransactionsByCampaignId(campaignId)
	if err != nil {
		msg := gin.H{"error": err}
		response := helper.APIResponse(http.StatusNotFound, "failed", "failed to get campaign transactions", msg)
		c.JSON(http.StatusNotFound, response)
		return
	}

	formatter := transaction.CampaignTransactionsFormatter(transactionsList)
	response := helper.APIResponse(http.StatusOK, "success", "success to get campaign transactions", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *handler) GetUserTransactions(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(users.Users)
	userId := currentUser.Id
	userTransactions, err := h.service.GetTransactionByUserId(userId)
	if err != nil {
		msg := gin.H{"error": err}
		response := helper.APIResponse(http.StatusNotFound, "failed", "failed to get user transactions", msg)
		c.JSON(http.StatusNotFound, response)
		return
	}
	formatter := transaction.UserTrabsactionsFormatter(userTransactions)
	response := helper.APIResponse(http.StatusOK, "success", "success to get user transactions", formatter)
	c.JSON(http.StatusOK, response)
}
