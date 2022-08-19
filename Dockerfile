FROM 1.19.0-alpine3.15

ARG SERVICE_PORT=3000
ARG DEBUG=false

WORKDIR /api

# add sources
ADD ./api /api/src/
ADD ./config /api/src/
ADD ./log /api/src/
ADD ./middleware /api/src/
ADD ./resources /api/
ADD ./routes /api/src/
ADD ./utils /api/src/
ADD ./go.mod /api/src/

# add config
ADD ./config.json /api/

# build
WORKDIR /api/src
RUN go mod tidy
RUN go build -o project-api main.go
RUN cp /api/src/project-api /api

WORKDIR /api
RUN rm -rf /api/src

ENTRYPOINT [ "/api/project-api" ]