package main

import (
	"fmt"
	"golang-mongo-graphql-002/internal/api"
	"golang-mongo-graphql-002/internal/api/generated"
	"golang-mongo-graphql-002/internal/issue"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
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

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &api.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

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

// func main() {
// newIssue := issue.Issue{
// 	CreatedAt:   time.Now(),
// 	UpdatedAt:   time.Now(),
// 	Code:        "TEST_CODE1",
// 	Title:       "Test title One.",
// 	Description: "This is a test description.",
// 	Completed:   false,
// }
// createdIssue, createIssueErr := issue.CreateIssue(newIssue)
// if createIssueErr != nil {
// 	panic(createIssueErr)
// }

// fmt.Printf("%+v\n", createdIssue)

// newIssues := []issue.Issue{
// 	{
// 		CreatedAt:   time.Now(),
// 		UpdatedAt:   time.Now(),
// 		Code:        "TEST_CODE2",
// 		Title:       "Test title Two.",
// 		Description: "This is a test description.",
// 		Completed:   false,
// 	},
// 	{
// 		CreatedAt:   time.Now(),
// 		UpdatedAt:   time.Now(),
// 		Code:        "TEST_CODE3",
// 		Title:       "Test title Three.",
// 		Description: "This is a test description.",
// 		Completed:   false,
// 	},
// }
// createdIssues, createIssuesErr := issue.CreateIssues(newIssues)
// if createIssuesErr != nil {
// 	panic(createIssuesErr)
// }

// fmt.Printf("%+v\n", createdIssues)

// filter := issue.Issue{
// 	Code: createdIssue.Code,
// }
// foundIssues, findIssuesErr := issue.FindIssues(filter)
// if findIssuesErr != nil {
// 	panic(findIssuesErr)
// }

// query := bson.D{bson.E{"code", bson.D{{"$in", bson.A{createdIssue.Code, createdIssues[0].Code, createdIssues[1].Code}}}}}

// foundIssues, findIssuesErr = issue.FindIssues(query)
// if findIssuesErr != nil {
// 	panic(findIssuesErr)
// }

// fmt.Printf("%+v\n", foundIssues)

// updateIssuesResponse, updateIssuesErr := issue.UpdateIssues(query, issue.Issue{Completed: true})
// if updateIssuesErr != nil {
// 	panic(updateIssuesErr)
// }

// fmt.Printf("%+v\n", updateIssuesResponse)

// foundIssues, findIssuesErr = issue.FindIssues(query)
// if findIssuesErr != nil {
// 	panic(findIssuesErr)
// }

// fmt.Printf("%+v\n", foundIssues)

// deletedIssues, deleteIssuesErr := issue.DeleteIssues(query)
// if deleteIssuesErr != nil {
// 	panic(deleteIssuesErr)
// }

// fmt.Printf("%+v\n", deletedIssues)

// println("Done")
// }
