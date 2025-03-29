FROM golang:1.24.1-alpine3.20 as build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY main.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:3.21

COPY --from=build /app/main ./main

CMD ["./main"]

