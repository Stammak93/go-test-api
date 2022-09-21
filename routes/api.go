package routes

import (
	"context"
	"net/http"
	pg_test "test/db-connect"
	"test/models"

	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
)

var DBInstance *bun.DB = pg_test.DB

var messages = make([]models.Message, 0)

// returns all messages posted to route api/test
func testRoute(c echo.Context) error {
	return c.JSON(http.StatusOK, messages)
}

// returns whatever you post as json
// data structure must be of type models.Message 
func testRouteTwo(c echo.Context) error {

	newMessage := new(models.Message)
	if err := c.Bind(newMessage); err != nil {
		return err
	}

	messages = append(messages, *newMessage)

	return c.JSON(http.StatusOK, newMessage)
}

// adds a user to db
// data structure must be of type models.User
func testDbInput(c echo.Context) error {

	ctx := context.Background()
	newUser := new(models.User)
	if err := c.Bind(newUser); err != nil {
		return err
	}

	user := &models.User{Name: newUser.Name, Banana: newUser.Banana}
	DBInstance.NewInsert().Model(user).Exec(ctx)

	return c.JSON(http.StatusOK, newUser)
}


func ApiRoutes(echo *echo.Echo) {

	echo.GET("/api/test", testRoute)
	echo.POST("/api/test", testRouteTwo)
	echo.POST("/db/test", testDbInput)
}