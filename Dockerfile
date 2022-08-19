FROM golang:1.19.0-alpine3.15

ENV SERVICE_PORT=3000
ENV DEBUG=false

WORKDIR /api

ADD ./config.json /api/
ADD ./resources /api/

# add execution file
ADD ./project-api /api/

ENTRYPOINT [ "/api/project-api -port=${SERVICE_PORT} -debug=${DEBUG}" ]