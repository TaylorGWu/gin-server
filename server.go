package main

import (
	"gin-server/gateway"
	"gin-server/front"
	"gin-server/protocol"
	"github.com/gin-gonic/gin"
)

type Front1 struct {

}

func (front Front1) BindMethod() (string) {
	return protocol.HTTP_GET_METHOD
}

func (front Front1) BindPath() (string) {
	return "/test1"
}

func (front Front1) Serve(context *gin.Context) {

}

func main() {
	fronts := make([]front.FrontInterface, 1)
	front1 := Front1{}

	myGateway := gateway.Gateway{
		Fronts: fronts,
	}
	myGateway.AddFronts(front1)
	myGateway.Run()
}
