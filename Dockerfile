FROM golang:1.18-alpine

WORKDIR /app

COPY . .

RUN go install -v ./...

RUN go build -o ./accounts-service ./src/main.go

EXPOSE 8080

CMD [ "./accounts-service" ]