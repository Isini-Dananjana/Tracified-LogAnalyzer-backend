package log_controller

import (
	//"io/ioutil"
	//"archive/zip"
	"archive/zip"
	"bytes"

	//"flag"
	"fmt"
	"io"
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
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type Loglist struct {
	UserName string   `json:"userName"`
	Project  string   `json:"project"`
	Logs     []string `json:"logs"`
}

type LogContent struct {
	FileName string `json:"filename"`
	Content  string `json:"content"`
}

func GetFileList(user string, project string) Loglist {

	var files []string
	// user :="tharindu"
	//project := "project1"
	root := "logs/" + user + "/" + project
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

func GetLogfileContent(user string, project string, Logs string) LogContent {

	root := "logs/" + user + "/" + project

	data, err := ioutil.ReadFile(root + "/" + Logs + ".txt")
	if err != nil {
		panic(err)
	}

	var dataT = string(data)

	logcontent := LogContent{
		FileName: Logs,
		Content:  dataT,
	}

	return logcontent

}

func unzipLogfile(Logs string) {

	fmt.Println("temp/" + Logs + ".txt.zip")

	zipReader, err := zip.OpenReader("temp/" + Logs + ".txt.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer zipReader.Close()

	// Iterate through each file/dir found in
	for _, file := range zipReader.Reader.File {
		// Open the file inside the zip archive
		// like a normal file
		zippedFile, err := file.Open()
		if err != nil {
			log.Fatal(err)
		}
		defer zippedFile.Close()

		// Specify what the extracted file name should be.
		// You can specify a full path or a prefix
		// to move it to a different directory.
		// In this case, we will extract the file from
		// the zip to a file of the same name.
		targetDir := "./temp"
		extractedFilePath := filepath.Join(
			targetDir,
			file.Name,
		)

		// Extract the item (or create directory)
		if file.FileInfo().IsDir() {
			// Create directories to recreate directory
			// structure inside the zip archive. Also
			// preserves permissions
			log.Println("Creating directory:", extractedFilePath)
			os.MkdirAll(extractedFilePath, file.Mode())
		} else {
			// Extract regular file since not a directory
			log.Println("Extracting file:", file.Name)

			// Open an output file for writing
			outputFile, err := os.OpenFile(
				extractedFilePath,
				os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
				file.Mode(),
			)
			if err != nil {
				log.Fatal(err)
			}
			defer outputFile.Close()

			// "Extract" the file by copying zipped file
			// contents to the output file
			_, err = io.Copy(outputFile, zippedFile)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func GetLogfileContentTest(user string, project string, Logs string) LogContent {

	bucket := "leadl/logs/" + user + "/" + project + "/"
	//bucket := "leadl/logs/Isini/99xIT/"
	/*
		TODO:change extension to config
	*/
	item := Logs + ".txt.zip"

	//fmt.Print(bucket+item)

	file, err := os.Create("temp/" + item)
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
			Bucket: aws.String(bucket),
			Key:    aws.String(item),
		})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Downloaded", file.Name(), numBytes, "bytes")

	unzipLogfile(Logs)

	data, err := ioutil.ReadFile("temp/" + Logs + ".txt")
	if err != nil {
		panic(err)
	}

	var dataT = string(data)

	logcontent := LogContent{
		FileName: Logs,
		Content:  dataT,
	}

	/*
		TODO:Move download to model class
		TODO:Handle download time
	*/
	return logcontent

}

//  func DownloadObject(sess *session.Session, filename *string, bucket *string) error {
//     // snippet-start:[s3.go.download_object.create]

//     file, err := os.Create("temp/"+*filename)
//     // snippet-end:[s3.go.download_object.create]
//     if err != nil {
//         return err
//     }

//      defer file.Close()

// 	//clean tempory files

//     // snippet-start:[s3.go.download_object.call]
//     downloader := s3manager.NewDownloader(sess)

//     _, err = downloader.Download(file,
//         &s3.GetObjectInput{
//             Bucket: bucket,
//             Key:    filename,
//         })
//     // snippet-end:[s3.go.download_object.call]
//     if err != nil {
//         return err
//     }

//     return nil
// }

//TODO fill these in!
const (
	S3_REGION = "ap-south-1"
	S3_BUCKET = "leadl"
)

func UplaodLogFiles(path string, inputfile multipart.File) {

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

	/*
		Create a abstract inteface (FileStorage Interface)
	*/
	err = AddFileToS3(s, path, fileBytes)
	if err != nil {
		log.Fatal(err)
	}
}

// // AddFileToS3 will upload a single file to S3, it will require a pre-built aws session
// // and will set file info like content type and encryption on the uploaded file.
func AddFileToS3(s *session.Session, fileDir string, fileBytes []byte) error {

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
		Bucket: aws.String(S3_BUCKET),
		Key:    aws.String(fileDir + ".zip"),
		ACL:    aws.String("private"),
		Body:   bytes.NewReader(fileBytes),
		//ContentLength:        aws.Int64(len(fileBytes)),
		ContentType:          aws.String(http.DetectContentType(fileBytes)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
	})
	return err
}

type Update struct {
	UserName    string `json:"userName"`
	ProjectName string `json:"project"`
	Data        string `json:"data"`
}

func HandleUpdateData(update Update) {

	fmt.Println(update.UserName)
	fmt.Println(update.ProjectName)
	fmt.Println(update.Data)

}
