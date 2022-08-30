package userRepo

import (
	"chat/domain"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoRepository struct {
	Collection *mongo.Collection
}

func NewMongoRepository(clc *mongo.Collection) domain.UserRepository {
	return &mongoRepository{
		Collection: clc,
	}

}

func (m *mongoRepository) Create(user domain.User) error {
	_, err := m.Collection.InsertOne(context.Background(), user)

	if err != nil {
		return err
	}

	return nil
}

func (m *mongoRepository) Read(username string) (domain.User, error) {
	filter := bson.D{{"username", username}}
	cur, err := m.Collection.Find(context.Background(), filter)

	if err != nil {
		return domain.User{}, err
	}

	defer cur.Close(context.Background())

	var results []domain.User

	for cur.Next(context.Background()) {

		var result domain.User

		if err = cur.Decode(&result); err != nil {
			return domain.User{}, err
		}

		results = append(results, result)

	}
	if results != nil {
		return results[0], nil
	}

	return domain.User{}, errors.New("this username doesn't exist. ")
}
