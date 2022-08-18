FROM golang:1.18.4-alpine3.16 AS build-env

ENV GO111MODULE=on

WORKDIR /workspace/engine

COPY ./gateway/ ./gateway
COPY ./proxy/ ./proxy
COPY ./manager/ ./manager
COPY ./engine.go ./engine.go
COPY ./go.mod ./go.mod

RUN go mod tidy
RUN go build -o ./application .

FROM alpine:3.16

WORKDIR /app

COPY --from=build-env /workspace/engine/application /engine

ENV MANAGER_PORT=8081
ENV PROXY_PORT=8080
ENV GATEWAY_PORT=8443

ENV NATS_URL=localhost
ENV NATS_USERNAME=""
ENV NATS_PASSWORD=""
ENV DB_HOST=""
ENV DB_USERNAME=""
ENV DB_PASSWORD=""
ENV DB_LOCATION=""
ENV CERTIFICATE_FILE="/cert/localhost.crt"
ENV CERTIFICATE_KEY="/cert/localhost.key"

EXPOSE ${MANAGER_PORT}
EXPOSE ${PROXY_PORT}
EXPOSE ${GATEWAY_PORT}

CMD /engine \
    --manager_port=${MANAGER_PORT} --proxy_port=${PROXY_PORT} --gateway_port=${GATEWAY_PORT} \
    --nats_url=${NATS_URL} --nats_username=${NATS_USERNAME} --nats_password=${NATS_PASSWORD} \
    --db_host=${DB_HOST} --db_username=${DB_USERNAME} --db_password=${DB_PASSWORD}  --db_location=${DB_LOCATION} \
    --certificateFile=${CERTIFICATE_FILE} --certificateKey=${CERTIFICATE_KEY}