package dao

import "Score_System/model"

func (d *DAO) CreateEvaluation(evaluation *model.Evaluation) {
	d.DB.Create(evaluation)
}

func (d *DAO) GetEvaluationByItemID(itemID string) []model.Evaluation {
	var evaluations []model.Evaluation
	d.DB.Where("item_id = ?", itemID).Find(&evaluations)
	return evaluations
}

func (d *DAO) GetEvaluationByUserID(id uint) []model.Evaluation {
	var evaluations []model.Evaluation
	d.DB.Where("user_id = ?", id).Find(&evaluations)
	return evaluations
}

func (d *DAO) GetEvaluations() []model.Evaluation {
	var evaluations []model.Evaluation
	d.DB.Find(&evaluations)
	return evaluations
}

func (d *DAO) GetEvaluationByID(id uint) model.Evaluation {
	var evaluation model.Evaluation
	d.DB.Where("id = ?", id).First(&evaluation)
	return evaluation
}
