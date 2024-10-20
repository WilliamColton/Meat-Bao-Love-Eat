package controller

import (
	"Score_System/model"
	"Score_System/utils"
	"github.com/gin-gonic/gin"
)

func (c *Controller) CreateItem(context *gin.Context) {
	var item model.Item
	err := context.BindJSON(&item)
	if err != nil {
		return
	}
	c.dao.CreateItem(&item)
	context.JSON(200, item)
}

func (c *Controller) GetItemByID(context *gin.Context) {
	idString := context.Param("id")
	id := utils.StringToUint(idString)
	item := c.dao.GetItemByID(id)
	context.JSON(200, item)
}

func (c *Controller) GetEvaluationsByItemID(context *gin.Context) {
	idString := context.Param("id")
	id := utils.StringToUint(idString)
	evaluations := c.dao.GetEvaluationsByItemID(id)
	context.JSON(200, evaluations)
}

func (c *Controller) GetScoreByItemID(context *gin.Context) {
	idString := context.Param("id")
	id := utils.StringToUint(idString)
	evaluations := c.dao.GetEvaluationsByItemID(id)

	n := 0
	var r int
	for i, v := range evaluations {
		if v.Recommended == true {
			n += 1
			r = n / (i + 1)
		}
	}
	context.JSON(200, r)
}

func (c *Controller) GetRandomItemsByLimit(context *gin.Context) {
	limitString := context.Query("limit")
	limit := utils.StringToInt(limitString)
	evaluations := c.dao.GetRandomItemsByLimit(limit)
	context.JSON(200, evaluations)
}
