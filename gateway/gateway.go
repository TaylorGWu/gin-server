package gateway

import (
	"github.com/gin-gonic/gin"
	"gin-server/front"
	"fmt"
)

type Gateway struct {
	Fronts []front.FrontInterface
	router *gin.Engine
}

func (gateway *Gateway) AddFronts(front front.FrontInterface) {
	gateway.Fronts = append(gateway.Fronts, front)
}

func (gateway *Gateway) Run() {
	gateway.router = gin.New()
	fmt.Print(len(gateway.Fronts))
	for _, front := range gateway.Fronts {
		path := front.BindPath()
		method := front.BindMethod()

		gateway.router.Handle(method, path,  front.Serve)
	}

	gateway.router.Run("localhost:39999")
}
