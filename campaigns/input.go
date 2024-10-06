package campaigns

type GetCampaignDetailInput struct {
	Id int `uri:id binding:"required"`
}
