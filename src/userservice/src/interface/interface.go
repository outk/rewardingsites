package interface

import "github.com/outk/rewardingsites/src/userservice/src/repository"

type UserService interface {
    AddUser(context.Context, *repository.AddUserRequest) (*repository.AddUserReply, error)
}