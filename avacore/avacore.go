package avacore

import (
	"encoding/json"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func CheckEnv() {
	_, ok := os.LookupEnv("AVABOT_NAME")
	if !ok {
		log.Fatal("AVABOT_NAME must be present in the environment.")
	}
}

func getBotFromEnvironment() *Avabot {
	CheckEnv()
	avabot := Avabot{
		Id:   os.Getenv("AVABOT_ID"),
		Name: os.Getenv("AVABOT_NAME"),
		Bio:  os.Getenv("AVABOT_BIO"),
	}

	return &avabot
}

func GetBotInfo() []byte {
	avabot := getBotFromEnvironment()
	avabotJsonBytes, err := json.Marshal(avabot)
	if err != nil {
		log.Fatal(err)
	}
	return avabotJsonBytes
}
