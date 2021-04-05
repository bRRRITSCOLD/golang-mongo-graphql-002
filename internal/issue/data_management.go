package issue

import (
	"context"

	"golang-mongo-graphql-002/internal/mongodb"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//CreateIssue - Insert a new document in the collection.
func CreateIssue(issue Issue) (Issue, error) {
	//Get MongoDB connection using connectionhelper.
	client, err := mongodb.GetMongoClient()
	if err != nil {
		return Issue{}, err
	}
	//Create a handle to the respective collection in the database.
	collection := client.Database(mongodb.DB).Collection(mongodb.ISSUES)
	//Perform InsertOne operation & validate against the error.
	insertOneResult, err := collection.InsertOne(context.TODO(), issue)
	if err != nil {
		return Issue{}, err
	}
	//Return success without any error.
	return Issue{
		ID:          insertOneResult.InsertedID.(primitive.ObjectID).Hex(),
		CreatedAt:   issue.CreatedAt,
		UpdatedAt:   issue.UpdatedAt,
		Title:       issue.Title,
		Code:        issue.Code,
		Description: issue.Description,
		Completed:   issue.Completed,
	}, nil
}

//CreateIssues - Insert multiple documents at once in the collection.
func CreateIssues(issues []Issue) ([]Issue, error) {
	var newIssues []Issue
	//Map struct slice to interface slice as InsertMany accepts interface slice as parameter
	insertableIssues := make([]interface{}, len(issues))
	for i, v := range issues {
		insertableIssues[i] = v
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
	for i := 0; i < len(insertManyResponse.InsertedIDs); i++ {
		newIssue := issues[i]
		newIssueID := insertManyResponse.InsertedIDs[i]
		newIssues = append(newIssues, Issue{
			ID:          newIssueID.(primitive.ObjectID).Hex(),
			CreatedAt:   newIssue.CreatedAt,
			UpdatedAt:   newIssue.UpdatedAt,
			Title:       newIssue.Title,
			Code:        newIssue.Code,
			Description: newIssue.Description,
			Completed:   newIssue.Completed,
		})
	}
	//Return success without any error.
	return newIssues, nil
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
