package repository

import (
	"context"
	"database/sql"
	"fmt"

	"git.foxminded.com.ua/grpc/grpc-server/interal/config"
	"git.foxminded.com.ua/grpc/grpc-server/interal/domain/models"
	"git.foxminded.com.ua/grpc/grpc-server/interal/logger"
)

type UserRepository struct {
	conf *config.Config
	db   *sql.DB
	l    *logger.Logger
}

func NewUserRepository(conf *config.Config, db *sql.DB, l *logger.Logger) *UserRepository {
	return &UserRepository{
		conf: conf,
		db:   db,
		l:    l,
	}
}

func (ur *UserRepository) Create(ctx context.Context, firstName, lastName, email, password string) (uint32, error) {
	sqlStatement := fmt.Sprintf("INSERT INTO users (first_name, last_name, email, password) VALUES ('%s', '%s', '%s', '%s')",
		firstName, lastName, email, password)
	fmt.Println("\n\n\n\n", sqlStatement)
	result, err := ur.db.ExecContext(ctx, sqlStatement)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	ur.l.Info.Printf("Created user with id = %d\n", id)
	return uint32(id), nil
}

func (ur *UserRepository) FetchByEmail(ctx context.Context, email string) (*models.User, error) {
	u := &models.User{}
	uaa := make([]models.User, 0)
	sqlStatement := fmt.Sprint("SELECT * FROM users;")
	// WHERE email = '%s' AND deleted_at IS NULL LIMIT 1
	fmt.Println("\n\n\n", sqlStatement)
	rows, err := ur.db.QueryContext(ctx, sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.Password, &u.CreatedAt, &u.UpdatedAt, &u.DeleteAt); err != nil {
			return nil, err
		}
		uaa = append(uaa, *u)
	}

	fmt.Println("\n\n\n", uaa)

	if err := rows.Err(); err != nil {
		return nil, err
	}

	ur.l.Info.Printf("Fetched by email user with id = %d\n", u.ID)
	return u, nil
}

func (ur *UserRepository) FetchByID(ctx context.Context, id uint32) (*models.User, error) {
	u := &models.User{}
	sqlStatement := "SELECT 1 FROM users WHERE id = ? AND deleted_at IS NULL LIMIT 1"
	rows, err := ur.db.QueryContext(ctx, sqlStatement, ur.conf.DBName, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&u.ID, u.FirstName, &u.LastName, &u.Email, &u.CreatedAt, &u.UpdatedAt, &u.DeleteAt); err != nil {
			return nil, err
		}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	ur.l.Info.Printf("Fetched by id user with id = %d\n", u.ID)
	return u, nil
}

func (ur *UserRepository) FetchUsers(ctx context.Context, limit, page int32) ([]models.User, error) {
	u := make([]models.User, limit)
	offset := (page - 1) * limit

	sqlStatement := "SELECT * FROM ? OFFSET ? LIMIT ? WHERE deleted_at IS NULL"
	rows, err := ur.db.QueryContext(ctx, sqlStatement, ur.conf.DBName, offset, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var i int
	for rows.Next() {
		if err = rows.Scan(&u[i].ID, u[i].FirstName, &u[i].LastName, &u[i].Email, &u[i].CreatedAt, &u[i].UpdatedAt, &u[i].DeleteAt); err != nil {
			return nil, err
		}
		i++
	}

	if rows.Err() != nil {
		return nil, err
	}
	ur.l.Info.Printf("Fetched users with ids from %d to %d", u[0].ID, u[len(u)].ID)

	return u, nil
}

func (ur *UserRepository) Update(ctx context.Context, u *models.User) (uint32, error) {
	fetchedUser, err := ur.FetchByID(ctx, uint32(u.ID))
	if err != nil {
		return 0, err
	}

	if u.FirstName == "" {
		u.FirstName = fetchedUser.FirstName
	}
	if u.LastName == "" {
		u.LastName = fetchedUser.LastName
	}
	if u.Email == "" {
		u.Email = fetchedUser.Email
	}
	if u.Password == "" {
		u.Password = fetchedUser.Password
	}

	sqlStatement := "UPDATE ? SET first_name = ?, last_name = ?, email = ?, password = ?, update_at = now() WHERE id = ? AND deleted_at IS NULL"
	result, err := ur.db.ExecContext(ctx, sqlStatement, ur.conf.DBName, u.FirstName, u.LastName, u.Email, u.Password)
	if err != nil {
		return 0, err
	}
	id, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	ur.l.Info.Printf("Updated user with id = %d", id)

	return uint32(id), nil
}

func (ur *UserRepository) Delete(ctx context.Context, id uint32) (uint32, error) {
	sqlStatement := "UPDATE ? SET deleted_at = now() WHERE id = ? AND deleted_at IS NULL"
	result, err := ur.db.ExecContext(ctx, sqlStatement, ur.conf.DBName, id)
	if err != nil {
		return 0, err
	}
	resultID, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	ur.l.Info.Printf("Soft deleted user with id = %d", id)
	return uint32(resultID), nil
}
