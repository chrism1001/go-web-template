FROM golang:1.23.1-alpine
WORKDIR /app

RUN go install github.com/a-h/templ/cmd/templ@latest

COPY go.* ./
RUN go version
RUN go mod download
RUN go mod tidy

COPY . .

RUN go build -o main main.go

EXPOSE 8080

CMD ["./main"]
