package main

import (
	"context"
	"function-extractor/pkg/extractor"
	"github.com/jackc/pgx/v5"
	"os"
	"path/filepath"
	"regexp"
)

func main() {
	var functions []extractor.Functions

	fileInfo, _ := os.Stat(startingDir)
	if fileInfo.IsDir() {
		functions = readDir(startingDir)
	}

	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	tx, err := conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		panic(err)
	}

	for _, c := range functions {
		_, err := conn.Exec(context.Background(), "insert into test_db.public.functions (filepath, func_name, body) values ($1,$2,$3);", c.File, c.Name, c.Body)
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

func readDir(path string) (rtnMe []extractor.Functions) {
	files, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if file.IsDir() {
			// Recursively descend into subdirectories
			rtnMe = append(rtnMe, readDir(filepath.Join(path, file.Name()))...)
		} else if filepath.Ext(file.Name()) == ".go" && !isTestFile(file.Name()) {
			functions, err := extractor.ExtractFunctionsFromFile(filepath.Join(path, file.Name()))
			if err != nil {
				panic(err)
			}
			rtnMe = append(rtnMe, functions...)
		}
	}
	return rtnMe
}

func isTestFile(name string) bool {
	re := regexp.MustCompile(`_test\.go$`)
	return re.MatchString(name)
}
