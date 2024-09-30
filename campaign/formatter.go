package campaign

type CampaignFormat struct {
	CampaignName     string `json:"campaign_name"`
	ShortDescription string `json:"short_description"`
	Description      string `json:"description"`
	GoalAmount       uint   `json:"goal_amount"`
	Slug             string `json:"slug"`
	ImageUrl         string `json:"image_url"`
}

func CreateCampignFormatter(campaign Campaigns) CampaignFormat {
	campaignFormat := CampaignFormat{
		CampaignName:     campaign.CampaignName,
		ShortDescription: campaign.ShortDescription,
		Description:      campaign.Description,
		GoalAmount:       campaign.GoalAmount,
		Slug:             campaign.Slug,
		ImageUrl:         "",
	}
	return campaignFormat

}
