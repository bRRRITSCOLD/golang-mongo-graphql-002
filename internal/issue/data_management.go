package issue

import (
	"context"

	"golang-mongo-graphql-002/internal/mongodb"

	"go.mongodb.org/mongo-driver/bson"
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
func CreateIssues(issues []Issue) error {
	//Map struct slice to interface slice as InsertMany accepts interface slice as parameter
	insertableIssues := make([]interface{}, len(issues))
	for i, v := range issues {
		insertableIssues[i] = v
	}
	//Get MongoDB connection using connectionhelper.
	client, err := mongodb.GetMongoClient()
	if err != nil {
		return err
	}
	//Create a handle to the respective collection in the database.
	collection := client.Database(mongodb.DB).Collection(mongodb.ISSUES)
	//Perform InsertMany operation & validate against the error.
	_, err = collection.InsertMany(context.TODO(), insertableIssues)
	if err != nil {
		return err
	}
	//Return success without any error.
	return nil
}

//FindIssues - Get All issues for collection
func FindIssue(filter Issue) (Issue, error) {
	// result := Issue{}
	// //Define filter query for fetching specific document from collection
	// // v := reflect.ValueOf(filter)
	// // typeOfS := v.Type()

	// // var bsonDFilter bson.D
	// // for i := 0; i < v.NumField(); i++ {
	// // 	bsonDFilter = append(bsonDFilter, bson.E{"code", v.Field(i).Interface()})
	// // 	fmt.Printf("Field: %s\tValue: %v\n", typeOfS.Field(i).Name, v.Field(i).Interface())
	// // }

	// // if len(pivot.Base) > 0 {
	// //   setElements = append(setElements, bson.E{"base", pivot.Base})
	// // }
	// // if len(pivot.Email) > 0 {
	// //     setElements = append(setElements, bson.E{"email", pivot.Email})
	// // }

	// // setMap := bson.D{
	// //     {"$set", setElements},
	// // }
	// // bsonFilter := bson.D{primitive.E{Key: "code", Value: code}}
	// // marshaledFilter, err := bson.Marshal(filter)
	// // if err != nil {
	// // 	return result, err
	// // }
	// //Get MongoDB connection using connectionhelper.
	// client, err := mongodb.GetMongoClient()
	// if err != nil {
	// 	return result, err
	// }
	// //Create a handle to the respective collection in the database.
	// collection := client.Database(mongodb.DB).Collection(mongodb.ISSUES)
	// //Perform FindOne operation & validate against the error.
	// err = collection.FindOne(context.TODO(), filter).Decode(&result)
	// if err != nil {
	// 	return result, err
	// }
	// //Return result without any error.
	// return result, nil
	//Define filter query for fetching specific document from collection
	filter := bson.D{{}} //bson.D{{}} specifies 'all documents'
	issues := []Issue{}
	//Get MongoDB connection using connectionhelper.
	client, err := connectionhelper.GetMongoClient()
	if err != nil {
		return issues, err
	}
	//Create a handle to the respective collection in the database.
	collection := client.Database(connectionhelper.DB).Collection(connectionhelper.ISSUES)
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
