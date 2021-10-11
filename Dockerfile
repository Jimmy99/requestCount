FROM golang:1.16

ADD . /app
WORKDIR /app
RUN go build -o app
CMD ["/app/app"]

EXPOSE 5000

