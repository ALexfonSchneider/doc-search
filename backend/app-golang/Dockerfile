FROM golang:1.21.6-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main ./cmd/service/service.go

CMD ["./main"]