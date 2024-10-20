package dao

import "Score_System/model"

func (d *DAO) CreateUser(user model.User) {
	d.DB.Create(&user)
}

func (d *DAO) GetPasswordByUserName(userName string) string {
	var user model.User
	d.DB.Where("user_name = ?", userName).First(&user)
	return user.Password
}
