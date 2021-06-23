package controller

import (
	//"io/ioutil"
	//"archive/zip"

	//"flag"
	"fmt"
	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/datamodels"
	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/repository"
	"io/ioutil"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	//"io/ioutil"

	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/models"
	filestorageHandler "github.com/TharinduBalasooriya/LogAnalyzerBackend/src/util/filestorage"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

var logrepo repository.LogRepository

func LogGetFileContent(user string, project string, log string) LogContent {

	bucket := "leadl/logs/" + user + "/" + project + "/"
	
	/*
		TODO:change extension to config
	*/
	item := log + os.Getenv("BUCKET_ITEM_EXT")

	//fmt.Print(bucket+item)

	object := filestorageHandler.AWS_S3_Object{
		Bucket: bucket,
		Item:   item,
	}

	data := models.Log_GetContent(object, log)

	var dataT = string(data)

	logcontent := LogContent{
		FileName: log,
		Content:  dataT,
	}

	/*
		
		TODO:Handle download time
	*/
	return logcontent

}

const (
	S3_REGION = "ap-south-1"
	S3_BUCKET = "leadl"
)

func LogSaveDetails(userName string, projectName string,logFileName string){

	logfile := datamodels.Log{
		Username: userName,
		LogFileName: logFileName,
		ProjectName: projectName,
		LastUpdate: time.Now().String(),

	}

	//res := []bson.M{}

	 exist,res := logrepo.CheckLogExist(logfile)

	if exist{


	fmt.Println("Already Exist")
	fmt.Println(res)



	}else{

		results,err:=models.Log_Save_Details(logfile);

		if err != nil{
			log.Fatal(err)

		}

		id := results.(primitive.ObjectID);
		fmt.Println("Successfully inserted" + id.String())

	}


	

}

func LogUploadFiles(path string, inputfile multipart.File) {

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

	/*
	 Create a file storage type object
	*/

	//S3 type object
	s3 := filestorageHandler.AWS_S3{
		Session:   s,
		Filepath:  path,
		FileBytes: fileBytes,
	}

	models.Log_uploadFiles(s3)

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
