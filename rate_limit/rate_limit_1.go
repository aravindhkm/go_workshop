package ratelimit

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

type ClientLimiter struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

var (
	mu        sync.Mutex
	clients   = make(map[string]*ClientLimiter)
	rateLimit = rate.Limit(5) // 5 requests per second
	burst     = 2
)

func getLimiter(ip string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	if client, exists := clients[ip]; exists {
		client.lastSeen = time.Now()
		return client.limiter
	}

	limiter := rate.NewLimiter(rateLimit, burst)
	clients[ip] = &ClientLimiter{
		limiter:  limiter,
		lastSeen: time.Now(),
	}

	return limiter
}

func cleanupOldClients() {
	for {
		time.Sleep(time.Minute)
		mu.Lock()
		for ip, client := range clients {
			if time.Since(client.lastSeen) > 3*time.Minute {
				delete(clients, ip)
			}
		}
		mu.Unlock()
	}
}

func rateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr // Customize IP detection as needed
		fmt.Println("ip address", ip)
		limiter := getLimiter(ip)

		if !limiter.Allow() {
			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func RateLimitOne() {
	go cleanupOldClients()

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "OK")
	})

	http.ListenAndServe(":8080", rateLimitMiddleware(mux))
	fmt.Println("RateLimitOne")
}
