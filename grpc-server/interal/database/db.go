package database

import (
	"database/sql"

	"fmt"

	"git.foxminded.com.ua/grpc/grpc-server/interal/config"

	// _ "github.com/go-mysql-org/go-mysql/driver"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func InitDatabase(conf *config.Config) (*sql.DB, error) {

	//?charset=utf8mb4&parseTime=True&loc=Local
	// mysqlconfig := fmt.Sprintf("")
	// dsn := fmt.Sprintf("root:%s@tcp(%s%s)/%s", conf.DBPassword, conf.DBHost, conf.DBAddr, conf.DBName)
	// fmt.Printf("\n\n\n%v\n\n\n", dsn)
	// sqlcfg := mysql.Config{
	// 	User: "root",
	// 	// Passwd: conf.DBPassword,
	// 	Passwd: "password",
	// 	Net:    "tcp",
	// 	// Addr:                 "172.19.0.2" + conf.DBAddr,
	// 	Addr: "localhost:3333",
	// 	// DBName:               conf.DBName,
	// 	DBName:               "api",
	// 	AllowNativePasswords: true,
	// 	CheckConnLiveness:    true,
	// 	MaxAllowedPacket:     5,
	// }

	// dsn := sqlcfg.FormatDSN()
	// fmt.Println(dsn)
	db, err := sql.Open("mysql", "root:password@tcp(172.17.0.2:3306)/api")
	if err != nil {
		return nil, err
	}

	// driverConnector, err := mysql.NewConnector(&sqlcfg)
	// if err != nil {
	// 	return nil, fmt.Errorf("mysql.NewConnector: %v", err)
	// }

	// db.SetMaxOpenConns(5)
	// db.SetMaxIdleConns(5)
	// db.SetConnMaxLifetime(time.Minute * 5)

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("db.Ping(): %v", err)
	}

	// migr, err := migrate.New("file:///usr/grpc-server/schema/", dsn)
	migr, err := migrate.New("file:///home/bopvlk/go/src/git.foxminded.com.ua/grpc/grpc-server/schema", "user:user@tcp(172.17.0.2:3306)/api")
	if err != nil {
		return nil, fmt.Errorf("migrate.New(): %v", err)
	}

	if err := migr.Up(); err != nil {
		return nil, err
	}
	return db, nil
}
