package routes

import (
	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/api"
	"github.com/gorilla/mux"
)

func LogRoutes() *mux.Router {
	var router = mux.NewRouter()
	router = mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/{user}/{project}", api.GetAllLog).Methods("GET")
	router.HandleFunc("/api/{user}/{project}/{logfileName}",api.GetLogFile).Methods("GET")

	return router
}
