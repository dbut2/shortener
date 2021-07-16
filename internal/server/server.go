package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dbut2/shortener/pkg/app"
	"github.com/dbut2/shortener/pkg/config"
	"github.com/dbut2/shortener/pkg/model"
	"github.com/gin-gonic/gin"
)

func Run(c config.Config) error {
	a := app.New(c)
	fmt.Println("a: ", a)

	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, "pong")
		})

		api.POST("/shorten", func(c *gin.Context) {
			var r model.Shorten
			err := c.ShouldBindJSON(&r)
			if err != nil {
				log.Println(err.Error())
				e(c, 403, err.Error())
			}
			a.Shorten(r)
		})
	}

	r.GET("/:code", func(c *gin.Context) {
		code := c.Param("code")
		s := a.Lengthen(code)
		c.Redirect(http.StatusTemporaryRedirect, s.Url)
	})

	return r.Run(c.Address)
}

func e(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"error":   code,
		"message": message,
	})
}
