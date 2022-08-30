package domain

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type DataBase struct {
	User   *mongo.Collection
	Pv     *mongo.Collection
	Gp     *mongo.Collection
	Msg    *mongo.Collection
	Client *mongo.Client
}

type UseCases struct {
	User UserUseCase
	Nats NatsUseCase
}

type Repositories struct {
	User UserRepository
	Nats NatsRepository
}
