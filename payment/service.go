package payment

import (
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"github.com/muhamadrizkiariffadillah/CrowdFunding-Golang-NuxtJS/users"
)

type Service interface {
	GetPaymentUrl(transaction Transaction, user users.Users) string
}
type service struct {
}

func PaymenService() *service {
	return &service{}
}
func (s *service) GetPaymentUrl(transaction Transaction, user users.Users) string {
	var c = snap.Client{}
	c.Env = midtrans.Sandbox
	c.ServerKey = ""

	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  "",
			GrossAmt: int64(transaction.Amount),
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: user.FullName,
			Email: user.Email,
		},
	}

	snapResp, _ := c.CreateTransaction(req)
	return snapResp.RedirectURL
}
