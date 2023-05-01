FROM golang:1.20.3-alpine3.17 
# as build-server
WORKDIR /usr/grpc-server
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o ./bin/grpc-server ./cmd/sevice-server/

ENV WAIT_VERSION 2.7.2
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/$WAIT_VERSION/wait /wait
RUN chmod +x /wait

# CMD ./bin/grpc-server

# FROM alpine:3.17.3
# WORKDIR /usr/grpc-server
# COPY --from=build-server /usr/grpc-server/bin/grpc-server /usr/grpc-server/bin/
# COPY --from=build-server /usr/grpc-server/schema/ /usr/grpc-server/schema/
# CMD ./bin/grpc-server
