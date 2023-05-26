package main

import (
	"bufio"
	"context"
	"github.com/jackc/pgx/v5"
	"os"
)

func main() {
	pathFile := ""
	fileType := ""

	conn, err := pgx.Connect(context.Background(), "postgres://test_user:test_password@127.0.0.1:5432/test_db")

	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	tx, err := conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		panic(err)
	}

	file, _ := os.Open(pathFile)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		curr := scanner.Text()

		body, err := os.ReadFile(curr)
		if err != nil {
			_ = tx.Rollback(ctx)
			panic(err)
		}

		_, err = conn.Exec(context.Background(), "insert into files (abspath, body, type) values ($1,$2,$3);",
			curr, string(body), fileType)
		if err != nil {
			_ = tx.Rollback(ctx)
			panic(err)
		}
	}
	err = tx.Commit(ctx)
	if err != nil {
		panic(err)
	}

}
