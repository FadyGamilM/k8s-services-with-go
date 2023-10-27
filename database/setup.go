package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

var conn *pgxpool.Pool

// InitDB initialize the database pool and returns error
func InitDB(connString string) error {
	var err error
	conn, err = pgxpool.New(context.Background(), connString)
	if err != nil {
		return err
	}
	return nil
}

func GetTime(ctx *gin.Context) time.Time {
	var tm time.Time

	err := conn.QueryRow(ctx, "SELECT NOW() as now;").Scan(&tm)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	return tm
}
