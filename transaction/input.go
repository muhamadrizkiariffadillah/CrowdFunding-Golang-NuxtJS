package transaction

import "github.com/muhamadrizkiariffadillah/CrowdFunding-Golang-NuxtJS/users"

type CreateTransactionInput struct {
	Amount     uint `json:"amount"`
	CampaignId int  `json:"campaign_id"`
	User       users.Users
}
