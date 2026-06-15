FROM golang:1.26

WORKDIR /app
COPY . .

RUN go build -o bin/goterp ./cmd/goterp

CMD ["make", "test"]
