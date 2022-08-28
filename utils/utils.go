package utils

import (
	"bufio"
	"chat/domain"
	"chat/internal/nats"
	natsRepo "chat/internal/nats/repository/mongo"
	natsUsecase "chat/internal/nats/usecase"
	"chat/internal/user/repository/mongo"
	"chat/internal/user/usecase"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"
	"strings"
)

func ConnNats() *nats.Client {
	c, err := nats.New()
	if err != nil {
		log.Println("Connecting to message broker failed")
	}
	return c
}

func Connection() domain.DataBase {
	dbUser := "root"
	dbPass := "root"
	dbName := "chat_db"
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+dbUser+":"+dbPass+"@localhost:27017/"))
	if err != nil {
		panic(err)
	}
	//defer func() {
	//	if err = client.Disconnect(ctx); err != nil {
	//		panic(err)
	//	}
	//}()
	// Ping the primary
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	userCollection := client.Database(dbName).Collection("users")

	Db := domain.DataBase{
		Collection: userCollection,
		Client:     client,
	}

	fmt.Println("Successfully connected and pinged.")
	return Db
}

func NewRepository(Db domain.DataBase) domain.Repositories {
	repository := domain.Repositories{
		User: userRepo.NewMongoRepository(Db.Collection),
		Nats: natsRepo.NewMongoRepository(Db.Collection),
	}
	return repository
}

func NewUseCase(repo domain.Repositories) domain.UseCases {
	usecase := domain.UseCases{
		User: userUseCase.NewUseCase(repo),
		Nats: natsUsecase.NewUseCase(repo),
	}
	return usecase
}

func Read() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	return text
}
