FROM golang:1.18.1-alpine3.15 AS build-env

ENV GO111MODULE=on

WORKDIR /workspace/manager

COPY ./.proto/ ./.proto
COPY ./model/ ./model
COPY ./repository/ ./repository
COPY ./service/ ./service
COPY ./utils/ ./utils
COPY ./manager.go ./manager.go
COPY ./go.mod ./go.mod

RUN go mod tidy
RUN go build -o ./application .

FROM alpine:3.15

WORKDIR /app

COPY --from=build-env /workspace/manager/application /manager

ENV PASSWORD="123456"
ENV MANAGER_URL="https://localhost:8081"

CMD /manager --password ${PASSWORD} --managerUrl ${MANAGER_URL}