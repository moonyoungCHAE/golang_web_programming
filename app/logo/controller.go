package logo

import (
	"crypto/md5"
	"fmt"
	"github.com/labstack/echo"
	"os"
)

type Controller struct {
}

func NewController() *Controller {
	return &Controller{}
}

func (controller Controller) Get(c echo.Context) error {
	url := "./assets/memebership.png"
	file, err := os.Stat(url)
	if err != nil {
		return echo.ErrInternalServerError
	}

	modifiedTime := file.ModTime()
	etag := fmt.Sprintf("%x", md5.Sum([]byte(modifiedTime.String())))
	c.Response().Header().Set("ETag", etag)
	return c.File(url)
}
