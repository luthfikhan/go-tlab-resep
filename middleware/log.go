package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Log(ctx *gin.Context) {
	ctx.Next()

	logrus.WithFields(map[string]any{
		"method": ctx.Request.Method,
		"path":   ctx.Request.URL.Path,
		"query":  ctx.Request.URL.Query(),
		"code":   ctx.Writer.Status(),
	}).Info("Request Info")
}
