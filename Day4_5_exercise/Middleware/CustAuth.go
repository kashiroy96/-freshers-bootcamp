package Middleware

import (
	"Day4_5_exercise/Service"
	utils "Day4_5_exercise/Utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func BasicAuthCustomer(c *gin.Context) {

	userName, password, hasAuth := c.Request.BasicAuth()
	res, err := Service.NewCustomer(utils.GetRepo()).IsCustomerAuthenticated(c)
	if hasAuth && err == nil && res.Name == userName && res.Password == password {
		c.Writer.Header().Set(userName+" with "+password, "is authenticated")
	} else {
		c.Writer.Header().Set("Authentication-Info", "Basic realm=Restricted")
		c.JSON(http.StatusBadRequest, gin.H{"Error": "please provide correct credentials first"})
		c.Abort()
		return
	}
}
