FROM golang:1.20.3-alpine3.17  as build-server
WORKDIR /usr/service-user
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o ./bin/service-user ./cmd/sevice-user/
CMD ./bin/service-user

FROM alpine:3.17.3
WORKDIR /usr/service-user
ENV WAIT_VERSION 2.7.2
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/$WAIT_VERSION/wait /wait
RUN chmod +x /wait
COPY --from=build-server /usr/service-user/bin/service-user /usr/service-user/bin/
COPY --from=build-server /usr/service-user/schema/ /usr/service-user/schema/
CMD ./bin/service-user
