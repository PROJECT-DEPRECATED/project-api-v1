FROM golang:1.19.0-alpine3.15

WORKDIR /api
COPY . .

RUN rm Dockerfile docker-compose.yml

RUN apk update
RUN apk add git
RUN apk add ca-certificates

RUN go mod tidy
RUN go build -o /api/project-api /api/main.go

ENTRYPOINT [ "/api/project-api" ]
