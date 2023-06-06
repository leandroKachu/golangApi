package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	StringDeConexao = ""

	Port      = 0
	SecretKey []byte
)

func Carregar() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	Port, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		Port = 9000
	}

	// StringDeConexao := fmt.Sprintf("postgresql://%s:%s@localhost:5432/%s?%s&sslmode=disable",
	// 	os.Getenv("DB_USER"),
	// 	os.Getenv("DB_PASSWORD"),
	// 	os.Getenv("DB_NAME"),
	// 	"charset=utf8&parseTime=True&loc=Local",
	// )

	SecretKey = []byte(os.Getenv("SECRET_KEY"))

	fmt.Println(StringDeConexao)

}
