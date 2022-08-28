package main

import (
	natsHandler "chat/internal/nats/handler"
	"chat/internal/user/userHandler"
	"chat/utils"
	"fmt"
)

func main() {
	//Connect to Nats
	c := utils.ConnNats()

	//Connect to Database
	Db := utils.Connection()

	//Get repositories
	Repos := utils.NewRepository(Db)

	//Get usecases
	UseCases := utils.NewUseCase(Repos)

	//

	natshandler := natsHandler.NewHandler(UseCases.Nats)
	userhandler := userHandler.NewHandler(UseCases.User)

	fmt.Println("1: SignUp.\n2: SingIn.")
	opt := utils.Read()
	switch opt {
	case "1":
		if err := userhandler.SignUp(); err != nil {
			fmt.Println(err)
		}
		fallthrough
	case "2":
		if err := userhandler.SignIn(); err != nil {
			fmt.Println("1: Send private message.\n2: Join group.\n3: Create a group.")
			option := utils.Read()
			switch option {
			case "1":
				sub := natshandler.JoinPv()
				c.Sub(sub)
				for {
					if msg.Text != "" {
						text := NewMessage(msg)
						Pub(nc, text)
					}
				}
			case "2":
				sub := natshandler.JoinGp()
				c.Sub(sub)
			case "3":
				sub := natshandler.CreateGp()
				c.Sub(sub)
			}
		}
	}
}
