package transaction

import (
	"gorm.io/gorm"
)

type Repository interface {
	GetByCampaignId(campaignId int) ([]Transactions, error)
	GetByUserId(userId int) ([]Transactions, error)
	Save(transaction Transactions) (Transactions, error)
	Update(transaction Transactions) (Transactions, error)
}

type repository struct {
	db *gorm.DB
}

func TransactionsRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetByCampaignId(campaignId int) ([]Transactions, error) {
	var transactions []Transactions

	err := r.db.Preload("User").Where("campaign_id = ?", campaignId).Order("id desc").Find(&transactions).Error

	if err != nil {
		return []Transactions{}, err
	}
	return transactions, nil
}
func (r *repository) GetByUserId(userId int) ([]Transactions, error) {
	var transactions []Transactions
	err := r.db.Preload("Campaigns.CampaignImage").Where("user_id = ?", userId).Find(&transactions).Error
	if err != nil {
		return []Transactions{}, err
	}
	return transactions, nil
}
func (r *repository) Save(transaction Transactions) (Transactions, error) {
	err := r.db.Create(&transaction).Error
	if err != nil {
		return Transactions{}, err
	}
	return transaction, nil
}

func (r *repository) Update(transaction Transactions) (Transactions, error) {
	err := r.db.Save(&transaction).Error
	if err != nil {
		return Transactions{}, err
	}
	return transaction, nil
}
