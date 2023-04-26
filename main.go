package main

import (
	"fmt"

	"encoding/json"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

func main() {
	// Set up AWS credentials
	sess, err := session.NewSession(&aws.Config{
		Region:   aws.String("us-west-2"),
		Endpoint: aws.String("https://sns.us-west-2.amazonaws.com"),
	})

	if err != nil {
		fmt.Println("Failed to create AWS session", err)
		return
	}

	// Connect to the SNS service
	svc := sns.New(sess)

	// Set up the SNS topic ARN to publish to

	topicArn := os.Getenv("SNS_TOPIC_ARN")
	if topicArn == "" {
		fmt.Println("Missing SNS_TOPIC_ARN environment variable")
		return
	}

	// Construct the JSON message to publish
	message := map[string]int{"a": 1}
	jsonMessage, err := json.Marshal(message)
	if err != nil {
		fmt.Println("Failed to marshal JSON message", err)
		return
	}

	// Publish the message to the topic
	_, err = svc.Publish(&sns.PublishInput{
		Message:  aws.String(string(jsonMessage)),
		TopicArn: aws.String(topicArn),
	})

	if err != nil {
		fmt.Println("Failed to publish message to SNS", err)
		return
	}

	fmt.Println("Message published to SNS")
}
