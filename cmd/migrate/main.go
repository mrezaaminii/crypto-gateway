package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const migrationDir = "migrations"

func main() {
	if len(os.Args) < 2 {
		usage()
	}

	switch os.Args[1] {
	case "create" : {
		if len(os.Args) < 3 {
			usage()
		}
		create(os.Args[2])
	}
	case "up", "down", "refresh" : 
		run(os.Args[1])
	case "force": {
	if len(os.Args) < 3 {
			usage()
		}
		force(os.Args[2])
}
default: {
	usage()
}
	}
}

func create(name string) {

	if err := os.MkdirAll(migrationDir, 0755); err != nil {
		exit(err)
	}

	version := time.Now().UTC().Format("20060102150405")

	for _, kind := range []string{"up", "down"} {
		filePath := filepath.Join(migrationDir, fmt.Sprintf("%s_%s.%s.sql", version, name, kind))
		if err := os.WriteFile(filePath, []byte(""), 0644); err != nil {
			exit(err)
		}
		fmt.Println("created: ", filePath)
	}

}

func run(command string) {
	m := newMigrate()
	defer m.Close()

	var err error
	switch command {
		case "up":
			err = m.Up()
		case "down" : 
			err = m.Steps(-1)
		case "refresh": {
			err = m.Down()
				if err == nil || errors.Is(err, migrate.ErrNoChange) {
					err = m.Up()
				}
		}
	}

	if errors.Is(err, migrate.ErrNoChange) {
		fmt.Println("Nothing to migrate")
		return
	}

	if err != nil {
		exit(err)
	}

	fmt.Println("Done: ", command)
}

func force(v string) {
	version, err := strconv.Atoi(v)

	if err != nil {
		exit(fmt.Errorf("invalid version %q", v))
	}

	m := newMigrate()
	defer m.Close()

	if err := m.Force(version); err != nil {
		exit(err)
	}

	fmt.Println("Forced version:", version)
}

func newMigrate() *migrate.Migrate {
	dbURL := os.Getenv("DATABASE_URL")

	if dbURL == "" {
		exit(errors.New("DATABASE_URL is not set"))
	}

	m, err := migrate.New("file://"+migrationDir, dbURL)

	if err != nil {
		exit(err)
	}

	return m
}

func usage() {
	fmt.Println(`Usage: 
	go run ./cmd/migrate create <name>
	go run ./cmd/migrate up
	go run ./cmd/migrate down
	go run ./cmd/migrate refresh
	go run ./cmd/migrate force <version>   (use -1 for "no version")`)
	os.Exit(1)
}

func exit(err error) {
	fmt.Println("Error:", err)
	os.Exit(1)
}