package middlewares

import (
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/Alvesafk/blog-2/back/internal/handlers"
	"golang.org/x/time/rate"
)

func SecurityHeadersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		w.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains")

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func RateLimiterMiddleware(next http.Handler) http.Handler {
	type visitor struct {
		limiter  *rate.Limiter
		lastSeen time.Time
	}

	var (
		mu       sync.Mutex
		visitors = make(map[string]*visitor)
	)

	go func() {
		for {
			time.Sleep(time.Minute)
			mu.Lock()

			for ip, visitor := range visitors {
				if time.Since(visitor.lastSeen) > 3*time.Minute {
					delete(visitors, ip)
				}
			}

			mu.Unlock()
		}
	}()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if ip == "127.0.0.1" || ip == "::1" {
			if realIp := r.Header.Get("X-Real-IP"); realIp != "" {
				ip = realIp
			}
		}

		mu.Lock()
		if _, found := visitors[ip]; !found {
			visitors[ip] = &visitor{limiter: rate.NewLimiter(5, 10)}
		}

		visitors[ip].lastSeen = time.Now()

		if !visitors[ip].limiter.Allow() {
			mu.Unlock()

			handlers.Response{
				Message: "Too many requests, try again later.",
			}.WriteJSON(w, http.StatusTooManyRequests)

			return
		}

		mu.Unlock()

		next.ServeHTTP(w, r)
	})
}
