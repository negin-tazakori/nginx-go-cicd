# Stage1: Build
FROM golang:1.25-alpine AS build

WORKDIR /app

COPY go.mod .
COPY main.go .

RUN CGO_ENABLED=0 go build -o server .



# Stige2: RUNtime
FROM alpine:3.20

COPY --from=build /app/server  /server

EXPOSE 8080

CMD ["/server"]
