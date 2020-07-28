FROM golang
WORKDIR "/app"
COPY go.* ./
RUN go mod download
COPY ./ ./
RUN go build -o main main.go
CMD ["./main"]