package awshelper

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/dynamodb"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

func InitSessionCredentials() *session.Session {
	sess, err := session.NewSession(&aws.Config{
		// TODO: be able to configure region
		Region:      aws.String("us-east-1"),
		Credentials: credentials.NewSharedCredentials("", "rufous"),
	})
	if err != nil {
		log.Fatal(err)
	}

	return sess
}

func HandleDynamoError(err error) error {
	aerr, ok := err.(awserr.Error)
	if ok {
		switch aerr.Code() {
		case dynamodb.ErrCodeProvisionedThroughputExceededException:
			fmt.Println(dynamodb.ErrCodeProvisionedThroughputExceededException, aerr.Error())
		case dynamodb.ErrCodeResourceNotFoundException:
			fmt.Println(dynamodb.ErrCodeResourceNotFoundException, aerr.Error())
		case dynamodb.ErrCodeRequestLimitExceeded:
			fmt.Println(dynamodb.ErrCodeRequestLimitExceeded, aerr.Error())
		case dynamodb.ErrCodeInternalServerError:
			fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
		default:
			fmt.Println(aerr.Error())
		}
	} else {
		fmt.Println(err.Error())
		return err
	}
	return aerr
}
