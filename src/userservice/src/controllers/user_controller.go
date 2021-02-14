package controllers

import (
	"log"

	"github.com/outk/rewardingsites/src/userservice/src/domain"
	"github.com/outk/rewardingsites/src/userservice/src/interfaces/db"
	"github.com/outk/rewardingsites/src/userservice/src/usecase"
)

type UserController struct {
	Interactor usecase.UserInteractor
}

func NewUserController(mongoDBHandler db.MongoDBHandler) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			UserRepository: &db.UserRepository{
				MongoDBHandler: mongoDBHandler,
			},
		},
	}
}

func (controller *UserController) AddUser(user domain.User) (err error) {
	err = controller.Interactor.AddUser(user)
	if err != nil {
		log.Fatalln(err)
	}
	return
}
