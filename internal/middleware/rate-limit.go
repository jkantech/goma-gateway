package middleware

/*
Copyright 2024 Jonas Kaninda

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
import (
	"errors"
	"fmt"
	"github.com/go-redis/redis_rate/v10"
	"github.com/gorilla/mux"
	"github.com/jkaninda/goma-gateway/pkg/logger"
	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
	"net/http"
	"time"
)

// RateLimitMiddleware limits request based on the number of tokens peer minutes.
func (rl *TokenRateLimiter) RateLimitMiddleware() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !rl.Allow() {
				logger.Error("Too many requests from IP: %s %s %s", getRealIP(r), r.URL, r.UserAgent())
				//RespondWithError(w, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), basicAuth.ErrorInterceptor)

				// Rate limit exceeded, return a 429 Too Many Requests response
				w.WriteHeader(http.StatusTooManyRequests)
				_, err := w.Write([]byte(fmt.Sprintf("%d Too many requests, API rate limit exceeded. Please try again later", http.StatusTooManyRequests)))
				if err != nil {
					return
				}
				return
			}
			// Proceed to the next handler if rate limit is not exceeded
			next.ServeHTTP(w, r)
		})
	}
}

// RateLimitMiddleware limits request based on the number of requests peer minutes.
func (rl *RateLimiter) RateLimitMiddleware() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			clientIP := getRealIP(r)
			logger.Debug("rate limiter: clientID: %s, redisBased: %s", clientIP, rl.RedisBased)
			if rl.RedisBased {
				err := redisRateLimiter(clientIP, rl.Requests)
				if err != nil {
					logger.Error("Redis Rate limiter error: %s", err.Error())
					RespondWithError(w, http.StatusTooManyRequests, fmt.Sprintf("%d Too many requests, API rate limit exceeded. Please try again later", http.StatusTooManyRequests), rl.ErrorInterceptor)
					return
				}
				// Proceed to the next handler if rate limit is not exceeded
				next.ServeHTTP(w, r)
			} else {
				rl.mu.Lock()
				client, exists := rl.ClientMap[clientIP]
				if !exists || time.Now().After(client.ExpiresAt) {
					client = &Client{
						RequestCount: 0,
						ExpiresAt:    time.Now().Add(rl.Window),
					}
					rl.ClientMap[clientIP] = client
				}
				client.RequestCount++
				rl.mu.Unlock()

				if client.RequestCount > rl.Requests {
					logger.Error("Too many requests from IP: %s %s %s", clientIP, r.URL, r.UserAgent())
					//Update Origin Cors Headers
					if allowedOrigin(rl.Origins, r.Header.Get("Origin")) {
						w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
					}
					RespondWithError(w, http.StatusTooManyRequests, fmt.Sprintf("%d Too many requests, API rate limit exceeded. Please try again later", http.StatusTooManyRequests), rl.ErrorInterceptor)
					return
				}
			}
			// Proceed to the next handler if rate limit is not exceeded
			next.ServeHTTP(w, r)
		})
	}
}
func redisRateLimiter(clientIP string, rate int) error {
	ctx := context.Background()

	res, err := limiter.Allow(ctx, clientIP, redis_rate.PerMinute(rate))
	if err != nil {
		return err
	}
	if res.Remaining == 0 {
		return errors.New("rate limit exceeded")
	}

	return nil
}
func InitRedis(addr, password string) {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
	})
	limiter = redis_rate.NewLimiter(Rdb)
}
