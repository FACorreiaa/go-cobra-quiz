FROM golang:alpine

WORKDIR /app

COPY . /app

RUN go build -o go-quiz-cli

ENTRYPOINT ["./go-quiz-cli"]
