package comment

import (
	"context"
	"time"

	"golang-mongo-graphql-002/internal/mongodb"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//CreateComments - Insert multiple documents at once in the collection.
func CreateComments(comments []Comment) ([]Comment, error) {
	//Map struct slice to interface slice as InsertMany accepts interface slice as parameter
	var newComments []Comment
	for _, v := range comments {
		newComment := MapToComment(v)
		timeNow := time.Now()
		newComment.CommentID = uuid.New().String()
		newComment.CreatedAt = timeNow
		newComment.UpdatedAt = timeNow
		newComments = append(newComments, newComment)
	}

	//Map struct slice to interface slice as InsertMany accepts interface slice as parameter
	var insertableComments []interface{}
	// insertableComments := make([]interface{}, len(newComments))
	for _, v := range newComments {
		insertableComments = append(insertableComments, v)
	}

	//Get MongoDB connection using connectionhelper.
	client, err := mongodb.GetMongoClient()
	if err != nil {
		return newComments, err
	}

	//Create a handle to the respective collection in the database.
	collection := client.Database(mongodb.DB).Collection(mongodb.COMMENTS)

	//Perform InsertMany operation & validate against the error.
	insertManyResponse, err := collection.InsertMany(context.TODO(), insertableComments)
	if err != nil {
		return newComments, err
	}

	// create items to be returned
	var returnableComments []Comment
	for i := 0; i < len(insertManyResponse.InsertedIDs); i++ {
		newComment := MapToComment(newComments[i])
		newCommentID := insertManyResponse.InsertedIDs[i]
		newComment.ID = newCommentID.(primitive.ObjectID).Hex()
		returnableComments = append(returnableComments, newComment)
	}

	//Return success without any error.
	return returnableComments, nil
}

//CreateComment - Insert a new document in the collection.
func CreateComment(cmt Comment) (Comment, error) {
	// call existing create comments func
	createdComments, createCommentsErr := CreateComments([]Comment{
		cmt,
	})
	if createCommentsErr != nil {
		return Comment{}, createCommentsErr
	}

	//Return success without any error.
	return createdComments[0], nil
}

//FindComments - Get All comments that match a criteria for collection
func FindComments(filter interface{}) ([]Comment, error) {
	comments := []Comment{}

	//Get MongoDB connection using connectionhelper.
	client, err := mongodb.GetMongoClient()
	if err != nil {
		return comments, err
	}

	//Create a handle to the respective collection in the database.
	collection := client.Database(mongodb.DB).Collection(mongodb.COMMENTS)

	//Perform Find operation & validate against the error.
	cur, findError := collection.Find(context.TODO(), filter)
	if findError != nil {
		return comments, findError
	}

	//Map result to slice
	for cur.Next(context.TODO()) {
		t := Comment{}

		err := cur.Decode(&t)
		if err != nil {
			return comments, err
		}

		comments = append(comments, t)
	}

	// once exhausted, close the cursor
	cur.Close(context.TODO())

	if len(comments) == 0 {
		return comments, mongo.ErrNoDocuments
	}
	return comments, nil
}

func DeleteComments(filter interface{}) (int64, error) {
	//Get MongoDB connection using connectionhelper.
	client, err := mongodb.GetMongoClient()
	if err != nil {
		return 0, err
	}

	//Create a handle to the respective collection in the database.
	collection := client.Database(mongodb.DB).Collection(mongodb.COMMENTS)

	//Perform DeleteMany operation & validate against the error.
	deletedComments, err := collection.DeleteMany(context.TODO(), filter)
	if err != nil {
		return 0, err
	}

	return deletedComments.DeletedCount, nil
}

func UpdateComments(filter interface{}, update Comment) (int64, error) {
	//Get MongoDB connection using connectionhelper.
	client, err := mongodb.GetMongoClient()
	if err != nil {
		return 0, err
	}

	// create update
	commentUpdate := MapToComment(update)
	commentUpdate.UpdatedAt = time.Now()

	//Create a handle to the respective collection in the database.
	collection := client.Database(mongodb.DB).Collection(mongodb.COMMENTS)

	//Perform DeleteMany operation & validate against the error.
	updateManyResponse, err := collection.UpdateMany(context.TODO(), filter, bson.D{
		bson.E{"$set", commentUpdate},
	})
	if err != nil {
		return 0, err
	}

	return updateManyResponse.ModifiedCount, nil
}
