package web_handler

/*
 * this is the FrontInterface which could be implemented by developer
 * developer can designed new struct and new struct can implement those two method
 */

type FrontInterface interface {
	BindMethod() string
	BindPath() string
}

func (front *FrontInterface) BindMethod() {
	return "GET"
}
