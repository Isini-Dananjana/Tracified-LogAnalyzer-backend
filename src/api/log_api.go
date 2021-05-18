package api

//API to handle main log operation
//Get file name , Get content

import (
	"encoding/json"
	"fmt"
	//"io/ioutil"
	//"log"
	"net/http"
	//"os"
	

	log_controller "github.com/TharinduBalasooriya/LogAnalyzerBackend/src/controller"
	"github.com/gorilla/mux"
)

func GetAllLog(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	logs := log_controller.GetFileList(params["user"], params["project"])
	json.NewEncoder(w).Encode(logs)

}

func GetLogFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	log := log_controller.GetLogfileContent(params["user"], params["project"], params["logfileName"])
	json.NewEncoder(w).Encode(log)

}

func HandleLogFileUpload(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	fmt.Println("File Upload Endpoint Hit")

	r.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	

	// read all of the contents of our uploaded file into a
	//
	// }

	//aws upload path
	fullFilePath := "logs/" + params["user"] + "/" +params["project"] + "/" + params["log"]
	//open new file
	// filetowrite, err := os.OpenFile(
	// 	fullFilePath,
	// 	os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
	// 	0666,
	// )
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()


	// bytesWritten, err := filetowrite.Write(fileBytes)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Printf("Wrote %d bytes.\n", bytesWritten)
	// fmt.Fprintf(w, "Successfully Uploaded File\n")

	// fmt.Println(params["log"])

	log_controller.UplaodLogFiles(fullFilePath,file) 

}
