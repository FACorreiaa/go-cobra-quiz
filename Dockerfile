FROM golang:alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -v -o api

FROM golang:alpine

WORKDIR /app

COPY --from=builder /app/api .

EXPOSE 8080

CMD ["./api", "--config-source=file","--config-file=config/config_dev.yml"]



