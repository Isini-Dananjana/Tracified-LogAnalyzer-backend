package routes

import (
	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/api"
	"github.com/gorilla/mux"
	//"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/middleware"
	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/websocket"

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
	//router.HandleFunc("/api/uploads/{user}/{project}/{log}", api.HandleLogFileUpload).Methods("POST")
	router.HandleFunc("/api/uploads/", api.HandleLogFileUpload).Methods("POST")

	//read the log file content
	//router.HandleFunc("/api/{user}/{project}/{logfileName}", api.GetLogFileContent).Methods("GET")

	//get log file content v2

	router.HandleFunc("/api/v2/content/{fileId}",api.GetLogFileContentv2).Methods("GET")

	//catch the log file updates



	router.HandleFunc("/api/updates",api.HandleFileUpdates).Methods("POST")

	//GetLogsByUserandProject

	router.HandleFunc("/logapi/{user}/{project}",api.GetLogListByUsernProject).Methods("GET")
	router.HandleFunc("/api/logs/getByProject/{user}/{project}",api.GetLogListByUsernProject).Methods("GET")
	router.HandleFunc("/ws",websocket.WSPage).Methods("GET")



	//router.Use(middleware.LoggingMiddleware)


	return router
}
