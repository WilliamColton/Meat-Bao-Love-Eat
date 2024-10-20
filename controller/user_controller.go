package controller

import (
	"Score_System/middleware"
	"Score_System/model"
	"Score_System/utils"
	"github.com/gin-gonic/gin"
)

func (c *Controller) UserRegister(context *gin.Context) {
	var user model.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		return
	}
	c.dao.CreateUser(user)
	context.JSON(200, "register success")
}

func (c *Controller) UserLogin(context *gin.Context) {
	var user model.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		return
	}
	password := c.dao.GetPasswordByUserName(user.UserName)
	if password == user.Password {
		tokenString, _ := c.Jwt.CreateToken(middleware.NewClaims(user.UserName))
		context.JSON(200, gin.H{
			"token": tokenString,
		})
	} else {
		context.JSON(200, "login fail")
	}
}

func (c *Controller) GetEvaluationsByUserID(context *gin.Context) {
	idString := context.Param("id")
	id := utils.StringToUint(idString)
	evaluations := c.dao.GetEvaluationByUserID(id)
	context.JSON(200, evaluations)
}
