package payments

import (
	"strconv"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type service struct{

}

type Service interface{
	GetPaymentUrl(transaction Order, user User) (string, error)
}

func NewService() *service {
	return &service{}
}


func (s *service) GetPaymentUrl(transaction Order, user User) (string, error) {
	midtrans.ServerKey = ""
	midtrans.ClientKey = ""
	midtrans.Environment = midtrans.Sandbox
	snapReq := &snap.Request{
		TransactionDetails : midtrans.TransactionDetails{
			OrderID : strconv.Itoa(transaction.ID), 
			GrossAmt: int64(transaction.Amount),
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: user.Name,
			Email: user.Email,
		},
		CreditCard : &snap.CreditCardDetails{
			Secure : true,
		}, 
	}

	snapResp,err := snap.CreateTransactionUrl(snapReq)
	return snapResp, err
}