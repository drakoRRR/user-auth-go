FROM golang:1.22.0-alpine

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

EXPOSE ${PORT}

CMD ["go", "run", "cmd/main.go"]
