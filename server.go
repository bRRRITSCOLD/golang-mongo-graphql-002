package main

import (
	"fmt"
	"golang-mongo-graphql-002/internal/issue"
	"time"
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
	createdIssue, createIssueErr := issue.CreateIssue(issue.Issue{
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Code:        "TEST_CODE2",
		Title:       "Test title.",
		Description: "This is a test description.",
		Completed:   false,
	})
	if createIssueErr != nil {
		panic(createIssueErr)
	}

	fmt.Printf("%+v\n", createdIssue)

	filter := issue.Issue{
		Code: createdIssue.Code,
	}
	foundIssue, findIssueErr := issue.FindIssue(filter)
	if findIssueErr != nil {
		panic(findIssueErr)
	}

	fmt.Printf("%+v\n", foundIssue)

	println("Done")
}
