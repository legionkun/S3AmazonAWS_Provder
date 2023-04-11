package usercase

import (
	"bytes"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3Provider struct {
	BucketName string
	Region     string
	ApiKey     string
	SecrecKey  string
	Domain     string
	session    *session.Session
}

func NewS3Provider(bucketName string, region string, apikey string, secrecKey string, domain string) *S3Provider {
	provider := S3Provider{
		BucketName: bucketName,
		Region:     region,
		ApiKey:     apikey,
		SecrecKey:  secrecKey,
		Domain:     domain,
	}

	S3Session, err := session.NewSession(&aws.Config{
		Region: aws.String(provider.Region),
		Credentials: credentials.NewStaticCredentials(
			provider.ApiKey,
			provider.SecrecKey,
			"",
		),
	})
	if err != nil {
		log.Fatalln(err)
	}
	provider.session = S3Session
	return &provider
}

func (provider *S3Provider) SaveFileUpload(data []byte, dst string) map[string]interface{} {
	fileBytes := bytes.NewReader(data)
	fileType := http.DetectContentType(data)
	_, err := s3.New(provider.session).PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(provider.BucketName),
		Key:         aws.String(dst),
		ACL:         aws.String("private"),
		ContentType: aws.String(fileType),
		Body:        fileBytes,
	})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			errAWS := map[string]interface{}{
				"Code":    aerr.Code(),
				"Error":   aerr.Error(),
				"Message": aerr.Message(),
			}
			log.Println(err.Error())
			return errAWS
		}
	}
	return nil
}

func (provider *S3Provider) SaveFileUploadforClientbyURL(dst string) (string, error) {
	data, err := s3.New(provider.session).PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String(provider.BucketName),
		Key:    aws.String(dst),
		ACL:    aws.String("private"),
	})
	if err != nil {
		log.Fatalln(err)
	}
	return data.Presign(time.Second * 60)
}

func SetUps3ProviderConfig() *S3Provider {
	BucketName := os.Getenv("S3BucketName")
	Region := os.Getenv("S3Region")
	ApiKey := os.Getenv("S3ApiKey")
	Secrect := os.Getenv("S3SecrecKey")
	Domain := os.Getenv("S3Domain")
	return NewS3Provider(BucketName, Region, ApiKey, Secrect, Domain)
}
