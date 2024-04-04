package controller

import (
	"golang-echo-mvc-framework/responsegraph"
	"golang-echo-mvc-framework/model"
	"github.com/labstack/echo"
	"net/http"
)

type ExampleController struct {
	model model.ExampleModel
}

// Get Example Controller
func (ExampleController ExampleController) GetPostsController(c echo.Context) error {
	posts := ExampleController.model.GetPosts()
	res := responsegraph.ResponseGeneric{
		Status:  "Success",
		Message: "Posts Loaded",
		Data:    posts,
	}
	return c.JSON(http.StatusOK, res)
}
