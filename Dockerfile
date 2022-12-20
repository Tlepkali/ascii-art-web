FROM golang:1.19

WORKDIR /ascii-art-web

COPY . .

EXPOSE 8080

RUN go build -o main.go .

CMD ["./main.go"]