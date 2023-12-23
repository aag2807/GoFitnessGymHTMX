package controller

import (
	"github.com/GoGym/src/domain/service"
	"github.com/GoGym/src/utils"
	lib "github.com/aag2807/triplex-to-go"
)

type DashboardController struct {
	Arguments lib.Arguments
	State     lib.State
	Renderer  *utils.PartialRenderer
}

func NewDashboardController() *LoginController {
	return &LoginController{
		Arguments:   lib.Arguments{},
		State:       lib.State{},
		Renderer:    utils.NewPartialRenderer(),
		userService: service.NewLoginService(),
	}
}
