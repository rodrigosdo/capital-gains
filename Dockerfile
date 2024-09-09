FROM golang:1.23-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o ./bin/capital-gains ./cmd/cli/main.go

ENTRYPOINT ["./bin/capital-gains"]
