package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mongo_db/models"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const connectionString = "mongodb+srv://<USERNAME>:<PASSWORD>@cluster0.oflzy5n.mongodb.net/?appName=Cluster0"
const dbName = "netflix"    // database name
const colName = "watchList" // collection name

// MOST IMPORTANT
/*We create a global MongoDB collection reference so all controller functions can directly read/write to that specific collection. */
var collection *mongo.Collection

// connect with mongo db
/* init is a specialized method that only run one time when the entire application is going to be executed
 */
func init() {
	client, err := mongo.Connect(options.Client().ApplyURI(connectionString))
	CheckError(err)
	fmt.Println("Mongodb connection suceess")
	// now i need a reference of the database buz i have dbName and colName, for this we have created "collection line 16"
	collection = client.Database(dbName).Collection(colName)
	fmt.Println("Collection instance is ready")
}

// insert one move in database
// insertOneMovie func name is small letter so we are not exporting it
func insertOneMovie(movie models.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	CheckError(err)
	// print the id of the movie which is inserted
	fmt.Println("The id of the recently added movie is ", inserted.InsertedID)
}

// update one movie
func updateOneMovie(movieId string) {
	// how to convert this movieId into something that mongodb understand buz in mongodb id => _id
	id, err := primitive.ObjectIDFromHex(movieId) // convert string in objectid which is understood by the mongodb
	CheckError(err)
	filter := bson.M{"_id": id}

	update := bson.M{"$set": bson.M{
		"watched": true,
	}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	CheckError(err)
	fmt.Println(result.ModifiedCount)

}

// delete one movie
func deleteOneMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)
	CheckError(err)

	filter := bson.M{"_id": id}

	result, err := collection.DeleteOne(context.Background(), filter)
	CheckError(err)
	fmt.Println("Deleted item :", result.DeletedCount)

}

// delete all movie

func deleteAllMovie() {
	filter := bson.D{{}}
	result, err := collection.DeleteMany(context.Background(), filter)
	CheckError(err)
	fmt.Println(result.DeletedCount)
}

// get all the movies from mongodb

func getAllMovies() []bson.M {
	// cursor is a type of mongodb object on which we have to loop through
	// cursor is not the actual value rather than a pointer
	cursor, err := collection.Find(context.Background(), bson.M{})
	CheckError(err)
	defer cursor.Close(context.Background())
	var movies []bson.M
	for cursor.Next(context.Background()) {
		var eachValue bson.M
		if err = cursor.Decode(&eachValue); err != nil {
			log.Fatal(err)
		}
		fmt.Println(eachValue) // type will bson
		movies = append(movies, eachValue)
	}
	return movies
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

// Now we will create actual helper/ controller files

func GetMyAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	// first decode the data send by user in json form using newdecorder
	w.Header().Set("content-Type", "application/json")
	w.Header().Set("Allow-Content-Allow-Methods", "POST")

	var movie models.Netflix

	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)

	// encode the data in json form and send it back
	json.NewEncoder(w).Encode(movie)
}

func MarkedAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	w.Header().Set("Allow-Content-Allow-Methods", "PUT")

	params := mux.Vars(r)

	updateOneMovie(params["id"])
	// json response
	json.NewEncoder(w).Encode(params["id"]) // just sending the id that it is watched
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	w.Header().Set("Allow-Content-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteOneMovie(params["id"])

	// json response that the movie is deleted
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	w.Header().Set("Allow-Content-Allow-Methods", "DELETE")

	deleteAllMovie()

	json.NewEncoder(w).Encode(map[string]string{
		"message": "All movie deleted",
	})
}
