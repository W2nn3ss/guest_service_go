package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"runtime"
	"time"
)

type CustomResponseWriter struct {
	gin.ResponseWriter
	Headers    map[string]string
	StatusCode int
}

func (w *CustomResponseWriter) WriteHeader(statusCode int) {
	log.Printf("WriteHeader called with status: %d", statusCode)
	for key, value := range w.Headers {
		log.Printf("Setting header %s: %s", key, value)
		w.Header().Set(key, value)
	}

	w.Header().Set("X-Debug-Time", w.Headers["X-Debug-Time"])
	w.Header().Set("X-Debug-Memory", w.Headers["X-Debug-Memory"])

	w.ResponseWriter.WriteHeader(statusCode)
}

func DebugMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		crw := &CustomResponseWriter{ResponseWriter: c.Writer, Headers: make(map[string]string), StatusCode: 200}
		c.Writer = crw

		log.Println("Calling c.Next()...")
		c.Next()

		duration := time.Since(start).Milliseconds()
		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)
		memoryUsage := memStats.Alloc / 1024

		crw.Headers["X-Debug-Time"] = fmt.Sprintf("%dms", duration)
		crw.Headers["X-Debug-Memory"] = fmt.Sprintf("%dkb", memoryUsage)

		log.Println("Headers set after c.Next()")

		if crw.StatusCode != 0 {
			crw.WriteHeaderNow()
		}

		for key, values := range crw.Header() {
			for _, value := range values {
				log.Printf("Header before response sent - %s: %s", key, value)
			}
		}
	}
}
