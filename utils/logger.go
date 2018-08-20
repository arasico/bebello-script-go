package utils

import (
	"fmt"
	"github.com/getsentry/raven-go"
	"log"
	"net/http"
	"time"
)

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}

// LogIt captures error and sends it to Sentry
func LogIt(err error, message string, body ...[]byte) error {
	if err == nil {
		return nil
	}
	if len(body) > 0 {
		raven.CaptureError(err, map[string]string{
			"message": message,
			"body":    string(body[0]),
		})
	} else {
		raven.CaptureError(err, map[string]string{
			"message": message,
		})
	}
	fmt.Println("err: "+message+" --> ", err)
	return err
}
