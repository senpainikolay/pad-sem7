FROM golang:latest

RUN mkdir /app

ENV PORT=$port

COPY . /app
WORKDIR /app

RUN export GO111MODULE=on
RUN go mod tidy

EXPOSE $port
CMD go run ./cmd 

