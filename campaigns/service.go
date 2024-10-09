package campaigns

import (
	"fmt"

	"github.com/gosimple/slug"
)

type Service interface {
	GetCampaigns(userId int) ([]Campaigns, error)
	GetCampaignById(input GetCampaignDetailInput) (Campaigns, error)
	CreateCampaign(input CreateCampaignInput) (Campaigns, error)
}

type services struct {
	r Repository
}

func CampaignServices(r Repository) *services {
	return &services{r}
}

func (s *services) GetCampaigns(userId int) ([]Campaigns, error) {
	if userId != 0 {
		campaigns, err := s.r.FindByUserId(userId)
		if err != nil {
			return []Campaigns{}, err
		}
		return campaigns, nil
	}
	campaigns, err := s.r.FindAll()
	if err != nil {
		return []Campaigns{}, err
	}
	return campaigns, nil
}

func (s *services) GetCampaignById(input GetCampaignDetailInput) (Campaigns, error) {
	campaign, err := s.r.FindById(input.Id)
	if err != nil {
		return Campaigns{}, err
	}
	return campaign, nil
}
func (s *services) CreateCampaign(input CreateCampaignInput) (Campaigns, error) {
	campaign := Campaigns{
		CampaignName:     input.CampaignName,
		ShortDescription: input.ShortDescription,
		Description:      input.Description,
		GoalAmount:       input.GoalAmount,
		Perks:            input.Perks,
		UserId:           input.User.Id,
	}
	slugCandidate := fmt.Sprintf("%v %s", input.User.Id, input.CampaignName)
	campaign.Slug = slug.Make(slugCandidate)

	newCampaign, err := s.r.Save(campaign)
	if err != nil {
		return Campaigns{}, err
	}
	return newCampaign, nil
}
