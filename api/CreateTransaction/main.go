package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse
type Request events.APIGatewayProxyRequest

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context, req Request) (Response, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())

	if err != nil {
		log.Fatal("Config: cannot load configuration")
	}

	dynamodbClient := dynamodb.NewFromConfig(cfg)

	transactionRepo := TransactionRepository{dynamodbClient}

	transaction := Transaction{
		Timestamp: strconv.FormatInt(time.Now().Unix(), 10),
	}

	//parse request
	err = json.Unmarshal([]byte(req.Body), &transaction)

	if err != nil {
		return Response{StatusCode: http.StatusBadRequest}, err
	}

	err = transactionRepo.Create(transaction)

	if err != nil {
		return Response{StatusCode: http.StatusInternalServerError}, err
	}
	// json.HTMLEscape(&buf, body)

	// resp := Response{
	// 	StatusCode:      200,
	// 	IsBase64Encoded: false,
	// 	Body:            buf.String(),
	// 	Headers: map[string]string{
	// 		"Content-Type": "application/json",
	// 	},
	// }
	return Response{StatusCode: http.StatusCreated}, nil
}

func main() {
	lambda.Start(Handler)
}
