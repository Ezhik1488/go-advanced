package middleware

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wrapper := &WrapperWriter{
			ResponseWriter: w,
			StatusCode:     http.StatusOK,
		}

		next.ServeHTTP(wrapper, r)
		log.WithFields(log.Fields{
			"Method":        r.Method,
			"URL":           r.URL.String(),
			"StatusCode":    wrapper.StatusCode,
			"Duration (ms)": time.Since(start).Milliseconds(),
		}).Info("Completed Request")
	})
}
