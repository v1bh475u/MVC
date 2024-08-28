FROM golang:latest

RUN mkdir /app

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod tidy

COPY . .

EXPOSE 8080
RUN go build -o mvc cmd/main.go
# RUN go run config/admincred.go
CMD ["/app/mvc"]