package transport

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"errors"
	"log"
	"unicode/utf8"
	"os"

	"github.com/andrew221293/transactions-stori-challenge/internal/entity"

	"github.com/labstack/echo/v4"
	"github.com/aws/aws-lambda-go/events"
	"github.com/labstack/echo/v4/middleware"
)

type (
	Router struct {
		*echo.Echo
		Address string
		Handler EchoHandler
	}
	EchoHandler struct {
		StoriUseCases UseCases
	}
	UseCases struct {
		Stori StoriUsecase
	}
)

//StoriUsecase implement the methods of usecase (business logic)
type StoriUsecase interface {
	ValidateTransaction(
		ctx context.Context,
		transactions []entity.Transaction) (entity.TransactionHistory, error)
}

//Lambda Start
func (r *Router) Start() {
	base := r.Group("/custom-endpoints")
	user := os.Getenv("BASIC_AUTH_USER")
	pass := os.Getenv("BASIC_AUTH_PASSWORD")

	base.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == user && password == pass {
			return true, nil
		}
		return false, entity.CustomError{
			Err:      fmt.Errorf("basic auth failed"),
			HTTPCode: http.StatusUnauthorized,
			Code:     "e6807c42-3568-41de-a15f-fe0f073ab657",
		}
	}))

	//Transactions endpoints
	transaction := base.Group("/transactions")
	transaction.GET("", r.Handler.Transactions)

}


//LocalHost routing
func (r *Router) LocalHost() error {
	base := r.Group("/custom-endpoints")
	user := os.Getenv("BASIC_AUTH_USER")
	pass := os.Getenv("BASIC_AUTH_PASSWORD")

	base.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == user && password == pass {
			return true, nil
		}
		return false, entity.CustomError{
			Err:      fmt.Errorf("basic auth failed"),
			HTTPCode: http.StatusUnauthorized,
			Code:     "e6807c42-3568-41de-a15f-fe0f073ab657",
		}
	}))

	//Transactions endpoints
	transaction := base.Group("/transactions")
	transaction.GET("", r.Handler.Transactions)

	return r.Echo.Start(r.Address)
}

func (r *Router) LambdaHandler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	newReq, err := proxyEventToHTTPRequest(req)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err //improve this
	}
	respWriter := NewProxyResponseWriter()
	r.ServeHTTP(http.ResponseWriter(respWriter), newReq)

	proxyResp, err := respWriter.GetProxyResponse()
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return proxyResp, nil
}

func proxyEventToHTTPRequest(req events.APIGatewayProxyRequest) (*http.Request, error) {
	decodedBody := []byte(req.Body)
	if req.IsBase64Encoded {
		base64Body, err := base64.StdEncoding.DecodeString(req.Body)
		if err != nil {
			return nil, err
		}
		decodedBody = base64Body
	}

	queryString := ""
	if len(req.QueryStringParameters) > 0 {
		queryString = "?"
		queryCnt := 0
		for q := range req.QueryStringParameters {
			if queryCnt > 0 {
				queryString += "&"
			}
			queryString += url.QueryEscape(q) + "=" + url.QueryEscape(req.QueryStringParameters[q])
			queryCnt++
		}
	}

	path := req.Path
	httpRequest, err := http.NewRequest(
		strings.ToUpper(req.HTTPMethod),
		path+queryString,
		bytes.NewReader(decodedBody),
	)

	if err != nil {
		fmt.Printf("Could not convert request %s:%s to http.Request\n", req.HTTPMethod, req.Path)
		log.Println(err)
		return nil, err
	}

	for h := range req.Headers {
		httpRequest.Header.Add(h, req.Headers[h])
	}

	return httpRequest, nil
}

type ProxyResponseWriter struct {
	headers http.Header
	body    []byte
	status  int
}

func NewProxyResponseWriter() *ProxyResponseWriter {
	return &ProxyResponseWriter{
		headers: make(http.Header),
		status:  http.StatusOK,
	}
}

// Header implementation from the http.ResponseWriter interface.
func (r *ProxyResponseWriter) Header() http.Header {
	return r.headers
}

// Write sets the response body in the object. If no status code
// was set before with the WriteHeader method it sets the status
// for the response to 200 OK.
func (r *ProxyResponseWriter) Write(body []byte) (int, error) {
	r.body = body
	if r.status == -1 {
		r.status = http.StatusOK
	}

	return len(body), nil
}

// WriteHeader sets a status code for the response. This method is used
// for error responses.
func (r *ProxyResponseWriter) WriteHeader(status int) {
	r.status = status
}

func (r *ProxyResponseWriter) GetProxyResponse() (events.APIGatewayProxyResponse, error) {
	if len(r.headers) == 0 {
		return events.APIGatewayProxyResponse{}, errors.New("No headers generated for response")
	}

	var output string
	isBase64 := false

	if utf8.Valid(r.body) {
		output = string(r.body)
	} else {
		output = base64.StdEncoding.EncodeToString(r.body)
		isBase64 = true
	}

	proxyHeaders := make(map[string]string)

	for h := range r.headers {
		proxyHeaders[h] = r.headers.Get(h)
	}

	return events.APIGatewayProxyResponse{
		StatusCode:      r.status,
		Headers:         proxyHeaders,
		Body:            output,
		IsBase64Encoded: isBase64,
	}, nil
}
