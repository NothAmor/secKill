package mysqlrepo

import (
	"secKill/app/common"
	"secKill/app/entity"
)

type UserDBImpl struct{}

func (udb *UserDBImpl) GetUsersByName(name string) (users []entity.User, err error) {
	err = common.DB.Where("username = ?", name).Find(&users).Error
	return
}

func (udb *UserDBImpl) GetUsersByNameAndPassword(name, password string) {
	var users []entity.User
	common.DB.Where("username = ? AND password = ?", name, password).Find(&users)
}
func (udb *UserDBImpl) CreateUser(user *entity.User) (err error) {
	err = common.DB.Create(user).Error
	return
}
