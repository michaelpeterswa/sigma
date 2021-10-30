FROM golang:1.16

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./cmd/sigma
RUN go build ./cmd/sigma

CMD ./sigma