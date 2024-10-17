package transaction

import (
	"time"
)

type CampaignTransactionFormat struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

type UserTransactionsFormat struct {
	Id        int            `json:"id"`
	Amount    int            `json:"amount"`
	Status    string         `json:"status"`
	CreatedAt time.Time      `json:"created_at"`
	Campaign  CampaignFormat `json:"campaign"`
}

type CampaignFormat struct {
	Name     string `json:"campaign_name"`
	ImageUrl string `json:"image"`
}

func CampaignTransactionFormatter(transaction Transactions) CampaignTransactionFormat {
	format := CampaignTransactionFormat{
		Id:        transaction.Id,
		Name:      transaction.User.FullName,
		Amount:    int(transaction.Amount),
		CreatedAt: transaction.CreatedAt,
	}
	return format
}

func CampaignTransactionsFormatter(transaction []Transactions) []CampaignTransactionFormat {
	if len(transaction) == 0 {
		return []CampaignTransactionFormat{}
	}

	var formats []CampaignTransactionFormat

	for _, format := range transaction {
		formatter := CampaignTransactionFormatter(format)
		formats = append(formats, formatter)
	}
	return formats
}

func UserTransactionFormatter(transaction Transactions) UserTransactionsFormat {
	format := UserTransactionsFormat{
		Id:        transaction.Id,
		Amount:    int(transaction.Amount),
		Status:    transaction.Status,
		CreatedAt: transaction.CreatedAt,
		Campaign: CampaignFormat{
			Name: transaction.Campaign.CampaignName,
		},
	}

	if len(transaction.Campaign.CampaignImage) > 0 {
		format.Campaign.ImageUrl = transaction.Campaign.CampaignImage[0].FileName
	}
	return format
}

func UserTrabsactionsFormatter(transaction []Transactions) []UserTransactionsFormat {
	if len(transaction) == 0 {
		return []UserTransactionsFormat{}
	}
	var formats []UserTransactionsFormat

	for _, format := range transaction {
		formatter := UserTransactionFormatter(format)
		formats = append(formats, formatter)
	}
	return formats
}
