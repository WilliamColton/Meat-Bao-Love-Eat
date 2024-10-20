package controller

import (
	"Score_System/model"
	"Score_System/utils"
	"github.com/gin-gonic/gin"
)

func (c *Controller) CreateLevel(context *gin.Context) {
	var level model.Level
	err := context.BindJSON(&level)
	if err != nil {
		return
	}
	c.dao.CreateLevel(&level)
	context.JSON(200, level)
}

func (c *Controller) GetLevels(context *gin.Context) {
	levels := c.dao.GetLevels()
	context.JSON(200, levels)
}

func (c *Controller) GetLevelByID(context *gin.Context) {
	idString := context.Param("id")
	id := utils.StringToUint(idString)
	level := c.dao.GetLevelByID(id)
	context.JSON(200, level)
}

func (c *Controller) GetItemsByLevelID(context *gin.Context) {
	idString := context.Param("id")
	id := utils.StringToUint(idString)
	items := c.dao.GetItemsByLevelID(id)
	context.JSON(200, items)
}
