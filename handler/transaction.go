package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/muhamadrizkiariffadillah/CrowdFunding-Golang-NuxtJS/helper"
	"github.com/muhamadrizkiariffadillah/CrowdFunding-Golang-NuxtJS/transaction"
)

type handler struct {
	service transaction.Service
}

func TransactionsHandler(service transaction.Service) *handler {
	return &handler{service}
}

func (h *handler) GetCampaignTransactions(c *gin.Context) {
	campaignId, _ := strconv.Atoi(c.Query("campaign_id"))

	transactions, err := h.service.GetTransactionsByCampaignId(campaignId)
	if err != nil {
		msg := gin.H{"error": err}
		response := helper.APIResponse(http.StatusNotFound, "failed", "failed to get campaign transactions", msg)
		c.JSON(http.StatusNotFound, response)
		return
	}

	// todo: formatter
	// formatter
	response := helper.APIResponse(http.StatusOK, "success", "success to get campaign transactions", transactions)
	c.JSON(http.StatusOK, response)

}
