package config

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func Init() {
	var myEnv map[string]string
	myEnv, err := godotenv.Read()

	s3Bucket := myEnv["S3_BUCKET"]
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Print(s3Bucket)

	// now do something with s3 or whatever
}
