FROM golang:latest
LABEL maintainer="Andres Romo <andres.roqa93@gmail.com>"
RUN apt-get update
WORKDIR /go/src/transactions-stori-challenge
COPY go.mod .
COPY go.sum .
COPY . .
RUN go install
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    MONGO_DATABASE=stori \
    BASIC_AUTH_PASSWORD=AdminStori1234 \
    BASIC_AUTH_USER=storiAdmin \
    MONGO_USER=andres_romo \
    MONGO_PASSWORD=L7ojQOMBllDqLEIW \
    MONGO_HOST="cluster0.qfcrs.mongodb.net/stori?retryWrites\=true&w\=majority" \
    BACKEND_HOST=localhost:8080 \
    SENDGRID_API_KEY=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 \
    PORT=8080
RUN go build
ENTRYPOINT ["./transactions-stori-challenge"]