// this is primary port
package core

import "errors"

type OrderService interface {
	CreateOrder(order Order) error
}

//this is business logic implementation of the OrderService
type OrderServiceImpl struct {
	repo OrderRepository // OrderRepository is the secondary port
}

func NewOrderService(repo OrderRepository) OrderService {
	return &OrderServiceImpl{repo: repo} // return an instance of OrderServiceImpl
}

func (s *OrderServiceImpl) CreateOrder(order Order) error {
	// business logic for creating an order
	if order.Total <= 0 {
		return errors.New("order total must be greater than zero") // validate order total
	}
	
	if err := s.repo.Save(order); err != nil {
		return err // save order using the repository
	}
	return nil // return nil if order is created successfully
}