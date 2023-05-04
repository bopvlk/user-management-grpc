package database

import (
	"database/sql"

	"fmt"

	"git.foxminded.com.ua/grpc/service-user/interal/apperrors"
	"git.foxminded.com.ua/grpc/service-user/interal/config"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	_ "github.com/go-sql-driver/mysql"
)

func InitDatabase(conf *config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf("root:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=true&loc=Local&multiStatements=true", conf.DBPassword, conf.DBAddr, conf.DBName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, apperrors.DbOpenErr.AppendMessage(err)
	}

	err = db.Ping()
	if err != nil {
		return nil, apperrors.DbPingErr.AppendMessage(err)
	}

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		return nil, apperrors.MigrateDriverErr.AppendMessage(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		conf.MigrateFileAddr,
		"mysql", driver)
	if err != nil {
		return nil, apperrors.MigrateCreationErr.AppendMessage(err)
	}
	m.Up()

	return db, nil
}
