package services

import (
	"context"
	"log"

	"github.com/outk/rewardingsites/src/userservice/src/controllers"
	"github.com/outk/rewardingsites/src/userservice/src/domain"
	"github.com/outk/rewardingsites/src/userservice/src/infrastructure"
	pb "github.com/outk/rewardingsites/src/userservice/src/pb"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func (s *UserService) AddUser(ctx context.Context, req *pb.AddUserRequest) (rep *pb.AddUserReply, err error) {
	// TODO: Before add user to database, must check the user is unique.
	user := domain.User{
		ID:         req.User.Id,
		FirstName:  req.User.FirstName,
		MiddleName: req.User.MiddleName,
		LastName:   req.User.LastName,
		Gender:     req.User.Gender,
		Birthday: domain.Birthday{
			Year:  req.User.Birthday.Year,
			Month: req.User.Birthday.Month,
			Day:   req.User.Birthday.Day,
		},
		Email:    req.User.Email,
		Password: req.User.Password,
	}

	userController := controllers.NewUserController(infrastructure.NewMongoDBHandler())

	err = userController.AddUser(user)
	if err != nil {
		log.Fatalln(err)
	}

	return &pb.AddUserReply{
		Successed: true,
	}, nil
}
