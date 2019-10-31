package main

import "fmt"
import "github.com/aws/aws-sdk-go/aws"
import "github.com/aws/aws-sdk-go/aws/session"
import "github.com/aws/aws-sdk-go/service/sns"
import "context"
import "github.com/aws/aws-lambda-go/lambda"
import "github.com/aws/aws-lambda-go/events"
import "encoding/json"

type SNSMessage struct {
	OriginationNumber          string `json:"originationNumber"`
	DestinationNumber          string `json:"DestinationNumber"`
	MessageKeyword             string `json:"messageKeyword"`
	MessageBody                string `json:"messageBody"`
	InboundMessageId           string `json:"inboundMessageId"`
	PreviousPublishedMessageId string `json:"previousPublishedMessageId"`
}

func main() {
	lambda.Start(HandleRequest)
}

func HandleRequest(ctx context.Context, snsEvent events.SNSEvent) {
	for _, record := range snsEvent.Records {
		snsRecord := record.SNS
		fmt.Printf("[%s %s] Message = %s \n", record.EventSource, snsRecord.Timestamp, snsRecord.Message)
		message := SNSMessage{}
		_ = json.Unmarshal([]byte(snsRecord.Message), &message)

		var phone string = message.OriginationNumber
		fmt.Printf("OK got phone number from message: %s\n", phone)
	}

	return
}
