package listeners

import (
	"awesomeProject/Project/OMS/service"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

type SqsConsumer struct {
	QueueURL  string
	SqsClient *sqs.Client
}

func NewSqsConsumer(queueURL string, ctx context.Context) (*SqsConsumer, error) {
	if queueURL == "" {
		return nil, fmt.Errorf("no QUEUE_URL specified")
	}

	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}

	return &SqsConsumer{
		QueueURL:  queueURL,
		SqsClient: sqs.NewFromConfig(cfg),
	}, nil
}

func (c *SqsConsumer) StartConsumer(ctx context.Context) {
	for {
		receiveMessages := &sqs.ReceiveMessageInput{
			QueueUrl:            aws.String(c.QueueURL),
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
		time.Sleep(1 * time.Second)
	}
}
func (c *SqsConsumer) ProcessMessage(ctx context.Context, msg *types.Message) error {
	fmt.Println("Received message:", *msg.Body)
	service.GetOrders(*msg.Body)
	return nil
}

func (c *SqsConsumer) DeleteMessage(ctx context.Context, receiptHandle *string) {
	_, err := c.SqsClient.DeleteMessage(ctx, &sqs.DeleteMessageInput{
		QueueUrl:      aws.String(c.QueueURL),
		ReceiptHandle: receiptHandle,
	})

	if err != nil {
		log.Println("Error deleting message:", err)
	} else {
		fmt.Println("Message deleted successfully")
	}
}
