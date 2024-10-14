package campaigns

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Campaigns, error)
	FindByUserId(userId int) ([]Campaigns, error)
	FindById(campaignId int) (Campaigns, error)
	Save(campaign Campaigns) (Campaigns, error)
	Update(campaign Campaigns) (Campaigns, error)
	SaveImages(campaignImages CampaignImages) (CampaignImages, error)
	MarkAllImagesAsNonPrimary(campaignID int) (bool, error)
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

	err := r.db.Where("user_id = ?", userId).Preload("CampaignImage", "campaign_images.is_primary = true").Find(&campaigns).Error

	if err != nil {
		return []Campaigns{}, err
	}

	return campaigns, nil
}
func (r *repository) FindById(campaignId int) (Campaigns, error) {
	var campaign Campaigns
	err := r.db.Where("id = ?", &campaignId).Preload("User").Preload("CampaignImage").Find(&campaign).Error
	if err != nil {
		return Campaigns{}, err
	}
	return campaign, nil
}

func (r *repository) Save(campaign Campaigns) (Campaigns, error) {
	err := r.db.Create(&campaign).Error
	if err != nil {
		return Campaigns{}, err
	}
	return campaign, nil
}

func (r *repository) Update(campaign Campaigns) (Campaigns, error) {
	err := r.db.Save(&campaign).Error
	if err != nil {
		return Campaigns{}, err
	}
	return campaign, nil
}

func (r *repository) SaveImages(campaignImages CampaignImages) (CampaignImages, error) {
	err := r.db.Create(&campaignImages).Error
	if err != nil {
		return CampaignImages{}, err
	}
	return campaignImages, nil
}

func (r *repository) MarkAllImagesAsNonPrimary(campaignId int) (bool, error) {
	err := r.db.Model(&CampaignImages{}).Where("campaign_id = ?", campaignId).Update("is_primary", false).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
