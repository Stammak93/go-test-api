package main

import (
	"context"
	"log"
	"net/http"
	pg_test "test/db-connect"
	"test/models"
	"test/routes"

	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun/extra/bundebug"
)


func main() {
	
	e := echo.New()
	ctx := context.Background()
	db := pg_test.Connect_Db()

	db.NewCreateTable().Model((*models.User)(nil)).Exec(ctx)
	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	
	routes.ApiRoutes(e)

	server := http.Server{
		Addr: "localhost:4000",
		Handler: e,
	}
	

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}
	log.Println("Listening on port 4000")
}
