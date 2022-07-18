package Models

import (
	"errors"
)

type Retailer struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type Product struct {
	Id         uint   `json:"id"`
	RetailerID uint   `json:"retailerID" gorm:"foreign_key"`
	Name       string `json:"name"`
	Price      uint   `json:"price"`
	Quantity   uint   `json:"quantity"`
}

var (
	errProductId       = errors.New("invalid ProductId")
	errRetailerId      = errors.New("invalid retailerId")
	errProductName     = errors.New("invalid product")
	errProductPrice    = errors.New("invalid price")
	errProductQuantity = errors.New("invalid quantity")
	ErrRetName         = errors.New("wrong retailer name ")
	ErrRetPassword     = errors.New("wrong password")
)

func ProductValidation(p *Product) error {

	switch {
	case p.Id < 0:
		return errProductId
	case p.Name == "":
		return errProductName
	case p.Price <= 0:
		return errProductPrice
	case p.Quantity <= 0:
		return errProductQuantity
	default:
		return nil
	}
}

func RetailerValidation(r *Retailer) error {
	switch {
	case r.Id < 0:
		return errRetailerId
	case r.Name == "":
		return ErrRetName
	case r.Password == "":
		return ErrRetPassword
	default:
		return nil
	}
}
