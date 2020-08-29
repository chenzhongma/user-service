package main

import (
	"fmt"
	database "github.com/chenzhongma/user-service/db"
	"github.com/chenzhongma/user-service/handler"
	pb "github.com/chenzhongma/user-service/proto/user"
	"github.com/micro/go-micro/v2"
	"log"
	repository "github.com/chenzhongma/user-service/repo"
)
func main()  {
	db,err:=database.CreateConnection()
	defer db.Close()
	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}
	db.AutoMigrate(&pb.User{})
	repo:=&repository.UserRepository{db}
	srv:=micro.NewService(
		micro.Name("laracom.user.service"),
		micro.Version("latest"),
		)
	srv.Init()
	pb.RegisterUserServiceHandler(srv.Server(),&handler.UserService{
		Repo: repo,
	})
	if err:=srv.Run();err!=nil{
		fmt.Println(err)
	}
}

