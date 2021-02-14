package interface

type UserService interface {
    AddUser(context.Context, *AddUserRequest) (*AddUserReply, error)
}