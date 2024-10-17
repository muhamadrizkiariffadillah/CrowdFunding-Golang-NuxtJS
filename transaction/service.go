package transaction

type Service interface {
	GetTransactionsByCampaignId(campaignId int) ([]Transactions, error)
	GetTransactionByUserId(userId int) ([]Transactions, error)
}

type service struct {
	repo Repository
}

func TransactionsServices(repo Repository) *service {
	return &service{repo}
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
