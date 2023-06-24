FROM golang:1.20.5

WORKDIR /app

RUN apt-get update && apt-get install -y git

COPY go.mod go.sum ./
RUN go mod download

EXPOSE 8080

CMD ["go", "run", "main.go"]
