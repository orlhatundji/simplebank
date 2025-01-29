# Build stage
FROM golang:1.23.5-alpine3.21
WORKDIR /app
COPY . .
RUN go build -o main main.go

# Run stage
FROM alpine:3.21
WORKDIR /app
COPY --from=0 /app/main .
COPY app.env .

EXPOSE 8080
CMD ["/app/main"]