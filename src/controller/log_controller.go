package log_controller

import (
	//"io/ioutil"
	//"archive/zip"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	//"io/ioutil"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)
type Loglist struct {
	UserName string   `json:"userName"`
	Project  string   `json:"project"`
	Logs     []string `json:"logs"`
}

type LogContent struct {

    FileName string `json:"filename"`
	Content string  `json:"content"`
	
}



func GetFileList(user string, project string) Loglist {
	
	
 
var files []string
   // user :="tharindu"
    //project := "project1"
    root := "logs/"+user+"/"+project
    //root:= "../logs/" + user + "/" + project
    err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
        files = append(files, info.Name())
        return nil
    })
    if err != nil {
        panic(err)
    }
    

	loglist := Loglist{
		UserName: user,
		Project:  project,
		Logs:     files,
	}

   
	return loglist
}

func GetLogfileContent(user string, project string ,Logs string) LogContent{

   
   root := "logs/"+user+"/"+project
   
    data, err := ioutil.ReadFile(root+"/"+Logs + ".txt")
	if err != nil {
	panic(err)
	 }

	var dataT = string(data)

	
     logcontent := LogContent{
            FileName: Logs,
            Content: dataT,

     }

     return logcontent

	
}


//TODO fill these in!
const (
	S3_REGION = "ap-south-1"
	S3_BUCKET = "loganalyzertracifi"
)

func UplaodLogFiles(path string , inputfile multipart.File)  {

	// byte array
	fileBytes, err := ioutil.ReadAll(inputfile)
	 if err != nil {
	fmt.Println(err)

	}

	// Create a single AWS session (we can re use this if we're uploading many files)
	s, err := session.NewSession(&aws.Config{Region: aws.String(S3_REGION)})
	if err != nil {
		log.Fatal(err)
	}
	

	// Upload



	err = AddFileToS3(s, path, fileBytes)
	if err != nil {
		log.Fatal(err)
	}
}

// // AddFileToS3 will upload a single file to S3, it will require a pre-built aws session
// // and will set file info like content type and encryption on the uploaded file.
func AddFileToS3(s *session.Session, fileDir string ,fileBytes[] byte) error {

	//     // Open the file for use
	// file, err := os.Open(fileDir)
	// if err != nil {
	// 	return err
	// }
	// defer file.Close()

	//     // Get file size and read the file content into a buffer
	// fileInfo, _ := file.Stat()
	// var size int64 = fileInfo.Size()
	// buffer := make([]byte, size)
	// file.Read(buffer)


	
	//     // Config settings: this is where you choose the bucket, filename, content-type etc.
	//     // of the file you're uploading.
	_, err := s3.New(s).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(S3_BUCKET),
		Key:                  aws.String(fileDir+".zip"),
		ACL:                  aws.String("private"),
		Body:                 bytes.NewReader(fileBytes),
		//ContentLength:        aws.Int64(len(fileBytes)),
		ContentType:          aws.String(http.DetectContentType(fileBytes)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
	})
	return err
}



