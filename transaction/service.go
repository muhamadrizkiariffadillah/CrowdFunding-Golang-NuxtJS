package transaction

import (
	"fmt"
	"strings"
	"time"

	"github.com/muhamadrizkiariffadillah/CrowdFunding-Golang-NuxtJS/campaigns"
	"github.com/muhamadrizkiariffadillah/CrowdFunding-Golang-NuxtJS/payment"
)

type Service interface {
	GetTransactionsByCampaignId(campaignId int) ([]Transactions, error)
	GetTransactionByUserId(userId int) ([]Transactions, error)
	CreateTrasaction(input CreateTransactionInput) (Transactions, error)
}

type service struct {
	repo               Repository
	campaignRepository campaigns.Repository
	paymentService     payment.Service
}

func TransactionsServices(repo Repository, campaignRepository campaigns.Repository, paymentService payment.Service) *service {
	return &service{repo, campaignRepository, paymentService}
}

func (s *service) GetTransactionsByCampaignId(campaignId int) ([]Transactions, error) {
	transactions, err := s.repo.GetByCampaignId(campaignId)
	if err != nil {
		return []Transactions{}, err
	}
	return transactions, nil
}
func (s *service) GetTransactionByUserId(userId int) ([]Transactions, error) {
	transactions, err := s.repo.GetByUserId(userId)
	if err != nil {
		return []Transactions{}, err
	}
	return transactions, nil
}
func (s *service) CreateTrasaction(input CreateTransactionInput) (Transactions, error) {
	transactions := Transactions{
		UserId:     input.User.Id,
		CampaignId: input.CampaignId,
		Status:     "Pending",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	// create unique
	campaign, err := s.campaignRepository.FindById(input.CampaignId)
	if err != nil {
		return Transactions{}, err
	}

	words := strings.Split(campaign.CampaignName, " ")

	var firstLatters []string

	for _, word := range words {
		firstLatter := strings.ToUpper(string(word[0]))
		firstLatters = append(firstLatters, firstLatter)
	}

	currentDate := transactions.CreatedAt.Format("060102")

	transactions.Code = fmt.Sprintf("%v-%v-%v", currentDate, firstLatters, input.User.Id)

	newTranscation, err := s.repo.Save(transactions)

	if err != nil {
		return Transactions{}, err
	}

	paymentTransaction := payment.Transaction{
		Id:     newTranscation.Id,
		Amount: int(newTranscation.Amount),
	}
	paymentUrl := s.paymentService.GetPaymentUrl(paymentTransaction, input.User)

	newTranscation.PaymentUrl = paymentUrl

	newTranscation, err = s.repo.Update(newTranscation)
	if err != nil {
		return Transactions{}, err
	}
	return newTranscation, nil

}
