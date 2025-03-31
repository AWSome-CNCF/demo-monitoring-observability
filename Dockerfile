FROM golang:1.24.1-alpine3.20 as build

WORKDIR /app

COPY ./src/go.mod ./src/go.sum ./
RUN go mod download

COPY ./src/ ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main .

FROM alpine:3.21

WORKDIR /app

COPY --from=build /app/main /app/main

CMD ["./main"]

