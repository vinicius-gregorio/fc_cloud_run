FROM golang:1.23.1
WORKDIR /app

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o server cmd/main.go

CMD [ "./server" ]