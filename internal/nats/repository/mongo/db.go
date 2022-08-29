package natsRepo

import (
	"chat/domain"
	natsDomain "chat/domain/nats"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoRepository struct {
	Pv  *mongo.Collection
	Gp  *mongo.Collection
	Msg *mongo.Collection
}

func NewMongoRepository(c domain.DataBase) natsDomain.NatsRepository {
	return &mongoRepository{
		Pv:  c.Pv,
		Gp:  c.Gp,
		Msg: c.Msg,
	}

}

func (m mongoRepository) CreatePv(form natsDomain.Private) error {
	_, err := m.Pv.InsertOne(context.Background(), form)

	if err != nil {
		return err
	}

	return nil
}

func (m mongoRepository) ReadPv(sub string) (natsDomain.Private, error) {
	filter := bson.D{{"sub", sub}}
	cur, err := m.Pv.Find(context.Background(), filter)

	if err != nil {
		return natsDomain.Private{}, err
	}

	defer cur.Close(context.Background())

	var results []natsDomain.Private

	for cur.Next(context.Background()) {

		var result natsDomain.Private

		if err = cur.Decode(&result); err != nil {
			return natsDomain.Private{}, err
		}

		results = append(results, result)

	}
	if results != nil {
		return results[0], nil
	}
	return natsDomain.Private{}, errors.New("not found! ")
}

func (m mongoRepository) CreateGp(form natsDomain.Group) error {
	_, err := m.Gp.InsertOne(context.Background(), form)

	if err != nil {
		return err
	}

	return nil
}

func (m mongoRepository) ReadGp(sub string) (natsDomain.Group, error) {
	filter := bson.D{{"sub", sub}}
	cur, err := m.Gp.Find(context.Background(), filter)

	if err != nil {
		return natsDomain.Group{}, err
	}

	defer cur.Close(context.Background())

	var results []natsDomain.Group

	for cur.Next(context.Background()) {

		var result natsDomain.Group

		if err = cur.Decode(&result); err != nil {
			return natsDomain.Group{}, err
		}

		results = append(results, result)

	}
	if results != nil {
		return results[0], nil
	}
	return natsDomain.Group{}, errors.New("not found! ")
}

func (m mongoRepository) CreateMsg(form natsDomain.Message) error {
	_, err := m.Msg.InsertOne(context.Background(), form)

	if err != nil {
		return err
	}

	return nil
}

func (m mongoRepository) ReadMsg(id primitive.ObjectID) ([]natsDomain.Message, error) {
	filter := bson.D{{"chatid", id}}
	cur, err := m.Msg.Find(context.Background(), filter)

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
	return nil, nil
}
