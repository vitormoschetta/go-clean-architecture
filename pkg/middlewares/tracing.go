package middlewares

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	CorrelationIDHeader = "X-Correlation-ID"
	CorrelationKey      = ContextKey(CorrelationIDHeader)
)

type ContextKey string

func Tracing(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		traceID := r.Header.Get(CorrelationIDHeader)
		if traceID == "" {
			traceID = generateTraceID()
			log.Println("New trace ID generated:", traceID)
		}
		ctx := context.WithValue(r.Context(), CorrelationKey, traceID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetTraceID(ctx context.Context) string {
	traceID, ok := ctx.Value(CorrelationKey).(string)
	if !ok {
		return ""
	}
	return traceID
}

func generateTraceID() string {
	pid := os.Getpid()
	timestamp := time.Now().UnixNano()
	return fmt.Sprintf("%d-%d", pid, timestamp)
}
