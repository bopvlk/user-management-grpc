package server

import (
	"log"
	"net"

	"git.foxminded.com.ua/grpc/service-user/interal/config"
	"git.foxminded.com.ua/grpc/service-user/interal/database"
	"git.foxminded.com.ua/grpc/service-user/interal/logger"
	"git.foxminded.com.ua/grpc/service-user/interal/repository"
	"git.foxminded.com.ua/grpc/service-user/proto/pb"
	"google.golang.org/grpc"
)

func Run() {
	l := logger.InitLoger()

	conf, err := config.InitConfig()
	if err != nil {
		l.Err.Fatal(err)
	}
	l.Info.Print("Config setup is successful")

	db, err := database.InitDatabase(conf)
	if err != nil {
		l.Err.Fatal(err)
	}
	defer db.Close()
	l.Info.Print("Initialization database is successful")

	lis, err := net.Listen("tcp", conf.GRCPPort)
	if err != nil {
		l.Err.Fatal(err)
	}
	defer lis.Close()
	l.Info.Print("Listening at port ", conf.GRCPPort)

	userService := newUserService(repository.NewUserRepository(conf, db, l))

	grpcServer := grpc.NewServer()
	l.Info.Print("GRPC Service-user runing...")
	pb.RegisterApiServiceServer(grpcServer, userService)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}

}
