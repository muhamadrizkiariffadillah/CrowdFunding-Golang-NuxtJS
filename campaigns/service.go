package campaigns

type Service interface {
	GetCampaigns(userId int) ([]Campaigns, error)
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
