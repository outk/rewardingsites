package usecase

import (
	"context"
	"time"

	"github.com/outk/rewardingsites/src/userservice/src/model"
	"github.com/outk/rewardingsites/src/userservice/src/repository"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func (s *model.UserService) AddUser(c *context.Context, req *repository.AddUserRequest) (rep *repository.AddUserReply, err error) {
	uri := "mongodb+srv://<username>:<password>@<cluster-address>/test?w=majority"
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	return rep, nil
}
