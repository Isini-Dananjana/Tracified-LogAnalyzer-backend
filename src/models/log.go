package models

import (

	//Importing file storage utility
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	filestorageHandler "github.com/TharinduBalasooriya/LogAnalyzerBackend/src/util/filestorage"
)

/*
This package containes all business logic log file

*/

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

/*
	Upload a file
*/
func Log_uploadFiles(fs filestorageHandler.FileStorage) {

	err := fs.AddFiles() // calling add files function of the file storage
	if err != nil {
		log.Fatal(err)
	}

}

func Log_GetContent(file_object filestorageHandler.File, logfileName string) []byte {

	fileExtension := os.Getenv("FILE_EXT")

	err := file_object.GetContent()
	if err != nil {
		log.Fatal(err)
	}
	unzipLogfile(logfileName)

	data, err := ioutil.ReadFile("temp/" + logfileName + fileExtension)
	if err != nil {
		panic(err)
	}

	return data

}
