package graph

import (
	"gin-graphql/todo"
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"net/http"
)

type RequestParams struct {
	Query     string                 `json:"query"`
	Operation string                 `json:"operation"`
	Variables map[string]interface{} `json:"variables"`
}

func TodoGraphRouter(c *gin.Context) {
	var requestObject RequestParams
	if err := c.ShouldBindJSON(&requestObject); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := graphql.Do(graphql.Params{
		Context:        c,
		Schema:         todo.TodoRootSchema,
		RequestString:  requestObject.Query,
		VariableValues: requestObject.Variables,
		OperationName:  requestObject.Operation,
	})

	if len(result.Errors) > 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": result.Errors})
		return
	}

	c.JSON(http.StatusOK, result)
}
