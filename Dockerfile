FROM golang:1.19.0-alpine3.15

WORKDIR /api

ADD ./config.json /api/
ADD ./resources /api/

# add execution file
ADD ./project-api /api/

ENTRYPOINT [ "./project-api" ]