package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

func main() {
	// Set up AWS credentials
	sess, err := session.NewSession(&aws.Config{
		Region:   aws.String("YOUR_REGION"),
		Endpoint: aws.String("https://sns.us-west-2.amazonaws.com"),
	})

	if err != nil {
		fmt.Println("Failed to create AWS session", err)
		return
	}

	// Connect to the SNS service
	svc := sns.New(sess)

	// Set up the SNS topic ARN to publish to
	topicArn := "YOUR_TOPIC_ARN"

	// Publish the message to the topic
	_, err = svc.Publish(&sns.PublishInput{
		Message:  aws.String("hello"),
		TopicArn: aws.String(topicArn),
	})

	if err != nil {
		fmt.Println("Failed to publish message to SNS", err)
		return
	}

	fmt.Println("Message published to SNS")
}
