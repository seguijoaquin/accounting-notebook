package main

import (
	"log"
	"net/http"
	"os"

	"github.com/seguijoaquin/accounting-notebook/cmd/api/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := run(port); err != nil {
		log.Fatal("Error running server", err)
	}
}

func run(port string) error {
	router := gin.Default()

	health := HealthChecker{}

	accountHandler := handler.NewAccountHandler()
	transactionHandler := handler.NewTransactionHandler()

	mapRoutes(router, health, accountHandler, transactionHandler)

	return router.Run(":" + port)
}

func mapRoutes(r *gin.Engine, health HealthChecker, account handler.Account, transaction handler.Transaction) {
	r.GET("/ping", health.PingHandler)

	v1 := r.Group("/api/v1")
	{
		v1.GET("/", account.HandleBalance)
		v1.GET("/transactions", transaction.HandleFetchAllTransactions)
		v1.GET("/transactions/:id", transaction.HandleFindTransactionByID)
		v1.POST("/transactions", transaction.HandleNewTransaction)
	}
}

// HealthChecker struct provides the handler for a health check endpoint.
type HealthChecker struct{}

// PingHandler returns a successfull pong answer to all HTTP requests.
func (h HealthChecker) PingHandler(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
