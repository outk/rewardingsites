package usecase

import "github.com/outk/rewardingsites/src/userservice/src/domain"

type UserRepository interface {
	Add(domain.User) error
}
