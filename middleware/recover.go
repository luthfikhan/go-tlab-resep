package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	model_web "github.com/luthfikhan/go-tlab-resep/models/web"
	"github.com/sirupsen/logrus"
)

func Recover() gin.HandlerFunc {
	return gin.CustomRecovery(func(ctx *gin.Context, err interface{}) {
		logrus.WithFields(map[string]any{
			"error": err,
		}).Error("Recover from panic")
		ctx.JSON(http.StatusInternalServerError, model_web.InternalServerError)

		gin.RecoveryWithWriter(gin.DefaultErrorWriter)
	})
}
