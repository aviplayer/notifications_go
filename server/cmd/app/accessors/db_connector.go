package accessors

import (
	"context"
	"fmt"

	. "github.com/jackc/pgx/v4/pgxpool"

	"andrew.com/notifications/cmd/app/config"
)

var (
	dbString = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.Dbname)
)

type DbConnection struct {
	connection *Pool
}

func CreateConnection() DbConnection {
	dbPool, err := Connect(context.Background(), dbString)
	if err != nil {
		panic(fmt.Sprintf("Unable to connect to database: %v\n", err))
	}
	return DbConnection{dbPool}
}

func (conn DbConnection) getConnection() (*Conn, error) {
	con, err := conn.connection.Acquire(context.Background())
	return con, err
}
