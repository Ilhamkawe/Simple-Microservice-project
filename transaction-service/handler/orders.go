package handler

import "transaction-service/orders"

type orderHandler struct {
	orderService orders.Service
}

func NewOrderHandler(orderService orders.Service) *orderHandler {
	return &orderHandler{orderService}
}
