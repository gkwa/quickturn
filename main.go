package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sns-publish",
	Short: "Publish a message to an SNS topic",
	Run:   publishMessage,
}

var (
	topicArn string
	message  string
)

func init() {
	rootCmd.PersistentFlags().StringVar(&topicArn, "topic-arn", "", "The ARN of the SNS topic to publish to")
	rootCmd.PersistentFlags().StringVar(&message, "message", "", "The JSON message to publish")
	rootCmd.MarkPersistentFlagRequired("topic-arn")
	rootCmd.MarkPersistentFlagRequired("message")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func publishMessage(cmd *cobra.Command, args []string) {
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

	// Publish the message to the topic
	_, err = svc.Publish(&sns.PublishInput{
		Message:  aws.String(message),
		TopicArn: aws.String(topicArn),
	})

	if err != nil {
		fmt.Println("Failed to publish message to SNS", err)
		return
	}

	fmt.Println("Message published to SNS")
}
