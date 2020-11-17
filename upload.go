package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/byroni/rufous/awshelper"
)

func upload(sess *session.Session) error {
	cyan.Println("Uploading project folder to lambda")

	svc := dynamodb.New(sess)

	input := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"jobID": {
				S: aws.String("testerJobId"),
			},
		},
		TableName: aws.String("rufous-dev-table"),
	}

	result, err := svc.GetItem(input)
	if err != nil {
		err := awshelper.HandleDynamoError(err)
		if err != nil {
			log.Fatal(err)
		}
		return nil
	}

	fmt.Println(result)
	return nil
}
