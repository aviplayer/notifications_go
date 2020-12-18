package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/pressly/goose"

	"andrew.com/migrations/cmd/config"
)

var (
	flags    = flag.NewFlagSet("goose", flag.ExitOnError)
	dbString = flags.String(
		"dbString",
		fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			config.Host, config.Port, config.User, config.Password, config.Dbname), "connection string")
	dir = flags.String("dir", "./migrations", "directory with migration files")
)

type Error interface{}

func logErr(err Error, message string) {
	if err != nil {
		log.Fatalf(message, err)
	}
}

func main() {
	if err := flags.Parse(os.Args[1:]); err != nil {
		log.Fatalf("unreal error: %v\n", err)
	}
	args := flags.Args()

	if len(args) < 1 {
		flags.Usage()
		return
	}

	command := args[0]

	log.Println("Connecting..." + *dbString)

	db, err := goose.OpenDBWithDriver("postgres", *dbString)
	logErr(err, "goose: failed to open DB: %v\n")

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("goose: failed to close DB: %v\n", err)
		}
	}()

	arguments := []string{}
	if len(args) > 2 {
		arguments = append(arguments, args[2:]...)
	}

	if err := goose.Run(command, db, *dir, arguments...); err != nil {
		log.Fatalf("goose %v: %v", command, err)
	}
	log.Println("BUILD SUCCESSFUL...")
}
