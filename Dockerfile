FROM golang:1.19.4-alpine3.16

WORKDIR /usr/local/api/src
COPY . .

RUN apk add make

RUN ./configure
RUN make
RUN cp ./project-api ../ && cp ./config.json ../

WORKDIR /usr/local/api
RUN rm -rf ./src

ENTRYPOINT [ "/usr/local/api/project-api" ]
