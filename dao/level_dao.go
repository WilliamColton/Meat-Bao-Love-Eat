package dao

import "Score_System/model"

func (d *DAO) CreateLevel(level *model.Level) {
	d.DB.Create(level)
}

func (d *DAO) GetLevels() []model.Level {
	var levels []model.Level
	d.DB.Find(&levels)
	return levels
}

func (d *DAO) GetLevelByID(id uint) model.Level {
	var level model.Level
	d.DB.Where("id = ?", id).First(&level)
	return level
}

func (d *DAO) GetItemsByLevelID(id uint) []model.Item {
	var items []model.Item
	d.DB.Where("level_id = ?", id).Find(&items)
	return items
}
