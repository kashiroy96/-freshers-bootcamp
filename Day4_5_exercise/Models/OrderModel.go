package Models

import (
	"errors"
)

type Customer struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type Order struct {
	Id         uint   `json:"id"`
	CustomerID uint   `json:"customerID" gorm:"foreign_key"`
	ProductID  uint   `json:"productID"`
	Quantity   uint   `json:"quantity"`
	Message    string `json:"message"`
}

var (
	errOrderId      = errors.New("invalid orderId")
	errCustomerId   = errors.New("invalid retailerId")
	errMessage      = errors.New("invalid Message")
	ErrCustName     = errors.New("wrong customer name ")
	ErrCustPassword = errors.New("wrong password")
)

func OrderValidation(order *Order) error {

	switch {
	case order.Id < 0:
		return errOrderId
	case order.CustomerID < 0:
		return errCustomerId
	case order.ProductID < 0:
		return errProductId
	case order.Quantity <= 0:
		return errProductQuantity
	case order.Message == "":
		return errMessage
	default:
		return nil
	}
}

func CustomerValidation(c *Customer) error {

	switch {
	case c.Id < 0:
		return errCustomerId
	case c.Name == "":
		return ErrCustName
	case c.Password == "":
		return ErrCustPassword
	default:
		return nil
	}
}
