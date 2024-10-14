package campaigns

import (
	"errors"
	"fmt"
	"time"

	"github.com/gosimple/slug"
)

type Service interface {
	GetCampaigns(userId int) ([]Campaigns, error)
	GetCampaignById(campaignId int) (Campaigns, error)
	CreateCampaign(input CreateCampaignInput) (Campaigns, error)
	UpdateCampaign(campaignId int, inputData CreateCampaignInput) (Campaigns, error)
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

func (s *services) GetCampaignById(campaignId int) (Campaigns, error) {
	campaign, err := s.r.FindById(campaignId)
	if err != nil {
		return Campaigns{}, err
	}
	return campaign, nil
}

func (s *services) CreateCampaign(input CreateCampaignInput) (Campaigns, error) {
	campaign := Campaigns{
		UserId:           input.User.Id,
		CampaignName:     input.CampaignName,
		ShortDescription: input.ShortDescription,
		Description:      input.Description,
		CurrentAmount:    0,
		GoalAmount:       input.GoalAmount,
		Perks:            input.Perks,
		BackerCount:      0,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}
	slugCandidate := fmt.Sprintf("%v %s", input.User.Id, input.CampaignName)
	campaign.Slug = slug.Make(slugCandidate)

	newCampaign, err := s.r.Save(campaign)
	if err != nil {
		return Campaigns{}, err
	}
	return newCampaign, nil
}

func (s *services) UpdateCampaign(campaignId int, inputData CreateCampaignInput) (Campaigns, error) {
	campaign, err := s.r.FindById(campaignId)
	if err != nil {
		return Campaigns{}, nil
	}

	if campaign.User.Id != inputData.User.Id {
		return Campaigns{}, errors.New("unauthorized action")
	}

	campaign.CampaignName = inputData.CampaignName
	campaign.ShortDescription = inputData.ShortDescription
	campaign.Description = inputData.Description
	campaign.GoalAmount = inputData.GoalAmount
	campaign.Perks = inputData.Perks

	updatedCampaign, err := s.r.Update(campaign)

	if err != nil {
		return Campaigns{}, err
	}

	return updatedCampaign, nil
}
