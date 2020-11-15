package awshelper

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

func InitSessionCredentials() *session.Session {
	sess, err := session.NewSession(&aws.Config{
		// TODO: be able to configure region
		Region:      aws.String("us-east-2"),
		Credentials: credentials.NewSharedCredentials("", "rufous"),
	})
	if err != nil {
		log.Fatal(err)
	}

	return sess
}
