package infrastructure

import (
	"context"
	"log"
	"time"

	"github.com/outk/rewardingsites/src/userservice/src/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDBHandler struct {
	Ctx    context.Context
	Client *mongo.Client
}

func NewMongoDBHandler() *MongoDBHandler {
	uri := "mongodb+srv://<username>:<password>@<cluster-address>/test?w=majority"
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Fatalln(err)
		}
	}()
	// Ping the primary
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatalln(err)
	}

	mongoDBHandler := new(MongoDBHandler)
	mongoDBHandler.Ctx = ctx
	mongoDBHandler.Client = client

	return mongoDBHandler
}

func (handler *MongoDBHandler) Add(user domain.User) error {
	client := handler.Client.Database("rewardingsites")
	usersCollection := client.Collection("Users")

	_, err := usersCollection.InsertOne(handler.Ctx, bson.M{
		"ID":          user.ID,
		"first_name":  user.FirstName,
		"middle_name": user.MiddleName,
		"last_name":   user.LastName,
		"gender":      user.Gender,
		"birthday": bson.M{
			"year":  user.Birthday.Year,
			"month": user.Birthday.Month,
			"day":   user.Birthday.Day,
		},
		"email":    user.Email,
		"password": user.Password,
	})
	if err != nil {
		log.Fatalln(err)
	}

	return err
}
