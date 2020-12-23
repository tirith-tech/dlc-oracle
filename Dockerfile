FROM golang:alpine as build

COPY . /usr/local/go/src/github.com/tirith-tech/dlc-oracle
WORKDIR /usr/local/go/src/github.com/tirith-tech/dlc-oracle
RUN go mod download
RUN go build

CMD ["go run util/createkey.go"]