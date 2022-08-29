package main

import (
	natsDomain "chat/domain/nats"
	natsHandler "chat/internal/nats/handler"
	"chat/internal/user/userHandler"
	"chat/utils"
	"fmt"
	"time"
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

	fmt.Println("1: SignUp.\n2: SignIn.")
	opt := utils.Read()
	switch opt {
	case "1":
		if err := userhandler.SignUp(); err != nil {
			fmt.Println(err)
		}
		fallthrough
	case "2":
		user, err := userhandler.SignIn()
		if err != nil {
			fmt.Println("err")
		}
		fmt.Println("1: Send private message.\n2: Join group.\n3: Create a group.")
		option := utils.Read()
		switch option {
		case "1":
			form := natshandler.JoinPv(user.UserName)
			c.Sub(user.UserName)
			c.Sub(form.Sub)
			for {
				text := utils.Read()
				text = user.UserName + ": " + text
				msg := natsDomain.Message{
					Message: text,
					Sub:     form.Sub,
					ChatID:  form.ID,
					Time:    time.Now(),
				}
				c.Pub(msg)
				if err := natshandler.CreateMsg(msg); err != nil {
					fmt.Println(err)
				}
			}
		case "2":
			form := natshandler.JoinGp()
			c.Sub(form.Sub)
			for {
				text := utils.Read()
				text = user.UserName + ": " + text
				msg := natsDomain.Message{
					Message: text,
					Sub:     form.Sub,
					ChatID:  form.ID,
					Time:    time.Now(),
				}
				c.Pub(msg)
				if err := natshandler.CreateMsg(msg); err != nil {
					fmt.Println(err)
				}
			}
		case "3":
			form, err := natshandler.CreateGp()
			if err != nil {
				fmt.Println(err)
			}
			c.Sub(form.Sub)
			for {
				text := utils.Read()
				text = user.UserName + ": " + text
				msg := natsDomain.Message{
					Message: text,
					Sub:     form.Sub,
					ChatID:  form.ID,
					Time:    time.Now(),
				}
				c.Pub(msg)
				if err := natshandler.CreateMsg(msg); err != nil {
					fmt.Println(err)
				}
			}
		}
	}
}
