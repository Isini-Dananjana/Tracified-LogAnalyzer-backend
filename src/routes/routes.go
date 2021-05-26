package routes

import (
	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/api"
	"github.com/gorilla/mux"
)

func LogRoutes() *mux.Router {
	var router = mux.NewRouter()
	router = mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/{user}/{project}", api.GetAllLog).Methods("GET")
	//router.HandleFunc("/api/{user}/{project}/{logfileName}", api.GetLogFile).Methods("GET")
	router.HandleFunc("/api/uploads/{user}/{project}/{log}", api.HandleLogFileUpload).Methods("POST")
	//router.HandleFunc("/api/uploads/", api.HandleLogFileUpload).Methods("POST")	
	router.HandleFunc("/api/{user}/{project}/{logfileName}", api.GetLogFileContent).Methods("GET")

	/*
		Catch updates on log files
	*/

	router.HandleFunc("/api/updates",api.HandleFileUpdates).Methods("POST")

	return router
}
