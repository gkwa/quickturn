package main

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/spf13/cobra"
)

func main() {
	// Set up the root command and its flags
	rootCmd := &cobra.Command{
		Use:   "sns-publish",
		Short: "Publish a message to an SNS topic",
		Run:   run,
	}
	rootCmd.PersistentFlags().StringP("message", "m", "", "The JSON message to publish")

	// Execute the root command
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
	}
}

func run(cmd *cobra.Command, args []string) {
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
	topicArn := "arn:aws:sns:us-west-2:123456789012:MyTopic"

	// Get the JSON message from the command-line argument
	message, err := cmd.PersistentFlags().GetString("message")
	if err != nil {
		fmt.Println("Failed to get message flag", err)
		return
	}

	// Unmarshal the JSON message
	var jsonMessage interface{}
	err = json.Unmarshal([]byte(message), &jsonMessage)
	if err != nil {
		fmt.Println("Failed to unmarshal JSON message", err)
		return
	}

	// Marshal the JSON message
	jsonBytes, err := json.Marshal(jsonMessage)
	if err != nil {
		fmt.Println("Failed to marshal JSON message", err)
		return
	}

	// Publish the message to the topic
	_, err = svc.Publish(&sns.PublishInput{
		Message:  aws.String(string(jsonBytes)),
		TopicArn: aws.String(topicArn),
	})

	if err != nil {
		fmt.Println("Failed to publish message to SNS", err)
		return
	}

	fmt.Println("Message published to SNS")
}
