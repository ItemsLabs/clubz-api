FROM golang:1.22-alpine

LABEL maintainer="Sergey Shcherbina <trytoheal@gmail.com>"

WORKDIR /app

RUN apk update && apk add gcc musl-dev git

# Copy go mod files first for better caching
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o clubz-api .

CMD ["./clubz-api"]