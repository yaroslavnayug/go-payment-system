package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	// postgres connect
	connection, err := pgxpool.Connect(context.Background(), os.Getenv("POSTGRESQL_URL"))
	if err != nil {
		panic(fmt.Errorf("unable to connect to database: %v", err))
	}

	// prepare sql file
	mainMigration := filepath.Join("cmd", "migration", "sql", "main.sql")
	c, err := ioutil.ReadFile(mainMigration)
	if err != nil {
		panic(fmt.Errorf("couldn't read main.sql file: %v", err))
	}
	sql := string(c)

	fmt.Println("Start main migration")
	_, err = connection.Exec(context.Background(), sql)
	if err != nil {
		panic(fmt.Errorf("couldn't exec main.sql file: %v", err))
	}
	fmt.Println("End main migration")
}
