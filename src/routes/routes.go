package routes

import (
	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/api"
	"github.com/gorilla/mux"
	//"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/middleware"

)

func LogRoutes() *mux.Router {
	var router = mux.NewRouter()
	router = mux.NewRouter().StrictSlash(true)

	//Get All Log files

	//TODO :  Configure to work with mongodb

	/*
	*
	* TODO:Proper api naming convention
	 */

	router.HandleFunc("/api/logs/{user}/", api.GetAllLog).Methods("GET")

	//getAllProjetcs
	router.HandleFunc("/api/projects/{user}/", api.GetAllProjects).Methods("GET")


	//upload file
	router.HandleFunc("/api/uploads/{user}/{project}/{log}", api.HandleLogFileUpload).Methods("POST")

	//read the log file content
	router.HandleFunc("/api/{user}/{project}/{logfileName}", api.GetLogFileContent).Methods("GET")

	//catch the log file updates



	router.HandleFunc("/api/updates",api.HandleFileUpdates).Methods("POST")

	//router.Use(middleware.LoggingMiddleware)


	return router
}
