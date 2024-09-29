package campaign

import "github.com/muhamadrizkiariffadillah/CrowdFunding-Golang-NuxtJS/users"

type CreateCampaignInput struct {
	CampaignName string `json:"campaign_name" binding:"required"`
	ShortDescription string `json:"short_description" binding:"required"`
	Description string `json:"description" binding:"required"`
	GoalAmount uint `json:"goal_amount" binding:"required"`
	Perks string `json:"perks" binding:"required"`
	User users.Users `json:"user" binding:"required"`
}