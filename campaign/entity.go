package campaign

import "time"


type Campaigns struct {
	Id int
	UserId int
	CampaignName string
	ShortDescription string
	Description string
	GoalAmount uint
	CurrentAmount uint
	Perks string
	BackerCount uint
	Slug string
	UpdatedAt time.Time
	CreatedAt time.Time

}