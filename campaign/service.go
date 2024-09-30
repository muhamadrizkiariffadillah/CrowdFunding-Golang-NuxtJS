package campaign

import (
	"fmt"
	"time"

	"github.com/gosimple/slug"
)

type Service interface {
	CreateCampaign(input CreateCampaignInput) (Campaigns, error)
}

type campaignServices struct {
	r Repository
}

func CampaignServices(r Repository) *campaignServices {
	return &campaignServices{r}
}

func (s *campaignServices) CreateCampaign(input CreateCampaignInput) (Campaigns, error) {
	campaign := Campaigns{
		UserId:           input.User.Id,
		CampaignName:     input.CampaignName,
		ShortDescription: input.ShortDescription,
		Description:      input.Description,
		GoalAmount:       input.GoalAmount,
		CurrentAmount:    0,
		Perks:            input.Perks,
		BackerCount:      0,
		UpdatedAt:        time.Now(),
		CreatedAt:        time.Now(),
	}
	slugCandidate := fmt.Sprintf("%v %v", input.User.Id, input.CampaignName)
	campaign.Slug = slug.Make(slugCandidate)
	newCampaign, err := s.r.Save(campaign)
	if err != nil {
		return Campaigns{}, err
	}
	return newCampaign, nil
}
