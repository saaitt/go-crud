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
func (p ProductHandler) Disable(c echo.Context) error {
	req := request.DisableOrDeleteProductRequest{}
	if err := c.Bind(&req); err != nil {
		return echo.ErrInternalServerError
	}
	if err := p.Service.Disable(req); err != nil {
		return echo.ErrInternalServerError
	}
	return c.JSON(http.StatusOK, "product is disabled")
}
func (p ProductHandler) ListActive(c echo.Context) error {
	pageNo, err := strconv.Atoi(c.Param("page_no"))
	if err != nil {
		fmt.Println(err)
		return echo.ErrInternalServerError
	}
	resp, err := p.Service.ListActive(pageNo)
	if err != nil {
		fmt.Println(err)
		return echo.ErrInternalServerError
	}
	return c.JSON(http.StatusOK, resp)
}
func (p ProductHandler) Delete(c echo.Context) error {
	req := request.DisableOrDeleteProductRequest{}
	if err := c.Bind(&req); err != nil {
		return echo.ErrInternalServerError
	}
	if err := p.Service.Delete(req); err != nil {
		return echo.ErrInternalServerError
	}
	return c.JSON(http.StatusOK, "product is deleted")
}

func (p ProductHandler) Find(c echo.Context) error {
	query := c.QueryParam("q")
	resp, err := p.Service.Find(query)
	if err != nil {
		return echo.ErrInternalServerError
	}
	return c.JSON(http.StatusOK, resp)
}