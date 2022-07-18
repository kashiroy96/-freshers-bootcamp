package utils

import (
	"Day4_5_exercise/Config"
	"Day4_5_exercise/Models"
)

type InterfaceRepo interface {
	IsProductPresent(id string, product *Models.Product) error
	InsertProduct(product *Models.Product)
	FindProducts(products *[]Models.Product)
	UpdateProduct(product *Models.Product)
	DeleteProduct(product *Models.Product) uint

	IsRetailerPresent(id string, retailer *Models.Retailer) error
	DoInsertRetailer(retailer *Models.Retailer)
	FindRetailers(retailers *[]Models.Retailer)
	FindRetailerAllOrders(id string) (Orders *[]Models.Order, er error)

	IsCustomerPresent(id string, customer *Models.Customer) error
	InsertCustomer(customer *Models.Customer)
	DeleteCustomer(customer *Models.Customer) uint
	FindCustomers(customer *[]Models.Customer)

	IsOrderPresent(id string, order *Models.Order) error
	InsertOrder(Order *Models.Order)
	DeleteOrder(Order *Models.Order) error
	IsCustomerOrdered(id string, Orders *[]Models.Order) (er error)
	FindAllOrders(orders *[]Models.Order) error
}

type Repo struct {
}

func GetRepo() InterfaceRepo {
	return &Repo{}
}

//----------------Retailer--------------------------------------------------------------//

func (repo *Repo) IsRetailerPresent(id string, retailer *Models.Retailer) error {
	return Config.DB.Where("id = ?", id).First(retailer).Error
}

func (repo *Repo) DoInsertRetailer(retailer *Models.Retailer) {
	Config.DB.Create(retailer)
}

func (repo *Repo) FindRetailers(retailers *[]Models.Retailer) {
	Config.DB.Find(retailers)
}

//----------------Product--------------------------------------------------------------//

func (repo *Repo) IsProductPresent(id string, product *Models.Product) error {
	return Config.DB.Where("id = ?", id).Find(product).Error
}

func (repo *Repo) InsertProduct(product *Models.Product) {
	Config.DB.Create(product)
}

func (repo *Repo) FindProducts(products *[]Models.Product) {
	Config.DB.Find(products)
}

func (repo *Repo) UpdateProduct(product *Models.Product) {
	Config.DB.Save(product)
}

func (repo *Repo) DeleteProduct(product *Models.Product) uint {
	return uint(Config.DB.Delete(product).RowsAffected)
}

//----------------Customer--------------------------------------------------------------//

func (repo *Repo) IsCustomerPresent(id string, customer *Models.Customer) error {
	return Config.DB.Where("id= ?", id).First(customer).Error
}

func (repo *Repo) InsertCustomer(customer *Models.Customer) {
	Config.DB.Create(customer)
}

func (repo *Repo) FindCustomers(customers *[]Models.Customer) {
	Config.DB.Find(customers)
}

func (repo *Repo) DeleteCustomer(customer *Models.Customer) uint {
	return uint(Config.DB.Delete(customer).RowsAffected)
}

//----------------Order--------------------------------------------------------------//

func (repo *Repo) InsertOrder(Order *Models.Order) {
	Config.DB.Create(Order)
}

func (repo *Repo) DeleteOrder(Order *Models.Order) error {
	return Config.DB.Delete(Order).Error
}

func (repo *Repo) IsOrderPresent(id string, order *Models.Order) error {
	return Config.DB.Where("id= ?", id).First(order).Error
}

func (repo *Repo) FindAllOrders(orders *[]Models.Order) error {
	return Config.DB.Find(orders).Error
}

func (repo *Repo) IsCustomerOrdered(id string, Orders *[]Models.Order) (er error) {
	return Config.DB.Where("customer_id = ?", id).Find(Orders).Error
}

func (repo *Repo) FindRetailerAllOrders(id string) (Orders *[]Models.Order, er error) {
	var retailerProducts []Models.Product
	Config.DB.Where("retailer_id = ?", id).Find(&retailerProducts)

	var orders []Models.Order
	for i := 0; i < len(retailerProducts); i++ {
		var tmpOrder []Models.Order
		err := Config.DB.Where("product_id = ?", retailerProducts[i].Id).Find(&tmpOrder).Error
		for j := 0; j < len(tmpOrder); j++ {
			orders = append(orders, tmpOrder[j])
		}
		if err != nil {
			var tmpOrders []Models.Order
			return &tmpOrders, err
		}
	}
	return &orders, nil
}
