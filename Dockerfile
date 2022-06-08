## Build
FROM golang:1.18.2-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY lib ./lib
COPY server ./server
COPY Makefile ./

RUN make build

## Deploy
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /go/bin/server /server

EXPOSE 8080

ENTRYPOINT ["/server"]
