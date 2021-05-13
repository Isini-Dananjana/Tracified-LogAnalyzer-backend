package main

import (
	"log"
	"net/http"

	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/routes"
)

func main() {

	//Starting the API server
	router := routes.LogRoutes()
	http.Handle("/api/", router)

	//Starting the FileServer
	//fs := http.FileServer(http.Dir("server/webapps/play_maths"))
	//http.Handle("/", fs)

	log.Println("Listening...")
	log.Fatal(http.ListenAndServe(":3000", nil))

}
