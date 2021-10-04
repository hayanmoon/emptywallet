package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type Transaction struct {
	Id       string `dynamodbav:"id"`
	Username string `dynamodbav:"username"`
	Amount   string `dynamodbav:"amount"`
}

// type DbRepo interface {
// 	Create(t Transaction) error
// }

type TransactionRepository struct {
	db *dynamodb.Client
}

func (t *TransactionRepository) Create(transaction Transaction) {
	//condition for existing name attribute
	nameNotExist := expression.AttributeNotExists(expression.Name("username"))
	builder := expression.NewBuilder().WithCondition(nameNotExist)
	exp, _ := builder.Build()

	item, _ := attributevalue.MarshalMap(transaction)

	input := &dynamodb.PutItemInput{
		TableName:                aws.String("transaction"),
		ConditionExpression:      exp.Condition(), // fail put if username already exist
		ExpressionAttributeNames: exp.Names(),
		Item:                     item,
	}

	something, err := t.db.PutItem(context.TODO(), input)
	fmt.Println(something, err)

	if err != nil {
		fmt.Println("Transaction: cannot save transaction")
	}
}
