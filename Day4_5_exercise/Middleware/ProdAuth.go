package Middleware

import (
	"Day4_5_exercise/Service"
	utils "Day4_5_exercise/Utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func BasicAuthRetailer(c *gin.Context) {
	// Get the Basic Authentication credentials
	userName, password, hasAuth := c.Request.BasicAuth()

	res, err := Service.NewProduct(utils.GetRepo()).IsRetailerAuthenticated(c)
	if hasAuth && err == nil && res.Name == userName && res.Password == password {
		c.Writer.Header().Set(userName+" with "+password, "is authenticated")
	} else {
		c.Writer.Header().Set("Authentication-Info", "Basic realm=Restricted")
		c.JSON(http.StatusBadRequest, gin.H{"Error": "please provide correct credentials first"})
		c.Abort()
		return
	}
}
