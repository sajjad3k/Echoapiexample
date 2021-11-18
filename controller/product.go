package controller

//Controller Logic for all the API routes

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/sajjad3k/echoapione/model"
)

//Fetching all the products

func Getproducts(c echo.Context) error {
	//var resp model.Response
	/*resp, err := model.Fetchallprod()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, resp)
	} */
	resp, err := model.Getallprod()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, resp)
	}
	return c.JSON(http.StatusOK, resp)
}

//Fetching a single product using the code
func Getproduct(c echo.Context) error {
	var resp model.Response
	code := c.Param("code")
	/* resp, err := model.Fetchaskedprod(code)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, resp)
	} */
	resp, err := model.Getoneprod(code)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, resp)
	}
	return c.JSON(http.StatusOK, resp)
}

//Creating a new product
func Createproduct(c echo.Context) error {
	var resp model.Response
	var prod model.Product
	c.Bind(&prod)
	/*resp,err := model.Postnewprod(prod)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, resp)
	} */
	resp, err := model.Creatprod(prod)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, resp)
	}
	return c.JSON(http.StatusOK, resp)
}

//Updating the product details
func Updateproduct(c echo.Context) error {
	var resp model.Response
	var prod model.Product
	code := c.Param("code")
	c.Bind(&prod)
	/* resp, err := model.Updateaskedprod(code, prod)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, resp)
	} */
	resp, err := model.Updateprod(code, prod)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, resp)
	}
	return c.JSON(http.StatusOK, resp)
}

//Deleting the product details
func Deleteproduct(c echo.Context) error {
	var resp model.Response
	code := c.Param("code")
	/* resp, err := model.Deleteaskedprod(code)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, resp)
	} */
	resp, err := model.Deleteprod(code)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, resp)
	}
	return c.JSON(http.StatusOK, resp)
}
