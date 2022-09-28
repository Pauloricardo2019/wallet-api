package s3

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"log"
	"wallet-api/adapter/config"
	"wallet-api/internal/model"
)

func UploadAlbumFile(userID string, albumID string, fileName string, fileBody []byte) (*string, error) {

	cfg := config.GetConfig()

	client, err := getClient()
	if err != nil {
		return nil, err
	}

	_, err = client.PutObject(context.TODO(),
		&s3.PutObjectInput{
			Bucket: aws.String(cfg.Bucket),
			Key:    aws.String("albums/" + userID + "/" + albumID + "/" + fileName),
			Body:   bytes.NewReader(fileBody),
			ACL:    "public-read",
		})
	if err != nil {
		return nil, err
	}
	//urlHost := fmt.Sprintf("%v.%v.amazonaws.com", &cfg.Bucket, &cfg.Region)
	urlPhoto := fmt.Sprintf("/albums/%s/%s/%s", userID, albumID, fileName)

	return &urlPhoto, nil
}

func UploadUserFile(fileName string, fileBody []byte) (*string, error) {
	cfg := config.GetConfig()

	client, err := getClient()
	if err != nil {
		return nil, err
	}

	_, err = client.PutObject(context.TODO(),
		&s3.PutObjectInput{
			Bucket: aws.String(cfg.Bucket),
			Key:    aws.String("users/profile/" + fileName),
			Body:   bytes.NewReader(fileBody),
			ACL:    "public-read",
		})
	if err != nil {
		return nil, err
	}
	//urlHost := fmt.Sprintf("%v.%v.amazonaws.com", &cfg.Bucket, &cfg.Region)
	urlPhoto := fmt.Sprintf("/user/profile/%s", fileName)

	return &urlPhoto, nil
}

func ListFiles(photos []model.Photo, userID, albumID string) error {
	client, err := getClient()
	if err != nil {
		return err
	}
	cfg := config.GetConfig()

	// Get the first page of results for ListObjectsV2 for a bucket
	output, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String(cfg.Bucket),
		Prefix: aws.String("albums/" + userID + "/" + albumID),
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("first page results:")
	for _, object := range output.Contents {
		//log.Printf("key=%s size=%d", aws.ToString(object.Key), object.Size)
		for _, photo := range photos {
			if *photo.UrlImage != aws.ToString(object.Key) {
				return err
			}
		}
	}

	return nil
}

func getClient() (*s3.Client, error) {
	cfg := config.GetConfig()

	awsCfg, err := awsConfig.LoadDefaultConfig(context.TODO(),
		awsConfig.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID:     cfg.AwsAccessKeyID,
				SecretAccessKey: cfg.AwsSecretAccessKey,
			},
		}))

	if err != nil {
		return nil, err
	}
	awsCfg.Region = cfg.Region

	// Create an Amazon S3 service client
	client := s3.NewFromConfig(awsCfg)
	return client, err
}
