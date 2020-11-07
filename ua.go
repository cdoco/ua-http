package main

import (
	stdContext "context"
	"fmt"
	"net/http"
	"regexp"
	"time"

	"github.com/labstack/echo/v4"
	mw "github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

var config Config

type (
	// Ua struct
	Ua struct {
		UserAgent string `json:"user-agent"`
		Platform  string `json:"platform"`
		OS        string `json:"os"`
		Device    string `json:"device"`
		Browser   string `json:"browser"`
	}

	// Config struct
	Config struct {
		Os []Regexp `mapstructure:"os"`
	}

	// Regexp struct
	Regexp struct {
		Name   string `mapstructure:"name"`
		Regexp string `mapstructure:"regexp"`
	}
)

func init() {
	// viper config
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("ser config load err:" + err.Error())
	}

	viper.Unmarshal(&config)
}

// Handler ua-http
func handler(c echo.Context) error {
	// Ua return
	var u = new(Ua)
	u.UserAgent = c.Request().UserAgent()

	for _, value := range config.Os {
		match, _ := regexp.MatchString(value.Regexp, u.UserAgent)
		if match {
			u.OS = value.Name
		}
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

	e.GET("/", handler)
	e.Logger.Fatal(e.Start(":6080"))
}
