package controller

import (
	"Score_System/dao"
	"Score_System/middleware"
)

type Controller struct {
	Jwt *middleware.JWT
	dao *dao.DAO
}

func NewController(jwt *middleware.JWT, dao *dao.DAO) *Controller {
	return &Controller{Jwt: jwt, dao: dao}
}
