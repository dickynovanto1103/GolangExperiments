FROM golang:1.12.7
WORKDIR "/app"
COPY go* ./
RUN go mod download
COPY ./ ./
RUN go build -o main main.go
CMD ["./main"]