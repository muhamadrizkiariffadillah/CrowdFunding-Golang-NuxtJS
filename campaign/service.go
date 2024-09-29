package campaign

import (
	"fmt"
	"time"

	"github.com/gosimple/slug"
)

type Service interface {
	CreateCampaign(input CreateCampaignInput)(Campaigns,error)
}

type campaignServices struct {
	r Repository
}

func CampaignServices(r Repository)*campaignServices{
	return &campaignServices{r}
}

func (s *campaignServices) CreateCampaign(input CreateCampaignInput)(Campaigns,error)  {
	campaign := Campaigns{
		CampaignName: input.CampaignName,
		ShortDescription: input.ShortDescription,
		Description: input.Description,
		GoalAmount: input.GoalAmount,
		Perks: input.Perks,
		UserId: input.User.Id,
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}
	slugCandidate := fmt.Sprintf("%v %v",input.CampaignName,input.User.Id)
	campaign.Slug = slug.Make(slugCandidate)
	newCampaign,err := s.r.Save(campaign)
	if err != nil {
		return Campaigns{},err
	}
	return newCampaign,nil
}