FROM golang:latest

WORKDIR /user-interaction-service

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build

CMD ["./user-interaction-service"]