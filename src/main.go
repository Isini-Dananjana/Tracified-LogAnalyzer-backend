package main

import (
	"log"
	"net/http"

	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/routes"
	"github.com/joho/godotenv"
	
)


func LoadEnv(){

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	
}
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
