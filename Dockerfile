FROM golang:1.23.3

WORKDIR /app

COPY go.mod ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o main ./cmd/bss/

CMD ["/app/main"]
