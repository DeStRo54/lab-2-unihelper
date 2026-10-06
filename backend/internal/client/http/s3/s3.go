package s3

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Config struct {
	Secret_key string `envconfig:"SECRET_KEY"`
	Access_key string `envconfig:"ACCESS_KEY"`
	Bucketname string `envconfig:"BUCKETNAME"`
	Region     string `envconfig:"REGION"`
	Endpoint   string `envconfig:"ENDPOINT"`
}

var ErrNotConfigured = errors.New("S3 storage is not configured")

type S3Client struct {
	client *s3.Client
	config *S3Config
}

func NewS3Client(s3Cfg *S3Config) *S3Client {
	if s3Cfg.Access_key == "" || s3Cfg.Secret_key == "" || s3Cfg.Bucketname == "" {
		return &S3Client{config: s3Cfg}
	}
	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion(s3Cfg.Region),
		config.WithBaseEndpoint(s3Cfg.Endpoint),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(s3Cfg.Access_key, s3Cfg.Secret_key, "")),
	)
	if err != nil {
		panic(err)
	}
	s3Client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.UsePathStyle = true
	})
	_, err = s3Client.HeadBucket(context.TODO(), &s3.HeadBucketInput{
		Bucket: &s3Cfg.Bucketname,
	})
	if err != nil {
		panic(err)
	}

	return &S3Client{client: s3Client, config: s3Cfg}
}
