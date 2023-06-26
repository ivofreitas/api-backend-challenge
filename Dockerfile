FROM golang:1.18-alpine as builder

WORKDIR /app

COPY . .

RUN go get github.com/swaggo/swag/gen@v1.16.1
RUN go install github.com/swaggo/swag/cmd/swag@v1.16.1
RUN swag init ./main.go
RUN go build -o main.bin main.go
FROM alpine as release

WORKDIR /app

COPY --from=builder /app/main.bin /app/main.bin

EXPOSE 8080

ENTRYPOINT ["./main.bin"]
