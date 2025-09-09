package logging

import (
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var Logger zerolog.Logger

// InitLogger sets up zerolog to write to app.log and returns the file for closing.
func InitLogger(logPath string) (*os.File, error) {
	logFile, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	zerolog.TimeFieldFormat = time.RFC3339

	// ConsoleWriter for human-readable logs in terminal
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}

	// MultiWriter: JSON to file, human-readable to terminal
	multiWriter := io.MultiWriter(logFile, consoleWriter)

	Logger = zerolog.New(multiWriter).With().Timestamp().Logger()
	log.Logger = Logger // set global logger

	return logFile, nil
}

// ZerologMiddleware logs HTTP requests using zerolog.
func ZerologMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		latency := time.Since(start)
		status := c.Writer.Status()

		requestID, _ := c.Get("request_id")
		username, _ := c.Get("username")

		event := Logger.Info().
			Str("method", c.Request.Method).
			Str("path", c.Request.URL.Path).
			Int("status", status).
			Dur("latency", latency).
			Str("client_ip", c.ClientIP())

		if requestID != nil {
			event = event.Str("request_id", requestID.(string))
		}
		if username != nil {
			event = event.Str("username", username.(string))
		}
		if len(c.Errors) > 0 {
			event.Str("errors", c.Errors.String())
		}
		// event.Msg("request completed")
	}
}
