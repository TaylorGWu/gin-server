package front

import (
	"github.com/gin-gonic/gin"
)

/*
 * this is the FrontInterface which could be implemented by developer
 * developer can designed new struct and new struct can implement those two method as well as others
 * Handle Method is the key method which include user defines method
 */

type FrontInterface interface {
	BindMethod() string
	BindPath() string
	Serve(context *gin.Context)
}