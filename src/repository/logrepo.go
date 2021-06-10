package repository

import (
	"context"
	"fmt"
	"log"

	"time"

	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/datamodels"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// type LogRepository interface {
// 	SaveLog(log datamodels.Log)
// 	GetAll()
// 	GetByUserName(user string)
// 	GetByProjectname(projectname string)
// }
//
var log_collection = new (mongo.Collection)
const LogsCollection = "Logs"
func init(){

	fmt.Println("Database Connection Established")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	log_collection = client.Database("leadldb").Collection(LogsCollection)


}
type LogRepository struct{}
func (l *LogRepository) SaveLog(log datamodels.Log)(interface{}, error){

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, err := log_collection.InsertOne(ctx, log)
	fmt.Println("Inserted a single document: ", result.InsertedID)
	return result.InsertedID, err


	
}
