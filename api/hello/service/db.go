package service

import (
	"fmt"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

type Transaction struct {
	Id     string
	Name   string
	Amount string
}

// type DbRepo interface {
// 	Create(t Transaction) error
// }

type TransactionRepository struct {
	Db *dynamodb.DynamoDB
}

func (t *TransactionRepository) Create(transaction Transaction) {
	//condition for existing name attribute
	nameNotExist := expression.AttributeNotExists(expression.Name("Name"))
	builder := expression.NewBuilder().WithCondition(nameNotExist)
	exp, err := builder.Build()

	if err != nil {
		fmt.Println("Transaction: cannot build expression")
	}

	// condition.
	input := &dynamodb.PutItemInput{
		ConditionExpression: exp.Condition(),
	}

	_, err = t.Db.PutItem(input)

	if err != nil {
		fmt.Println("Transaction: cannot save transaction")
	}
}
