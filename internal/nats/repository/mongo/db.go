package natsRepo

import (
	"chat/domain"
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

func NewMongoRepository(c domain.DataBase) domain.NatsRepository {
	return &mongoRepository{
		Pv:  c.Pv,
		Gp:  c.Gp,
		Msg: c.Msg,
	}

}

func (m mongoRepository) CreatePv(form domain.Private) error {
	_, err := m.Pv.InsertOne(context.Background(), form)

	if err != nil {
		return err
	}

	return nil
}

func (m mongoRepository) ReadPv(sub string) (domain.Private, error) {
	filter := bson.D{{"sub", sub}}
	cur, err := m.Pv.Find(context.Background(), filter)

	if err != nil {
		return domain.Private{}, err
	}

	defer cur.Close(context.Background())

	var results []domain.Private

	for cur.Next(context.Background()) {

		var result domain.Private

		if err = cur.Decode(&result); err != nil {
			return domain.Private{}, err
		}

		results = append(results, result)

	}
	if results != nil {
		return results[0], nil
	}
	return domain.Private{}, errors.New("not found! ")
}

func (m mongoRepository) CreateGp(form domain.Group) error {
	_, err := m.Gp.InsertOne(context.Background(), form)

	if err != nil {
		return err
	}

	return nil
}

func (m mongoRepository) ReadGp(sub string) (domain.Group, error) {
	filter := bson.D{{"sub", sub}}
	cur, err := m.Gp.Find(context.Background(), filter)

	if err != nil {
		return domain.Group{}, err
	}

	defer cur.Close(context.Background())

	var results []domain.Group

	for cur.Next(context.Background()) {

		var result domain.Group

		if err = cur.Decode(&result); err != nil {
			return domain.Group{}, err
		}

		results = append(results, result)

	}
	if results != nil {
		return results[0], nil
	}
	return domain.Group{}, errors.New("not found! ")
}

func (m mongoRepository) CreateMsg(form domain.Message) error {
	_, err := m.Msg.InsertOne(context.Background(), form)

	if err != nil {
		return err
	}

	return nil
}

func (m mongoRepository) ReadMsg(id primitive.ObjectID) ([]domain.Message, error) {
	filter := bson.D{{"chatid", id}}
	cur, err := m.Msg.Find(context.Background(), filter)

	if err != nil {
		return nil, err
	}

	defer cur.Close(context.Background())

	var results []domain.Message

	for cur.Next(context.Background()) {

		var result domain.Message

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
