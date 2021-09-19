package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"log"
	"os"
)

func GetQueueURL(sess *session.Session, queue string) (*sqs.GetQueueUrlOutput, error) {
	svc := sqs.New(sess)

	log.Println("queue_name: " + queue)
	result, err := svc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: &queue,
	})
	if err != nil {
		return nil, err
	}

	return result, nil
}

func SendMsg(sess *session.Session, queueURL *string, details ContactDetails) error {
	svc := sqs.New(sess)

	b, _ := json.Marshal(details)
	fmt.Println(string(b))

	_, err := svc.SendMessage(&sqs.SendMessageInput{
		MessageBody: aws.String(string(b)),
		QueueUrl:    queueURL,
	})
	if err != nil {
		return err
	}

	return nil
}

func sendSQS(details ContactDetails) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)
	if err != nil {
		fmt.Println("Got an error creating session")
		fmt.Println(err)
		return
	}

	result, err := GetQueueURL(sess, os.Getenv("QUEUE_NAME"))
	if err != nil {
		fmt.Println("Got an error getting the queue URL:")
		fmt.Println(err)
		return
	}

	queueURL := result.QueueUrl

	err = SendMsg(sess, queueURL, details)
	if err != nil {
		fmt.Println("Got an error sending the message:")
		fmt.Println(err)
		return
	}

	fmt.Println("Sent message to queue ")
}
