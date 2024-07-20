FROM golang:1.22.3

WORKDIR /main/server

COPY . .

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
