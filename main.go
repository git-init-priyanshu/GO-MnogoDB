package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"go-MongoDB/routes"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var mongoClient *mongo.Client
var uri string
var db string
var collection string

func init() {
	// Loading env file
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file provided")
	}

	// Getting env variables
	uri = os.Getenv("MONGO_URI")
	db = os.Getenv("DB_NAME")
	collection = os.Getenv("COLLECTION_NAME")

	if uri == "" || db == "" || collection == "" {
		log.Fatal("Env variables not provided")
	}

	// Connecting to mongodb
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("Connection failed", err)
	}
	// Pinging the db
	if err := client.Ping(context.Background(), readpref.Primary()); err != nil {
		log.Fatal("Ping failed", err)
	}

	fmt.Println("Connected to DB")
	mongoClient = client // ?
}

func main() {
	defer func() {
		if err := mongoClient.Disconnect(context.TODO()); err != nil {
			log.Fatal("Cannot dissconnect", err)
		}
	}()

	// Get collection
	coll := mongoClient.Database(db).Collection(collection)

  // Get employee routes
  empRoutes := routes.EmpService{MongoCollection: coll}

	// Init router
	router := mux.NewRouter()

	// Routes
	router.HandleFunc("/health", healthHandler).Methods(http.MethodGet)
	router.HandleFunc("/employee/{id}", empRoutes.GetEmployee).Methods(http.MethodGet)
	router.HandleFunc("/employees", empRoutes.GetAllEmployees).Methods(http.MethodGet)
	router.HandleFunc("/employee", empRoutes.CreateEmployee).Methods(http.MethodPost)
	router.HandleFunc("/employee/{id}", empRoutes.UpdateEmployee).Methods(http.MethodPut)
	router.HandleFunc("/employees/{id}", empRoutes.DeleteEmpoyee).Methods(http.MethodDelete)
	router.HandleFunc("/employees", empRoutes.DeleteAllEmpoyees).Methods(http.MethodDelete)

	fmt.Println("Server is running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK..."))
}
