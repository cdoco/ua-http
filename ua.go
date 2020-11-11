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
		Os      []Regexp `mapstructure:"os"`
		Device  []Regexp `mapstructure:"device"`
		Browser []Regexp `mapstructure:"browser"`
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
		fmt.Println("config load err:" + err.Error())
	}

	viper.Unmarshal(&config)
}

// Handler ua-http
func handler(c echo.Context) error {

	var ua = c.QueryParam("ua")
	if ua == "" {
		ua = c.Request().UserAgent()
	}

	// Ua return
	var u = &Ua{
		UserAgent: ua,
		Platform:  "PC",
		OS:        "unknown",
		Device:    "unknown",
		Browser:   "unknown",
	}

	// Set os name
	for _, value := range config.Os {
		match, _ := regexp.MatchString(value.Regexp, u.UserAgent)
		if match {
			u.OS = value.Name
			break
		}
	}

	// Set platform
	if u.OS == "IOS" || u.OS == "Android" || u.OS == "Windows Phone" || u.OS == "BlackBerry" {
		u.Platform = "Mobile"
	}

	// Set device name
	for _, value := range config.Device {
		match, _ := regexp.MatchString(value.Regexp, u.UserAgent)
		if match {
			u.Device = value.Name
			break
		}
	}

	// Set browser Name
	for _, value := range config.Browser {
		match, _ := regexp.MatchString(value.Regexp, u.UserAgent)
		if match {
			u.Browser = value.Name
			break
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
	e.Logger.Fatal(e.Start(":5080"))
}
