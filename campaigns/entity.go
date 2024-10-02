package campaigns

import (
	"time"

	"github.com/muhamadrizkiariffadillah/CrowdFunding-Golang-NuxtJS/users"
)

type Campaigns struct {
	Id               int
	UserId           int
	CampaignName     string
	ShortDescription string
	Description      string
	GoalAmount       uint
	CurrentAmount    uint
	Perks            string
	BackerCount      uint
	Slug             string
	UpdatedAt        time.Time
	CreatedAt        time.Time
	CampaignImage    []CampaignImages `gorm:"foreignKey:campaign_id"`
	User             users.Users      `gorm:"foreignKey:user_id"`
}

type CampaignImages struct {
	Id         int
	CampaignId int `gorm:"foreignKey:campaign_id"`
	FileName   string
	IsPrimary  bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
