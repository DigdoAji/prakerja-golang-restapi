package controllers

import (
	"fmt"
	"net/http"
	"ujk-golang/configs"
	comicbase "ujk-golang/models/base"
	comicdatabase "ujk-golang/models/comic/database"
	comicrequest "ujk-golang/models/comic/request"
	comicresponse "ujk-golang/models/comic/response"

	"github.com/labstack/echo/v4"
)


func GetComicsController(c echo.Context) error {
	var comics []comicdatabase.Comic

	result := configs.DB.Find(&comics)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, comicbase.BaseResponse{
			Status: false,
			Message: "Failed get all data comics",
			Data: nil,
		})
	}

	return c.JSON(http.StatusOK, comicbase.BaseResponse{
		Status: true,
		Message: "Success get all data comics",
		Data: comics,
	})
}


func AddComicController(c echo.Context) error {
	var comicAdd comicrequest.ComicAdd
	c.Bind(&comicAdd)

	if comicAdd.Title == "" {
		return c.JSON(http.StatusBadRequest, comicbase.BaseResponse{
			Status: false,
			Message: "title cannot be empty",
			Data: nil,
		})
	}

	var comicDatabase comicdatabase.Comic
	comicDatabase.Title = comicAdd.Title
	comicDatabase.Author = comicAdd.Author
	comicDatabase.Genre = comicAdd.Genre
	comicDatabase.Category = comicAdd.Category
	comicDatabase.Date_Published = comicAdd.Date_Published
	comicDatabase.Completed = comicAdd.Completed

	fmt.Println("Adding new comic in database", comicDatabase)

	result := configs.DB.Create(&comicDatabase)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, comicbase.BaseResponse{
			Status: false,
			Message: "Failed add new comic",
			Data: nil,
		})
	}

	var comicResponse comicresponse.ComicResponse
	comicResponse.MapFromDatabaseComic(comicDatabase)

	return c.JSON(http.StatusCreated, comicbase.BaseResponse{
		Status: true,
		Message: "Success add new comic",
		Data: comicResponse,
	})
}


func GetComicDetailController(c echo.Context) error {
	comicID := c.Param("id")
	var comic comicdatabase.Comic

	result := configs.DB.First(&comic, comicID)

	if result.Error != nil {
		return c.JSON(http.StatusNotFound, comicbase.BaseResponse{
			Status: false,
			Message: "Detail comic not found",
			Data: nil,
		})
	}

	var comicResponse comicresponse.ComicResponse
	comicResponse.MapFromDatabaseComic(comic)

	return c.JSON(http.StatusOK, comicbase.BaseResponse{
		Status: true,
		Message: "Success get comic details",
		Data: comicResponse,
	})
}


func UpdateComicController(c echo.Context) error {
	comicID := c.Param("id")
	var comic comicdatabase.Comic

	result := configs.DB.First(&comic, comicID)

	if result.Error != nil {
		return c.JSON(http.StatusNotFound, comicbase.BaseResponse{
			Status: false,
			Message: "Detail comic not found",
			Data: nil,
		})
	}

	var comicUpdate comicrequest.ComicUpdate
	c.Bind(&comicUpdate)

	if comicUpdate.Title == "" {
		return c.JSON(http.StatusBadRequest, comicbase.BaseResponse{
			Status: false,
			Message: "title cannot be empty",
			Data: nil,
		})
	}

	var comicDatabase comicdatabase.Comic
	comicDatabase.Title = comicUpdate.Title
	comicDatabase.Author = comicUpdate.Author
	comicDatabase.Genre = comicUpdate.Genre
	comicDatabase.Category = comicUpdate.Category
	comicDatabase.Date_Published = comicUpdate.Date_Published
	comicDatabase.Completed = comicUpdate.Completed

	fmt.Println("Updating info comic in database", comicDatabase)

	result = configs.DB.Save(&comic)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, comicbase.BaseResponse{
			Status: false,
			Message: "Failed update data comic",
			Data: nil,
		})
	}

	var comicResponse comicresponse.ComicResponse
	comicResponse.MapFromDatabaseComic(comic)

	return c.JSON(http.StatusOK, comicbase.BaseResponse{
		Status: true,
		Message: "Success to update data comic",
		Data: comicResponse,
	})
}

func DeleteComicController(c echo.Context) error {
	comicID := c.Param("id")
	var comic comicdatabase.Comic

	result := configs.DB.First(&comic, comicID)

	if result.Error != nil {
		return c.JSON(http.StatusNotFound, comicbase.BaseResponse{
			Status: false,
			Message: "Detail comic not found",
			Data: nil,
		})
	}

	result = configs.DB.Delete(&comic)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, comicbase.BaseResponse{
			Status: false,
			Message: "Failed delete data comic",
			Data: nil,
		})
	}

	return c.JSON(http.StatusOK, comicbase.BaseResponse{
		Status: true,
		Message: "Success delete data comic",
		Data: nil,
	})
}
