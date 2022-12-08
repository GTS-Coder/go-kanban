FROM golang:1.17-alpine3.14 AS builder
WORKDIR /app
COPY  go.mod .
COPY  go.sum .
RUN go mod download
COPY . .

RUN go build -o ./out/dist/ .

CMD ./out/dist
