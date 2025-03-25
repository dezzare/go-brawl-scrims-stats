FROM golang:1.23.3-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o main ./cmd/bss/

EXPOSE 5000

CMD ["/app/main"]
