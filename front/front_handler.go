package front

import (
	"github.com/gin-gonic/gin"
)

type FrontHandler struct{
	front FrontInterface
}

func (frontHandler *FrontHandler) Handle (context *gin.Context) {
	frontHandler.front.Serve(context)
}

