FROM golang:1.10-alpine

RUN mkdir /app

ADD . /app/

WORKDIR /app

RUN go build -o main .

EXPOSE 8089

CMD ["/app/main"]
