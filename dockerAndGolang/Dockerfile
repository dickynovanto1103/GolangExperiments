FROM golang

COPY . /go/src
WORKDIR /go/src

RUN go build main.go

CMD ["./main"]