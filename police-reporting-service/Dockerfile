FROM golang:latest

RUN mkdir /app

ENV PORT=$port
ENV PORT2=$port2


COPY . /app
WORKDIR /app

RUN export GO111MODULE=on
RUN go mod tidy

EXPOSE $port  $port2
CMD go run ./cmd 

