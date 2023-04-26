package main

import (
	"encoding/json"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"log"
)

type Message struct {
	Default string `json:"default"`
}

type Person struct {
	Name string `json:"name"`
}

func main() {
	cfg, _ := external.LoadDefaultAWSConfig()
	snsClient := sns.New(cfg)

	person := Person{
		Name: "Felix Kjellberg",
	}
	personStr, _ := json.Marshal(person)

	message := Message{
		Default: string(personStr),
	}
	messageBytes, _ := json.Marshal(message)
	messageStr := string(messageBytes)

	req := snsClient.PublishRequest(&sns.PublishInput{
		TopicArn:         aws.String("arn:aws:sns:us-west-2:193048895737:hello1"),
		Message:          aws.String(messageStr),
		MessageStructure: aws.String("json"),
	})

	res, err := req.Send()
	if err != nil {
		log.Fatal(err)
	}

	log.Print(res)
}
