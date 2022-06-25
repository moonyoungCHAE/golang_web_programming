package app

import (
	"comento_git_practice/app/membership"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"reflect"
)

type Config struct {
	Controller membership.Controller
}

func DefaultConfig() *Config {
	data := map[string]membership.Membership{}
	service := membership.NewService(*membership.NewRepository(data))
	controller := membership.NewController(*service)
	return &Config{
		Controller: *controller,
	}
}

func NewEcho(config Config) *echo.Echo {
	e := echo.New()
	controller := config.Controller
	e.Use(myMiddleMembershipLogger)
	e.GET("/memberships/:id", controller.GetByID)
	e.GET("/memberships", controller.GetMembers)
	e.POST("/memberships", controller.Create)
	e.DELETE("/memberships/:id", controller.RemoveByID)

	return e
}

func myMiddleMembershipLogger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var err error
		var clog []byte
		if err = next(c); err != nil {
			return err
		}
		req := c.Request()
		res := c.Response()
		data := c.Get("filter")
		if data != nil {

			t := reflect.ValueOf(data)
			e := t.Elem()
			m := make(map[string]interface{})
			for i := 0; i < e.NumField(); i++ {
				mValue := e.Field(i)
				mType := e.Type().Field(i)
				if _, ok := mType.Tag.Lookup("logging"); ok {
					m[mType.Name] = mValue.Interface()
				}
			}
			clog, _ = json.Marshal(m)
		}
		fmt.Println(
			c.RealIP(),                               // Request host IP
			req.Method,                               // Request Method
			req.Host,                                 // Request Host
			req.Header.Get(echo.HeaderContentLength), // Request Data Size
			res.Status,                               // Response status
			res.Size,                                 // Response Data Size
			string(clog),
		)
		return err
	}
}
