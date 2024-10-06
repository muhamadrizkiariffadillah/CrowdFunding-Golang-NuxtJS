package campaigns

import "strings"

type CampaignFormat struct {
	CampaignName     string `json:"campaign_name"`
	ShortDescription string `json:"short_description"`
	Description      string `json:"description"`
	CurrentAmount    uint   `json:"current_amount"`
	GoalAmount       uint   `json:"goal_amount"`
	Slug             string `json:"slug"`
	ImageUrl         string `json:"image_url"`
}

type CampaignDetailFormat struct {
	CampaignName     string                `json:"campaign_name"`
	ShortDescription string                `json:"short_description"`
	Description      string                `json:"description"`
	CurrentAmount    uint                  `json:"current_amount"`
	GoalAmount       uint                  `json:"goal_amount"`
	Slug             string                `json:"slug"`
	ImageUrl         string                `json:"image_url"`
	Perks            []string              `json:"perks"`
	User             CampaignUserFormat    `json:"user"`
	CampaignImages   []CampaignImageFormat `json:"images"`
}

type CampaignUserFormat struct {
	FullName string `json:"full_name"`
	ImageUrl string `json:"image_url"`
}

type CampaignImageFormat struct {
	FileName  string `json:"file"`
	IsPrimary bool   `json:"is_primary"`
}

func CampaignFormatter(campaign Campaigns) CampaignFormat {
	formatCampaign := CampaignFormat{
		CampaignName:     campaign.CampaignName,
		ShortDescription: campaign.ShortDescription,
		Description:      campaign.Description,
		CurrentAmount:    campaign.CurrentAmount,
		GoalAmount:       campaign.GoalAmount,
		Slug:             campaign.Slug,
		ImageUrl:         "",
	}

	if len(campaign.CampaignImage) > 0 {
		formatCampaign.ImageUrl = campaign.CampaignImage[0].FileName
	}
	return formatCampaign

}

func CampaignsFormatter(campaigns []Campaigns) []CampaignFormat {
	if len(campaigns) == 0 {
		return []CampaignFormat{}
	}

	var formatCampaigns []CampaignFormat

	for _, campaign := range campaigns {
		formatcampaign := CampaignFormatter(campaign)
		formatCampaigns = append(formatCampaigns, formatcampaign)
	}

	return formatCampaigns
}

func GetCampaignDetailFormatter(campaign Campaigns) CampaignDetailFormat {
	format := CampaignDetailFormat{
		CampaignName:     campaign.CampaignName,
		ShortDescription: campaign.ShortDescription,
		Description:      campaign.Description,
		CurrentAmount:    campaign.CurrentAmount,
		GoalAmount:       campaign.GoalAmount,
		ImageUrl:         "",
		Slug:             campaign.Slug,
		User: CampaignUserFormat{
			FullName: campaign.User.FullName,
			ImageUrl: campaign.User.AvatarFileName,
		},
	}
	if len(campaign.CampaignImage) > 0 {
		format.ImageUrl = campaign.CampaignImage[0].FileName
	}

	var perks []string

	for _, perk := range strings.Split(campaign.Perks, ",") {
		perks = append(perks, strings.TrimSpace(perk))
	}

	format.Perks = perks

	var images []CampaignImageFormat

	for _, image := range campaign.CampaignImage {
		campaignFormatter := CampaignImageFormat{
			FileName:  image.FileName,
			IsPrimary: image.IsPrimary,
		}
		images = append(images, campaignFormatter)
	}

	format.CampaignImages = images

	return format
}
