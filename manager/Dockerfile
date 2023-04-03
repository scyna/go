FROM golang:1.18.4-alpine3.16 AS build-env

ENV GO111MODULE=on

WORKDIR /workspace/manager

COPY . .

RUN go mod tidy

RUN go build -o ./application .

FROM alpine:3.16

WORKDIR /app

COPY --from=build-env /workspace/manager/application /manager

ENV PASSWORD="123456"
ENV MANAGER_URL="https://localhost:8081"

CMD /manager --password ${PASSWORD} --managerUrl ${MANAGER_URL}