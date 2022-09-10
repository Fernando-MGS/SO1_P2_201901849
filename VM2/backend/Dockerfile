# syntax=docker/dockerfile:1
#Version go
FROM golang:1.19

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /docker-backend-go

##DEPLOY
EXPOSE 4000

CMD [ "/docker-backend-go" ]