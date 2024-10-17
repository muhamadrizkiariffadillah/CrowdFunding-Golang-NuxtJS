package transaction

import (
	"time"

	"github.com/muhamadrizkiariffadillah/CrowdFunding-Golang-NuxtJS/campaigns"
	"github.com/muhamadrizkiariffadillah/CrowdFunding-Golang-NuxtJS/users"
)

type Transactions struct {
	Id         int
	UserId     int `gorm:"foreignKey:user_id"`
	CampaignId int `gorm:"foreignKey:campaign_id"`
	Amount     uint
	Status     string
	Code       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	User       users.Users
	Campaign   campaigns.Campaigns
}
