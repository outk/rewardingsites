package db

import "github.com/outk/rewardingsites/src/userservice/src/domain"

type MongoDBHandler interface {
	Add(domain.User) error
}
