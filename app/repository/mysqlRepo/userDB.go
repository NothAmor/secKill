package mysqlrepo

import (
	"secKill/app/entity"
)

type userDB interface {
	GetUsersByName(name string) (users []entity.User, err error)

	GetUsersByNameAndPassword(name, password string)

	CreateUser(user *entity.User) (err error)
}
