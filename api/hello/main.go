package main

import (
	"context"
	"log"
	"net/http"

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
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("ap-southeast-1"))

	if err != nil {
		log.Fatal("Config: cannot load configuration")
	}

	client := dynamodb.NewFromConfig(cfg)

	transactionRepo := TransactionRepository{client}

	transaction := Transaction{
		Id:       "1",
		Username: "ivan",
		Amount:   "1234",
	}

	transactionRepo.Create(transaction)

	// session, err := session.NewSession()
	// client := dynamodb.New(session)
	// repo := TransactionRepository{client}
	// repo.Create(transaction)
	// var buf bytes.Buffer

	// body, err := json.Marshal(map[string]interface{}{
	// 	"message": "Go Serverless v1.0! Your function executed successfully!",
	// })
	// if err != nil {
	// 	return Response{StatusCode: 404}, err
	// }
	// json.HTMLEscape(&buf, body)

	// resp := Response{
	// 	StatusCode:      200,
	// 	IsBase64Encoded: false,
	// 	Body:            buf.String(),
	// 	Headers: map[string]string{
	// 		"Content-Type": "application/json",
	// 	},
	// }
	return Response{StatusCode: http.StatusOK}, nil
}

func main() {
	lambda.Start(Handler)
}
