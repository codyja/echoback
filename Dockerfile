FROM golang:1.18.1-alpine3.15

WORKDIR /app

COPY *.go ./

RUN go build -o /echoback main.go

EXPOSE 8080

CMD [ "/echoback" ]