#Build Stage
FROM golang:1.21.5-alpine3.19 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

#Run Stage
FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 8000
CMD [ "/app/main" ]