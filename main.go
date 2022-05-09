package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/andrew221293/transactions-stori-challenge/internal/entity"
	"github.com/andrew221293/transactions-stori-challenge/internal/store"
	"github.com/andrew221293/transactions-stori-challenge/internal/transport"
	"github.com/andrew221293/transactions-stori-challenge/internal/usecase"

	"github.com/labstack/echo/v4"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/labstack/echo/v4/middleware"
)

var genericErrResponse = entity.ResponseError{
	Error: "something went wrong",
	Code:  "959a1908-62f0-4cad-afc0-d9b4300085db",
}

func main() {
	isLocalEnvironment := os.Getenv("_LAMBDA_SERVER_PORT") == "" && os.Getenv("_AWS_LAMBDA_RUNTIME_API") == ""
	loc, _ := time.LoadLocation("America/New_York")
	time.Local = loc // -> this is setting the global timezone
	
	ctx := context.Background()
	mongoUser := os.Getenv("MONGO_USER")
	mongoPass := os.Getenv("MONGO_PASSWORD")
	mongoHost := os.Getenv("MONGO_HOST")
	address := os.Getenv("BACKEND_HOST")

	// uri to connect with mongoDB
	uri := fmt.Sprintf(
		"mongodb+srv://%s:%s@%s", //?maxPoolSize=%s",
		mongoUser,
		mongoPass,
		mongoHost,
	)

	// initialize store (DB)
	storiStore, err := store.NewStoriStore(ctx, uri)
	if err != nil {
		log.Fatalf("failed to initialize the stori store: %v", err)
	}

	e := echo.New()
	e.HTTPErrorHandler = customErrorHandler
	e.Use(middleware.CORS())
	e.Use(middleware.Recover())
	e.Pre(middleware.RemoveTrailingSlash())

	router := &transport.Router{
		Echo:    e,
		Address: address,
		Handler: transport.EchoHandler{
			StoriUseCases: transport.UseCases{
				Stori: usecase.StoriUseCase{
					Store: storiStore,
				},
			},
		},
	}
	
	if !isLocalEnvironment {
		router.Start()
		lambda.Start(router.LambdaHandler)
	} else {
		go func() {
			if err := router.LocalHost(); err != nil {
				e.Logger.Infof("Shutting down ser")
			}
		}()

		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt)
		<-quit
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := e.Shutdown(ctx); err != nil {
			e.Logger.Fatal(err)
		}
	}
}

func customErrorHandler(err error, e echo.Context) {
	var ce entity.CustomError
	if errors.As(err, &ce) {
		e.JSON(ce.HTTPCode, ce.ToResponseError()) // nolint: errcheck
		return
	}
	e.JSON(http.StatusInternalServerError, genericErrResponse) // nolint: errcheck
}
