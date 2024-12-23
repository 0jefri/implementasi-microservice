package services

import (
	"log"
	"tiket-online/models"
	"tiket-online/repository"
)

type OrderService struct {
	repo *repository.OrderRepository
}

func NewOrderService(repo *repository.OrderRepository) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) CreateOrder(order models.Order) error {
	order.TotalPrice = float64(order.Qty) * order.Price
	log.Printf("[DEBUG] Calculated TotalPrice: %f for Order: %+v\n", order.TotalPrice, order)
	return s.repo.Create(order)
}

func (s *OrderService) GetOrderDetail(id string) (models.Order, error) {
	return s.repo.GetByID(id)
}

func (s *OrderService) UpdateOrder(id string, updatedOrder models.Order) error {
	updatedOrder.TotalPrice = float64(updatedOrder.Qty) * updatedOrder.Price
	return s.repo.Update(id, updatedOrder)
}
