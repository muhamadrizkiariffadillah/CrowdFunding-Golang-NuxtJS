package campaigns

type CampaignFormat struct {
	CampaignName     string `json:"campaign_name"`
	ShortDescription string `json:"short_description"`
	Description      string `json:"description"`
	CurrentAmount    uint   `json:"current_amount"`
	GoalAmount       uint   `json:"goal_amount"`
	Slug             string `json:"slug"`
	ImageUrl         string `json:"image_url"`
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
