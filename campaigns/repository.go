package campaigns

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Campaigns, error)
	FindByUserId(userId int) ([]Campaigns, error)
}

type repository struct {
	db *gorm.DB
}

func CampaignRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Campaigns, error) {
	var campaigns []Campaigns
	err := r.db.Preload("CampaignImages", "campaign_images.is_primary = true").Find(&campaigns).Error
	if err != nil {
		return []Campaigns{}, err
	}
	return campaigns, nil
}

func (r *repository) FindByUserId(userId int) ([]Campaigns, error) {
	var campaigns []Campaigns

	err := r.db.Where("user_id = ?", userId).Preload("CampaignImages", "campaign_images.is_primary = true").Find(&campaigns).Error

	if err != nil {
		return []Campaigns{}, err
	}

	return campaigns, nil
}
