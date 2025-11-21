package rest

import (
	"context"
	"fmt"
	"net/http"
	"time"

	_ "github.com/lib/pq"
	"github.com/rs/cors"

	"dm_loanservice/drivers/goconf"
	"dm_loanservice/drivers/logger"
	"dm_loanservice/internal/endpoint"
)

func RunServer(
	ctx context.Context,
	e endpoint.Endpoints,
	port string,
) error {

	// initialize routes
	r := initRoutes(ctx, e)

	hosts := []string{"http://localhost:3000", "http://localhost:3002", "https://th.grasp.co.uk", "https://www.dmortgages.com", "https://dmortgages.com"}

	co := cors.New(cors.Options{
		AllowedOrigins: hosts,
		AllowedMethods: []string{"GET", "PATCH", "POST", "PUT", "OPTIONS", "DELETE"},
		AllowedHeaders: []string{
			"Content-Type", "Content-Length", "Accept-Encoding", "X-Csrf-Token", "Authorization",
			"x-api-key", "Cookie-Type", "Accept-Language", "Origin", "Accept", "Accept-Encoding",
			"User-Agent", "Authorization", "Referer",
		},
		ExposedHeaders:   []string{"Content-Length", "Access-Control-Allow-Origin", "X-Csrf-Token"},
		AllowCredentials: true,
		Debug:            true,
	})

	handler := co.Handler(r)

	// initialize rest server
	server := initServer(handler, port)

	logger.LogInfo("status-REST", "listening", "port", port)
	return server.ListenAndServe()
}

func initServer(r http.Handler, port string) *http.Server {
	addr := fmt.Sprintf(":%s", port)
	server := &http.Server{
		Addr:         addr,
		WriteTimeout: time.Duration(goconf.Config().GetInt("rest.write_timeout")) * time.Second,
		ReadTimeout:  time.Duration(goconf.Config().GetInt("rest.read_timeout")) * time.Second,
		IdleTimeout:  time.Duration(goconf.Config().GetInt("rest.idle_timeout")) * time.Second,
		Handler:      r,
	}
	return server
}
