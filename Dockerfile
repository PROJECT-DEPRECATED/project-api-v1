FROM golang:1.19.3-alpine3.16

WORKDIR /api/src
COPY . .

RUN go mod tidy
RUN go build -o /api/project-api /api/src/main.go
RUN cp ./config.json ../

WORKDIR /api
RUN rm -rf ./src

ENTRYPOINT [ "/api/project-api" ]
