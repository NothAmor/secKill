package user

import (
	"context"
	"errors"
	"secKill/app/entity"
	proto "secKill/app/proto/user"
	mysqlRepo "secKill/app/repository/mysqlRepo"
	"secKill/utils/logger"
)

func Register(ctx context.Context, args proto.CreateUserReq, reply *proto.CreateUserResp) (err error) {
	tag := "logic.user.register"

	users, err := mysqlRepo.GetUsersByName(args.Name)
	if err != nil {
		logger.Ex(ctx, tag, "getUsersByName failed. name:[%+v], err:[%+v]", args.Name, err)
		return
	}
	if len(users) > 0 {
		err = errors.New("用户名已存在")
		logger.Ix(ctx, tag, "user already exist. name:[%+v], users:[%+v]", args.Name, users)
		return
	}

	err = mysqlRepo.CreateUser(&entity.User{
		Name:     args.Name,
		Password: args.Password,
	})
	if err != nil {
		logger.Ex(ctx, tag, "createUser failed. name:[%+v], err:[%+v]", args.Name, err)
		return
	}

	return
}

func Login(ctx context.Context, args proto.UserLoginReq, reply *proto.UserLoginResp) (err error) {
	tag := "logic.user.login"

	users, err := mysqlRepo.GetUsersByNameAndPassword(args.Name, args.Password)
	if err != nil {
		logger.Ex(ctx, tag, "getUsersByNameAndPassword failed. name:[%+v], err:[%+v]", args.Name, err)
		return
	}
	if len(users) == 0 {
		err = errors.New("用户名与密码不匹配")
		logger.Ix(ctx, tag, "user and password not match. name:[%+v]", args.Name)
		return
	}

	reply.Token = users[0].Name
	return
}
