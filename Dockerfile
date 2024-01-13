FROM golang:1.21

COPY . /app
WORKDIR /app
RUN go install

RUN GOOS=linux GOARCH=amd64 go build main.go

CMD ["./main"]