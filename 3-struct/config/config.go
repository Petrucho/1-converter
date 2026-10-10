package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Encrypter struct {
	Key string
}

func NewEncrypter() *Encrypter {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Не удалось загрузить ENV-файл!")
	}
	key := os.Getenv("KEY")
	if key == "" {
		panic("Не передан параметр KEY в переменные окружения!")
	}
	return &Encrypter{
		Key: key,
	}
}
