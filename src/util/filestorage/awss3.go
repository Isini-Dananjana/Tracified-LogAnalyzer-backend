package filestorageHandler

import (
	"bytes"
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type AWS_S3 struct {
	Session         *session.Session
	Filepath  string
	FileBytes []byte
}
type AWS_S3_Object struct {
	Bucket string
	Item string
}
const (
	S3_REGION = "ap-south-1"
	S3_BUCKET = "leadl"
)

func (fs AWS_S3) AddFiles() error {
	_, err := s3.New(fs.Session).PutObject(&s3.PutObjectInput{
		Bucket: aws.String(S3_BUCKET),
		Key:    aws.String(fs.Filepath + os.Getenv("ARCHIVED_EXT")),
		ACL:    aws.String("private"),
		Body:   bytes.NewReader(fs.FileBytes),
		//ContentLength:        aws.Int64(len(fileBytes)),
		ContentType:          aws.String(http.DetectContentType(fs.FileBytes)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
	})
	return err
}


func (obj AWS_S3_Object) GetContent() error{


	file, err := os.Create("temp/" + obj.Item)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	// Initialize a session in us-west-2 that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String("ap-south-1")},
	)

	downloader := s3manager.NewDownloader(sess)

	numBytes, err := downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(obj.Bucket),
			Key:    aws.String(obj.Item),
		})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Downloaded", file.Name(), numBytes, "bytes")

 	return err;
}
