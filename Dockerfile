FROM golang:1.22.1-alpine

WORKDIR /

COPY . .

RUN go mod download

RUN go build -o main .

CMD ["/main"]