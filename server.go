package main

import (
	"fmt"
	"golang-mongo-graphql-002/internal/issue"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// import (
// 	"github.com/gin-gonic/gin"

// 	"github.com/99designs/gqlgen/graphql/handler"
// 	"github.com/99designs/gqlgen/graphql/playground"
// )

// // Defining the Graphql handler
// func graphqlHandler() gin.HandlerFunc {
// 	// NewExecutableSchema and Config are in the generated.go file
// 	// Resolver is in the resolver.go file
// 	h := handler.NewDefaultServer(NewExecutableSchema(Config{Resolvers: &Resolver{}}))

// 	return func(c *gin.Context) {
// 		h.ServeHTTP(c.Writer, c.Request)
// 	}
// }

// // Defining the Playground handler
// func playgroundHandler() gin.HandlerFunc {
// 	h := playground.Handler("GraphQL", "/query")

// 	return func(c *gin.Context) {
// 		h.ServeHTTP(c.Writer, c.Request)
// 	}
// }

// func main() {
// 	// Setting up Gin
// 	r := gin.Default()
// 	r.POST("/query", graphqlHandler())
// 	r.GET("/", playgroundHandler())
// 	r.Run()
// }

func main() {
	newIssue := issue.Issue{
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Code:        "TEST_CODE1",
		Title:       "Test title One.",
		Description: "This is a test description.",
		Completed:   false,
	}
	createdIssue, createIssueErr := issue.CreateIssue(newIssue)
	if createIssueErr != nil {
		panic(createIssueErr)
	}

	fmt.Printf("%+v\n", createdIssue)

	newIssues := []issue.Issue{
		{
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Code:        "TEST_CODE2",
			Title:       "Test title Two.",
			Description: "This is a test description.",
			Completed:   false,
		},
		{
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Code:        "TEST_CODE3",
			Title:       "Test title Three.",
			Description: "This is a test description.",
			Completed:   false,
		},
	}
	createdIssues, createIssuesErr := issue.CreateIssues(newIssues)
	if createIssuesErr != nil {
		panic(createIssuesErr)
	}

	fmt.Printf("%+v\n", createdIssues)

	filter := issue.Issue{
		Code: createdIssue.Code,
	}
	foundIssues, findIssuesErr := issue.FindIssues(filter)
	if findIssuesErr != nil {
		panic(findIssuesErr)
	}
	query := bson.D{bson.E{"code", bson.D{{"$in", bson.A{createdIssue.Code, createdIssues[0].Code, createdIssues[1].Code}}}}}
	foundIssues, findIssuesErr = issue.FindIssues(query)
	if findIssuesErr != nil {
		panic(findIssuesErr)
	}

	fmt.Printf("%+v\n", foundIssues)

	deleteFilter := issue.Issue{
		Description: createdIssue.Description,
	}
	deletedIssues, deleteIssuesErr := issue.DeleteIssues(deleteFilter)
	if deleteIssuesErr != nil {
		panic(deleteIssuesErr)
	}

	fmt.Printf("%+v\n", deletedIssues)

	println("Done")
}
