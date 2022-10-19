package main

import (
	"algorithm-exercises/4-twitter-bot/model"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Start...")

	myCredential := model.Credentials{

		AccessToken:       os.Getenv("ACCESS_TOKEN"),
		AccessTokenSecret: os.Getenv("ACCESS_TOKEN_SECRET"),
		ConsumerKey:       os.Getenv("CONSUMER_KEY"),
		ConsumerSecret:    os.Getenv("CONSUMER_SECRET"),
	}

	client, err := model.GetClient(&myCredential)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(client)
}
