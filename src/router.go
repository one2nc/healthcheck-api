package healthcheck

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/projectdiscovery/gologger"
)

func InitializeRouter() {
	serverAddr := os.Getenv("SERVER_ADDR")
	if serverAddr == "" {
		serverAddr = "localhost:8080"
	}

	router := gin.Default()
	InitializeRoutes(router)
	srv := &http.Server{
		Addr:    serverAddr,
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			gologger.Info().Msgf(err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 5 seconds.
	gologger.Info().Msg("Shutdown server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		gologger.Fatal().Msg(err.Error())
	}
}
