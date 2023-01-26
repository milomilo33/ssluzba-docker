FROM golang:1.19-alpine

ENV GO111MODULE=on

WORKDIR /app

COPY . .

RUN go get

RUN go build -o /uns

EXPOSE 8080

CMD [ "/uns" ]