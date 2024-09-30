package campaign

import "gorm.io/gorm"

type Repository interface {
	Save(campaign Campaigns) (Campaigns, error)
}

type repository struct {
	db *gorm.DB
}

func CampaignRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(camp Campaigns) (Campaigns, error) {
	err := r.db.Save(&camp).Error
	if err != nil {
		return Campaigns{}, err
	}
	return camp, nil
}
