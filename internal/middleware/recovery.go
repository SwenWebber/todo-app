package middleware

import (
	"log"
	"net/http"
	"os"
	"runtime/debug"
)

type Recovery struct {
	logger *log.Logger
}

func NewRecovery() *Recovery {
	logger := log.New(os.Stdout, "Recovery: ", log.Ldate|log.Ltime|log.Lshortfile)

	return &Recovery{
		logger: logger,
	}
}

func (rec *Recovery) RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				rec.logger.Printf("Panic:%v\n%s", err, debug.Stack())
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}

		}()
		next.ServeHTTP(w, r)
	})
}
