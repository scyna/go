FROM golang:1.19.1-alpine3.16 AS build-env

ENV GO111MODULE=on

WORKDIR  /workspace/vinfast



WORKDIR  /workspace/vinfast

RUN go mod tidy
RUN go get github.com/scyna/go/scyna@latest

RUN go build -o ./application .

FROM alpine:3.16
WORKDIR /app

COPY --from=build-env /workspace/vinfast/account/application ./application

ENV PASSWORD="123456789aA@#"
ENV MANAGER_URL="https://localhost:8081"

CMD /app/application --password ${PASSWORD} --managerUrl ${MANAGER_URL}