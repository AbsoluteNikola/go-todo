package config

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/ichtrojan/thoth"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"os"
)

func Database() *sql.DB {
	logger, _ := thoth.Init("log")

	user, exist := os.LookupEnv("DB_USER")

	if !exist {
		logger.Log(errors.New("DB_USER not set in .env"))
		log.Fatal("DB_USER not set in .env")
	}

	pass, exist := os.LookupEnv("DB_PASS")

	if !exist {
		logger.Log(errors.New("DB_PASS not set in .env"))
		log.Fatal("DB_PASS not set in .env")
	}

	host, exist := os.LookupEnv("DB_HOST")

	if !exist {
		logger.Log(errors.New("DB_HOST not set in .env"))
		log.Fatal("DB_HOST not set in .env")
	}

	port, exist := os.LookupEnv("DB_PORT")

	if !exist {
		logger.Log(errors.New("DB_PORT not set in .env"))
		log.Fatal("DB_PORT not set in .env")
	}

	dbname, exist := os.LookupEnv("DB_NAME")

	if !exist {
		logger.Log(errors.New("DB_NAMET not set in .env"))
		log.Fatal("DB_PORT not set in .env")
	}

	credentials := fmt.Sprintf("sslmode=disable host=%s port=%s user=%s password=%s dbname=%s",
		host, port, user, pass, dbname)

	fmt.Println(credentials)

	database, err := sql.Open("postgres", credentials)

	if err != nil {
		logger.Log(err)
		log.Fatal(err)
	} else {
		fmt.Println("Database Connection Successful")
	}

	if err != nil {
		fmt.Println(err)
	}

	return database
}
