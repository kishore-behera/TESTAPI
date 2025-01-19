package main

import (
    "context"
    "log"
    "net/http"
    "os"
    "testapi/handlers"
    "github.com/gorilla/mux"
    "github.com/joho/godotenv"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
    // Load .env file
    if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }

    // Connect to MongoDB
    client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
    if err != nil {
        log.Fatal(err)
    }
    defer client.Disconnect(context.Background())

    // Initialize collection
    collection := client.Database(os.Getenv("DATABASE_NAME")).Collection("users")
    handlers.InitHandler(collection)

    // Initialize router
    router := mux.NewRouter()

    // Define routes
    router.HandleFunc("/users", handlers.CreateUser).Methods("POST")
    router.HandleFunc("/users", handlers.GetUsers).Methods("GET")
    router.HandleFunc("/users/{id}", handlers.GetUser).Methods("GET")
    router.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")
    router.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")

    // Start server
    log.Println("Server is running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}
