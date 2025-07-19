package ratelimit

import (
	"fmt"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/ulule/limiter/v3"
	redisStore "github.com/ulule/limiter/v3/drivers/store/redis"
)

func RateExampleTwo() {
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	store, err := redisStore.NewStoreWithOptions(redisClient, limiter.StoreOptions{})
	if err != nil {
		panic(err)
	}

	rateConfig := limiter.Rate{
		Period: 1 * time.Minute,
		Limit:  100,
	}
	instance := limiter.New(store, rateConfig)

	middleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			key := r.RemoteAddr
			context, err := instance.Get(r.Context(), key)
			if err != nil {
				http.Error(w, "Internal error", http.StatusInternalServerError)
				return
			}

			w.Header().Set("X-RateLimit-Limit", fmt.Sprint(context.Limit))
			w.Header().Set("X-RateLimit-Remaining", fmt.Sprint(context.Remaining))
			w.Header().Set("X-RateLimit-Reset", fmt.Sprint(context.Reset))

			if context.Reached {
				http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
				return
			}

			next.ServeHTTP(w, r)
		})
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "OK from Redis-limited API")
	})

	http.ListenAndServe(":8080", middleware(mux))
}
