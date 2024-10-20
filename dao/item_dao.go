package dao

import "Score_System/model"

func (d *DAO) CreateItem(m *model.Item) {
	d.DB.Create(m)
}

func (d *DAO) GetItemByID(id uint) model.Item {
	var item model.Item
	d.DB.Where("id = ?", id).First(&item)
	return item
}

func (d *DAO) GetEvaluationsByItemID(id uint) []model.Evaluation {
	var evaluations []model.Evaluation
	d.DB.Where("item_id = ?", id).Find(&evaluations)
	return evaluations
}

func (d *DAO) GetRandomItemsByLimit(limit int) []model.Item {
	var items []model.Item
	d.DB.Order("RANDOM()").Limit(limit).Find(&items)
	return items
}
