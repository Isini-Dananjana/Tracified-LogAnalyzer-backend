package api

//API to handle main log operation
//Get file name , Get content

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/controller"
)



func GetAllLog(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	logs := log_controller.GetFileList(params["user"],params["project"])
	json.NewEncoder(w).Encode(logs)
	
}

func GetLogFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	log:= log_controller.GetLogfileContent(params["user"],params["project"],params["logfileName"])
	json.NewEncoder(w).Encode(log)


}



