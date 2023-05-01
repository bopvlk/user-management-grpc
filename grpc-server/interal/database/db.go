package database

import (
	"database/sql"

	"fmt"

	"git.foxminded.com.ua/grpc/grpc-server/interal/config"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	_ "github.com/go-sql-driver/mysql"
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

	//allowNativePasswords=true
	db, err := sql.Open("mysql", "root:qwerty@tcp(mariadb:3306)/api")
	if err != nil {
		return nil, fmt.Errorf(" sql.Open(): %v", err)
	}

	// _, err = db.Exec("CREATE DATABASE `api`")
	// if err != nil {
	// 	return nil, fmt.Errorf("db.Exec(\"CREATE DATABASE...\"): %v", err)
	// }

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("db.Ping(): %v", err)
	}

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		return nil, fmt.Errorf("mysql.WithInstance(): %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file:///usr/grpc-server/schema/",
		"mysql", driver)
	if err != nil {
		return nil, fmt.Errorf("migrate.NewWithDatabaseInstance(): %v", err)
	}
	m.Up()

	// migr, err := migrate.New("file:///usr/grpc-server/schema/", "user:user@tcp(172.17.0.2:3306)/api")
	// migr, err := migrate.New("file:///home/bopvlk/go/src/git.foxminded.com.ua/grpc/grpc-server/schema", "user:user@tcp(172.17.0.2:3306)/api")
	// if err != nil {
	// 	return nil, fmt.Errorf("migrate.New(): %v", err)
	// }

	// if err := migr.Up(); err != nil {
	// 	return nil, err
	// }
	return db, nil
}
