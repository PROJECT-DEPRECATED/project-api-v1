FROM golang:1.19.1-alpine3.16

WORKDIR /api/src

COPY . .

RUN apk update
RUN apk add git
RUN apk add ca-certificates

RUN go mod tidy
RUN go build -o /api/project-api /api/src/main.go

RUN cp ./config.json ../

WORKDIR /api

RUN rm -rf ./src

ENTRYPOINT [ "/api/project-api" ]
