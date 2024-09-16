FROM golang:1.23.1-alpine
WORKDIR /app
  
RUN apk --no-cache add curl \
  && curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-x64 \
  && mv tailwindcss-linux-x64 tailwindcss && chmod +x tailwindcss

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

COPY tailwind.config.js ./

RUN go build -o main main.go

EXPOSE 8080

CMD ["./main"]
