FROM golang:1.18-alpine as build

RUN apk update && apk add make

WORKDIR /app

COPY ./api/go.mod ./
COPY ./api/go.sum ./

RUN go mod download

COPY ./api/main.go ./main.go
COPY ./api/cmd ./cmd
COPY ./api/pkg ./pkg

RUN mkdir -p /app/bin
RUN go build -o ./bin/server main.go

FROM alpine:latest

WORKDIR /

COPY --from=build /app/bin/server /server
COPY ./.env /.env

EXPOSE 9090

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

USER appuser

ENTRYPOINT ["/server"]