package Service

import (
	"Day4_5_exercise/Models"
	"Day4_5_exercise/Utils"
	"errors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"strconv"
)

type CustomerInterface interface {
	InsertCustomer(c *gin.Context)
	GetCustomers(c *gin.Context)
	DeleteCustomer(c *gin.Context)
	InsertOrder(c *gin.Context)
	GetAllOrders(c *gin.Context)
	GetCustomerOrders(c *gin.Context)
	IsCustomerAuthenticated(c *gin.Context) (customer *Models.Customer, err error)
}

type CustomerHandler struct {
	repo utils.InterfaceRepo
}

func NewCustomer(repo utils.InterfaceRepo) *CustomerHandler {
	return &CustomerHandler{repo}
}

func (ch *CustomerHandler) InsertCustomer(c *gin.Context) {
	var customer Models.Customer

	err := c.ShouldBindJSON(&customer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	err1 := Models.CustomerValidation(&customer)
	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err1.Error()})
		return
	}
	ch.repo.InsertCustomer(&customer)
	c.JSON(http.StatusOK, gin.H{"New Customer inserted": &customer})
}

func (ch CustomerHandler) GetCustomers(c *gin.Context) {
	var customers []Models.Customer
	ch.repo.FindCustomers(&customers)
	c.JSON(http.StatusOK, &customers)
}

func (ch *CustomerHandler) DeleteCustomer(c *gin.Context) {
	var customer Models.Customer
	err := ch.repo.IsCustomerPresent(c.Param("id"), &customer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	err1 := ch.repo.DeleteCustomer(&customer)
	//good or bad ?
	if err1 < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Given product cannot be deleted because it doesn't exist"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Product is Deleted Successfully": &customer})
}

func (ch *CustomerHandler) InsertOrder(c *gin.Context) {
	var order Models.Order
	err := c.ShouldBindJSON(&order)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	err1 := Models.OrderValidation(&order)
	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err1.Error()})
		return
	}

	var customer Models.Customer

	err2 := ch.repo.IsCustomerPresent(strconv.Itoa(int(order.CustomerID)), &customer)

	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid Customer"})
		return
	}

	ch.repo.InsertOrder(&order)
	c.JSON(http.StatusOK, gin.H{"Order placed Successfully": &order})
}

func (ch *CustomerHandler) GetAllOrders(c *gin.Context) {
	var order []Models.Order
	ch.repo.FindAllOrders(&order)
	c.JSON(http.StatusOK, order)
}

func (ch *CustomerHandler) GetCustomerOrders(c *gin.Context) {
	var customerExist Models.Customer
	err := ch.repo.IsCustomerPresent(c.Param("customerID"), &customerExist)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid Customer"})
		return
	}
	var customerOrders []Models.Order
	err1 := ch.repo.IsCustomerOrdered(c.Param("customerID"), &customerOrders)
	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Customer have not any orders"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Customer orders": customerOrders})
}

func (ch *CustomerHandler) IsCustomerAuthenticated(c *gin.Context) (*Models.Customer, error) {
	var customer Models.Customer
	err := ch.repo.IsCustomerPresent(c.Param("customerID"), &customer)

	if err != nil {
		var invalidCustomer Models.Customer
		return &invalidCustomer, errors.New("customer is not authenticated")
	}
	return &customer, nil
}
