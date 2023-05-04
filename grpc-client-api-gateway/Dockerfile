FROM golang:1.20.3-alpine3.17  as build-client
WORKDIR /usr/grpc-client-api-gateway
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o ./bin/grpc-client-api-gateway ./cmd/grpc-gateway/
CMD ./bin/grpc-client-api-gateway

FROM alpine:3.17.3
WORKDIR /usr/grpc-client-api-gateway
COPY --from=build-client /usr/grpc-client-api-gateway/bin/grpc-client-api-gateway /usr/grpc-client-api-gateway/bin/
CMD ./bin/grpc-client-api-gateway