package orders

import (
	"transaction-service/helper"
	"transaction-service/payments"
)

type Service interface{}

type service struct {
	repository Repository
	paymentService payments.Service
}

func NewService(repository Repository, paymentRepository payments.Service) *service {
	return &service{repository, paymentRepository}
}

func (s *service) Create(input CreateOrderInput) (Orders, error) {
	order := Orders{}

	order.Status = input.Status
	order.UserID = input.UserID
	order.CourseID = input.CourseID

	newOrder, err := s.repository.Create(order)

	if err != nil {
		return newOrder, err
	}

	paymentOrder := payments.Order{
		ID: newOrder.ID,
		Amount: input.Amount,
	}

	user, err := helper.GetUserByID(newOrder.UserID) 

	if err != nil {
		return newOrder,err 
	}

	paymentUrl , err := s.paymentService.GetPaymentUrl(paymentOrder, payments.User{Email: user.Data.Email, Name: user.Data.Name})

	newOrder.SnapUrl = paymentUrl

	newOrder, err = s.repository.Update(newOrder)

	if err != nil {
		return newOrder, err 
	}

	return newOrder, nil

}

func (s *service) _()