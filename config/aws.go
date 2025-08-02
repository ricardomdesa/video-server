package config

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

func NewAWSSession(cfg *Env) (*session.Session, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"), // A região é necessária, mas não importa no LocalStack.
		Endpoint: aws.String(cfg.S3Endpoint),
		S3ForcePathStyle: aws.Bool(true), // Necessário para compatibilidade com LocalStack
		Credentials: credentials.NewStaticCredentials(cfg.AWSAccessKeyID, cfg.AWSSecretAccessKey, ""),
	})
	if err != nil {
		return nil, err
	}

	return sess, nil
}