package anyController

import (
	"context"
	"fmt"
	"log"
	"time"

	anymodel "gihub.com/cmrohityadav/go/mongodbapi/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/v2/bson"
)

const connectionString = "mongodb://localhost:27017"
const dbName = "firstDbForGo"
const collectionName = "watchList"

var collection *mongo.Collection

func init() {
	// Define MongoDB client options
	clientOptions := options.Client().ApplyURI(connectionString)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}

	// Ping the database to ensure a successful connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Could not ping MongoDB:", err)
	}

	fmt.Println("MongoDB connection successful")

	// Get the collection reference
	collection = client.Database(dbName).Collection(collectionName)
	fmt.Println("Collection instance is ready")
}

// MongoDB helpers - file

//insert 1 record

func insertOneMovie(movie anymodel.Netflix){
	inserted,err:=collection.InsertOne(context.Background(),movie)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 movie in  db with id: ",inserted.InsertedID)

}

//update 1 record

func updateOneMovie(movieId string){
	id,_:=primitive.ObjectIDFromHex(movieId)
	filter:=bson.M{"_id":id}
	update:=bson.M{"$set":bson.M{"watched":true}}
	result,err:=collection.UpdateOne(context.Background(),filter,update)
	if err !=nil{
		log.Fatal(err)
	}
	fmt.Println("modified count : ",result.ModifiedCount)
}

//delete 1 record
func deleteOneMovie(movieId string){
	id,_:=primitive.ObjectIDFromHex(movieId)
	filter:=bson.M{"-id":id}
	deleteCount,err:=collection.DeleteOne(context.Background(),filter)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("Movie got delete with delete count :",deleteCount)

}

// delete all records from mongodb
func deleteAllMovie(){
	// filter:=bson.D{{}}
	deleteCount,err:=collection.DeleteMany(context.Background(),bson.D{{}},nil)
	if err !=nil{
		log.Fatal(err)
	}
	fmt.Println("Number of movies delete: ",deleteCount.DeletedCount)


}





