package httpserver

import (
	"log"
	"os"

	"git.foxminded.com.ua/grpc/grpc-client-api-gateway/interal/apperrors"
	"git.foxminded.com.ua/grpc/grpc-client-api-gateway/interal/config"
	clients "git.foxminded.com.ua/grpc/grpc-client-api-gateway/interal/service-clients"
	"github.com/fatih/color"
	"google.golang.org/grpc"
)

type httpServer struct {
	userDAO clients.DAO
	log     *loger
}

type loger struct {
	info *log.Logger
	warn *log.Logger
	err  *log.Logger
}

func Run() {
	l := initLoger()

	conf, err := config.InitConfig()
	if err != nil {
		l.err.Fatal(err)
	}

	conn, err := initGRPCClientConn(conf)
	if err != nil {
		l.err.Fatal(err)
	}
	defer conn.Close()

	srv := httpServer{
		log:     l,
		userDAO: clients.NewUserService(conf, conn),
	}

	httpRouter := initRouter(&srv)
	httpRouter.Run(conf.HTTPPort)

}

func initLoger() *loger {
	warn := log.New(os.Stderr, color.HiYellowString("[ WARN  ]"), log.Ltime|log.Lshortfile)
	info := log.New(os.Stderr, color.HiGreenString("[ INFO  ]"), log.Ltime|log.Lshortfile)
	err := log.New(os.Stderr, color.HiRedString("[ ERROR ]"), log.Ltime|log.Lshortfile)

	return &loger{
		info: info,
		warn: warn,
		err:  err,
	}
}

func initGRPCClientConn(conf *config.Config) (*grpc.ClientConn, error) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("grpc-server"+conf.GRPCPort /*grpc.WithTransportCredentials(insecure.NewCredentials())*/, grpc.WithInsecure())
	if err != nil {
		return nil, apperrors.ClientConnectionGRPCServer.AppendMessage(err)
	}
	return conn, nil
}
