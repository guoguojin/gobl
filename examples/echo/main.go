package main

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"

	"gitlab.com/gobl/gobl/pkg/bootstrap"
	"gitlab.com/gobl/gobl/pkg/service"

	"gitlab.com/gobl/gobl/pkg/cmd"
	"gitlab.com/gobl/gobl/pkg/logger"
)

const (
	usage = `echo`
	short = `echo service`
	long  = `echo service echos whatever it is given`
)

type echoState struct {
	service.State
}

var eState *echoState

func initHTTP(ctx context.Context, state service.State) error {
	log := logger.New(logger.ApplicationLogLevel(), logger.ConfiguredLumberjackLogger())

	r := chi.NewRouter()
	r.Use(middleware.RequestID, middleware.RealIP, middleware.Logger, middleware.Recoverer, middleware.Timeout(time.Minute))

	r.Get("/echo/{message}", func(w http.ResponseWriter, r *http.Request) {
		message := chi.URLParam(r, "message")
		if message == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(nil)

			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(message))
	})

	go func() {
		log.Info("Starting echo server on port: 8080")
		if err := http.ListenAndServe(":8080", r); err != nil {
			log.Error("HTTP Server terminated", zap.Error(err))
		}
	}()

	return nil
}

func main() {
	app := bootstrap.New().
		AddInitFunc(initHTTP)

	cmd.SetCliProperties(usage, short, long)

	eState = &echoState{}

	cmd.Execute(context.Background(), app, eState)
}
