package controller

import (
	"Score_System/model"
	"Score_System/utils"
	"github.com/gin-gonic/gin"
)

func (c *Controller) GetEvaluations(context *gin.Context) {
	evaluations := c.dao.GetEvaluations()
	context.JSON(200, evaluations)
}

func (c *Controller) CreateEvaluation(context *gin.Context) {
	var evaluation model.Evaluation
	err := context.BindJSON(&evaluation)
	if err != nil {
		return
	}
	c.dao.CreateEvaluation(&evaluation)
	context.JSON(200, evaluation)
}

func (c *Controller) GetEvaluationByID(context *gin.Context) {
	idString := context.Param("id")
	id := utils.StringToUint(idString)
	evaluation := c.dao.GetEvaluationByID(id)
	context.JSON(200, evaluation)
}
