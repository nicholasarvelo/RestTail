package main

import (
	"github.com/nicholasarvelo/resttail"
	"os"
)

func main() {
	url := os.Getenv("TESTRAIL_URL")
	if url == "" {
		url = "http://localhost:7070/testrail"
	}
	username := os.Getenv("TESTRAIL_USERNAME")
	password := os.Getenv("TESTRAIL_TOKEN")

	client := resttail.NewClient(url, username, password)

	err := client.GenerateCustom()
	if err != nil {
		panic(err)
	}

}