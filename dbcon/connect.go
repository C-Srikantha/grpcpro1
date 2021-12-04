package dbcon

import (
	"context"

	"github.com/go-pg/pg"
)

func Connection() (*pg.DB, error) {
	address := &pg.Options{
		User:     "postgres",
		Password: "codecraft",
		Addr:     ":8080",
		Database: "grpcpractise",
	}
	conn := pg.Connect(address)
	_, err := conn.ExecContext(context.Background(), "SELECT 1")
	return conn, err
}
