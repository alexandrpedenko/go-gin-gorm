package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func GlobalErrorHandler(r *gin.Engine) {
	r.Use(func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				log.Info().Msg("Recovered from panic: ")

				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Internal Server Error",
				})
			}
		}()

		c.Next()
	})
}
