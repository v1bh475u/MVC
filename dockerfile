FROM golang:latest

RUN mkdir /app

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod tidy

COPY . .

EXPOSE 8080

CMD ["go", "run", "cmd/main.go"]