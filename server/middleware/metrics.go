package middleware

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hiagomf/bank-api/server/config"
	"github.com/hiagomf/bank-api/server/oops"
	"github.com/hiagomf/bank-api/server/utils"
	"go.uber.org/zap"
)

// Metricas é um middleware que disponibiliza os dados de uma requisição
// nos logs
func Metricas() gin.HandlerFunc {
	return gin.LoggerWithConfig(gin.LoggerConfig{
		Formatter: func(params gin.LogFormatterParams) string {
			output := map[string]interface{}{
				"date":              time.Now(),
				"status_code":       params.StatusCode,
				"client_ip":         params.ClientIP,
				"method":            params.Method,
				"path":              params.Path,
				"latency":           float64(params.Latency) / float64(time.Millisecond),
				"client_user_agent": params.Request.UserAgent(),
				"log_type":          "acesso",
				"router":            "external",
			}

			exists := true
			if !exists {
				output["user_id"] = 0
				output["username"] = "Usuário não logado"
				output["user"] = "Indefinido"
			}

			if v, set := params.Keys["error"]; set {
				var (
					err = v.(error)
					e   *oops.Error
				)
				if errors.As(err, &e) {
					output["error"] = e.Error()
					output["error_code"] = e.Code
					output["trace"] = e.Trace
					output["cause"] = e.Err.Error()
					output["log_type"] = "erro"
				}
			}

			if params.StatusCode == 500 {
				output["fatal_error"] = "Possible panic"
				output["log_type"] = "error"
			}

			if v, set := params.Keys["metrics"]; set {
				if v2, ok := v.(*utils.JSONB); ok {
					output["matric_failed"] = v2
				}
			}

			if strings.Split(params.Path, "/")[0] == "api" {
				output["router"] = "internal"
			}

			b, err := json.Marshal(output)
			if err != nil {
				log.Printf("%+v\n", output)
			}

			return string(b) + string('\n')
		},
		Output: func() io.Writer {
			accessFile, err := os.OpenFile(config.GetConfig().AccessLogDirectory, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0777)
			if err != nil {
				log.Fatal(err)
			}

			gin.DisableConsoleColor()
			return io.MultiWriter(accessFile)
		}(),
	})
}

// GinZap adiciona um middleware customizado do zap
func GinZap(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		t1 := time.Now()
		c.Next()
		latency := time.Since(t1)

		fields := []zap.Field{
			zap.Time("date", time.Now()),
			zap.Int("status_code", c.Writer.Status()),
			zap.String("client_ip", c.ClientIP()),
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.Float64("latency", float64(latency)/float64(time.Millisecond)),
			zap.String("client_user_agent", c.Request.UserAgent()),
			zap.String("log_type", "access"),
			zap.String("request_id", c.Value("RID").(string)),
		}

		var (
			usuarioID      = int64(0)
			uNome, usuario = "Usuário não logado", "Indefinido"
		)

		if strings.Split(c.Request.URL.Path, "/")[0] == "api" {
			zap.String("router", "internal")
		} else {
			zap.String("router", "external")
		}

		fields = append(fields, []zap.Field{
			zap.Int64("user_id", usuarioID),
			zap.String("username", uNome),
			zap.String("user", usuario),
		}...)

		Erro := false
		if v, set := c.Keys["error"]; set {
			var (
				err = v.(error)
				e   *oops.Error
			)
			if errors.As(err, &e) {
				fields = append(fields, []zap.Field{
					zap.Int("error_code", e.Code),
					zap.String("error", e.Error()),
					zap.String("cause", e.Err.Error()),
					zap.Strings("trace", e.Trace),
				}...)
				Erro = true
			}
		}

		if Erro {
			logger.Error("tratamento da requisição falhou", fields...)
		} else {
			logger.Info("requisição tratada", fields...)
		}
	}
}
