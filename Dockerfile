FROM golang:1.20.6

WORKDIR /app

RUN apt-get update && apt-get install -y git

COPY go.mod go.sum ./
RUN go install -tags mysql github.com/golang-migrate/migrate/v4/cmd/migrate@latest
RUN go mod download

EXPOSE 8080

CMD ["go", "run", "main.go"]
