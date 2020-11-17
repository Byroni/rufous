package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/byroni/rufous/awshelper"
)

func list(sess *session.Session) error {
	svc := dynamodb.New(sess)

	input := &dynamodb.ScanInput{
		TableName: aws.String("rufous-dev-table"),
	}

	result, err := svc.Scan(input)
	if err != nil {
		err := awshelper.HandleDynamoError(err)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println(result)
	return nil
}
