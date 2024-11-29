package middleware

import (
	"log"
	"net/http"
	"os"
	"time"
)

type Logger struct {
	logger *log.Logger
}

func NewLogger() *Logger {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)

	return &Logger{
		logger: logger,
	}

}

func (l *Logger) LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		//log request details
		l.logger.Printf("Started %s %s", r.Method, r.URL.Path)
		//calling next handler
		next.ServeHTTP(w, r)
		//log req duration
		l.logger.Printf("Completed in %v", time.Since(start))
	})
}
