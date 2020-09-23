package http

import (
	"context"
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"

	"github.com/golangid/menekel"
	articleUcase "github.com/golangid/menekel"
	"github.com/labstack/echo/v4"

	validator "gopkg.in/go-playground/validator.v9"
)

// ResponseError will hold the response error structs
type ResponseError struct {
	Message string `json:"message"`
}

type articleHandler struct {
	AUsecase articleUcase.ArticleUsecase
}

// InitArticleHandler will initialize the article's HTTP handler
func InitArticleHandler(e *echo.Echo, us articleUcase.ArticleUsecase) {
	handler := &articleHandler{
		AUsecase: us,
	}

	e.GET("/articles", handler.FetchArticle)
	e.POST("/articles", handler.Store)
	e.GET("/articles/:id", handler.GetByID)
	e.DELETE("/articles/:id", handler.Delete)
}

func (a *articleHandler) FetchArticle(c echo.Context) error {
	numS := c.QueryParam("num")
	num, _ := strconv.Atoi(numS)
	cursor := c.QueryParam("cursor")
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	listAr, nextCursor, err := a.AUsecase.Fetch(ctx, cursor, int64(num))
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	c.Response().Header().Set(`X-Cursor`, nextCursor)
	return c.JSON(http.StatusOK, listAr)
}

func (a *articleHandler) GetByID(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	id := int64(idP)
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	art, err := a.AUsecase.GetByID(ctx, id)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, art)
}

func isRequestValid(m *menekel.Article) (bool, error) {
	validate := validator.New()

	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (a *articleHandler) Store(c echo.Context) error {
	var article menekel.Article
	err := c.Bind(&article)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := isRequestValid(&article); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err = a.AUsecase.Store(ctx, &article)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, article)
}

func (a *articleHandler) Delete(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	id := int64(idP)
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err = a.AUsecase.Delete(ctx, id)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)
	switch err {
	case menekel.ErrInternalServerError:
		return http.StatusInternalServerError
	case menekel.ErrNotFound:
		return http.StatusNotFound
	case menekel.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
