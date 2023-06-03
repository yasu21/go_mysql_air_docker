# buildステージ
FROM golang:1.20.1-alpine3.17 AS build
WORKDIR /app
COPY go.mod go.sum main.go ./
RUN go mod download
COPY . ./
RUN GOOS=linux go build -ldflags "-w -s" -o main /app/main.go

# deployステージ
FROM alpine:3.17 AS deploy
WORKDIR /app
COPY --from=build /app/main .
CMD [ "/app/main" ]

# localステージ
FROM golang:1.20.1-alpine3.17 AS local
RUN apk add --no-cache git
WORKDIR /app
COPY . ./
RUN go mod tidy && \
    go install github.com/cosmtrek/air@v1.44.0
CMD ["air", "-c", ".air.toml"]
#CMD ["go", "run", "main.go"]