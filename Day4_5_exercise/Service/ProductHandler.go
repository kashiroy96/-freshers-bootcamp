package Service

import (
	"Day4_5_exercise/Models"
	"Day4_5_exercise/Utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"strconv"
)

type ProductInterface interface {
	InsertProduct(c *gin.Context)
	GetProduct(c *gin.Context)
	GetProducts(c *gin.Context)
	DeleteProduct(c *gin.Context)
	UpdateProduct(c *gin.Context)

	InsertRetailer(c *gin.Context)
	GetRetailers(c *gin.Context)
	IsRetailerAuthenticated(c *gin.Context) (retails *Models.Retailer, err error)
}

type ProductHandler struct {
	repo utils.InterfaceRepo
}

func NewProduct(repo utils.InterfaceRepo) *ProductHandler {
	return &ProductHandler{repo}
}

// Retailer

func (ph ProductHandler) InsertRetailer(c *gin.Context) {

	var retailer Models.Retailer
	err := c.ShouldBindJSON(&retailer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	err1 := Models.RetailerValidation(&retailer)
	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err1.Error()})
		return
	}
	ph.repo.DoInsertRetailer(&retailer)
	c.JSON(http.StatusOK, &retailer)
}

func (ph ProductHandler) GetRetailers(c *gin.Context) {
	var retailers []Models.Retailer
	ph.repo.FindRetailers(&retailers)
	c.JSON(http.StatusOK, retailers)
}

// Product

func (ph ProductHandler) InsertProduct(c *gin.Context) {
	var product Models.Product
	err := c.ShouldBindJSON(&product)
	fmt.Println("product", product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	err1 := Models.ProductValidation(&product)
	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err1.Error()})
		return
	}

	var retailer Models.Retailer

	fmt.Println("hi", c.Param("retailerID"))
	err2 := ph.repo.IsRetailerPresent(strconv.Itoa(int(product.RetailerID)), &retailer)
	fmt.Println(retailer.Id)
	if err2 != nil {
		fmt.Println(err2)
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Unauthorized retailer"})
		return
	}

	ph.repo.InsertProduct(&product)
	c.JSON(http.StatusOK, gin.H{"Product Inserted Successfully": product})
}

func (ph ProductHandler) GetProducts(c *gin.Context) {
	var products []Models.Product
	ph.repo.FindProducts(&products)
	c.JSON(http.StatusOK, products)
}

func (ph ProductHandler) GetProduct(c *gin.Context) {
	var product Models.Product
	fmt.Println("hi", c.Param("id"))
	err1 := ph.repo.IsProductPresent(c.Param("id"), &product)

	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err1.Error()})
		return
	}
	c.JSON(http.StatusOK, product)
}

func (ph ProductHandler) DeleteProduct(c *gin.Context) {
	var product Models.Product
	err := ph.repo.IsProductPresent(c.Param("id"), &product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	err1 := ph.repo.DeleteProduct(&product)
	//good or bad ?
	if err1 < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Given product cannot be deleted because it doesn't exist"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Product is Deleted Successfully": &product})
}

func (ph ProductHandler) UpdateProduct(c *gin.Context) {
	var productNew Models.Product
	err := c.ShouldBindJSON(&productNew)
	//fmt.Println("product", productNew)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	var product Models.Product
	//fmt.Println("hek", c.Param("id"))
	err1 := ph.repo.IsProductPresent(c.Param("id"), &product)
	//fmt.Println("product", product)
	if err1 != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": err1.Error()})
		return
	}
	ph.repo.UpdateProduct(&productNew)
	c.JSON(http.StatusOK, gin.H{"Product Updated Successfully": &productNew})
}

func (ph ProductHandler) GetRetailerTransactions(c *gin.Context) {
	allOrders, err := ph.repo.FindRetailerAllOrders(c.Param("retailerID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Total Retailer Deliveries": allOrders})
}

func (ph ProductHandler) IsRetailerAuthenticated(c *gin.Context) (*Models.Retailer, error) {
	var retailer Models.Retailer
	err := ph.repo.IsRetailerPresent(c.Param("retailerID"), &retailer)
	if err != nil {
		var invalidRetailer Models.Retailer
		return &invalidRetailer, errors.New("retailer is not authenticated")
	}
	return &retailer, nil
}
