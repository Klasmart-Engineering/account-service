FROM golang:1.18-alpine

WORKDIR /app

COPY . .

RUN go install -v ./...

RUN go build -o ./account-service ./main.go

EXPOSE 8080

CMD [ "./account-service" ]