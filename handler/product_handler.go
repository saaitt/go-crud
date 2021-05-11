package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/saaitt/go-crud/request"
	"github.com/saaitt/go-crud/service"
)

type ProductHandler struct {
	Service service.ProductService
}

func (p ProductHandler) Create(c echo.Context) error {
	req := request.CreateProductRequest{}
	if err := c.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}
	resp, err := p.Service.Create(req)
	if err != nil {
		return echo.ErrInternalServerError
	}
	return c.JSON(http.StatusOK, resp)
}

func (p ProductHandler) List(c echo.Context) error {
	pageNo, err := strconv.Atoi(c.Param("page_no"))
	if err != nil {
		fmt.Println(err)
		return echo.ErrInternalServerError
	}
	resp, err := p.Service.List(pageNo)
	if err != nil {
		fmt.Println(err)
		return echo.ErrInternalServerError
	}
	return c.JSON(http.StatusOK, resp)
}
