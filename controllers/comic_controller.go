package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"ujk-golang/configs"
	comicbase "ujk-golang/models/base"
	comicdatabase "ujk-golang/models/comic/database"
	comicrequest "ujk-golang/models/comic/request"
	comicresponse "ujk-golang/models/comic/response"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// Get All Comics
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

// Add New Comic
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

// Get Comic Detail by ID
func GetComicDetailController(c echo.Context) error {
	comicID := c.Param("id")
	var comic comicdatabase.Comic

	result := configs.DB.First(&comic, comicID)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, comicbase.BaseResponse{
				Status:  false,
				Message: "Comic not found",
				Data:    nil,
			})
		}
		return c.JSON(http.StatusInternalServerError, comicbase.BaseResponse{
			Status:  false,
			Message: "Failed to get comic details by ID",
			Data:    nil,
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

// Update Comic by ID
func UpdateComicController(c echo.Context) error {
	comicID := c.Param("id")
	var comic comicdatabase.Comic

	result := configs.DB.First(&comic, comicID)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, comicbase.BaseResponse{
				Status:  false,
				Message: "Comic not found",
				Data:    nil,
			})
		}
		return c.JSON(http.StatusInternalServerError, comicbase.BaseResponse{
			Status:  false,
			Message: "Failed to get comic details by ID",
			Data:    nil,
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

	comic.Title = comicUpdate.Title
	comic.Author = comicUpdate.Author
	comic.Genre = comicUpdate.Genre
	comic.Category = comicUpdate.Category
	comic.Date_Published = comicUpdate.Date_Published
	comic.Completed = comicUpdate.Completed

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

// Delete Comic by ID
func DeleteComicController(c echo.Context) error {
	comicID := c.Param("id")
	var comic comicdatabase.Comic

	result := configs.DB.First(&comic, comicID)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, comicbase.BaseResponse{
				Status:  false,
				Message: "Comic not found",
				Data:    nil,
			})
		}
		return c.JSON(http.StatusInternalServerError, comicbase.BaseResponse{
			Status:  false,
			Message: "Failed to get comic details by ID",
			Data:    nil,
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
