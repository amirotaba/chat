package natsRepo

import (
	natsDomain "chat/domain/nats"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoRepository struct {
	Collection *mongo.Collection
}

func NewMongoRepository(clc *mongo.Collection) natsDomain.NatsRepository {
	return &mongoRepository{
		Collection: clc,
	}

}

func (m mongoRepository) Create(message natsDomain.Message) error {
	_, err := m.Collection.InsertOne(context.Background(), message)

	if err != nil {
		return err
	}

	return nil
}

func (m mongoRepository) Read(sub string) ([]natsDomain.Message, error) {
	filter := bson.D{{"subject", sub}}
	cur, err := m.Collection.Find(context.Background(), filter)

	if err != nil {
		return nil, err
	}

	defer cur.Close(context.Background())

	var results []natsDomain.Message

	for cur.Next(context.Background()) {

		var result natsDomain.Message

		if err = cur.Decode(&result); err != nil {
			return nil, err
		}

		results = append(results, result)

	}
	if results != nil {
		return results, nil
	}
	return nil, errors.New("not found! ")
}
