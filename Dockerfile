FROM golang:1.25 AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/go-users-simple-crud

FROM alpine:latest

RUN apk --no-cache add ca-certificates curl

WORKDIR /app

COPY --from=build /app/go-users-simple-crud .

EXPOSE 8000 6060

CMD ["/app/go-users-simple-crud"]
