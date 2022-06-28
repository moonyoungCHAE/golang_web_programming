package logo

import (
	"crypto/md5"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"strings"
)

type Controller struct {
}

func NewController() *Controller {
	return &Controller{}
}

func (c Controller) Get(ctx echo.Context) error {
	url := "./golang_web_programming/assets/membership.png"

	file, err := os.Stat(url)
	if err != nil {
		return echo.ErrInternalServerError
	}

	modifiedTime := file.ModTime()
	etag := fmt.Sprintf("%x", md5.Sum([]byte(modifiedTime.String())))

	if match := ctx.Request().Header.Get("If-None-Match"); match != "" {
		if strings.Contains(match, etag) {
			ctx.Response().WriteHeader(http.StatusNotModified)
			return nil
		}
	}

	ctx.Response().Header().Set("ETag", etag)
	return ctx.File(url)
}
