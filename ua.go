package main

import (
	stdContext "context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	mw "github.com/labstack/echo/v4/middleware"
)

// Ua struct
type Ua struct {
	UserAgent string `json:"user-agent"`
	Platform  string `json:"platform"`
	OS        string `json:"os"`
	Device    string `json:"device"`
	Browser   string `json:"browser"`
}

// Handler ua-http
func ua(c echo.Context) error {

	// r, _ := regexp.Compile("MicroMessenger")
	// r.FindString(userAgent))

	u := &Ua{
		Platform:  "ZiHang Gao",
		UserAgent: c.Request().UserAgent(),
	}

	// return json
	return c.JSON(http.StatusOK, u)
}

func main() {
	e := echo.New()

	e.Pre(mw.RemoveTrailingSlash())
	e.Use(mw.Logger())
	e.Use(mw.Recover())
	e.Use(mw.Gzip())

	// Http Server
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set(echo.HeaderServer, "ua-http/1.0")

			// set context timeout
			ctx, cancel := stdContext.WithTimeout(
				c.Request().Context(),
				time.Millisecond*1500,
			)
			c.SetRequest(c.Request().WithContext(ctx))
			defer cancel()

			return next(c)
		}
	})

	e.GET("/", ua)
	e.Logger.Fatal(e.Start(":6080"))
}
