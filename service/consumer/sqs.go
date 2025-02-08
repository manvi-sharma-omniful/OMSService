package service

import (
	"context"
	"fmt"
	"log"
	"time"

	"awesomeProject/Project/OMS/service"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types" //
)

type SqsConsumer struct {
	QURL      string
	SqsClient *sqs.Client
}

func NewSqsConsumer(qURL string, ctx context.Context) (*SqsConsumer, error) {
	if qURL == "" {
		return nil, fmt.Errorf("no QUEUE_URL specified")
	}

	cnfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}

	return &SqsConsumer{
		QURL:      qURL,
		SqsClient: sqs.NewFromConfig(cnfg),
	}, nil
}

func (c *SqsConsumer) StartConsumer(ctx context.Context) {
	for {
		receiveMessages := &sqs.ReceiveMessageInput{
			QueueUrl:            aws.String(c.QURL),
			MaxNumberOfMessages: 10,
			WaitTimeSeconds:     5,
		}

		resp, err := c.SqsClient.ReceiveMessage(ctx, receiveMessages)
		if err != nil {
			log.Println("Error receiving messages:", err)
			continue
		}

		for _, message := range resp.Messages {
			if err := c.ProcessMessage(ctx, &message); err != nil {
				log.Println("Error processing message:", err)
			} else {
				c.DeleteMessage(ctx, message.ReceiptHandle)
			}
		}
		time.Sleep(time.Second * 2)
	}
}

func (c *SqsConsumer) ProcessMessage(ctx context.Context, msg *types.Message) error {
	service.GetOrders(*msg.Body)
	return nil
}

func (c *SqsConsumer) DeleteMessage(ctx context.Context, receiptHandle *string) {
	_, err := c.SqsClient.DeleteMessage(ctx, &sqs.DeleteMessageInput{
		QueueUrl:      aws.String(c.QURL),
		ReceiptHandle: receiptHandle,
	})

	if err != nil {
		log.Println("Error deleting message:", err)
	} else {
		fmt.Println("Message deleted successfully")
	}
}
