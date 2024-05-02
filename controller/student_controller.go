package controller

import (
	"belajar_rest_api/model"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type response struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func responseData(c echo.Context, status int, data interface{}, message string) error {

	return c.JSON(status, response{
		Data:    data,
		Message: message,
	})
}

func CreateStudent(c echo.Context) (err error) {
	var s model.Student

	err = c.Bind(&s)
	if err != nil {
		log.Println("failed parse request", err)
		return
	}

	return responseData(c, http.StatusCreated, s, "successfully create student")
}
