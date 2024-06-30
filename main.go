package main

import (
	"github/chaso-pa/gql-server/graph"
	"github/chaso-pa/gql-server/middleware"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const defaultPort = "8080"

func main() {
	middleware.LoadEnv()
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db := middleware.ConnectDB()
	router := gin.Default()

	router.POST("/graphql", graphqlHandler(db))
	router.GET("/", playgroundHandler())

	if os.Getenv("GIN_MODE") != "release" {
		router.Run("localhost:" + port)
	} else {
		router.Run(":" + port)
	}
}

func graphqlHandler(db *gorm.DB) gin.HandlerFunc {
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		DB: db,
	}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/graphql")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
