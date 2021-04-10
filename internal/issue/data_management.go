package issue

import (
	"context"
	"time"

	"golang-mongo-graphql-002/internal/mongodb"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//CreateIssues - Insert multiple documents at once in the collection.
func CreateIssues(issues []Issue) ([]Issue, error) {
	//Map struct slice to interface slice as InsertMany accepts interface slice as parameter
	var newIssues []Issue
	for _, v := range issues {
		newIssue := MapToIssue(v)
		timeNow := time.Now()
		newIssue.IssueID = uuid.New().String()
		newIssue.CreatedAt = timeNow
		newIssue.UpdatedAt = timeNow
		newIssues = append(newIssues, newIssue)
	}

	//Map struct slice to interface slice as InsertMany accepts interface slice as parameter
	var insertableIssues []interface{}
	// insertableIssues := make([]interface{}, len(newIssues))
	for _, v := range newIssues {
		insertableIssues = append(insertableIssues, v)
	}

	//Get MongoDB connection using connectionhelper.
	client, err := mongodb.GetMongoClient()
	if err != nil {
		return newIssues, err
	}

	//Create a handle to the respective collection in the database.
	collection := client.Database(mongodb.DB).Collection(mongodb.ISSUES)

	//Perform InsertMany operation & validate against the error.
	insertManyResponse, err := collection.InsertMany(context.TODO(), insertableIssues)
	if err != nil {
		return newIssues, err
	}

	// create items to be returned
	var returnableIssues []Issue
	for i := 0; i < len(insertManyResponse.InsertedIDs); i++ {
		newIssue := MapToIssue(newIssues[i])
		newIssueID := insertManyResponse.InsertedIDs[i]
		newIssue.ID = newIssueID.(primitive.ObjectID).Hex()
		returnableIssues = append(returnableIssues, newIssue)
	}

	//Return success without any error.
	return returnableIssues, nil
}

//CreateIssue - Insert a new document in the collection.
func CreateIssue(iss Issue) (Issue, error) {
	// call existing create issues func
	createdIssues, createIssuesErr := CreateIssues([]Issue{
		iss,
	})
	if createIssuesErr != nil {
		return Issue{}, createIssuesErr
	}

	//Return success without any error.
	return createdIssues[0], nil
}

//FindIssues - Get All issues that match a criteria for collection
func FindIssues(filter interface{}) ([]Issue, error) {
	issues := []Issue{}

	//Get MongoDB connection using connectionhelper.
	client, err := mongodb.GetMongoClient()
	if err != nil {
		return issues, err
	}

	//Create a handle to the respective collection in the database.
	collection := client.Database(mongodb.DB).Collection(mongodb.ISSUES)

	//Perform Find operation & validate against the error.
	cur, findError := collection.Find(context.TODO(), filter)
	if findError != nil {
		return issues, findError
	}

	//Map result to slice
	for cur.Next(context.TODO()) {
		t := Issue{}

		err := cur.Decode(&t)
		if err != nil {
			return issues, err
		}

		issues = append(issues, t)
	}

	// once exhausted, close the cursor
	cur.Close(context.TODO())

	if len(issues) == 0 {
		return issues, mongo.ErrNoDocuments
	}
	return issues, nil
}

func DeleteIssues(filter interface{}) (int64, error) {
	//Get MongoDB connection using connectionhelper.
	client, err := mongodb.GetMongoClient()
	if err != nil {
		return 0, err
	}

	//Create a handle to the respective collection in the database.
	collection := client.Database(mongodb.DB).Collection(mongodb.ISSUES)

	//Perform DeleteMany operation & validate against the error.
	deletedIssues, err := collection.DeleteMany(context.TODO(), filter)
	if err != nil {
		return 0, err
	}

	return deletedIssues.DeletedCount, nil
}

func UpdateIssues(filter interface{}, update Issue) (int64, error) {
	//Get MongoDB connection using connectionhelper.
	client, err := mongodb.GetMongoClient()
	if err != nil {
		return 0, err
	}

	// create update
	issueUpdate := MapToIssue(update)
	issueUpdate.UpdatedAt = time.Now()

	//Create a handle to the respective collection in the database.
	collection := client.Database(mongodb.DB).Collection(mongodb.ISSUES)

	//Perform DeleteMany operation & validate against the error.
	updateManyResponse, err := collection.UpdateMany(context.TODO(), filter, bson.D{
		bson.E{"$set", issueUpdate},
	})
	if err != nil {
		return 0, err
	}

	return updateManyResponse.ModifiedCount, nil
}
