# GO Fiber App
FROM golang:1.19.0
WORKDIR /app
COPY . .
RUN go mod download
CMD ["go", "run", "main.go"]
EXPOSE 80