package usecase

import "github.com/outk/rewardingsites/src/userservice/src/domain"

type UserInteractor struct {
	UserRepository UserRepository
}

func (interactor *UserInteractor) AddUser(user domain.User) (err error) {
	err = interactor.UserRepository.Add(user)
	return
}
