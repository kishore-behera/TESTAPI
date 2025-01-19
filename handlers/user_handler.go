package handlers

import (
    "context"
    "encoding/json"
    "net/http"
    "testapi/models"
    "github.com/gorilla/mux"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection

func InitHandler(c *mongo.Collection) {
    collection = c
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var user models.User
    json.NewDecoder(r.Body).Decode(&user)
    result, err := collection.InsertOne(context.Background(), user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(result)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var users []models.User
    cursor, err := collection.Find(context.Background(), bson.M{})
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer cursor.Close(context.Background())
    for cursor.Next(context.Background()) {
        var user models.User
        cursor.Decode(&user)
        users = append(users, user)
    }
    json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    id, _ := primitive.ObjectIDFromHex(params["id"])
    var user models.User
    err := collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    id, _ := primitive.ObjectIDFromHex(params["id"])
    var user models.User
    json.NewDecoder(r.Body).Decode(&user)
    update := bson.M{
        "$set": bson.M{
            "name":  user.Name,
            "email": user.Email,
            "age":   user.Age,
        },
    }
    result, err := collection.UpdateOne(context.Background(), bson.M{"_id": id}, update)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(result)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    id, _ := primitive.ObjectIDFromHex(params["id"])
    result, err := collection.DeleteOne(context.Background(), bson.M{"_id": id})
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(result)
}
