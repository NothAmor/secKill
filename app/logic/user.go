package logic

import (
	"context"
	"errors"
	"secKill/app/entity"
	"secKill/app/proto"
	mysqlrepo "secKill/app/repository/mysqlRepo"
	"secKill/utils/logger"
)

func Register(ctx context.Context, args proto.CreateUserReq, reply *proto.CreateUserResp) (err error) {
	tag := "logic.user.register"

	users, err := mysqlrepo.GetUsersByName(args.Name)
	if err != nil {
		logger.Ex(ctx, tag, "getUsersByName failed. name:[%+v], err:[%+v]", args.Name, err)
		return
	}
	if len(users) > 0 {
		err = errors.New("用户名已存在")
		logger.Ix(ctx, tag, "user already exist. name:[%+v], users:[%+v]", args.Name, users)
		return
	}

	err = mysqlrepo.CreateUser(&entity.User{
		Name:     args.Name,
		Password: args.Password,
	})
	if err != nil {
		logger.Ex(ctx, tag, "createUser failed. name:[%+v], err:[%+v]", args.Name, err)
		return
	}

	return
}
