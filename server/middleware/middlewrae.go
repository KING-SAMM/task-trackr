package middleware

import (
	"context"
	"fmt"
	"os"
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorrila/mux"

	// To work with environment file
	"github.com/joho/godotenv"
	
	//Monngodb drivers
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

// Initialisation function
func init() {
	loadTheEnv()
	createDBInstance()
}

func loadTheEnv() {
	err := godotenv.Load('.env')

	if err != nil {
		log.Fatal("Error loading the .env file")
	}
}

func createDBInstance() {
	connectionString := os.Getenv('DB_URI')
	dbName := os.Getenv('DB_NAME')
	collectionName := os.Getenv('DB_COLLECTION_NAME')

	clientOptions := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	// Handle the error
	if err != nil {
		log.Fatal(err)
	}

	// With the client, ping and check that all is OK
	err := client.Ping(context.TODO(), nil)

	// Handle the error from ping
	if err != nil {
		log.Fatal(err)
	}

	// If all is OK
	fmt.Println("Connected to mongoDB")

	collection := client.Database(dbName).Collection(collectionName)
	fmt.Println("'collection' instance created")
}

func GetAllTasks(w http.ResponseWriter, r *http.Request) {

}

func CreateTask(w http.ResponseWriter, r *http.Request) {

}

func TaskComplete(w http.ResponseWriter, r *http.Request) {

}

func UndoTask(w http.ResponseWriter, r *http.Request) {

}

func DeleteTask(w http.ResponseWriter, r *http.Request) {

}

func DeleteAllTasks(w http.ResponseWriter, r *http.Request) {

}