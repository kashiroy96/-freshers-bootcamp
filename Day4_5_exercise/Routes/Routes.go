package Routes

import (
	"Day4_5_exercise/Middleware"
	"Day4_5_exercise/Service"
	"Day4_5_exercise/Utils"
	"github.com/gin-gonic/gin"
)

// SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {

	Customer := Service.NewCustomer(utils.GetRepo())
	Product := Service.NewProduct(utils.GetRepo())

	r := gin.Default()
	grp1 := r.Group("retailer-api")
	{
		grp1.POST("retailer", Product.InsertRetailer)
		grp1.GET("retailers", Product.GetRetailers)

		grp1.POST("product/:retailerID", Middleware.BasicAuthRetailer, Product.InsertProduct)
		grp1.GET("product/:retailerID/:id", Middleware.BasicAuthRetailer, Product.GetProduct)
		grp1.GET("products", Product.GetProducts)
		grp1.DELETE("product/:retailerID/:id", Middleware.BasicAuthRetailer, Product.DeleteProduct)
		grp1.PUT("product/:retailerID/:id", Middleware.BasicAuthRetailer, Product.UpdateProduct)

		grp1.POST("customer", Customer.InsertCustomer)
		grp1.GET("customers", Middleware.BasicAuthCustomer, Customer.GetCustomers)
		grp1.DELETE("customer/:id", Middleware.BasicAuthCustomer, Customer.DeleteCustomer)
		grp1.POST("order/:customerID", Middleware.BasicAuthCustomer, Customer.InsertOrder)
		grp1.GET("orders", Customer.GetAllOrders)
		grp1.GET("orders/:customerID", Middleware.BasicAuthCustomer, Customer.GetCustomerOrders)
		grp1.GET("transactions/:retailerID", Middleware.BasicAuthRetailer, Product.GetRetailerTransactions)

	}
	return r
}
