package mysqlrepo

import (
	"secKill/app/common"
	"secKill/app/entity"
)

func GetUsersByName(name string) (users []entity.User, err error) {
	common.DB.Where("username = ?", name).Find(&users)
	return
}

func GetUsersByNameAndPassword(name, password string) (users []entity.User, err error) {
	common.DB.Where("username = ? and password = ?", name, password).Find(&users)
	return
}

func CreateUser(user *entity.User) (err error) {
	err = common.DB.Create(user).Error
	return
}
