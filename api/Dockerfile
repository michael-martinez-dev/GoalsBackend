FROM golang:1.18-alpine as build

RUN apk update && apk add make

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY ./main.go ./main.go
COPY ./cmd ./cmd
COPY ./pkg ./pkg

RUN mkdir -p /app/bin
RUN go build -o ./bin/server main.go

FROM alpine:latest

WORKDIR /

COPY --from=build /app/bin/server /server

EXPOSE 9090

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

USER appuser

ENTRYPOINT ["/server"]