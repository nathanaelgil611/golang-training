package services

import (
	"sesi-4/database"
	"sesi-4/model"
)

// var orderList = make([]model.Order)

type OrderServiceIface interface {
	Create(order model.Order)
	GetAll() []model.Order
	GetByID(orderID int) model.Order
	Delete(orderID int)
	Update(Order *model.Order, orderID int)
}

type OrderSvc struct {
	ListOrder []model.Order
}

func NewOrderService() OrderServiceIface {
	var list []model.Order
	return &OrderSvc{list}
}

func (u *OrderSvc) Create(order model.Order) {
	// u.ListOrder[order.OrderID] = *order
	err := database.CreateOrder(order)
	if err != nil {
		panic(err)
	}

	// fmt.Println(u)
}

func (u *OrderSvc) GetByID(orderID int) model.Order {
	order, err := database.GetOrder(orderID)
	if err != nil {
		panic(err)
	}
	return order
}

func (u *OrderSvc) GetAll() []model.Order {
	listOrder, err := database.GetAllOrder()
	if err != nil {
		panic(err)
	}
	return listOrder
}

func (u *OrderSvc) Delete(orderID int) {
	// delete(u.ListOrder, orderID)
	err := database.DeleteOrder(orderID)
	if err != nil {
		panic(err)
	}

}

func (u *OrderSvc) Update(order *model.Order, orderID int) {
	err := database.UpdateOrder(order, orderID)
	if err != nil {
		panic(err)
	}
}
