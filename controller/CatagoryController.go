package controller

import (
	"fmt"
	"golang-echo-mvc-framework/entity"
	"golang-echo-mvc-framework/model"
	"golang-echo-mvc-framework/utils"
	"net/http"

	"github.com/labstack/echo"
)

type CatagoryController struct {
	catagoryModel model.CatagoryModel
	userModel     model.UserModel
}

func (e *CatagoryController) AddCatagory(c echo.Context) error {
	var catagory = entity.Catagory{}
	err := c.Bind(&catagory)
	if err != nil {
		fmt.Printf("[CatagoryController.AddCatagory] error bind data %v \n", err)
		return utils.HandleError(c, http.StatusInternalServerError, "Oopss server has be wrong")
	}

	_, err = e.userModel.GetUserById(int(catagory.UserID))
	if err != nil {
		return utils.HandleError(c, http.StatusNotFound, err.Error())
	}
	mCatagory, err := e.catagoryModel.AddCatagory(&catagory)
	if err != nil {
		return utils.HandleError(c, http.StatusInternalServerError, err.Error())
	}
	return utils.HandleSuccess(c, mCatagory)
}

func (e *CatagoryController) ViewAllCatagory(c echo.Context) error {
	catagories, err := e.catagoryModel.GetCatagories()
	if err != nil {
		return utils.HandleError(c, http.StatusInternalServerError, err.Error())
	}
	// fmt.Println("catagories", catagories)
	return utils.HandleSuccess(c, catagories)
}
