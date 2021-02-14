package db

import "github.com/outk/rewardingsites/src/userservice/src/domain"

type UserRepository struct {
	MongoDBHandler
}

func (repo *UserRepository) AddUser(user domain.User) (err error) {
	err = repo.Add(user)
	if err != nil {
		panic(err)
	}

	return err
}
