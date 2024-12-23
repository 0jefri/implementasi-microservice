package repository

import (
	"errors"
	"log"
	"tiket-online/models"
)

type OrderRepository struct {
	orders map[string]models.Order
}

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{
		orders: make(map[string]models.Order),
	}
}

func (r *OrderRepository) Create(order models.Order) error {
	if _, exists := r.orders[order.ID]; exists {
		return errors.New("order already exists")
	}
	log.Printf("[DEBUG] Storing Order: %+v\n", order)
	r.orders[order.ID] = order
	return nil
}

func (r *OrderRepository) GetByID(id string) (models.Order, error) {
	order, exists := r.orders[id]
	if !exists {
		return models.Order{}, errors.New("order not found")
	}
	return order, nil
}

func (r *OrderRepository) Update(id string, updatedOrder models.Order) error {
	if _, exists := r.orders[id]; !exists {
		return errors.New("order not found")
	}
	r.orders[id] = updatedOrder
	return nil
}
